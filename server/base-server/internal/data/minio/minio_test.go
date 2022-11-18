package minio_test

import (
	"context"
	"fmt"
	"server/base-server/internal/conf"
	"server/base-server/internal/data/minio"
	"server/common/errors"
	"server/common/utils"
	"testing"

	"github.com/minio/madmin-go"

	"server/common/log"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gopkg.in/yaml.v2"
)

func confInit(t *testing.T) *conf.Bootstrap {
	c := config.New(
		config.WithSource(
			file.NewSource("../../../configs"),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := c.Load(); err != nil {
		t.Fatal(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		t.Fatal(err)
	}

	return &bc
}

func TestCreateDeleteBucket(t *testing.T) {
	conf := confInit(t)
	logger := log.DefaultLogger
	minio := minio.NewMinio(conf.Data, logger)

	// 创建已存在bucket
	bucketName := "bucketnameexist"
	err := minio.CreateBucket(bucketName)

	if err == nil {
		t.Fatal(err)
	} else if err != nil && !errors.IsError(errors.ErrorMinioBucketExisted, err) {
		t.Fatal(err)
	}

	// 创建不存在bucket
	id := utils.GetUUIDWithoutSeparator()
	bucketName = fmt.Sprintf("bucketnamenotexist%s", id)
	err = minio.CreateBucket(bucketName)
	if err != nil {
		t.Fatal(err)
	}

	// 删除已存在bucket
	err = minio.DeleteBucket(bucketName)
	if err != nil {
		t.Fatal(err)
	}

	// 删除不存在bucket
	bucketName = "bucketnamedelete"
	err = minio.DeleteBucket(bucketName)
	if err == nil {
		t.Fatal(err)
	} else if err != nil && !errors.IsError(errors.ErrorMinioBucketNotExist, err) {
		t.Fatal(err)
	}
}

func TestListObjects(t *testing.T) {
	conf := confInit(t)
	logger := log.DefaultLogger
	minio := minio.NewMinio(conf.Data, log.DefaultLogger)

	objectInfoList, err := minio.ListObjects("global", "codes/08ac4f251bbe4555bf5e1179a2d4f47c/", false)
	if err != nil {
		t.Fatal(err)
	}

	logger.Print(len(objectInfoList))
	logger.Print(objectInfoList[0].Name[len("codes/08ac4f251bbe4555bf5e1179a2d4f47c/"):])
}

func TestPresignedDownObject(t *testing.T) {
	conf := confInit(t)
	logger := log.DefaultLogger
	minio := minio.NewMinio(conf.Data, logger)

	url, err := minio.PresignedDownloadObject("data", "ib_logfile0", "http://192.168.202.72:8081?test=123")
	if err != nil {
		t.Fatal(err)
	}

	logger.Print(url.String())
}

func TestPresignedUploadObject(t *testing.T) {
	conf := confInit(t)
	logger := log.DefaultLogger
	minio := minio.NewMinio(conf.Data, logger)

	url, err := minio.PresignedUploadObject("data", "test.txt", "http://abc?test=123")
	if err != nil {
		t.Fatal(err)
	}

	logger.Print(url.String())
}

func TestUser(t *testing.T) {
	ctx := context.Background()
	madmClnt, err := madmin.New("192.168.202.73:31311", "minioadmin", "minioadmin", false)
	if err != nil {
		t.Fatal(err)
	}

	if err = madmClnt.AddCannedPolicy(ctx, "lfj", []byte(`{"Version": "2012-10-17","Statement": [{"Action": ["s3:*"],"Effect": "Allow","Resource": ["arn:aws:s3:::lfj"]}]}`)); err != nil {
		t.Fatal(err)
	}

	if err = madmClnt.AddUser(ctx, "lfj", "123456789"); err != nil {
		t.Fatal(err)
	}
	if err = madmClnt.SetPolicy(ctx, "lfj", "lfj", false); err != nil {
		panic(err)
	}
}
