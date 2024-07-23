package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	v1 "server/common/api/v1"
	"server/common/errors"
	"server/common/log"
	sftpgov2 "server/common/sftpgo/v2/openapi"
	"server/common/utils"
	"strings"
)

type UserService struct {
	api.UnimplementedUserServiceServer
	conf       *conf.Bootstrap
	log        *log.Helper
	data       *data.Data
	defaultPVS common.PersistentVolumeSourceExtender
}

const (
	UNLIMITED = 0
)

func NewUserService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.UserServiceServer {
	pvs, err := common.BuildStorageSource(conf.Storage)
	if err != nil {
		panic(err)
	}
	if pvs.Size() == 0 {
		panic("mod init failed, missing config [module.storage.source]")
	}

	return &UserService{
		conf:       conf,
		log:        log.NewHelper("UserService", logger),
		data:       data,
		defaultPVS: *pvs,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *api.ListUserRequest) (*api.ListUserReply, error) {
	usersTbl, err := s.data.UserDao.List(ctx, &model.UserList{
		SortBy:    req.SortBy,
		OrderBy:   req.OrderBy,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		FullName:  req.FullName,
		Email:     req.Email,
		SearchKey: req.SearchKey,
		Phone:     req.Phone,
		Status:    int32(req.Status),
		Desc:      req.Desc,
	})
	if err != nil {
		return nil, err
	}

	usersCount, err := s.data.UserDao.Count(ctx, &model.UserList{
		FullName:  req.FullName,
		Email:     req.Email,
		SearchKey: req.SearchKey,
		Status:    int32(req.Status),
	})
	if err != nil {
		return nil, err
	}

	users := make([]*api.UserItem, len(usersTbl))
	for idx, user := range usersTbl {
		bindInfo := make([]*api.Bind, 0)
		if user.Bind != nil {
			for _, a := range user.Bind {
				replyBind := new(api.Bind)
				replyBind.Platform = a.Platform
				replyBind.UserId = a.UserId
				replyBind.UserName = a.UserName
				bindInfo = append(bindInfo, replyBind)
			}
		}
		item := &api.UserItem{
			Id:            user.Id,
			FullName:      user.FullName,
			Email:         user.Email,
			Phone:         user.Phone,
			Gender:        api.GenderType(user.Gender),
			Status:        api.UserStatus(user.Status),
			Password:      user.Password,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
			Bind:          bindInfo,
			ResourcePools: user.ResourcePools,
			Desc:          user.Desc,
			Permission:    (*v1.UserPermission)(user.Permission),
			MinioUserName: user.MinioUserName,
			Buckets:       user.Buckets,
			EmailNotify:   user.EmailNotify,
		}
		users[idx] = item
	}

	return &api.ListUserReply{
		TotalSize: usersCount,
		Users:     users,
	}, nil
}

func (s *UserService) CheckOrInitUser(ctx context.Context, req *api.CheckOrInitUserRequest) (*api.CheckOrInitUserReply, error) {
	// to check or init user home storage bucket
	err := s.initUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.CheckOrInitUserReply{}, nil
}

func (s *UserService) FindUser(ctx context.Context, req *api.FindUserRequest) (*api.FindUserReply, error) {
	a := model.UserQuery{
		Id:    req.Id,
		Email: req.Email,
		Phone: req.Phone,
	}
	if req.Bind != nil {
		a.Bind = &model.Bind{
			Platform: req.Bind.Platform,
			UserId:   req.Bind.UserId,
			UserName: req.Bind.UserName,
		}
	}
	user, err := s.data.UserDao.Find(ctx, &a)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &api.FindUserReply{
			User: nil,
		}, nil
	}
	bindInfo := make([]*api.Bind, 0)
	if user.Bind != nil {
		for _, a := range user.Bind {
			replyBind := new(api.Bind)
			replyBind.Platform = a.Platform
			replyBind.UserId = a.UserId
			replyBind.UserName = a.UserName
			bindInfo = append(bindInfo, replyBind)
		}
	}
	reply := &api.FindUserReply{
		User: &api.UserItem{
			Id:            user.Id,
			FullName:      user.FullName,
			Email:         user.Email,
			Phone:         user.Phone,
			FtpUserName:   user.FtpUserName,
			Gender:        api.GenderType(user.Gender),
			Status:        api.UserStatus(user.Status),
			Password:      user.Password,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
			Bind:          bindInfo,
			ResourcePools: user.ResourcePools,
			Desc:          user.Desc,
			Permission:    (*v1.UserPermission)(user.Permission),
			MinioUserName: user.MinioUserName,
			Buckets:       user.Buckets,
			EmailNotify:   user.EmailNotify,
		},
	}

	return reply, nil
}

func (s *UserService) initUser(ctx context.Context, userId string) error {
	user, err := s.data.UserDao.Find(ctx, &model.UserQuery{Id: userId})
	if err != nil {
		return err
	}

	err = s.data.Minio.CreateBucket(common.GetUserBucket(userId))
	if err != nil {
		if !errors.IsError(errors.ErrorMinioBucketExisted, err) {
			return err
		}
	}

	// create user namespace
	_, err = s.data.Cluster.GetNamespace(ctx, userId)
	if err != nil {
		_, err = s.data.Cluster.CreateNamespace(ctx, userId)
		if err != nil {
			return err
		}
	}

	// create user storage pv
	pv := common.BuildStoragePersistentVolume(userId, s.defaultPVS.Capacity)
	pv.Spec.PersistentVolumeSource = s.defaultPVS.PersistentVolumeSource
	_, err = s.data.Cluster.GetPersistentVolume(ctx, pv.Name)
	if err != nil {
		_, err = s.data.Cluster.CreatePersistentVolume(ctx, pv)
		if err != nil {
			return err
		}
	}

	// create user storage pvc
	pvc := common.BuildStoragePersistentVolumeChaim(userId, userId, s.defaultPVS.Capacity)
	_, err = s.data.Cluster.GetPersistentVolumeClaim(ctx, pvc.Namespace, pvc.Name)
	if err != nil {
		_, err = s.data.Cluster.CreatePersistentVolumeClaim(ctx, pvc)
		if err != nil {
			return err
		}
	}

	for idx, storage := range conf.Storages {
		pv := common.BuildStoragePersistentVolume(common.GetExtraHomeVFName(idx, userId), storage.Requests)
		pv.Spec.PersistentVolumeSource = storage.StorageType.PersistentVolumeSource
		_, err = s.data.Cluster.GetPersistentVolume(ctx, pv.Name)
		if err != nil {
			_, err = s.data.Cluster.CreatePersistentVolume(ctx, pv)
			if err != nil {
				return err
			}
		}

		pvc := common.BuildStoragePersistentVolumeChaim(userId, common.GetExtraHomeVFName(idx, userId), storage.Requests)
		_, err = s.data.Cluster.GetPersistentVolumeClaim(ctx, pvc.Namespace, pvc.Name)
		if err != nil {
			_, err = s.data.Cluster.CreatePersistentVolumeClaim(ctx, pvc)
			if err != nil {
				return err
			}
		}
	}

	if user.FtpUserName != "" { // 额外的存储是新加的，通过用户登录构建SFTP虚拟文件夹
		utils.HandlePanic(ctx, func(i ...interface{}) {
			err := s.createOrUpdateFtpAccount(ctx, &CreateOrUpdateFtpAccountRequest{
				FtpUserName: user.FtpUserName,
				UserId:      userId,
			})
			log.Error(ctx, err)
		})()
	}

	return nil
}

func (s *UserService) AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserReply, error) {
	cond := model.UserQuery{
		Email: req.Email,
		Phone: req.Phone,
	}
	if req.Bind != nil {
		cond.Bind = &model.Bind{
			Platform: req.Bind.Platform,
			UserId:   req.Bind.UserId,
			UserName: req.Bind.UserName,
		}
	}
	existed, err := s.data.UserDao.Find(ctx, &cond)
	if err != nil {
		return nil, err
	}
	if existed != nil {
		return nil, errors.Errorf(nil, errors.ErrorUserAccountExisted)
	}

	password, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := model.UserAdd{}
	user.Email = req.Email
	user.Phone = req.Phone
	user.Password = password
	user.Id = utils.GetUUIDWithoutSeparator()
	user.FullName = req.FullName
	user.Gender = int32(req.Gender)
	user.Status = int32(api.UserStatus_ACTIVITY)
	user.Bind = cond.Bind
	user.ResourcePools = []string{s.conf.Service.Resource.DefaultPoolName}
	user.Desc = req.Desc
	u, err := s.data.UserDao.Add(ctx, &user)
	if err != nil {
		return nil, err
	}

	if err = s.initUser(ctx, u.Id); err != nil {
		s.log.Error(ctx, err)
		return nil, err
	}
	bindInfo := make([]*api.Bind, 0)
	if u.Bind != nil {
		for _, a := range u.Bind {
			replyBind := new(api.Bind)
			replyBind.Platform = a.Platform
			replyBind.UserId = a.UserId
			replyBind.UserName = a.UserName
			bindInfo = append(bindInfo, replyBind)
		}
	}
	reply := &api.AddUserReply{
		User: &api.UserItem{
			Id:       u.Id,
			FullName: u.FullName,
			Email:    u.Email,
			Phone:    u.Phone,
			Gender:   api.GenderType(u.Gender),
			Status:   api.UserStatus(u.Status),
			Password: u.Password,
			Bind:     bindInfo,
		},
	}

	reply.User.CreatedAt = u.CreatedAt.Unix()
	reply.User.UpdatedAt = u.UpdatedAt.Unix()

	return reply, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserReply, error) {
	userId := req.Id
	bindInfo := make([]*model.Bind, 0)
	if req.Bind != nil {
		for _, a := range req.Bind {
			reqBind := new(model.Bind)
			reqBind.Platform = a.Platform
			reqBind.UserId = a.UserId
			reqBind.UserName = a.UserName
			bindInfo = append(bindInfo, reqBind)
		}
	}
	user := model.UserUpdate{
		FullName:      req.FullName,
		Email:         req.Email,
		Phone:         req.Phone,
		Gender:        int32(req.Gender),
		Status:        int32(req.Status),
		ResourcePools: req.ResourcePools,
		Desc:          req.Desc,
		Permission:    (*model.Permission)(req.Permission),
		EmailNotify:   req.EmailNotify,
	}
	if len(bindInfo) > 0 {
		user.Bind = bindInfo
	}
	if req.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, err
		}
		user.Password = string(password)
	}
	result, err := s.data.UserDao.Update(ctx, &model.UserUpdateCond{Id: userId}, &user)
	if err != nil {
		return nil, err
	}

	bindInfo2 := make([]*api.Bind, 0)
	if result.Bind != nil {
		for _, a := range result.Bind {
			replyBind := new(api.Bind)
			replyBind.Platform = a.Platform
			replyBind.UserId = a.UserId
			replyBind.UserName = a.UserName
			bindInfo2 = append(bindInfo2, replyBind)
		}
	}
	return &api.UpdateUserReply{
		User: &api.UserItem{
			Id:       result.Id,
			FullName: result.FullName,
			Email:    result.Email,
			Phone:    result.Phone,
			Gender:   api.GenderType(result.Gender),
			Status:   api.UserStatus(result.Status),
			Password: result.Password,
			Bind:     bindInfo2,
			Desc:     result.Desc,
		},
	}, nil
}

func (s *UserService) ListUserInCond(ctx context.Context, req *api.ListUserInCondRequest) (*api.ListUserInCondReply, error) {
	users, err := s.data.UserDao.ListIn(ctx, &model.UserListIn{Ids: req.Ids})
	if err != nil {
		return nil, err
	}

	userItems := make([]*api.UserItem, len(users))
	for idx, user := range users {
		bindInfo := make([]*api.Bind, 0)
		if user.Bind != nil {
			for _, a := range user.Bind {
				replyBind := new(api.Bind)
				replyBind.Platform = a.Platform
				replyBind.UserId = a.UserId
				replyBind.UserName = a.UserName
				bindInfo = append(bindInfo, replyBind)
			}
		}
		item := &api.UserItem{
			Id:          user.Id,
			FullName:    user.FullName,
			Email:       user.Email,
			Phone:       user.Phone,
			Gender:      api.GenderType(user.Gender),
			Status:      api.UserStatus(user.Status),
			Password:    user.Password,
			CreatedAt:   user.CreatedAt.Unix(),
			UpdatedAt:   user.UpdatedAt.Unix(),
			Bind:        bindInfo,
			EmailNotify: user.EmailNotify,
		}
		userItems[idx] = item
	}
	return &api.ListUserInCondReply{
		Users: userItems,
	}, nil
}

func (s *UserService) ListUserConfigKey(ctx context.Context, req *api.ListUserConfigKeyRequest) (*api.ListUserConfigKeyReply, error) {
	return &api.ListUserConfigKeyReply{ConfigKeys: common.UserConfigKeys}, nil
}

func (s *UserService) GetUserConfig(ctx context.Context, req *api.GetUserConfigRequest) (*api.GetUserConfigReply, error) {
	config, err := s.data.UserDao.GetConfig(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &api.GetUserConfigReply{
		Config: config,
	}, nil
}

func (s *UserService) UpdateUserConfig(ctx context.Context, req *api.UpdateUserConfigRequest) (*api.UpdateUserConfigReply, error) {
	for k, v := range req.Config {
		in := false
		for _, i := range common.UserConfigKeys {
			if k == i.Key {
				in = true
				err := i.ValidateValue(v)
				if err != nil {
					return nil, err
				}
				break
			}
		}

		if !in {
			return nil, errors.Errorf(nil, errors.ErrorUserConfigKeyNotExist)
		}
	}

	err := s.data.UserDao.UpdateConfig(ctx, req.UserId, req.Config)
	if err != nil {
		return nil, err
	}

	return &api.UpdateUserConfigReply{}, nil
}

func (s *UserService) buildFtpHomeDir(ctx context.Context, userId string) string {
	return fmt.Sprintf("/minio/%s/%s", userId, common.USERHOME)
}

func (s *UserService) UpdateUserFtpAccount(ctx context.Context, req *api.UpdateUserFtpAccountRequest) (*api.UpdateUserFtpAccountReply, error) {
	user, err := s.data.UserDao.Find(ctx, &model.UserQuery{Id: req.UserId})
	if err != nil {
		return nil, err
	}
	err = s.createOrUpdateFtpAccount(ctx, &CreateOrUpdateFtpAccountRequest{
		FtpUserName: req.FtpUserName,
		Email:       user.Email,
		Password:    req.FtpPassword,
		UserId:      req.UserId,
	})
	if err != nil {
		return nil, err
	}

	_, err = s.data.UserDao.Update(ctx, &model.UserUpdateCond{Id: req.UserId}, &model.UserUpdate{FtpUserName: req.FtpUserName})
	if err != nil {
		return nil, err
	}

	return &api.UpdateUserFtpAccountReply{}, nil
}

func (s *UserService) UpdateUserMinioAccount(ctx context.Context, req *api.UpdateUserMinioAccountRequest) (*api.UpdateUserMinioAccountReply, error) {
	user, err := s.data.UserDao.Find(ctx, &model.UserQuery{Id: req.UserId})
	if err != nil {
		return nil, err
	}

	if user.MinioUserName != "" && user.MinioUserName != req.MinioUserName {
		return nil, errors.Errorf(nil, errors.ErrorUserChangeMinioUsernameForbidden)
	}
	err = s.data.Minio.CreateOrUpdateAccount(ctx, req.MinioUserName, req.MinioPassword)
	if err != nil {
		return nil, err
	}

	_, err = s.data.UserDao.Update(ctx, &model.UserUpdateCond{Id: req.UserId}, &model.UserUpdate{MinioUserName: req.MinioUserName})
	if err != nil {
		return nil, err
	}

	return &api.UpdateUserMinioAccountReply{}, nil
}

func (s *UserService) UpdateUserMinioBuckets(ctx context.Context, req *api.UpdateUserMinioBucketsRequest) (*api.UpdateUserMinioBucketsReply, error) {
	user, err := s.data.UserDao.Find(ctx, &model.UserQuery{Id: req.UserId})
	if err != nil {
		return nil, err
	}

	if user.MinioUserName == "" {
		return nil, errors.Errorf(nil, errors.ErrorUserMinioUsernameNotExist)
	}

	for _, b := range req.Buckets {
		isExist, err := s.data.Minio.BucketExists(ctx, b)
		if err != nil {
			return nil, err
		}
		if !isExist {
			err := s.data.Minio.CreateBucket(b)
			if err != nil {
				return nil, err
			}
		}
	}

	err = s.data.Minio.SetUserBucketsAccess(ctx, user.MinioUserName, req.Buckets)
	if err != nil {
		return nil, err
	}

	_, err = s.data.UserDao.Update(ctx, &model.UserUpdateCond{Id: req.UserId}, &model.UserUpdate{Buckets: req.Buckets})
	if err != nil {
		return nil, err
	}

	return &api.UpdateUserMinioBucketsReply{}, nil
}

type CreateOrUpdateFtpAccountRequest struct {
	FtpUserName string
	Email       string
	Password    string
	UserId      string
}

func (s *UserService) createOrUpdateFtpAccount(ctx context.Context, req *CreateOrUpdateFtpAccountRequest) error {
	var err error
	password := ""
	if req.Password != "" {
		password, err = utils.EncryptPassword(req.Password)
		if err != nil {
			return err
		}
	}

	fuser, err := s.data.Ftp.GetFtpUser(ctx, req.FtpUserName)
	if err != nil && !errors.IsError(errors.ErrorSFtpGOUserNotExist, err) {
		return err
	}

	if fuser == nil && req.Password != "" {
		fuser = sftpgov2.NewUser()
		//去掉minio中转
		//fileSystemConfig := s.newFileSystemConfig(req.HomeS3Bucket, req.HomeS3Object)
		permissions := map[string][]sftpgov2.Permission{
			"/": {sftpgov2.PERMISSION_STAR},
		}

		fuser.SetStatus(1)
		fuser.SetUid(0)
		fuser.SetGid(0)
		fuser.SetMaxSessions(UNLIMITED)
		fuser.SetQuotaSize(1)
		fuser.SetQuotaFiles(1)
		fuser.SetExpirationDate(UNLIMITED)
		fuser.SetPermissions(permissions)
		//fuser.SetFilesystem(*fileSystemConfig)
		fuser.SetHomeDir(common.GetTmpHomePath(req.UserId))
		fuser.SetUploadBandwidth(UNLIMITED)
		fuser.SetDownloadBandwidth(UNLIMITED)

		fuser.SetUsername(req.FtpUserName)
		fuser.SetEmail(req.Email)
		fuser.SetPassword(password)
		s.assignVF(fuser, req)

		_, err = s.data.Ftp.CreateFtpUser(ctx, fuser)
		if err != nil {
			return err
		}
	} else {
		needUpdate := false
		if req.FtpUserName != "" {
			fuser.SetUsername(req.FtpUserName)
		}
		if req.Email != "" {
			needUpdate = true
			fuser.SetEmail(req.Email)
		}
		if req.Password != "" {
			needUpdate = true
			fuser.SetPassword(password)
		}

		if !strings.HasPrefix(fuser.GetHomeDir(), common.GetTmpHomePath(req.UserId)) {
			needUpdate = true
			fileSystemConfig := sftpgov2.NewFilesystemConfig()
			fileSystemConfig.SetProvider(sftpgov2.FSPROVIDERS__0)
			fuser.SetFilesystem(*fileSystemConfig)
			fuser.SetHomeDir(common.GetTmpHomePath(req.UserId))
			fuser.SetQuotaSize(1)
			fuser.SetQuotaFiles(1)
		}

		update := s.assignVF(fuser, req)
		if update {
			needUpdate = true
		}

		if needUpdate {
			err := s.data.Ftp.UpdateFtpUser(ctx, *fuser, 1)
			if err != nil {
				return err
			}
			log.Infof(ctx, "%s update ftp", req.UserId)
		}

	}

	return nil
}

func (s *UserService) assignVF(fuser *sftpgov2.User, req *CreateOrUpdateFtpAccountRequest) bool {
	needUpdate := false
	vfMap := map[string]sftpgov2.VirtualFolder{}
	for _, vf := range fuser.VirtualFolders {
		vfMap[vf.GetName()] = vf
	}

	for i, _ := range conf.Storages {
		_, exist := vfMap[common.GetExtraHomeVFName(i, req.UserId)]
		if !exist {
			needUpdate = true
			folder := sftpgov2.NewVirtualFolder(common.GetExtraHomeVirtualPath(i))
			folder.SetName(common.GetExtraHomeVFName(i, req.UserId))
			folder.SetMappedPath(common.GetExtraHomeMappedPath(i, req.UserId))
			fuser.VirtualFolders = append(fuser.VirtualFolders, *folder)
		}
	}

	_, exist := vfMap[common.GetUserHomeVFName(req.UserId)]
	if !exist {
		needUpdate = true
		folder := sftpgov2.NewVirtualFolder(common.GetUserHomeVirtualPath())
		folder.SetName(common.GetUserHomeVFName(req.UserId))
		folder.SetMappedPath(common.GetUserHomeMappedPath(req.UserId))
		fuser.VirtualFolders = append(fuser.VirtualFolders, *folder)
	}

	return needUpdate
}
