package server

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"week04/app/book/internal/conf"
	v12 "week04/app/book/internal/service"
	"week04/app/book/service/v1"
)

type GRPCServer struct {
	gp     *errgroup.Group
	config conf.GRPCOptions
	Server *grpc.Server
}

// NewGRPCServer provide GRPCServer with conf.Options ...
func NewGRPCServer(group *errgroup.Group, options conf.Options, service *v12.GRPCBookService) *GRPCServer {
	srv := grpc.NewServer()
	v1.RegisterBookServer(srv, service)
	return &GRPCServer{gp: group, config: options.Server.GRPC, Server: srv}
}

func (g *GRPCServer) Serve(ctx context.Context) {
	g.gp.Go(func() error {
		fmt.Println("[BOOK API] grpc listen on", g.config.Addr)
		lis, err := net.Listen("tcp", g.config.Addr)
		if err != nil {
			return err
		}
		return g.Server.Serve(lis)
	})

	g.gp.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("[BOOK API] grpc server exit ...")
			g.Server.Stop()
			return ctx.Err()
		}
	})
}

// Stop stop the GRPCServer ...
func (g *GRPCServer) Stop() {
	g.gp.Go(func() error {
		fmt.Println("[BOOK API] grpc server stop ...")
		g.Server.GracefulStop()
		return nil
	})
}
