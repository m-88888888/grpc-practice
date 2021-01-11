package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/m-88888888/grpc-practice/pb/calc/proto"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const port = ":50051"

type ServerUnary struct {
	pb.UnimplementedCalcServer
}

func (s *ServerUnary) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumReply, error) {
	// Requestからmessageを取得
	a := in.GetA()
	b := in.GetB()
	fmt.Println(a, b)
	reply := fmt.Sprintf("%d + %d = %d", a, b, a+b)
	return &pb.SumReply{
		Message: reply,
	}, nil
}

func server() error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "missing port.")
	}
	s := grpc.NewServer()
	var server ServerUnary
	pb.RegisterCalcServer(s, &server)
	if err := s.Serve(listen); err != nil {
		return errors.Wrap(err, "Failed to start server.")
	}
	return nil
}

func main() {
	fmt.Println("gRPC server start.")
	//if err := server(); err != nil {
	//	log.Fatalf("%v", err)
	//}
	if err := product.server(); err != nil {
		log.Fatalf("%v", err)
	}
}
