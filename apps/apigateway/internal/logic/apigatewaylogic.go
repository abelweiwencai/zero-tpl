package logic

import (
	"context"

	"apigateway/internal/svc"
	"apigateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ApigatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApigatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) ApigatewayLogic {
	return ApigatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApigatewayLogic) Apigateway(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
