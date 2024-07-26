package data

import (
	"context"
	"server/base-server/internal/conf"
	"server/base-server/internal/data/cluster"
	"server/base-server/internal/data/dao"
	"server/base-server/internal/data/dao/algorithm_dao"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/dao/model/resources"
	"server/base-server/internal/data/ftp"
	"server/base-server/internal/data/influxdb"
	"server/base-server/internal/data/minio"
	"server/base-server/internal/data/prometheus"
	"server/base-server/internal/data/redis"
	"server/base-server/internal/data/registry"

	"server/common/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	UserDao               dao.UserDao
	AdminUserDao          dao.AdminUserDao
	AlgorithmDao          algorithm_dao.AlgorithmDao
	ResourceDao           dao.ResourceDao
	ResourceSpecDao       dao.ResourceSpecDao
	DevelopDao            dao.DevelopDao
	TrainJobDao           dao.TrainJobDao
	ModelDao              dao.ModelDao
	DatasetDao            dao.DatasetDao
	WorkspaceDao          dao.WorkspaceDao
	ImageDao              dao.ImageDao
	BillingDao            dao.BillingDao
	LableDao              dao.LableDao
	Cluster               cluster.Cluster
	Minio                 minio.Minio
	Registry              registry.ArtifactRegistry
	Redis                 redis.Redis
	Influxdb              influxdb.Influxdb
	ModelDeployDao        dao.ModelDeployDao
	PlatformStatisticsDao dao.PlatformStatisticsDao
	UserEndpointDao       dao.UserEndpointDao
	Prometheus            prometheus.Prometheus
	Ftp                   *ftp.Ftp
}

func NewData(bc *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	confData := bc.Data
	d := &Data{}

	db, err := dbInit(confData)
	if err != nil {
		return nil, nil, err
	}

	influxdb, err := influxdb.NewInfluxdb(confData)
	if err != nil {
		if !bc.App.IsDev {
			return nil, nil, err
		}
		log.Error(context.TODO(), err)
	}

	d.Minio = minio.NewMinio(confData, logger)
	d.Registry = registry.NewRegistry(confData, logger)
	redis, err := redis.NewRedis(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	d.Redis = redis
	cluster, clusterCancel, err := cluster.NewCluster(confData, logger)
	d.Cluster = cluster
	prometheus := prometheus.NewPrometheus(confData.Prometheus.BaseUrl)
	d.Prometheus = prometheus

	d.UserDao = dao.NewUserDao(db, logger)
	d.AdminUserDao = dao.NewAdminUserDao(db, logger)
	d.AlgorithmDao = algorithm_dao.NewAlgorithmDao(db, logger)
	d.ResourceDao = dao.NewResourceDao(db, logger)
	d.ResourceSpecDao = dao.NewResourceSpecDao(db, logger)
	d.DevelopDao = dao.NewDevelopDao(db, influxdb, logger)
	d.ModelDao = dao.NewModelDao(db, logger)
	d.ModelDeployDao = dao.NewModelDeployDao(db, influxdb, logger)
	d.DatasetDao = dao.NewDatasetDao(db, logger)
	d.WorkspaceDao = dao.NewWorkspaceDao(db, logger)
	d.ImageDao = dao.NewImageDao(db, logger)
	d.TrainJobDao = dao.NewTrainJobDao(db, influxdb, logger)
	d.BillingDao = dao.NewBillingDao(db, logger)
	d.LableDao = dao.NewLableDao(db, logger)
	d.PlatformStatisticsDao = dao.NewPlatformStatisticsDao(db)
	d.UserEndpointDao = dao.NewUserEndpointDao(db)
	d.Ftp = ftp.NewFtp(bc)

	return d, func() {
		clusterCancel()
		redis.Close()
	}, nil
}

func dbInit(confData *conf.Data) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(confData.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   log.DefaultGormLogger,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.AdminUser{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Algorithm{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.AlgorithmVersion{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.AlgorithmAccess{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.AlgorithmAccessVersion{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&resources.Resource{})
	if err != nil {
		return nil, err
	}
	idxRemove := "idx_resource_name"
	if db.Migrator().HasIndex(&resources.Resource{}, idxRemove) {
		db.Migrator().DropIndex(&resources.Resource{}, idxRemove)
	}
	err = db.AutoMigrate(&model.Workspace{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.WorkspaceUser{})
	if err != nil {
		return nil, err
	}
	err = db.SetupJoinTable(&model.Workspace{}, "Users", &model.WorkspaceUser{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Image{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.ImageAccess{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&resources.ResourceSpec{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.TrainJob{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.TrainJobTemplate{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.ModelDeploy{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Model{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.ModelVersion{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.ModelAccess{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.ModelVersionAccess{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Notebook{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.NotebookJob{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Dataset{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.DatasetVersion{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.DatasetAccess{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.DatasetVersionAccess{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.BillingOwner{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.BillingRechargeRecord{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.BillingPayRecord{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Lable{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.UserConfig{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.UserEndpoint{})
	if err != nil {
		return nil, err
	}

	return db, err
}
