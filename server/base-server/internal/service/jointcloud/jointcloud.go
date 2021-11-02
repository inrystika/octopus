package jointcloud

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/jointcloud"
	"server/common/errors"

	"github.com/jinzhu/copier"
)

type jointCloudService struct {
	api.UnimplementedJointCloudServiceServer
	conf *conf.Bootstrap
	data *data.Data
}

func NewJointCloudService(conf *conf.Bootstrap, data *data.Data) api.JointCloudServiceServer {
	s := &jointCloudService{
		conf: conf,
		data: data,
	}

	return s
}

func (s *jointCloudService) ListJointCloudDataset(ctx context.Context, req *api.ListJointCloudDatasetRequest) (*api.ListJointCloudDatasetReply, error) {
	reply, err := s.data.JointCloud.ListDataSet(ctx, &jointcloud.DataSetQuery{
		PageIndex: int(req.PageIndex),
		PageSize:  int(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	dataSets := make([]*api.ListJointCloudDatasetReply_DataSet, 0)
	for _, n := range reply.List {
		dataSet := &api.ListJointCloudDatasetReply_DataSet{}
		err := copier.Copy(dataSet, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		dataSets = append(dataSets, dataSet)
	}

	return &api.ListJointCloudDatasetReply{DataSets: dataSets}, nil
}

func (s *jointCloudService) ListJointCloudDatasetVersion(ctx context.Context, req *api.ListJointCloudDatasetVersionRequest) (*api.ListJointCloudDatasetVersionReply, error) {
	reply, err := s.data.JointCloud.ListDataSetVersion(ctx, &jointcloud.DataSetVersionQuery{
		PageIndex:   int(req.PageIndex),
		PageSize:    int(req.PageSize),
		DataSetCode: req.DataSetCode,
	})
	if err != nil {
		return nil, err
	}
	versions := make([]*api.ListJointCloudDatasetVersionReply_DataSetVersion, 0)
	for _, n := range reply.List {
		version := &api.ListJointCloudDatasetVersionReply_DataSetVersion{}
		err := copier.Copy(version, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		versions = append(versions, version)
	}

	return &api.ListJointCloudDatasetVersionReply{Versions: versions}, nil
}
