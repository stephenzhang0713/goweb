package logic

import (
	"context"

	"goweb/go-zero/demo/greet/internal/svc"
	"goweb/go-zero/demo/greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: "Hello go-zero",
	}, nil
}
