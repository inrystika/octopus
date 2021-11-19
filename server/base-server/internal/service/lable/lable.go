package lable

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"time"

	"server/common/errors"
	"server/common/log"
	"server/common/utils"

	"github.com/jinzhu/copier"
)

var LABLE_DEFAULT_DESC = map[api.Relegation]map[api.Type][]string{
	api.Relegation_LABLE_RELEGATION_DATASET: {
		api.Type_LABLE_TYPE_DATASET_TYPE:  {"图像", "视频", "音频", "文本"},
		api.Type_LABLE_TYPE_DATASET_APPLY: {"图像分类", "目标检测", "目标跟踪", "语义分割", "文本分类", "中文分词", "音频分类", "数据增强"},
	},
	api.Relegation_LABLE_RELEGATION_ALGORITHM: {
		api.Type_LABLE_TYPE_ALGORITHM_APPLY:     {"图像分类", "目标检测", "目标跟踪", "语义分割", "文本分类", "中文分词", "音频分类", "模型优化"},
		api.Type_LABLE_TYPE_ALGORITHM_FRAMEWORK: {"TensorFlow", "Pytorch", "MindSpore", "Keras"},
	},
}

type lableService struct {
	api.UnimplementedLableServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewLableService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) (api.LableServiceServer, error) {
	log := log.NewHelper("LableService", logger)

	s := &lableService{
		conf: conf,
		log:  log,
		data: data,
	}

	err := s.addDefaultLable(context.TODO())
	if err != nil {
		s.log.Errorf(context.TODO(), "add default lable failed, cause by:　%v", err)
		return nil, err
	}

	return s, nil
}

func (s *lableService) addDefaultLable(ctx context.Context) error {
	for relegationType, m := range LABLE_DEFAULT_DESC {
		for lableType, array := range m {
			for _, lableDesc := range array {
				lable, _ := s.data.LableDao.QueryLable(ctx, &model.LableQuery{
					RelegationType: int(relegationType),
					SourceType:     int(api.Source_LABLE_SOURCE_PRESET),
					LableType:      int(lableType),
					LableDesc:      lableDesc,
				})
				if lable != nil {
					continue
				}

				err := s.data.LableDao.AddLable(ctx, &model.Lable{
					Id:             utils.GetUUIDWithoutSeparator(),
					RelegationType: int(relegationType),
					SourceType:     int(api.Source_LABLE_SOURCE_PRESET),
					LableType:      int(lableType),
					LableDesc:      lableDesc,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *lableService) AddLable(ctx context.Context, req *api.AddLableRequest) (*api.AddLableReply, error) {
	item, err := s.data.LableDao.QueryLable(ctx, &model.LableQuery{
		RelegationType: int(req.RelegationType),
		LableType:      int(req.LableType),
		LableDesc:      req.LableDesc,
	})
	if item != nil {
		return nil, errors.Errorf(err, errors.ErrorLableRepeated)
	}

	if _, ok := LABLE_DEFAULT_DESC[api.Relegation(req.RelegationType)]; !ok {
		return nil, errors.Errorf(nil, errors.ErrorLableIllegal)
	}
	if _, ok := LABLE_DEFAULT_DESC[api.Relegation(req.RelegationType)][api.Type(req.LableType)]; !ok {
		return nil, errors.Errorf(nil, errors.ErrorLableIllegal)
	}

	lable := &model.Lable{
		Id:             utils.GetUUIDWithoutSeparator(),
		RelegationType: int(req.RelegationType),
		SourceType:     int(api.Source_LABLE_SOURCE_CUSTOMIZE),
		LableType:      int(req.LableType),
		LableDesc:      req.LableDesc,
	}

	err = s.data.LableDao.AddLable(ctx, lable)
	if err != nil {
		return nil, err
	}

	apiLable := &api.Lable{}
	err = copier.Copy(apiLable, lable)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	return &api.AddLableReply{
		Lable: apiLable,
	}, nil
}

func (s *lableService) ListLable(ctx context.Context, req *api.ListLableRequest) (*api.ListLableReply, error) {
	query := &model.LableListQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	lablebl, totalSize, err := s.data.LableDao.ListLable(ctx, query)
	if err != nil {
		return nil, err
	}

	lables := make([]*api.Lable, 0)
	for _, n := range lablebl {
		lable := &api.Lable{}
		err := copier.Copy(lable, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}

		lables = append(lables, lable)
	}

	return &api.ListLableReply{
		TotalSize: totalSize,
		Lables:    lables,
	}, nil
}

func (s *lableService) GetLable(ctx context.Context, req *api.GetLableRequest) (*api.GetLableReply, error) {
	lable, err := s.data.LableDao.GetLable(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	apiLable := &api.Lable{}
	err = copier.Copy(apiLable, lable)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	return &api.GetLableReply{
		Lable: apiLable,
	}, nil
}

func (s *lableService) DeleteLable(ctx context.Context, req *api.DeleteLableRequest) (*api.DeleteLableReply, error) {
	lable, err := s.data.LableDao.GetLable(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if lable.ReferTimes != 0 {
		return nil, errors.Errorf(err, errors.ErrorLableRefered)
	}
	if lable.SourceType != int(api.Source_LABLE_SOURCE_CUSTOMIZE) {
		return nil, errors.Errorf(err, errors.ErrorLableNotDelete)
	}

	err = s.data.LableDao.DeleteLable(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &api.DeleteLableReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

func (s *lableService) UpdateLable(ctx context.Context, req *api.UpdateLableRequest) (*api.UpdateLableReply, error) {
	lable, err := s.data.LableDao.GetLable(ctx, req.Lable.Id)
	if err != nil {
		return nil, err
	}

	item, err := s.data.LableDao.QueryLable(ctx, &model.LableQuery{
		RelegationType: lable.RelegationType,
		LableType:      lable.LableType,
		LableDesc:      req.Lable.LableDesc,
	})
	if item != nil && item.Id != req.Lable.Id {
		return nil, errors.Errorf(err, errors.ErrorLableRepeated)
	}

	modelLable := &model.Lable{}
	err = copier.Copy(modelLable, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	err = s.data.LableDao.UpdateLable(ctx, modelLable)
	if err != nil {
		return nil, err
	}

	return &api.UpdateLableReply{
		UpdatedAt: time.Now().Unix(),
	}, nil
}
