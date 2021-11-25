package grpc

import (
	"context"
	"fibonachi/internal/config"
	"fibonachi/internal/delivery"
	"fibonachi/internal/delivery/grpc/pb"
	"fibonachi/internal/service"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

type Server struct {
	fibonacci service.Fibonacci
	pb.FibonacciServer
}

func New(fibo service.Fibonacci) *Server {
	return &Server{
		fibonacci: fibo,
	}
}

func (s *Server) Run(quit chan bool) {
	cfg := config.New()
	listener, err := net.Listen("tcp", fmt.Sprint(":", cfg.GrpcPort))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFibonacciServer(grpcServer, s)
	fmt.Println("Grpc server started successfully")

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Println(err)
		}
		<-quit
		grpcServer.GracefulStop()
	}()
}

func (s *Server) Post(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	response := &pb.Response{}
	isValid := delivery.ValidateInputParameters(request.X, request.Y)
	if !isValid {
		logrus.Errorf("Validation failed for x: %d, y: %d", request.X, request.Y)
		response.Error = "Validation failed"
		return response, nil
	}
	resp, err := s.fibonacci.CalculateResult(ctx, request.X, request.Y)
	if err != nil {
		response.Error = err.Error()
	}
	response.Numbers = resp

	return response, nil
}
