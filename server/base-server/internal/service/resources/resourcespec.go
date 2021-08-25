package resources

import (
	"context"
	"encoding/json"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model/resources"
	"server/common/errors"

	"server/common/log"

	"k8s.io/apimachinery/pkg/api/resource"
)

type ResourceSpecService struct {
	api.UnimplementedResourceSpecServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewResourceSpecService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ResourceSpecServiceServer {

	return &ResourceSpecService{
		conf: conf,
		log:  log.NewHelper("ResourceSpecService", logger),
		data: data,
	}
}

func (rsvc *ResourceSpecService) ListResourceSpec(ctx context.Context, req *api.ListResourceSpecRequest) (*api.ResourceSpecList, error) {

	if req.PageIndex < 0 || req.PageSize < 0 {
		return &api.ResourceSpecList{}, errors.Errorf(nil, errors.ErrorListResourceSpec)
	}

	resourceSpecList := &api.ResourceSpecList{
		ResourceSpecs: make([]*api.ResourceSpec, 0),
	}

	rsList, err := rsvc.data.ResourceSpecDao.ListResourceSpec(req.PageIndex, req.PageSize)

	if err != nil {
		return &api.ResourceSpecList{}, errors.Errorf(err, errors.ErrorListResourceSpec)
	}

	for _, dbr := range rsList {

		var tempResourceQuantity map[string]string

		err = json.Unmarshal([]byte(dbr.ResourceQuantity), &tempResourceQuantity)

		if err != nil {
			return &api.ResourceSpecList{}, errors.Errorf(err, errors.ErrorListResourceSpec)
		}

		rspec := &api.ResourceSpec{
			Id:               dbr.Id,
			Name:             dbr.Name,
			Price:            dbr.Price,
			ResourceQuantity: tempResourceQuantity,
		}
		resourceSpecList.ResourceSpecs = append(resourceSpecList.ResourceSpecs, rspec)
	}

	return resourceSpecList, err
}

func (rsvc *ResourceSpecService) CreateResourceSpec(ctx context.Context, req *api.CreateResourceSpecRequest) (*api.CreateResourceSpecReply, error) {

	allResources, err := rsvc.data.ResourceDao.ListResource()
	if err != nil {
		return &api.CreateResourceSpecReply{}, errors.Errorf(err, errors.ErrorCreateResourceSpec)
	}

	resourceMap := make(map[string]*resources.Resource)

	for _, r := range allResources {
		resourceMap[r.Name] = r
	}

	for rName := range req.ResourceQuantity {
		if _, ok := resourceMap[rName]; !ok {
			return &api.CreateResourceSpecReply{}, errors.Errorf(nil, errors.ErrorCreateResourceSpec)
		}
	}

	for resName, quantity := range req.ResourceQuantity {
		resQuant, err := resource.ParseQuantity(quantity)
		if err != nil {
			return &api.CreateResourceSpecReply{}, errors.Errorf(err, errors.ErrorCreateResourceSpec)
		}

		if resName == "shm" {
			if _, ok := req.ResourceQuantity["memory"]; !ok {
				return &api.CreateResourceSpecReply{}, errors.Errorf(err, errors.ErrorCreateResourceSpec)
			}

			memQuant, err := resource.ParseQuantity(req.ResourceQuantity["memory"])
			if err != nil {
				return &api.CreateResourceSpecReply{}, errors.Errorf(err, errors.ErrorCreateResourceSpec)
			}
			//fmt.Println("one shm Quantity:",resQuant.Value())
			resQuant.Add(resQuant.DeepCopy())
			//fmt.Println("two shm Quantity:",resQuant.Value())
			//fmt.Println("one memory Quantity:",memQuant.Value())
			if memQuant.Cmp(resQuant.DeepCopy()) < 0 {
				//shm > 1/2 memory is a error
				return &api.CreateResourceSpecReply{}, errors.Errorf(err, errors.ErrorCreateResourceSpec)
			}
		}
	}

	resQuantityBytes, err := json.Marshal(req.ResourceQuantity)

	if err != nil {
		return &api.CreateResourceSpecReply{}, errors.Errorf(err, errors.ErrorCreateResourceSpec)
	}

	resQuantityStr := string(resQuantityBytes)

	cRq := &resources.CreateResourceSpecRequest{
		Name:             req.Name,
		Price:            req.Price,
		ResourceQuantity: resQuantityStr,
	}

	id, err := rsvc.data.ResourceSpecDao.CreateResourceSpec(cRq)

	if err != nil {
		return &api.CreateResourceSpecReply{}, errors.Errorf(err, errors.ErrorCreateResourceSpec)
	}

	return &api.CreateResourceSpecReply{Id: id}, nil
}

func (rsvc *ResourceSpecService) DeleteResourceSpec(ctx context.Context, req *api.DeleteResourceSpecRequest) (*api.DeleteResourceSpecReply, error) {

	id, err := rsvc.data.ResourceSpecDao.DeleteResourceSpec(req.Id)

	if err != nil {
		return &api.DeleteResourceSpecReply{}, errors.Errorf(err, errors.ErrorDeleteResourceSpec)
	}

	return &api.DeleteResourceSpecReply{Id: id}, nil
}

func (rsvc *ResourceSpecService) GetResourceSpec(ctx context.Context, req *api.GetResourceSpecRequest) (*api.GetResourceSpecReply, error) {
	dbr, err := rsvc.data.ResourceSpecDao.GetResourceSpec(req.Id)

	if err != nil {
		return &api.GetResourceSpecReply{}, errors.Errorf(err, errors.ErrorGetResourceSpec)
	}

	var tempResourceQuantity map[string]string

	err = json.Unmarshal([]byte(dbr.ResourceQuantity), &tempResourceQuantity)

	if err != nil {
		return &api.GetResourceSpecReply{}, errors.Errorf(err, errors.ErrorGetResourceSpec)
	}

	rspec := &api.ResourceSpec{
		Id:               dbr.Id,
		Name:             dbr.Name,
		Price:            dbr.Price,
		ResourceQuantity: tempResourceQuantity,
	}

	return &api.GetResourceSpecReply{ResourceSpec: rspec}, nil
}
