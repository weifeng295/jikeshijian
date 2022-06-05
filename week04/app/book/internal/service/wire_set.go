package service

import (
	"github.com/google/wire"
	"week04/app/book/internal/biz"
	"week04/app/book/internal/data/ent"
)

// ProvideSet for service package ...
var ProvideSet = wire.NewSet(NewHTTPBookRepo, NewGRPCBookRepo, NewGRPCBookService)

func NewHTTPBookRepo(client *ent.Client) biz.HTTPBookRepo {
	return &HTTPBookService{Client: client}
}

func NewGRPCBookRepo(client *ent.Client) biz.GRPCBookRepo {
	return &GRPCBookService{Client: client}
}

func NewGRPCBookService(client *ent.Client) *GRPCBookService {
	return &GRPCBookService{Client: client}
}

//https://github.com/qmdx00/Go-Camp/blob/master/week04/app/book/internal/service/wire_set.go
