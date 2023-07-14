package service

import (
	"context"
	"v1/internal/dao"
	"v1/internal/model"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	s := Service{ctx: ctx}
	s.dao = dao.New(model.DB)
	return s
}
