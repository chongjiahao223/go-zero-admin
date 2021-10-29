package logic

import (
	"api/internal/common/errorx"
	"api/internal/svc"
	"api/internal/types"
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"strings"
	"time"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminLoginLogic {
	return AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req types.LoginReq) (*types.LoginResq, error) {
	if len(strings.TrimSpace(req.Phone)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("手机号或密码不能为空,请求信息失败,参数:%s", reqStr)
		return nil, errorx.NewDefaultError("手机号或密码不能为空")
	}

	adminInfo, err := l.svcCtx.AdminModel.FindPhoneOne(req.Phone)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", req.Phone, err.Error())
		return nil, errorx.NewDefaultError("登录失败")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", req.Phone, err.Error())
		return nil, errorx.NewDefaultError("登录失败")
	}

	if req.Password != adminInfo.Password {
		logx.WithContext(l.ctx).Errorf("用户密码不正确,参数:%s", req.Password)
		return nil, errorx.NewDefaultError("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, adminInfo.Id)

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("登录失败")
	}

	return &types.LoginResq{
		Code: "000000",
		Message: "登录成功",
		Id: adminInfo.Id,
		Name: adminInfo.Name,
		Phone: adminInfo.Phone,
		AccessToken: jwtToken,
		AccessExipre: accessExpire,
	}, nil
}

func (l *AdminLoginLogic) getJwtToken(secretKey string, iat, seconds, adminId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["adminId"] = adminId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
