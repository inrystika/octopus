package minio

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"path/filepath"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/common/errors"
	"strings"
	"time"

	"github.com/minio/madmin-go"

	"server/common/log"

	miniogo "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio interface {
	// 创建桶
	CreateBucket(bucketName string) error
	// 删除桶
	DeleteBucket(bucketName string) error
	// 生成上传对象url
	PresignedUploadObject(bucketName string, objectName string, domain string) (*url.URL, error)
	// 生成下载对象url
	PresignedDownloadObject(bucketName string, objectName string, domain string) (*url.URL, error)
	// 查看对象
	ListObjects(bucketName string, objectPrefix string, recurSvie bool) ([]*ObjectInfo, error)
	// 查看对象是否存在
	ObjectExist(bucketName string, objectName string) (bool, error)
	// RemoveObject 删除对象
	RemoveObject(bucketName string, objectName string) (bool, error)

	CreateOrUpdateAccount(ctx context.Context, userName string, password string) error

	BucketExists(ctx context.Context, bucketName string) (bool, error)

	SetUserBucketsAccess(ctx context.Context, userName string, buckets []string) error
}

type ObjectInfo struct {
	Name         string
	LastModified int64
	Size         int64
	ContentType  string
}

type minio struct {
	log         *log.Helper
	conf        *conf.Data
	client      *miniogo.Client
	adminClient *madmin.AdminClient
}

func NewMinio(conf *conf.Data, logger log.Logger) Minio {
	client, err := miniogo.New(conf.Minio.Base.EndPoint, &miniogo.Options{
		Creds:  credentials.NewStaticV4(conf.Minio.Base.AccessKeyID, conf.Minio.Base.SecretAccessKey, ""),
		Secure: conf.Minio.Base.UseSSL,
	})
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioBucketInitFailed)
		panic(err)
	}

	adminClient, err := madmin.New(conf.Minio.Base.EndPoint, conf.Minio.Base.AccessKeyID, conf.Minio.Base.SecretAccessKey, false)
	if err != nil {
		panic(err)
	}

	minio := &minio{
		log:         log.NewHelper("Minio", logger),
		conf:        conf,
		client:      client,
		adminClient: adminClient,
	}

	// 创建默认的桶
	err = minio.CreateBucket(common.BUCKET)
	if err != nil {
		panic(err)
	}

	return minio
}

// 创建桶
func (m *minio) CreateBucket(bucketName string) error {
	ctx := context.Background()

	isExist, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckBucketExistFailed)
		return err
	}
	if isExist {
		err = errors.Errorf(err, errors.ErrorMinioBucketExisted)
		m.log.Warnw(ctx, err)
		return nil
	}

	err = m.client.MakeBucket(ctx, bucketName, miniogo.MakeBucketOptions{
		Region:        "us-east-1",
		ObjectLocking: false,
	})
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioMakeBucketFailed)
		return err
	}

	m.log.Infof(ctx, "successfully created mybucket, bucketName=%s", bucketName)
	return nil
}

// 删除桶
func (m *minio) DeleteBucket(bucketName string) error {
	ctx := context.Background()

	isExist, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckBucketExistFailed)
		return err
	}
	if !isExist {
		err = errors.Errorf(err, errors.ErrorMinioBucketNotExist)
		m.log.Errorw(ctx, err)
		return nil
	}

	err = m.client.RemoveBucket(ctx, bucketName)
	if err != nil {
		err := errors.Errorf(err, errors.ErrorMinioDeleteBucketFailed)
		return err
	}

	m.log.Infof(ctx, "successfully delete mybucket, bucketName=%s", bucketName)
	return nil
}

// 生成上传对象url
func (m *minio) PresignedUploadObject(bucketName string, objectName string, domain string) (*url.URL, error) {
	ctx := context.Background()

	uploadExpiry := time.Duration(m.conf.Minio.Business.UploadExpiry) * time.Second // duration单位是纳秒，所有得换算下
	uri, err := m.client.PresignedPutObject(ctx, bucketName, objectName, uploadExpiry)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioPresignedPutObjectFailed)
		return nil, err
	}

	m.log.Infof(ctx, "successfully PresignedPutObject, bucketName=%s|url=%s", bucketName, uri)

	if domain == "" {
		return uri, nil
	}

	domainUrl, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	uri.Path = path.Join(m.conf.Minio.Base.ProxyPath, uri.Path)
	uri.Host = domainUrl.Host
	uri.Scheme = domainUrl.Scheme
	m.log.Infof(ctx, "successfully PresignedPutObject change domain, bucketName=%s|url=%s", bucketName, uri)
	return uri, nil
}

// 生成下载对象url
func (m *minio) PresignedDownloadObject(bucketName string, objectName string, domain string) (*url.URL, error) {
	ctx := context.Background()

	reqParams := make(url.Values)
	paramKey := "response-content-disposition"
	paramVale := fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(objectName))
	reqParams.Set(paramKey, paramVale)
	downloadExpiry := time.Duration(m.conf.Minio.Business.DownloadExpiry) * time.Second // duration单位是纳秒，所有得换算下
	uri, err := m.client.PresignedGetObject(ctx, bucketName, objectName, downloadExpiry, reqParams)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioPresignedGetObjectFailed)
		return nil, err
	}

	m.log.Infof(ctx, "successfully PresignedGetObject, bucketName=%s|url=%s", bucketName, uri)

	if domain == "" {
		return uri, nil
	}

	domainUrl, err := url.Parse(domain)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorUrlParseFailed)
		return nil, err
	}
	uri.Path = path.Join(m.conf.Minio.Base.ProxyPath, uri.Path)
	uri.Host = domainUrl.Host
	uri.Scheme = domainUrl.Scheme
	m.log.Infof(ctx, "successfully PresignedGetObject change domain, bucketName=%s|url=%s", bucketName, uri)
	return uri, nil
}

// 查看对象
func (m *minio) ListObjects(bucketName string, objectPrefix string, recurSvie bool) ([]*ObjectInfo, error) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	opts := miniogo.ListObjectsOptions{
		Prefix:    objectPrefix,
		Recursive: recurSvie,
		MaxKeys:   0, // 暂时没用，先随便赋值
		UseV1:     false,
	}

	var objects []*ObjectInfo
	objectCh := m.client.ListObjects(ctx, bucketName, opts)
	for object := range objectCh {
		if object.Err != nil {
			err := errors.Errorf(object.Err, errors.ErrorMinioListObjectFailed)
			return nil, err
		}

		objects = append(objects, &ObjectInfo{
			Name:         object.Key,
			LastModified: object.LastModified.Unix(),
			Size:         object.Size,
			ContentType:  object.ContentType,
		})
	}

	m.log.Infof(ctx, "successfully ListObjects, bucketName=%s|objectPrefix=%s", bucketName, objectPrefix)
	return objects, nil
}

// 查看对象是否存在
func (m *minio) ObjectExist(bucketName string, objectName string) (bool, error) {
	ctx := context.Background()

	isExist, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckBucketExistFailed)
		return false, err
	}
	if !isExist {
		err = errors.Errorf(err, errors.ErrorMinioBucketNotExist)
		m.log.Errorw(ctx, err)
		return false, nil
	}

	_, err = m.client.StatObject(ctx, bucketName, objectName, miniogo.StatObjectOptions{})
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckObjectExistFailed)
		m.log.Errorw(ctx, err)
		return false, nil
	}

	return true, nil
}

// RemoveObject 删除对象
func (m *minio) RemoveObject(bucketName string, objectName string) (bool, error) {
	ctx := context.Background()

	m.log.Info(ctx, "RemoveObject param. bucketName:", bucketName, ", objectName:", objectName)
	// 安全校验逻辑，判断object至少位于3层路径之下，防止大范围删除动作
	if strings.Count(objectName, "/") < 3 {
		m.log.Error(ctx, "RemoveObject not safe Path depth. bucketName:", bucketName, ", objectName:", objectName)
		return true, nil
	}
	// 安全校验逻辑，判断object非如下路径
	if strings.Contains(objectName, fmt.Sprintf("%s/%s/", common.MODEL_FOLDER, common.PREAB_FOLDER)) ||
		strings.Contains(objectName, fmt.Sprintf("%s/%s/", common.CODE_FOLDER, common.PREAB_FOLDER)) ||
		strings.Contains(objectName, fmt.Sprintf("%s/%s/", common.DATASET_FOLDER, common.PREAB_FOLDER)) {

		m.log.Error(ctx, "RemoveObject not safe Path. bucketName:", bucketName, ", objectName:", objectName)
		return true, nil
	}
	time.Sleep(360 * time.Second)
	isExist, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckBucketExistFailed)
		m.log.Error(ctx, "RemoveObject check bucket exists failed. bucketName:", bucketName, ", error:", err)
		return false, err
	}
	if !isExist {
		m.log.Error(ctx, "RemoveObject bucket not exists. bucketName:", bucketName)
		return true, nil
	}
	removeOpts := miniogo.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	_, err = m.client.StatObject(ctx, bucketName, objectName, miniogo.StatObjectOptions{})
	if err != nil {
		objectPrefix := objectName + "/"
		listOpts := miniogo.ListObjectsOptions{
			Prefix:    objectPrefix,
			Recursive: true,
			MaxKeys:   0, // 暂时没用，先随便赋值
			UseV1:     false,
		}
		objectCh := m.client.ListObjects(ctx, bucketName, listOpts)
		for object := range objectCh {
			m.log.Info(ctx, "RemoveObject bucketName:", bucketName, ", objectName:", object.Key)
			err := m.client.RemoveObject(ctx, bucketName, object.Key, removeOpts)
			if err != nil {
				m.log.Error(ctx, "RemoveObject failed. objectName:", object.Key, ", error:", err)
				return false, errors.Errorf(err, errors.ErrorMinioRemoveObjectFailed)
			}
		}
	} else {
		m.log.Info(ctx, "RemoveObject bucketName:", bucketName, ", objectName:", objectName)
		err = m.client.RemoveObject(ctx, bucketName, objectName, removeOpts)
		if err != nil {
			m.log.Error(ctx, "RemoveObject failed. objectName:", objectName, ", error:", err)
			return false, errors.Errorf(err, errors.ErrorMinioRemoveObjectFailed)
		}
	}
	return true, nil
}

func (m *minio) CreateOrUpdateAccount(ctx context.Context, userName string, password string) error {
	err := m.adminClient.AddUser(ctx, userName, password)
	if err != nil {
		return errors.Errorf(err, errors.ErrorMinioCreateAccountFailed)
	}
	return nil
}

func (m *minio) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	isExist, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return false, errors.Errorf(err, errors.ErrorMinioCheckBucketExistFailed)
	}

	return isExist, nil
}

func (m *minio) SetUserBucketsAccess(ctx context.Context, userName string, buckets []string) error {
	rs := make([]string, 0)
	for _, b := range buckets {
		rs = append(rs, fmt.Sprintf(`"arn:aws:s3:::%s/*"`, b))
	}

	policy := fmt.Sprintf(`{"Version": "2012-10-17","Statement": [{"Action": ["s3:*"],"Effect": "Allow","Resource": [%s]}]}`, strings.Join(rs, ","))
	fmt.Println(policy)
	err := m.adminClient.AddCannedPolicy(ctx, userName, []byte(policy))
	if err != nil {
		return errors.Errorf(nil, errors.ErrorMinioOperationFailed)
	}

	err = m.adminClient.SetPolicy(ctx, userName, userName, false)
	if err != nil {
		return errors.Errorf(nil, errors.ErrorMinioOperationFailed)
	}

	return nil
}
