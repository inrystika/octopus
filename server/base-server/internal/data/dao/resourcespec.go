package dao

import (
	"gorm.io/gorm"
	"server/base-server/internal/data/dao/model/resources"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"
)

type ResourceSpecDao interface {
	ListResourceSpec(pageIndex int32, pageSize int32) ([]*resources.ResourceSpec, error)
	CreateResourceSpec(request *resources.CreateResourceSpecRequest) (string, error)
	DeleteResourceSpec(id string) (string, error)
	GetResourceSpec(id string) (*resources.ResourceSpec, error)
	UpdateResourceSpec(resource *resources.ResourceSpec) (string, error)
	GetResourceSpecIgnore(id string) (*resources.ResourceSpec, error)
}

type resourceSepcDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewResourceSpecDao(db *gorm.DB, logger log.Logger) ResourceSpecDao {
	return &resourceSepcDao{
		log: log.NewHelper("ResourceSpecDao", logger),
		db:  db,
	}
}

func (d *resourceSepcDao) ListResourceSpec(pageIndex int32, pageSize int32) ([]*resources.ResourceSpec, error) {
	db := d.db
	resourceSpec := make([]*resources.ResourceSpec, 0)
	if pageIndex < 0 || pageSize < 0 {
		return nil, errors.Errorf(nil, errors.ErrorDBFindFailed)
	} else if pageIndex == 0 && pageSize == 0 {

		if err := db.Find(&resourceSpec).Error; err != nil {
			return nil, err
		}
	} else {
		if pageIndex >= 1 {
			offset := (pageIndex - 1) * pageSize
			if err := db.Limit(int(pageSize)).Offset(int(offset)).Find(&resourceSpec).Error; err != nil {
				return nil, err
			}
		}
	}

	return resourceSpec, nil
}

func (d *resourceSepcDao) CreateResourceSpec(request *resources.CreateResourceSpecRequest) (string, error) {
	db := d.db
	id := utils.GetUUIDWithoutSeparator()

	resourceSpec := &resources.ResourceSpec{
		Id:               id,
		Name:             request.Name,
		Price:            request.Price,
		ResourceQuantity: request.ResourceQuantity,
	}

	if err := db.Create(&resourceSpec).Error; err != nil {
		return "", errors.Errorf(err, errors.ErrorDBCreateFailed)
	}

	return id, nil
}

//No Update Dao Fun for ResourceSpec

func (d *resourceSepcDao) DeleteResourceSpec(id string) (string, error) {
	db := d.db

	if err := db.Delete(&resources.ResourceSpec{Id: id}).Error; err != nil {
		return "", errors.Errorf(err, errors.ErrorDBDeleteFailed)
	}

	return id, nil
}

func (d *resourceSepcDao) GetResourceSpec(id string) (*resources.ResourceSpec, error) {
	db := d.db

	resourceSpec := &resources.ResourceSpec{Id: id}

	if err := db.Find(&resourceSpec).Error; err != nil {
		return nil, err
	}

	return resourceSpec, nil
}

func (d *resourceSepcDao) UpdateResourceSpec(resource *resources.ResourceSpec) (string, error) {
	db := d.db

	if err := db.Save(resource).Error; err != nil {
		return "", err
	}

	return resource.Id, nil
}

func (d *resourceSepcDao) GetResourceSpecIgnore(id string) (*resources.ResourceSpec, error) {
	db := d.db

	resourceSpec := &resources.ResourceSpec{Id: id}

	if err := db.Unscoped().Find(&resourceSpec).Error; err != nil {
		return nil, err
	}

	return resourceSpec, nil
}
