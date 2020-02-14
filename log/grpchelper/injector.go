package grpchelper

import (
	"context"

	"github.com/anz-bank/pkg/log"
	"google.golang.org/grpc"
)

// BuildLoggerInjectorMiddleware returns a middleware that injects logger to the grpc context.
// It accepts fields so that you can add fields and configurations. If fields has no logger,
// a standard logger will be injected.
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

// InfoMiddleware returns a middleware that does a non-verbose log.
func InfoMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	log.Infof(ctx, "Info: %s", info.FullMethod)
	return
}
