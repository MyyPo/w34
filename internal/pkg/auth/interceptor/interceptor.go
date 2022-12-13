package interceptor

import (
	"context"

	"github.com/MyyPo/w34.Go/internal/pkg/auth/jwt"
	"github.com/MyyPo/w34.Go/internal/statestore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	jwtManager      jwt.JWTManager
	redisClient     statestore.RedisClient
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(
	jwtManager jwt.JWTManager,
	redisClient statestore.RedisClient,
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
	_, err := i.jwtManager.ValidateJwtExtractClaims(accessToken, i.jwtManager.AccessPublicSignature)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "invalid access token")
	}

	return nil
}

// used to try to refresh the session
// func (i Interceptor) refresh(ctx context.Context, md metadata.MD) error {
// 	refreshArr := md["refresh_token"]
// 	if len(refreshArr) == 0 {
// 		return status.Errorf(codes.Unauthenticated, "no access or refresh token was provided by metadata")
// 	}
// 	refreshToken := refreshArr[0]
// 	claims, err := i.jwtManager.ValidateJwtExtractClaims(refreshToken, i.jwtManager.PathToRefreshPublicSignature)
// 	if err != nil {
// 		return status.Errorf(codes.Unauthenticated, "invalid refresh token")
// 	}

// 	tokenUserID := claims.Subject

// 	refreshTokenInDB, err := i.redisClient.GetToken(ctx, tokenUserID)
// 	if err != nil {
// 		return status.Errorf(codes.Internal, "failed to authorize")
// 	}
// 	// verify that the token is whitelisted
// 	if refreshToken != refreshTokenInDB {
// 		return status.Errorf(codes.Unauthenticated, "invalid refresh token")
// 	}

// 	err = i.redisClient.DeleteRefreshTokenStringID(ctx, tokenUserID)
// 	if err != nil {
// 		return status.Errorf(codes.Internal, "failed to authorize")
// 	}

// 	// create new tokens
// 	intUserID, _ := strconv.ParseInt(tokenUserID, 10, 32)
// 	int32UserID := int32(intUserID)
// 	newAccessToken, err := i.jwtManager.GenerateAccessToken(int32UserID)
// 	if err != nil {
// 		return status.Errorf(codes.Internal, "failed to authorize")
// 	}
// 	newRefreshToken, err := i.jwtManager.GenerateRefreshToken(int32UserID)
// 	if err != nil {
// 		return status.Errorf(codes.Internal, "failed to authorize")
// 	}

// 	err = i.redisClient.StoreRefreshTokenStringID(ctx, tokenUserID, newRefreshToken)
// 	if err != nil {
// 		return status.Errorf(codes.Internal, "failed to authorize")
// 	}

// 	return nil
// }
