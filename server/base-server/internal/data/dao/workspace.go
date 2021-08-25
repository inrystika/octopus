package dao

import (
	"context"
	"errors"
	"server/base-server/internal/data/dao/model"
	cerrors "server/common/errors"

	"server/common/log"

	"gorm.io/gorm"
)

type WorkspaceDao interface {
	List(ctx context.Context, condition *model.WorkspaceList) ([]*model.Workspace, error)
	Count(ctx context.Context, condition *model.WorkspaceList) (int64, error)
	Find(ctx context.Context, condition *model.WorkspaceQuery) (*model.Workspace, error)
	Add(ctx context.Context, workspace *model.WorkspaceAdd) (*model.Workspace, error)
	Update(ctx context.Context, workspace *model.WorkspaceUpdate) (*model.Workspace, error)
	Delete(ctx context.Context, workspace *model.WorkspaceDelete) (*model.Workspace, error)
	AddWorkspaceUsers(ctx context.Context, workspaceUsers *model.WorkspaceUserBatchAdd) error
	DelWorkspaceUsers(ctx context.Context, workspaceUsers *model.WorkspaceUserBatchDel) error
	ListWorkspaceUser(ctx context.Context, condition *model.WorkspaceUserList) ([]*model.User, error)
	ListUserWorkspace(ctx context.Context, condition *model.UserWorkspaceList) ([]*model.Workspace, error)
	ListIn(ctx context.Context, condition *model.WorkspaceListIn) ([]*model.Workspace, error)
}

type workspaceDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewWorkspaceDao(db *gorm.DB, logger log.Logger) WorkspaceDao {
	return &workspaceDao{
		log: log.NewHelper("WorkspaceDao", logger),
		db:  db,
	}
}

func (d *workspaceDao) List(ctx context.Context, condition *model.WorkspaceList) ([]*model.Workspace, error) {
	db := d.db
	workspaces := make([]*model.Workspace, 0)

	db = condition.Pagination(db)
	db = condition.Order(db)
	db = condition.Where(db)
	db = condition.Or(db)

	db.Find(&workspaces)
	return workspaces, nil
}

func (d *workspaceDao) Count(ctx context.Context, condition *model.WorkspaceList) (int64, error) {
	db := d.db
	var count int64

	db = condition.Where(db)
	db = condition.Or(db)

	db.Model(&model.Workspace{}).Count(&count)
	return count, nil
}

func (d *workspaceDao) Find(ctx context.Context, condition *model.WorkspaceQuery) (*model.Workspace, error) {
	db := d.db

	var workspace model.Workspace
	result := db.Where(&model.Workspace{Id: condition.Id}).First(&workspace)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &workspace, nil
}

func (d *workspaceDao) Add(ctx context.Context, workspace *model.WorkspaceAdd) (*model.Workspace, error) {
	if workspace.Id == "" {
		return nil, gorm.ErrPrimaryKeyRequired
	}

	db := d.db
	w := &model.Workspace{
		Id:      workspace.Id,
		Name:    workspace.Name,
		RPoolId: workspace.RPoolId,
	}

	result := db.Create(w)
	if result.Error != nil {
		return nil, result.Error
	}
	return w, nil
}

func (d *workspaceDao) Update(ctx context.Context, workspace *model.WorkspaceUpdate) (*model.Workspace, error) {
	if workspace.Id == "" {
		return nil, gorm.ErrPrimaryKeyRequired
	}

	result := d.db.Model(&model.Workspace{Id: workspace.Id}).Updates(&model.Workspace{
		Name:    workspace.Name,
		RPoolId: workspace.RPoolId,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return d.Find(ctx, &model.WorkspaceQuery{Id: workspace.Id})
}

func (d *workspaceDao) Delete(ctx context.Context, workspace *model.WorkspaceDelete) (*model.Workspace, error) {
	if w, err := d.Find(ctx, &model.WorkspaceQuery{Id: workspace.Id}); err != nil {
		return nil, err
	} else {
		if w == nil {
			return nil, nil
		}
		result := d.db.Delete(&model.Workspace{Id: w.Id})
		if result.Error != nil {
			return nil, result.Error
		}
		return w, nil
	}
}

func (d *workspaceDao) AddWorkspaceUsers(ctx context.Context, usersAppend *model.WorkspaceUserBatchAdd) error {
	if len(usersAppend.UserIds) < 1 {
		return nil
	}

	db := d.db

	uws := []model.WorkspaceUser{}
	userMap := make(map[string]string)
	for _, userId := range usersAppend.UserIds {
		if _, yes := userMap[userId]; yes {
			continue
		}
		userMap[userId] = userId
		uws = append(uws, model.WorkspaceUser{
			UserId:      userId,
			WorkspaceId: usersAppend.WorkspaceId,
		})
	}
	result := db.Create(&uws)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *workspaceDao) DelWorkspaceUsers(ctx context.Context, batchDel *model.WorkspaceUserBatchDel) error {
	if batchDel.WorkspaceId == "" {
		return cerrors.Errorf(nil, cerrors.ErrorDBPrimaryKeyEmpty)
	}
	db := d.db
	if len(batchDel.UserIds) > 0 {
		db = db.Where("user_id in ? and workspace_id = ?", batchDel.UserIds, batchDel.WorkspaceId)
	}
	wu := model.WorkspaceUser{}
	if result := db.Unscoped().Delete(&wu); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *workspaceDao) ListWorkspaceUser(ctx context.Context, cond *model.WorkspaceUserList) ([]*model.User, error) {
	workspace, err := d.Find(ctx, &model.WorkspaceQuery{Id: cond.WorkspaceId})
	if err != nil {
		return nil, err
	}
	if workspace == nil {
		return nil, nil
	}
	if err = d.db.Model(workspace).Association("Users").Find(&workspace.Users); err != nil {
		return nil, err
	}
	return workspace.Users, nil
}

func (d *workspaceDao) ListUserWorkspace(ctx context.Context, cond *model.UserWorkspaceList) ([]*model.Workspace, error) {
	var user model.User
	result := d.db.Where(&model.User{Id: cond.UserId}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := d.db.Model(&user).Association("Workspaces").Find(&user.Workspaces); err != nil {
		return nil, err
	}

	return user.Workspaces, nil
}

func (d *workspaceDao) ListIn(ctx context.Context, condition *model.WorkspaceListIn) ([]*model.Workspace, error) {
	if len(condition.Ids) < 1 {
		return nil, gorm.ErrMissingWhereClause
	}

	var workspaces []*model.Workspace
	result := d.db.Find(&workspaces, condition.Ids)
	if result.Error != nil {
		return nil, result.Error
	}

	return workspaces, nil
}
