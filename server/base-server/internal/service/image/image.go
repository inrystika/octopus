package image

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	pb "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/utils"
	"strings"
	"time"

	"server/common/log"

	docker "github.com/fsouza/go-dockerclient"
)

type ImageService struct {
	pb.UnimplementedImageServiceServer
	conf         *conf.Bootstrap
	log          *log.Helper
	data         *data.Data
	dockerClient *docker.Client
}

func NewImageService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) pb.ImageServiceServer {
	dockerClient, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
	return &ImageService{
		conf:         conf,
		log:          log.NewHelper("ImageService", logger),
		data:         data,
		dockerClient: dockerClient,
	}
}

func (s *ImageService) ListPreImage(ctx context.Context, req *pb.ListPreImageRequest) (*pb.ListPreImageReply, error) {
	imageCount, err := s.data.ImageDao.Count(ctx, &model.ImageList{
		IsPrefab:      int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_YES),
		SourceType:    int32(req.SourceType),
		ImageNameLike: req.ImageNameLike,
		NameVerLike:   req.NameVerLike,
		ImageVersion:  req.ImageVersion,
		Status:        int32(req.ImageStatus),
		SearchKey:     req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	imageDetails := []*pb.ImageDetail{}
	if imageCount > 0 {
		images, err := s.data.ImageDao.List(ctx, &model.ImageList{
			PageIndex:     req.PageIndex,
			PageSize:      req.PageSize,
			OrderBy:       req.OrderBy,
			SortBy:        req.SortBy,
			IsPrefab:      int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_YES),
			Status:        int32(req.ImageStatus),
			SourceType:    int32(req.SourceType),
			ImageNameLike: req.ImageNameLike,
			NameVerLike:   req.NameVerLike,
			ImageVersion:  req.ImageVersion,
			SearchKey:     req.SearchKey,
		})
		if err != nil {
			return nil, err
		}
		for _, image := range images {
			imageDetails = append(imageDetails, &pb.ImageDetail{
				Id:           image.Id,
				ImageName:    image.ImageName,
				ImageVersion: image.ImageVersion,
				ImageDesc:    image.ImageDesc,
				ImageAddr:    image.ImageAddr,
				ImageStatus:  pb.ImageStatus(image.Status),
				SourceType:   pb.ImageSourceType(image.SourceType),
				SpaceId:      image.SpaceId,
				UserId:       image.UserId,
				IsPrefab:     pb.ImageIsPrefab(image.IsPrefab),
				CreatedAt:    image.CreatedAt.Unix(),
				UpdatedAt:    image.UpdatedAt.Unix(),
			})
		}
	}

	return &pb.ListPreImageReply{
		TotalSize: imageCount,
		Images:    imageDetails,
	}, nil
}

func (s *ImageService) ListUserImage(ctx context.Context, req *pb.ListUserImageRequest) (*pb.ListUserImageReply, error) {
	imageCount, err := s.data.ImageDao.Count(ctx, &model.ImageList{
		IsPrefab:      int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_NO),
		UserId:        req.UserId,
		SpaceId:       req.SpaceId,
		SourceType:    int32(req.SourceType),
		ImageNameLike: req.ImageNameLike,
		NameVerLike:   req.NameVerLike,
		ImageVersion:  req.ImageVersion,
		Status:        int32(req.ImageStatus),
		SearchKey:     req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	imageDetails := []*pb.ImageDetail{}
	if imageCount > 0 {
		images, err := s.data.ImageDao.List(ctx, &model.ImageList{
			PageIndex:     req.PageIndex,
			PageSize:      req.PageSize,
			OrderBy:       req.OrderBy,
			SortBy:        req.SortBy,
			IsPrefab:      int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_NO),
			UserId:        req.UserId,
			SpaceId:       req.SpaceId,
			SourceType:    int32(req.SourceType),
			ImageNameLike: req.ImageNameLike,
			NameVerLike:   req.NameVerLike,
			ImageVersion:  req.ImageVersion,
			Status:        int32(req.ImageStatus),
			SearchKey:     req.SearchKey,
		})
		if err != nil {
			return nil, err
		}
		for _, image := range images {
			imageDetails = append(imageDetails, &pb.ImageDetail{
				Id:           image.Id,
				ImageName:    image.ImageName,
				ImageVersion: image.ImageVersion,
				ImageDesc:    image.ImageDesc,
				ImageAddr:    image.ImageAddr,
				ImageStatus:  pb.ImageStatus(image.Status),
				SourceType:   pb.ImageSourceType(image.SourceType),
				SpaceId:      image.SpaceId,
				UserId:       image.UserId,
				IsPrefab:     pb.ImageIsPrefab(image.IsPrefab),
				CreatedAt:    image.CreatedAt.Unix(),
				UpdatedAt:    image.UpdatedAt.Unix(),
			})
		}
	}

	return &pb.ListUserImageReply{
		TotalSize: imageCount,
		Images:    imageDetails,
	}, nil
}

func (s *ImageService) ListCommImage(ctx context.Context, req *pb.ListCommImageRequest) (*pb.ListCommImageReply, error) {
	iac, err := s.data.ImageDao.CountImageByAccess(ctx, &model.ImageAccessList{
		IsPrefab:      int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_NO),
		SpaceId:       req.SpaceId,
		SourceType:    int32(req.SourceType),
		ImageNameLike: req.ImageNameLike,
		NameVerLike:   req.NameVerLike,
		ImageVersion:  req.ImageVersion,
		Status:        int32(req.ImageStatus),
		SearchKey:     req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	imageDetails := []*pb.ImageDetail{}
	if iac > 0 {
		ias, err := s.data.ImageDao.ListImageByAccess(ctx, &model.ImageAccessList{
			PageIndex:     req.PageIndex,
			PageSize:      req.PageSize,
			OrderBy:       req.OrderBy,
			SortBy:        req.SortBy,
			IsPrefab:      int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_NO),
			SpaceId:       req.SpaceId,
			SourceType:    int32(req.SourceType),
			ImageNameLike: req.ImageNameLike,
			NameVerLike:   req.NameVerLike,
			ImageVersion:  req.ImageVersion,
			Status:        int32(req.ImageStatus),
			SearchKey:     req.SearchKey,
		})
		if err != nil {
			return nil, err
		}
		for _, image := range ias {
			imageDetails = append(imageDetails, &pb.ImageDetail{
				Id:           image.Id,
				ImageName:    image.ImageName,
				ImageVersion: image.ImageVersion,
				ImageDesc:    image.ImageDesc,
				ImageAddr:    image.ImageAddr,
				ImageStatus:  pb.ImageStatus(image.Status),
				SourceType:   pb.ImageSourceType(image.SourceType),
				SpaceId:      image.SpaceId,
				UserId:       image.UserId,
				IsPrefab:     pb.ImageIsPrefab(image.IsPrefab),
				CreatedAt:    image.CreatedAt.Unix(),
				UpdatedAt:    image.UpdatedAt.Unix(),
			})
		}
	}

	return &pb.ListCommImageReply{
		TotalSize: iac,
		Images:    imageDetails,
	}, nil
}

func (s *ImageService) ShareImage(ctx context.Context, req *pb.ShareImageRequest) (*pb.ShareImageReply, error) {
	image, err := s.data.ImageDao.Find(ctx, &model.ImageQuery{Id: req.ImageId})
	if err != nil {
		return nil, err
	}
	if image == nil {
		return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
	}
	// only share user image
	if image.IsPrefab != int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_NO) {
		return nil, errors.Errorf(nil, errors.ErrorImageOpForbidden)
	}
	// only share the image of self
	if image.UserId != req.UserId {
		return nil, errors.Errorf(nil, errors.ErrorImageOpForbidden)
	}
	// only share the image to the created space of image
	if image.SpaceId != req.SpaceId {
		return nil, errors.Errorf(nil, errors.ErrorImageOpForbidden)
	}

	ia, err := s.data.ImageDao.AddImageAccess(ctx, &model.ImageAccessAdd{
		Id:      utils.GetUUIDWithoutSeparator(),
		SpaceId: image.SpaceId,
		ImageId: image.Id,
		UserId:  image.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ShareImageReply{
		SharedAt: ia.CreatedAt.Unix(),
	}, nil
}

func (s *ImageService) CloseShareImage(ctx context.Context, req *pb.CloseShareImageRequest) (*pb.CloseShareImageReply, error) {
	_, err := s.data.ImageDao.DeleteImageAccess(ctx, &model.ImageAccessDel{
		SpaceId: req.SpaceId,
		ImageId: req.ImageId,
		UserId:  req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CloseShareImageReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

func (s *ImageService) AddImage(ctx context.Context, req *pb.AddImageRequest) (*pb.AddImageReply, error) {
	c, err := s.data.ImageDao.Count(ctx, &model.ImageList{
		UserId:       req.UserId,
		SpaceId:      req.SpaceId,
		ImageName:    req.ImageName,
		ImageVersion: req.ImageVersion,
	})
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return nil, errors.Errorf(nil, errors.ErrorImageExisted)
	}

	imageAdd := &model.ImageAdd{
		Id:           utils.GetUUIDWithoutSeparator(),
		IsPrefab:     int32(req.IsPrefab),
		UserId:       req.UserId,
		SpaceId:      req.SpaceId,
		SourceType:   int32(req.SourceType),
		ImageName:    req.ImageName,
		ImageVersion: req.ImageVersion,
		ImageAddr:    req.ImageAddr,
		ImageDesc:    req.ImageDesc,
	}
	if req.SourceType == pb.ImageSourceType_IMAGE_SOURCE_TYPE_UPLOADED {
		// 上传镜像类型，需上传，生成镜像地址
		im := model.Image{
			SpaceId:      imageAdd.SpaceId,
			UserId:       imageAdd.UserId,
			ImageName:    imageAdd.ImageName,
			ImageVersion: imageAdd.ImageVersion,
			IsPrefab:     imageAdd.IsPrefab,
		}
		imageAdd.ImageAddr = s.generateImageRepositoryPath(&im)
		imageAdd.Status = int32(pb.ImageStatus_IMAGE_STATUS_NO_MADE)
	} else if req.SourceType == pb.ImageSourceType_IMAGE_SOURCE_TYPE_REMOTE {
		// 远程镜像类型，无需上传，保存镜像地址
		if imageAdd.ImageAddr == "" {
			return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
		}
		imageAdd.Status = int32(pb.ImageStatus_IMAGE_STATUS_MADE)
	} else if req.SourceType == pb.ImageSourceType_IMAGE_SOURCE_TYPE_SAVED {
		im := model.Image{
			SpaceId:      imageAdd.SpaceId,
			UserId:       imageAdd.UserId,
			ImageName:    imageAdd.ImageName,
			ImageVersion: imageAdd.ImageVersion,
			IsPrefab:     imageAdd.IsPrefab,
		}
		imageAdd.ImageAddr = s.generateImageRepositoryPath(&im)
		imageAdd.Status = int32(pb.ImageStatus_IMAGE_STATUS_NO_MADE)
	} else {
		return nil, errors.Errorf(nil, errors.ErrorImageSourceType)
	}

	image, err := s.data.ImageDao.Add(ctx, imageAdd)
	if err != nil {
		return nil, err
	}
	return &pb.AddImageReply{
		ImageId:   image.Id,
		CreatedAt: image.CreatedAt.Unix(),
		ImageAddr: image.ImageAddr,
	}, nil
}

func (s *ImageService) DeleteImage(ctx context.Context, req *pb.DeleteImageRequest) (*pb.DeleteImageReply, error) {
	if req.IsPrefab != pb.ImageIsPrefab_IMAGE_IS_PREFAB_YES {
		_, err := s.data.ImageDao.DeleteImageAccess(ctx, &model.ImageAccessDel{
			SpaceId: req.SpaceId,
			ImageId: req.ImageId,
			UserId:  req.UserId,
		})
		if err != nil {
			return nil, err
		}
	}

	_, err := s.data.ImageDao.Delete(ctx, &model.ImageDel{
		Id: req.ImageId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteImageReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

func (s *ImageService) UpdateImage(ctx context.Context, req *pb.UpdateImageRequest) (*pb.UpdateImageReply, error) {
	image, err := s.data.ImageDao.Find(ctx, &model.ImageQuery{Id: req.ImageId})
	if err != nil {
		return nil, err
	}
	if image == nil {
		return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
	}

	imageUpdated := &model.ImageUpdate{
		ImageName:    req.ImageName,
		ImageVersion: req.ImageVersion,
		ImageDesc:    req.ImageDesc,
		Status:       int32(req.ImageStatus),
	}
	if image.SourceType == int32(pb.ImageSourceType_IMAGE_SOURCE_TYPE_REMOTE) {
		if imageUpdated.ImageAddr != "" {
			imageUpdated.ImageAddr = req.ImageAddr
		}
	}

	image, err = s.data.ImageDao.Update(ctx, &model.ImageUpdateCond{Id: req.ImageId}, imageUpdated)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateImageReply{
		UpdatedAt: image.UpdatedAt.Unix(),
	}, nil
}

func (s *ImageService) ConfirmUploadImage(ctx context.Context, req *pb.ConfirmUploadImageRequest) (*pb.ConfirmUploadImageReply, error) {
	image, err := s.data.ImageDao.Find(ctx, &model.ImageQuery{Id: req.ImageId})
	if err != nil {
		return nil, err
	}
	if image == nil {
		return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
	}
	if image.Status != int32(pb.ImageStatus_IMAGE_STATUS_NO_MADE) && image.Status != int32(pb.ImageStatus_IMAGE_STATUS_MADE_FAILED) {
		return nil, errors.Errorf(nil, errors.ErrorImageStatusMakeError)
	}

	go func() {
		_, err := s.data.ImageDao.Update(ctx, &model.ImageUpdateCond{Id: image.Id}, &model.ImageUpdate{
			Status: int32(pb.ImageStatus_IMAGE_STATUS_MAKING),
		})
		if err != nil {
			s.log.Errorw(ctx, err)
			return
		}

		// push image tar to registry
		err = s.importAndPushImage(image)
		if err != nil {
			s.log.Errorw(ctx, err)
			_, err = s.data.ImageDao.Update(ctx, &model.ImageUpdateCond{Id: image.Id}, &model.ImageUpdate{
				Status: int32(pb.ImageStatus_IMAGE_STATUS_MADE_FAILED),
			})
			if err != nil {
				s.log.Errorw(ctx, err)
			}
			return
		}

		_, err = s.data.ImageDao.Update(ctx, &model.ImageUpdateCond{Id: image.Id}, &model.ImageUpdate{
			Status: int32(pb.ImageStatus_IMAGE_STATUS_MADE),
		})
		if err != nil {
			s.log.Errorw(ctx, err)
		}
	}()
	// create async job to handle image.tar
	//err = s.data.Cluster.CreateAndListenJob(ctx, s.generateJobToHandleImageTar(image), func (e error) {
	//	_, err = s.data.ImageDao.Update(ctx, &model.ImageUpdateCond{Id: image.Id}, &model.ImageUpdate {
	//		Status:    int32(pb.ImageStatus_IMAGE_STATUS_MADE),
	//	})
	//	if err != nil {
	//		s.log.Errorw(err)
	//	}
	//})
	//if err != nil {
	//	s.log.Errorw("image[" + req.ImageId + "] handle err after finished upload, error info:", err)
	//	return nil, err
	//}

	return &pb.ConfirmUploadImageReply{UpdatedAt: time.Now().Unix()}, nil
}

func (s *ImageService) generateImageAddress(image *model.Image) string {
	if image == nil {
		return ""
	}

	var imageVersion string
	if image.ImageVersion != "" {
		imageVersion = ":" + image.ImageVersion
	}
	return fmt.Sprintf("%s%s", s.generateImageRepository(image), imageVersion)
}

func (s *ImageService) generateImageRepository(image *model.Image) string {
	if image == nil {
		return ""
	}

	return fmt.Sprintf("%s/%s", s.conf.Data.Harbor.Host, s.generateImageRepositoryPath(image))
}

func (s *ImageService) generateImageRepositoryPath(image *model.Image) string {
	if image == nil {
		return ""
	}

	// support the historical data inversion function
	if image.ImageType > 0 {
		if image.IsPrefab == int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_YES) {
			return fmt.Sprintf("%s/%s/%s/%v", common.PREAB_FOLDER, image.UserId, image.ImageName, image.ImageType)
		} else {
			return fmt.Sprintf("%s/%s/%s/%v", image.SpaceId, image.UserId, image.ImageName, image.ImageType)
		}
	}

	if image.IsPrefab == int32(pb.ImageIsPrefab_IMAGE_IS_PREFAB_YES) {
		return fmt.Sprintf("%s/%s/%s", common.PREAB_FOLDER, image.UserId, image.ImageName)
	} else {
		return fmt.Sprintf("%s/%s/%s", image.SpaceId, image.UserId, image.ImageName)
	}
}

func (s *ImageService) importAndPushImage(image *model.Image) error {
	registryHost := s.conf.Data.Harbor.Host
	imageRepository := s.generateImageRepository(image)
	//opts := docker.ImportImageOptions{
	//	Source:     path.Join(s.conf.Data.Minio.Base.MountPath, image.SourceFilePath),
	//	Repository: imageRepository,
	//	Tag:        image.ImageVersion,
	//}
	// load image
	imageTarSrc := path.Join(s.conf.Data.Minio.Base.MountPath, image.SourceFilePath)
	imageTar, err := os.Open(imageTarSrc)
	defer imageTar.Close()
	if err != nil {
		s.log.Errorf(context.Background(), "open docker image error, opt.source: %s", imageTarSrc)
		return err
	}
	var loadResult bytes.Buffer
	opts := docker.LoadImageOptions{
		OutputStream: &loadResult,
		InputStream:  imageTar,
	}
	err = s.dockerClient.LoadImage(opts)
	if err != nil {
		s.log.Errorf(context.Background(), "docker load error, opt.source: %s", imageTarSrc)
		return err
	}
	originImageName := strings.Split(loadResult.String(), " ")[2]
	originImageName = strings.Trim(originImageName, "\n")

	// re tag image
	tagOpts := docker.TagImageOptions{
		Repo: imageRepository,
		Tag:  image.ImageVersion,
	}
	err = s.dockerClient.TagImage(originImageName, tagOpts)
	if err != nil {
		s.log.Errorf(context.Background(), "docker tag error, origin image name: %s", originImageName)
		return err
	}

	// push image
	opt1 := docker.PushImageOptions{
		Name:     imageRepository,
		Registry: registryHost,
		Tag:      image.ImageVersion,
	}

	serverAddress := registryHost
	if s.conf.Data.Harbor.UseSSL {
		serverAddress = "https://" + serverAddress
	} else {
		serverAddress = "http://" + serverAddress
	}
	optAuth := docker.AuthConfiguration{
		Username:      s.conf.Data.Harbor.Username,
		Password:      s.conf.Data.Harbor.Password,
		ServerAddress: serverAddress,
	}
	err = s.dockerClient.PushImage(opt1, optAuth)
	if err != nil {
		s.log.Errorf(context.Background(), "docker push error, opt.name: %s, opt.registry: %s, opt.tag: %s", opt1.Name, opt1.Registry, opt1.Tag)
		return err
	}
	err = s.dockerClient.RemoveImage(opt1.Name)
	if err != nil {
		s.log.Errorf(context.Background(), "docker rm image error, opt.name: %s, opt.registry: %s, opt.tag: %s", opt1.Name, opt1.Registry, opt1.Tag)
	}
	err = s.dockerClient.RemoveImage(originImageName)
	if err != nil {
		s.log.Errorf(context.Background(), "docker rm origin image error, originImageName: %s, opt.registry: %s, opt.tag: %s", originImageName, opt1.Registry, opt1.Tag)
	}
	return nil
}

func (s *ImageService) GetImageUploadUrl(ctx context.Context, req *pb.GetImageUploadUrlRequest) (*pb.GetImageUploadUrlReply, error) {
	image, err := s.data.ImageDao.Find(ctx, &model.ImageQuery{Id: req.ImageId})
	if err != nil {
		return nil, err
	}
	if image == nil {
		return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
	}
	if image.SourceType != int32(pb.ImageSourceType_IMAGE_SOURCE_TYPE_UPLOADED) {
		return nil, errors.Errorf(nil, errors.ErrorImageSourceTypeToUpload)
	}
	if image.Status == int32(pb.ImageStatus_IMAGE_STATUS_MAKING) || image.Status == int32(pb.ImageStatus_IMAGE_STATUS_MADE) {
		return nil, errors.Errorf(nil, errors.ErrorImageOpForbidden)
	}

	uploadUrl, err := s.getUploadUrl(image, req.Domain, req.FileName)
	if err != nil {
		return nil, err
	}
	b, o := getTempMinioPath(image, req.FileName)
	_, err = s.data.ImageDao.Update(ctx, &model.ImageUpdateCond{Id: image.Id}, &model.ImageUpdate{
		SourceFilePath: fmt.Sprintf("%s/%s", b, o),
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetImageUploadUrlReply{
		UploadUrl: uploadUrl,
	}, nil
}

func (s *ImageService) getUploadUrl(image *model.Image, domain, fileName string) (string, error) {
	bucketName, objectName := getTempMinioPath(image, fileName)
	uri, err := s.data.Minio.PresignedUploadObject(bucketName, objectName, domain)
	if err != nil {
		return "", err
	}

	return uri.String(), nil
}

func getTempMinioPath(image *model.Image, fileName string) (bucketName string, objectName string) {
	bucketName = common.GetMinioBucket()
	objectName = common.GetMinioUploadImageObject(image.Id, image.ImageVersion, fileName)
	return
}

// 获取镜像详情
func (s *ImageService) FindImage(ctx context.Context, req *pb.FindImageRequest) (*pb.FindImageReply, error) {
	image, err := s.data.ImageDao.Find(ctx, &model.ImageQuery{
		Id: req.ImageId,
	})
	if err != nil {
		return nil, err
	}
	if image == nil {
		return &pb.FindImageReply{
			Image: nil,
		}, nil
	}
	accesses := make([]*pb.ImageAccess, len(image.Accesses))
	for idx, a := range image.Accesses {
		accesses[idx] = &pb.ImageAccess{
			ImageId: a.ImageId,
			UserId:  a.UserId,
			SpaceId: a.SpaceId,
		}
	}
	var imageFullAddr string
	if image.SourceType == int32(pb.ImageSourceType_IMAGE_SOURCE_TYPE_UPLOADED) {
		imageFullAddr = s.generateImageAddress(image)
	} else if image.SourceType == int32(pb.ImageSourceType_IMAGE_SOURCE_TYPE_SAVED) {
		imageFullAddr = s.generateImageAddress(image)
	} else {
		imageFullAddr = image.ImageAddr
	}
	reply := &pb.FindImageReply{
		Image: &pb.ImageDetail{
			Id:           image.Id,
			ImageName:    image.ImageName,
			ImageVersion: image.ImageVersion,
			ImageDesc:    image.ImageDesc,
			ImageAddr:    image.ImageAddr,
			ImageStatus:  pb.ImageStatus(image.Status),
			SourceType:   pb.ImageSourceType(image.SourceType),
			SpaceId:      image.SpaceId,
			UserId:       image.UserId,
			IsPrefab:     pb.ImageIsPrefab(image.IsPrefab),
			CreatedAt:    image.CreatedAt.Unix(),
			UpdatedAt:    image.UpdatedAt.Unix(),
		},
		ImageFullAddr: imageFullAddr,
		Accesses:      accesses,
	}

	return reply, nil
}

// 条件疲劳查询
func (s *ImageService) ListImageInCond(ctx context.Context, req *pb.ListImageInCondRequest) (*pb.ListImageInCondReply, error) {
	images, err := s.data.ImageDao.ListIn(ctx, &model.ImageListIn{Ids: req.Ids})
	if err != nil {
		return nil, err
	}

	imageItems := make([]*pb.ImageDetail, len(images))
	for idx, image := range images {
		item := &pb.ImageDetail{
			Id:           image.Id,
			ImageName:    image.ImageName,
			ImageVersion: image.ImageVersion,
			ImageDesc:    image.ImageDesc,
			ImageAddr:    image.ImageAddr,
			ImageStatus:  pb.ImageStatus(image.Status),
			SourceType:   pb.ImageSourceType(image.SourceType),
			SpaceId:      image.SpaceId,
			UserId:       image.UserId,
			IsPrefab:     pb.ImageIsPrefab(image.IsPrefab),
			CreatedAt:    image.CreatedAt.Unix(),
			UpdatedAt:    image.UpdatedAt.Unix(),
		}
		imageItems[idx] = item
	}
	return &pb.ListImageInCondReply{
		Images: imageItems,
	}, nil
}

// 条件镜像范围查询
func (s *ImageService) ListImageAccessInCond(ctx context.Context, req *pb.ListImageAccessInCondRequest) (*pb.ListImageAccessInCondReply, error) {
	imageAccesses, err := s.data.ImageDao.ListImageAccessIn(ctx, &model.ImageAccessListIn{Ids: req.ImageIds, SpaceId: req.SpaceId, UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	ias := []*pb.ImageAccess{}
	for _, ia := range imageAccesses {
		item := &pb.ImageAccess{
			ImageId: ia.ImageId,
			UserId:  ia.UserId,
			SpaceId: ia.SpaceId,
		}
		ias = append(ias, item)
	}
	return &pb.ListImageAccessInCondReply{
		Accesses: ias,
	}, nil
}
