package local_grpc

import (
	"context"

	"github.com/costiss/go-rest-grpc/internal/service"
	pb "github.com/costiss/go-rest-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type calculatorGrpcServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *calculatorGrpcServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	first := int(in.First)
	second := int(in.Second)
	result := service.Add(first, second)

	return &pb.AddReply{Result: int32(result)}, nil
}

func MakeGrpcServer() *grpc.Server {

	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterCalculatorServer(server, &calculatorGrpcServer{})

	return server
}
