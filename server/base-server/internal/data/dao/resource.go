package dao

import (
	"server/base-server/internal/data/dao/model/resources"
	"server/common/errors"
	"server/common/utils"

	"server/common/log"

	"gorm.io/gorm"
)

type ResourceDao interface {
	ListResource() ([]*resources.Resource, error)
	ListResourceAll() ([]*resources.Resource, error)
	CreateResource(request *resources.CreateResourceRequest) (string, error)
	GetResource(id string) (*resources.Resource, error)
	UpdateResource(resource *resources.Resource) (string, error)
	DeleteResource(id string) (string, error)
	DeleteResourceByName(name string) (string, error)
}

type resourceDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewResourceDao(db *gorm.DB, logger log.Logger) ResourceDao {
	return &resourceDao{
		log: log.NewHelper("ResourceDao", logger),
		db:  db,
	}
}

func (d *resourceDao) ListResource() ([]*resources.Resource, error) {
	db := d.db
	resources := make([]*resources.Resource, 0)

	if err := db.Find(&resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}

func (d *resourceDao) ListResourceAll() ([]*resources.Resource, error) {
	db := d.db
	resources := make([]*resources.Resource, 0)

	if err := db.Unscoped().Find(&resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}

func (d *resourceDao) CreateResource(request *resources.CreateResourceRequest) (string, error) {
	db := d.db
	id := utils.GetUUIDWithoutSeparator()

	resources := &resources.Resource{
		Id:          id,
		Name:        request.Name,
		Desc:        request.Desc,
		ResourceRef: request.ResourceRef,
	}

	if err := db.Create(&resources).Error; err != nil {
		return "", errors.Errorf(err, errors.ErrorDBCreateFailed)
	}

	return id, nil
}

func (d *resourceDao) UpdateResource(resource *resources.Resource) (string, error) {
	db := d.db

	if err := db.Save(resource).Error; err != nil {
		return "", err
	}

	return resource.Id, nil
}

func (d *resourceDao) GetResource(id string) (*resources.Resource, error) {
	db := d.db
	resource := &resources.Resource{}
	if err := db.Where("id=?", id).Find(resource).Error; err != nil {
		return &resources.Resource{}, err
	}

	return resource, nil
}

func (d *resourceDao) DeleteResource(id string) (string, error) {
	db := d.db
	resource := &resources.Resource{Id: id}

	if err := db.Delete(resource).Error; err != nil {
		return "", err
	}

	return id, nil
}

func (d *resourceDao) DeleteResourceByName(name string) (string, error) {
	db := d.db
	resource := &resources.Resource{}

	if err := db.Where("name = ?", name).First(resource).Error; err != nil {
		return "", err
	}

	if err := db.Delete(resource).Error; err != nil {
		return "", err
	}

	return name, nil
}
