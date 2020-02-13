package grpchelper

import (
	"context"

	"github.com/anz-bank/pkg/log"
	"google.golang.org/grpc"
)

func BuildLoggerInjectorMiddleware(fieldsWithLogger log.Fields) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if !fieldsWithLogger.HasLogger() {
			ctx = log.WithLogger(log.NewStandardLogger()).Onto(ctx)
		} else {
			ctx = fieldsWithLogger.Onto(ctx)
		}
		return handler(ctx, req)
	}
}

func InfoMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Infof(ctx, "Info: %s", info.FullMethod)
	return handler(ctx, req)
}
