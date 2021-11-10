package logic

import (
	"api/internal/common/errorx"
	"api/model"
	"context"
	"encoding/json"

	"api/internal/svc"
	"api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdminAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminAddLogic {
	return AdminAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminAddLogic) AdminAdd(req types.AddAdminReq) (*types.AddAdminResq, error) {
	_, err := l.svcCtx.AdminModel.Insert(model.Admins{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
		Status:   req.Status,
	})
	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加管理员信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加管理员失败")
	}

	return &types.AddAdminResq{
		Code:    "000000",
		Message: "添加管理员成功",
	}, nil
}
