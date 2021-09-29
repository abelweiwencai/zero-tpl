package logic

import (
	"context"

	"apiserver/internal/svc"
	"apiserver/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ApiserverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiserverLogic(ctx context.Context, svcCtx *svc.ServiceContext) ApiserverLogic {
	return ApiserverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiserverLogic) Apiserver(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
