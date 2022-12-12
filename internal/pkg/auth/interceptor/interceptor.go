package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type Interceptor struct {
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(accessibleRoles map[string][]string) Interceptor {
	return Interceptor{
		accessibleRoles: accessibleRoles,
	}
}

func (i Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		return handler(ctx, req)
	}

}
