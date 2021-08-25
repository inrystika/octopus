package service

import (
	"context"
	"encoding/json"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	"server/common/errors"
	"server/common/log"

	"github.com/golang/protobuf/ptypes/empty"
)

type NodeService struct {
	api.UnimplementedNodeServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewNodeService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.NodeServiceServer {
	return &NodeService{
		conf: conf,
		log:  log.NewHelper("NodeService", logger),
		data: data,
	}
}

func (nsvc *NodeService) ListNode(ctx context.Context, req *empty.Empty) (*api.NodeList, error) {
	reply, err := nsvc.data.NodeClient.ListNode(ctx, &empty.Empty{})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListNode)
	}

	apiReply := &api.NodeList{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListNode)
	}

	return apiReply, nil
}
