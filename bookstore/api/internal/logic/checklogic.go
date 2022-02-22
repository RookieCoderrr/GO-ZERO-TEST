package logic

import (
	"GO-ZERO-TEST/bookstore/rpc/check/checker"
	"context"

	"GO-ZERO-TEST/bookstore/api/internal/svc"
	"GO-ZERO-TEST/bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (*types.CheckResp, error) {
	// manual code start
	resp, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Book:  req.Book,
	})
	if err != nil {
		return nil, err
	}

	return &types.CheckResp{
		Found: resp.Found,
		Price: resp.Price,
	}, nil
	// manual code stop
}
