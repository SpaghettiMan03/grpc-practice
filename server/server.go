package main

import (
	"fmt"
	"grpc-practice/gen/api"
	"grpc-practice/handler"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(8*1024*1024),
		grpc.ConnectionTimeout(1*time.Minute),
	)

	api.RegisterImageUploadServiceServer(
		server,
		handler.NewImageUploadHandler(),
	)
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port:%v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("stopping gRPC server")
	server.GracefulStop()
}

//func main() {
//	port := 50051
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
//	if err != nil {
//		log.Fatalf("failed to listen:%v", err)
//	}
//
//	logger, err := zap.NewProduction()
//	if err != nil {
//		panic(err)
//	}
//	grpc_zap.ReplaceGrpcLoggerV2(logger)
//
//	server := grpc.NewServer(
//		grpc.UnaryInterceptor(
//			grpc_middleware.ChainUnaryServer(
//				grpc_zap.UnaryServerInterceptor(logger),
//				grpc_auth.UnaryServerInterceptor(auth),
//			),
//		),
//	)
//	api.RegisterPancakeBakerServiceServer(
//		server,
//		handler.NewBakerHandler(),
//	)
//	reflection.Register(server)
//
//	go func() {
//		log.Printf("start gRPC server port: %v", port)
//		server.Serve(lis)
//	}()
//
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt)
//	<-quit
//	log.Println("stopping gRPC server")
//	server.GracefulStop()
//}
//
//func auth(ctx context.Context) (context.Context, error) {
//	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
//	if err != nil {
//		return nil, err
//	}
//
//	if token != "hi/mi/tsu" {
//		return nil, status.Error(codes.Unauthenticated, "無効なトークンです")
//	}
//
//	return context.WithValue(ctx, "UserName", "God"), nil
//}
