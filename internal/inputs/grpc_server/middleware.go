package grpc_server

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	msg := "LOG INTERCEPTOR"
	if err != nil {
		log.Printf("%v method: %v req: %v err: %v", msg, info.FullMethod, req, err)
	} else {
		log.Printf("%v method: %v req: %v", msg, info.FullMethod, req)
	}
	return resp, err
}
