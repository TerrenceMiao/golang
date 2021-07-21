package server

import (
	"context"
	"fmt"

	pbh "github.com/terrencemiao/golang/protos/hello"
	pbc "github.com/terrencemiao/golang/protos/common"
)

type Server struct {
}

func (h *Server) Greet(ctx context.Context, req *pbh.GreetingRequest) (*pbh.GreetingResponse, error) {
	return &pbh.GreetingResponse{
		Greeting: fmt.Sprintf("Hello %s", req.GetName()),
	}, nil
}

func (h *Server) Bogus(ctx context.Context, req *pbc.BogusRequest) (*pbc.BogusResponse, error) {
	return &pbc.BogusResponse{
		Error: fmt.Sprintf("Bogus, %s", "thing go wrong!!!"),
	}, nil
}
