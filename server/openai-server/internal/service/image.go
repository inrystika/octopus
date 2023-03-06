package service

import (
	"context"
	innterapi "server/base-server/api/v1"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/log"
	"server/common/utils/collections/set"
	pb "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

	"github.com/jinzhu/copier"
)

type ImageService struct {
	pb.UnimplementedImageServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewImageService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) pb.ImageServiceServer {
	return &ImageService{
		conf: conf,
		log:  log.NewHelper("ImageService", logger),
		data: data,
	}
}

func (s *ImageService) ListPreImage(ctx context.Context, req *pb.ListPreImageRequest) (*pb.ListPreImageReply, error) {

	reply, err := s.data.ImageClient.ListPreImage(ctx, &innterapi.ListPreImageRequest{
		PageSize:      req.PageSize,
		PageIndex:     req.PageIndex,
		SortBy:        req.SortBy,
		OrderBy:       req.OrderBy,
		ImageNameLike: req.ImageNameLike,
		NameVerLike:   req.NameVerLike,
		SourceType:    innterapi.ImageSourceType(req.SourceType),
		ImageStatus:   innterapi.ImageStatus(req.ImageStatus),
		ImageVersion:  req.ImageVersion,
		SearchKey:     req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	images := make([]*pb.ImageDetail, len(reply.Images))
	for idx, image := range reply.Images {
		images[idx] = &pb.ImageDetail{
			Id:           image.Id,
			ImageName:    image.ImageName,
			ImageDesc:    image.ImageDesc,
			ImageStatus:  int32(image.ImageStatus),
			ImageAddr:    image.ImageAddr,
			ImageVersion: image.ImageVersion,
			SourceType:   int32(image.SourceType),
			SpaceId:      image.SpaceId,
			UserId:       image.UserId,
			CreatedAt:    image.CreatedAt,
			UpdatedAt:    image.UpdatedAt,
		}
	}

	return &pb.ListPreImageReply{
		TotalSize: reply.TotalSize,
		Images:    images,
	}, nil
}

func (s *ImageService) ListUserImage(ctx context.Context, req *pb.ListUserImageRequest) (*pb.ListUserImageReply, error) {
	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)

	reply, err := s.data.ImageClient.ListUserImage(ctx, &innterapi.ListUserImageRequest{
		PageSize:      req.PageSize,
		PageIndex:     req.PageIndex,
		SortBy:        req.SortBy,
		OrderBy:       req.OrderBy,
		UserId:        userId,
		SpaceId:       spaceId,
		ImageNameLike: req.ImageNameLike,
		NameVerLike:   req.NameVerLike,
		SourceType:    innterapi.ImageSourceType(req.SourceType),
		ImageStatus:   innterapi.ImageStatus(req.ImageStatus),
		ImageVersion:  req.ImageVersion,
		SearchKey:     req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	userIds := []string{}
	imageIds := []string{}
	images := make([]*pb.UserImage, len(reply.Images))
	for idx, image := range reply.Images {
		images[idx] = &pb.UserImage{
			IsShared: false,
			Image: &pb.ImageDetail{
				Id:           image.Id,
				ImageName:    image.ImageName,
				ImageDesc:    image.ImageDesc,
				ImageStatus:  int32(image.ImageStatus),
				ImageAddr:    image.ImageAddr,
				ImageVersion: image.ImageVersion,
				SourceType:   int32(image.SourceType),
				SpaceId:      image.SpaceId,
				UserId:       image.UserId,
				CreatedAt:    image.CreatedAt,
				UpdatedAt:    image.UpdatedAt,
			},
		}
		userIds = append(userIds, image.UserId)
		imageIds = append(imageIds, image.Id)
	}

	if len(userIds) > 0 {
		userIds = set.NewStrings(userIds...).Values()
		userReply, err := s.data.UserClient.ListUserInCond(ctx, &innterapi.ListUserInCondRequest{Ids: userIds})
		if err != nil {
			return nil, err
		}
		userMap := make(map[string]string)
		for _, u := range userReply.Users {
			userMap[u.Id] = u.FullName
		}

		for _, image := range images {
			image.Image.Username = userMap[image.Image.Id]
		}
	}

	if len(imageIds) > 0 {
		imageIds = set.NewStrings(imageIds...).Values()
		imageAccessReply, err := s.data.ImageClient.ListImageAccessInCond(ctx, &innterapi.ListImageAccessInCondRequest{ImageIds: imageIds})
		if err != nil {
			return nil, err
		}
		imageAccesses := set.NewStrings()
		for _, ia := range imageAccessReply.Accesses {
			imageAccesses.Add(ia.ImageId)
		}

		for _, image := range images {
			if imageAccesses.Contains(image.Image.Id) {
				image.IsShared = true
			}
		}
	}

	return &pb.ListUserImageReply{
		TotalSize: reply.TotalSize,
		Images:    images,
	}, nil
}

func (s *ImageService) ListCommImage(ctx context.Context, req *pb.ListCommImageRequest) (*pb.ListCommImageReply, error) {
	_, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)

	reply, err := s.data.ImageClient.ListCommImage(ctx, &innterapi.ListCommImageRequest{
		PageSize:      req.PageSize,
		PageIndex:     req.PageIndex,
		SortBy:        req.SortBy,
		OrderBy:       req.OrderBy,
		SpaceId:       spaceId,
		ImageNameLike: req.ImageNameLike,
		NameVerLike:   req.NameVerLike,
		SourceType:    innterapi.ImageSourceType(req.SourceType),
		ImageStatus:   innterapi.ImageStatus(req.ImageStatus),
		ImageVersion:  req.ImageVersion,
		SearchKey:     req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	userIds := []string{}
	images := make([]*pb.ImageDetail, len(reply.Images))
	for idx, image := range reply.Images {
		images[idx] = &pb.ImageDetail{
			Id:           image.Id,
			ImageName:    image.ImageName,
			ImageDesc:    image.ImageDesc,
			ImageAddr:    image.ImageAddr,
			ImageVersion: image.ImageVersion,
			ImageStatus:  int32(image.ImageStatus),
			SourceType:   int32(image.SourceType),
			SpaceId:      image.SpaceId,
			UserId:       image.UserId,
			CreatedAt:    image.CreatedAt,
			UpdatedAt:    image.UpdatedAt,
		}
		userIds = append(userIds, image.UserId)
	}

	if len(userIds) > 0 {
		userIds = set.NewStrings(userIds...).Values()
		userReply, err := s.data.UserClient.ListUserInCond(ctx, &innterapi.ListUserInCondRequest{Ids: userIds})
		if err != nil {
			return nil, err
		}
		userMap := make(map[string]string)
		for _, u := range userReply.Users {
			userMap[u.Id] = u.FullName
		}

		for _, image := range images {
			image.Username = userMap[image.UserId]
		}
	}

	return &pb.ListCommImageReply{
		TotalSize: reply.TotalSize,
		Images:    images,
	}, nil
}

func (s *ImageService) AddImage(ctx context.Context, req *pb.AddImageRequest) (*pb.AddImageReply, error) {
	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)

	reply, err := s.data.ImageClient.AddImage(ctx, &innterapi.AddImageRequest{
		ImageName:    req.ImageName,
		ImageDesc:    req.ImageDesc,
		ImageAddr:    req.ImageAddr,
		ImageVersion: req.ImageVersion,
		SourceType:   innterapi.ImageSourceType(req.SourceType),
		IsPrefab:     innterapi.ImageIsPrefab_IMAGE_IS_PREFAB_NO,
		UserId:       userId,
		SpaceId:      spaceId,
	})

	if err != nil {
		return nil, err
	}
	return &pb.AddImageReply{
		ImageId:   reply.ImageId,
		CreatedAt: reply.CreatedAt,
	}, nil
}

func (s *ImageService) UploadImage(ctx context.Context, req *pb.UploadImageRequest) (*pb.UploadImageReply, error) {
	reply, err := s.data.ImageClient.GetImageUploadUrl(ctx, &innterapi.GetImageUploadUrlRequest{
		ImageId:  req.ImageId,
		FileName: req.FileName,
		Domain:   req.Domain,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UploadImageReply{
		UploadUrl: reply.UploadUrl,
	}, nil
}

func (s *ImageService) DeleteImage(ctx context.Context, req *pb.DeleteImageRequest) (*pb.DeleteImageReply, error) {
	userId := commctx.UserIdFromContext(ctx)
	imageReply, err := s.data.ImageClient.FindImage(ctx, &innterapi.FindImageRequest{ImageId: req.ImageId})
	if err != nil {
		return nil, err
	}
	if imageReply.Image == nil {
		return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
	}
	if imageReply.Image.UserId != userId {
		return nil, errors.Errorf(nil, errors.ErrorImageOpForbidden)
	}

	reply, err := s.data.ImageClient.DeleteImage(ctx, &innterapi.DeleteImageRequest{
		ImageId:  req.ImageId,
		SpaceId:  imageReply.Image.SpaceId,
		UserId:   imageReply.Image.UserId,
		IsPrefab: innterapi.ImageIsPrefab_IMAGE_IS_PREFAB_NO,
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteImageReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

func (s *ImageService) UpdateImage(ctx context.Context, req *pb.UpdateImageRequest) (*pb.UpdateImageReply, error) {
	userId := commctx.UserIdFromContext(ctx)
	imageReply, err := s.data.ImageClient.FindImage(ctx, &innterapi.FindImageRequest{ImageId: req.ImageId})
	if err != nil {
		return nil, err
	}
	if imageReply.Image == nil {
		return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
	}
	if imageReply.Image.UserId != userId {
		return nil, errors.Errorf(nil, errors.ErrorImageOpForbidden)
	}
	reply, err := s.data.ImageClient.UpdateImage(ctx, &innterapi.UpdateImageRequest{
		ImageId:      req.ImageId,
		ImageName:    req.ImageName,
		ImageDesc:    req.ImageDesc,
		ImageAddr:    req.ImageAddr,
		ImageVersion: req.ImageVersion,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateImageReply{
		UpdatedAt: reply.UpdatedAt,
	}, nil
}

func (s *ImageService) ShareImage(ctx context.Context, req *pb.ShareImageRequest) (*pb.ShareImageReply, error) {
	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)
	reply, err := s.data.ImageClient.ShareImage(ctx, &innterapi.ShareImageRequest{ImageId: req.ImageId, UserId: userId, SpaceId: spaceId})
	if err != nil {
		return nil, err
	}
	return &pb.ShareImageReply{SharedAt: reply.SharedAt}, nil
}

func (s *ImageService) CloseShareImage(ctx context.Context, req *pb.CloseShareImageRequest) (*pb.CloseShareImageReply, error) {
	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)
	reply, err := s.data.ImageClient.CloseShareImage(ctx, &innterapi.CloseShareImageRequest{
		ImageId: req.ImageId,
		UserId:  userId,
		SpaceId: spaceId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CloseShareImageReply{CloseSharedAt: reply.CloseSharedAt}, nil
}

func (s *ImageService) ConfirmUploadImage(ctx context.Context, req *pb.ConfirmUploadImageRequest) (*pb.ConfirmUploadImageReply, error) {
	userId := commctx.UserIdFromContext(ctx)
	imageReply, err := s.data.ImageClient.FindImage(ctx, &innterapi.FindImageRequest{ImageId: req.ImageId})
	if err != nil {
		return nil, err
	}
	if imageReply.Image == nil {
		return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
	}
	if imageReply.Image.UserId != userId {
		return nil, errors.Errorf(nil, errors.ErrorImageOpForbidden)
	}
	reply, err := s.data.ImageClient.ConfirmUploadImage(ctx, &innterapi.ConfirmUploadImageRequest{ImageId: req.ImageId})
	if err != nil {
		return nil, err
	}
	return &pb.ConfirmUploadImageReply{UpdatedAt: reply.UpdatedAt}, nil
}

func (s *ImageService) FindImage(ctx context.Context, req *pb.FindImageRequest) (*pb.FindImageReply, error) {
	innerReq := &innterapi.FindImageRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.ImageClient.FindImage(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &pb.FindImageReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
