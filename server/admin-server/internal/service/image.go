package service

import (
	"context"
	pb "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innterapi "server/base-server/api/v1"
	"server/common/constant"
	commctx "server/common/context"
	"server/common/log"
	"server/common/utils/collections/set"
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
	reply, err := s.data.ImageClient.ListUserImage(ctx, &innterapi.ListUserImageRequest{
		PageSize:      req.PageSize,
		PageIndex:     req.PageIndex,
		SortBy:        req.SortBy,
		OrderBy:       req.OrderBy,
		ImageNameLike: req.ImageNameLike,
		SourceType:    innterapi.ImageSourceType(req.SourceType),
		ImageStatus:   innterapi.ImageStatus(req.ImageStatus),
		ImageVersion:  req.ImageVersion,
		SearchKey:     req.SearchKey,
		UserId:        req.UserId,
		SpaceId:       req.SpaceId,
	})
	if err != nil {
		return nil, err
	}

	userIds := []string{}
	workspaceIds := []string{}
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
		userIds = append(userIds, image.UserId)
		if constant.SYSTEM_WORKSPACE_DEFAULT != image.SpaceId {
			workspaceIds = append(workspaceIds, image.SpaceId)
		}
	}

	if len(userIds) > 0 {
		userIds = set.NewStrings(userIds...).Values()
		userReply, err := s.data.UserClient.ListUserInCond(ctx, &innterapi.ListUserInCondRequest{Ids: userIds})
		if err != nil {
			return nil, err
		}
		userMap := make(map[string]string)
		emailMap := make(map[string]string)
		for _, u := range userReply.Users {
			userMap[u.Id] = u.FullName
			emailMap[u.Id] = u.Email
		}

		for _, image := range images {
			image.Username = userMap[image.UserId]
			image.UserEmail = emailMap[image.UserId]
		}
	}
	if len(workspaceIds) > 0 {
		workspaceIds = set.NewStrings(workspaceIds...).Values()
		wkReply, err := s.data.WorkspaceClient.ListWorkspaceInCond(ctx, &innterapi.ListWorkspaceInCondRequest{Ids: workspaceIds})
		if err != nil {
			return nil, err
		}
		wkMap := make(map[string]string)
		for _, w := range wkReply.Workspaces {
			wkMap[w.Id] = w.Name
		}

		for _, image := range images {
			image.SpaceName = wkMap[image.SpaceId]
		}
	}

	return &pb.ListUserImageReply{
		TotalSize: reply.TotalSize,
		Images:    images,
	}, nil
}

func (s *ImageService) AddPreImage(ctx context.Context, req *pb.AddPreImageRequest) (*pb.AddPreImageReply, error) {
	userId := commctx.UserIdFromContext(ctx)
	reply, err := s.data.ImageClient.AddImage(ctx, &innterapi.AddImageRequest{
		ImageName:    req.ImageName,
		ImageDesc:    req.ImageDesc,
		ImageAddr:    req.ImageAddr,
		ImageVersion: req.ImageVersion,
		SourceType:   innterapi.ImageSourceType(req.SourceType),
		IsPrefab:     innterapi.ImageIsPrefab_IMAGE_IS_PREFAB_YES,
		UserId:       userId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddPreImageReply{
		ImageId:   reply.ImageId,
		CreatedAt: reply.CreatedAt,
	}, nil
}

func (s *ImageService) UploadPreImage(ctx context.Context, req *pb.UploadPreImageRequest) (*pb.UploadPreImageReply, error) {
	reply, err := s.data.ImageClient.GetImageUploadUrl(ctx, &innterapi.GetImageUploadUrlRequest{
		ImageId:  req.ImageId,
		FileName: req.FileName,
		Domain:   req.Domain,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UploadPreImageReply{
		UploadUrl: reply.UploadUrl,
	}, nil
}

func (s *ImageService) DeletePreImage(ctx context.Context, req *pb.DeletePreImageRequest) (*pb.DeletePreImageReply, error) {
	// todo 只能删除预置镜像

	reply, err := s.data.ImageClient.DeleteImage(ctx, &innterapi.DeleteImageRequest{
		ImageId:  req.ImageId,
		IsPrefab: innterapi.ImageIsPrefab_IMAGE_IS_PREFAB_YES,
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeletePreImageReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

func (s *ImageService) UpdateImage(ctx context.Context, req *pb.UpdatePreImageRequest) (*pb.UpdatePreImageReply, error) {
	// todo 只能更新预置镜像

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

	return &pb.UpdatePreImageReply{
		UpdatedAt: reply.UpdatedAt,
	}, nil
}

func (s *ImageService) ConfirmUploadPreImage(ctx context.Context, req *pb.ConfirmUploadPreImageRequest) (*pb.ConfirmUploadPreImageReply, error) {
	reply, err := s.data.ImageClient.ConfirmUploadImage(ctx, &innterapi.ConfirmUploadImageRequest{ImageId: req.ImageId})
	if err != nil {
		return nil, err
	}
	return &pb.ConfirmUploadPreImageReply{UpdatedAt: reply.UpdatedAt}, nil
}
