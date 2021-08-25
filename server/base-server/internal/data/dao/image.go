package dao

import (
	"context"
	"errors"
	"server/base-server/internal/data/dao/model"
	"server/common/log"

	"gorm.io/gorm"
)

type ImageDao interface {
	List(ctx context.Context, condition *model.ImageList) ([]*model.Image, error)
	Count(ctx context.Context, condition *model.ImageList) (int64, error)
	Find(ctx context.Context, condition *model.ImageQuery) (*model.Image, error)
	Add(ctx context.Context, image *model.ImageAdd) (*model.Image, error)
	Delete(ctx context.Context, image *model.ImageDel) (*model.Image, error)
	Update(ctx context.Context, condition *model.ImageUpdateCond, image *model.ImageUpdate) (*model.Image, error)
	AddImageAccess(context.Context, *model.ImageAccessAdd) (*model.ImageAccess, error)
	FindImageAccess(ctx context.Context, condition *model.ImageAccessQuery) (*model.ImageAccess, error)
	ListImageAccess(ctx context.Context, condition *model.ImageAccessList) ([]*model.ImageAccess, error)
	CountImageByAccess(ctx context.Context, condition *model.ImageAccessList) (int64, error)
	ListImageByAccess(ctx context.Context, condition *model.ImageAccessList) ([]*model.Image, error)
	DeleteImageAccess(ctx context.Context, image *model.ImageAccessDel) (*model.ImageAccess, error)
	CountImageAccess(ctx context.Context, condition *model.ImageAccessList) (int64, error)
	ListIn(ctx context.Context, condition *model.ImageListIn) ([]*model.Image, error)
	ListImageAccessIn(ctx context.Context, condition *model.ImageAccessListIn) ([]*model.ImageAccess, error)
}

type imageDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewImageDao(db *gorm.DB, logger log.Logger) ImageDao {
	return &imageDao{
		log: log.NewHelper("ImageDao", logger),
		db:  db,
	}
}

func (d *imageDao) List(ctx context.Context, condition *model.ImageList) ([]*model.Image, error) {
	db := d.db
	images := make([]*model.Image, 0)

	db = condition.Pagination(db)
	db = condition.Order(db)
	db = condition.Where(db)
	db = condition.Or(db)

	result := db.Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}

func (d *imageDao) Count(ctx context.Context, condition *model.ImageList) (int64, error) {
	db := d.db
	var count int64

	db = condition.Where(db)
	db = condition.Or(db)

	result := db.Model(&model.Image{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (d *imageDao) Find(ctx context.Context, condition *model.ImageQuery) (*model.Image, error) {
	db := d.db

	var image model.Image
	result := db.Preload("Accesses").Where(&model.Image{
		Id: condition.Id,
	}).First(&image)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &image, nil
}

func (d *imageDao) Add(ctx context.Context, imageAdd *model.ImageAdd) (*model.Image, error) {
	db := d.db

	i := model.Image{
		Id:           imageAdd.Id,
		ImageName:    imageAdd.ImageName,
		ImageVersion: imageAdd.ImageVersion,
		ImageType:    imageAdd.ImageType,
		ImageDesc:    imageAdd.ImageDesc,
		ImageAddr:    imageAdd.ImageAddr,
		SourceType:   imageAdd.SourceType,
		SpaceId:      imageAdd.SpaceId,
		UserId:       imageAdd.UserId,
		IsPrefab:     imageAdd.IsPrefab,
		Status:       imageAdd.Status,
	}

	result := db.Create(&i)
	if result.Error != nil {
		return nil, result.Error
	}

	return &i, nil
}

func (d *imageDao) Update(ctx context.Context, cond *model.ImageUpdateCond, imageUpdate *model.ImageUpdate) (*model.Image, error) {
	if cond.Id == "" {
		return nil, gorm.ErrPrimaryKeyRequired
	}

	condition := model.Image{
		Id: cond.Id,
	}

	result := d.db.Model(&condition).Updates(model.Image{
		SourceFilePath: imageUpdate.SourceFilePath,
		ImageName:      imageUpdate.ImageName,
		ImageVersion:   imageUpdate.ImageVersion,
		ImageType:      imageUpdate.ImageType,
		ImageDesc:      imageUpdate.ImageDesc,
		ImageAddr:      imageUpdate.ImageAddr,
		Status:         imageUpdate.Status,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return d.Find(ctx, &model.ImageQuery{
		Id: cond.Id,
	})
}

func (d *imageDao) Delete(ctx context.Context, imageDel *model.ImageDel) (*model.Image, error) {
	if i, err := d.Find(ctx, &model.ImageQuery{Id: imageDel.Id}); err != nil {
		return nil, err
	} else {
		if i == nil {
			return nil, nil
		}
		result := d.db.Delete(&model.Image{Id: i.Id})
		if result.Error != nil {
			return nil, result.Error
		}
		return i, nil
	}
}

func (d *imageDao) AddImageAccess(ctx context.Context, imageAccessAdd *model.ImageAccessAdd) (*model.ImageAccess, error) {
	db := d.db

	ia := model.ImageAccess{
		Id:      imageAccessAdd.Id,
		ImageId: imageAccessAdd.ImageId,
		SpaceId: imageAccessAdd.SpaceId,
		UserId:  imageAccessAdd.UserId,
	}

	result := db.Create(&ia)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ia, nil
}

func (d *imageDao) FindImageAccess(ctx context.Context, iaq *model.ImageAccessQuery) (*model.ImageAccess, error) {
	db := d.db
	var ia model.ImageAccess
	result := db.Where(&model.ImageAccess{
		Id:      iaq.Id,
		ImageId: iaq.ImageId,
		UserId:  iaq.UserId,
		SpaceId: iaq.SpaceId,
	}).Preload("Image").First(&ia)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &ia, nil
}

func (d *imageDao) DeleteImageAccess(ctx context.Context, iad *model.ImageAccessDel) (*model.ImageAccess, error) {
	if ia, err := d.FindImageAccess(ctx, &model.ImageAccessQuery{
		Id:      iad.Id,
		ImageId: iad.ImageId,
		UserId:  iad.UserId,
		SpaceId: iad.SpaceId,
	}); err != nil {
		return nil, err
	} else {
		if ia == nil {
			return nil, nil
		}
		result := d.db.Delete(&model.ImageAccess{Id: ia.Id})
		if result.Error != nil {
			return nil, result.Error
		}
		return ia, nil
	}
}

func (d *imageDao) ListImageAccess(ctx context.Context, condition *model.ImageAccessList) ([]*model.ImageAccess, error) {
	db := d.db
	db = condition.Pagination(db)
	db = condition.Order(db)
	db = condition.JoinImage(db)

	var imageAccesses []*model.ImageAccess
	result := db.Find(&imageAccesses)
	if result.Error != nil {
		return nil, result.Error
	}

	return imageAccesses, nil
}

func (d *imageDao) ListImageByAccess(ctx context.Context, condition *model.ImageAccessList) ([]*model.Image, error) {
	db := d.db
	db = condition.Pagination(db)
	db = condition.Order(db)
	db = condition.JoinImageAccess(db)

	var images []*model.Image
	result := db.Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}

func (d *imageDao) CountImageByAccess(ctx context.Context, condition *model.ImageAccessList) (int64, error) {
	db := d.db
	var count int64

	db = condition.JoinImageAccess(db)
	result := db.Model(&model.Image{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (d *imageDao) CountImageAccess(ctx context.Context, condition *model.ImageAccessList) (int64, error) {
	db := d.db
	var count int64

	db = condition.JoinImage(db)
	result := db.Model(&model.ImageAccess{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (d *imageDao) ListIn(ctx context.Context, condition *model.ImageListIn) ([]*model.Image, error) {
	if len(condition.Ids) < 1 {
		return nil, gorm.ErrMissingWhereClause
	}

	var images []*model.Image
	result := d.db.Find(&images, condition.Ids)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}

func (d *imageDao) ListImageAccessIn(ctx context.Context, condition *model.ImageAccessListIn) ([]*model.ImageAccess, error) {
	db := d.db
	db = condition.Where(db)

	var imageAccesses []*model.ImageAccess
	result := db.Find(&imageAccesses)
	if result.Error != nil {
		return nil, result.Error
	}

	return imageAccesses, nil
}
