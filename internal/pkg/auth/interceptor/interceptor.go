package interceptor

import (
	"context"

	"github.com/MyyPo/w34.Go/internal/pkg/auth/jwt"
	auth_redis "github.com/MyyPo/w34.Go/internal/pkg/auth/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	jwtManager      jwt.JWTManager
	redisClient     auth_redis.RedisClient
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(
	jwtManager jwt.JWTManager,
	redisClient auth_redis.RedisClient,
	accessibleRoles map[string][]string,
) Interceptor {

	return Interceptor{
		jwtManager:      jwtManager,
		redisClient:     redisClient,
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
		if err := i.authorize(ctx, info.FullMethod); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}

}

func (i Interceptor) authorize(ctx context.Context, method string) error {
	// if no roles speciified, then the route is accessible to everyone
	_, ok := i.accessibleRoles[method]
	if !ok {
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata was not provided")
	}
	accessArr := md["access_token"]
	if len(accessArr) == 0 {
		return status.Errorf(codes.Unauthenticated, "access token was not provided")
	}
	accessToken := accessArr[0]
	_, err := i.jwtManager.ValidateJwtExtractClaims(accessToken, i.jwtManager.PathToAccessPublicSignature)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "invalid access token")
	}

	return nil
}
