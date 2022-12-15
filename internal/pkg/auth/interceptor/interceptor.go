package interceptor

import (
	"context"
	"strconv"

	"github.com/MyyPo/w34.Go/internal/jwt"
	"github.com/MyyPo/w34.Go/internal/statestore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	jwtManager      jwt.JWTManager
	redisClient     statestore.Service
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(
	jwtManager jwt.JWTManager,
	redisClient statestore.Service,
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
		authCtx, err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(authCtx, req)
	}

}

// unique keys for ctx
type idKey string
type accKey string
type refrKey string

const userIdKey idKey = "user_id"
const accTokKey accKey = "access_token"
const refrTokKey refrKey = "refresh_token"

func (i Interceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	// if no roles speciified, then the route is accessible to everyone
	_, ok := i.accessibleRoles[method]
	if !ok {
		return nil, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata was not provided")
	}
	accessArr := md["access_token"]
	if len(accessArr) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "access token was not provided")

	}
	accessToken := accessArr[0]
	claims, err := i.jwtManager.ValidateJwtExtractClaims(accessToken, i.jwtManager.AccessPublicSignature)
	if err != nil {
		refreshArr := md["refresh_token"]
		if len(refreshArr) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "access and refresh tokens were not provied")
		}
		refreshToken := refreshArr[0]
		// try to refresh if there is a refresh token
		return i.refresh(ctx, refreshToken)
	}

	authCtx := context.WithValue(ctx, userIdKey, claims.Subject)

	return authCtx, nil
}

// used to try to refresh the session
func (i Interceptor) refresh(ctx context.Context, refreshToken string) (context.Context, error) {
	claims, err := i.jwtManager.ValidateJwtExtractClaims(refreshToken, i.jwtManager.RefreshPublicSignature)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token")
	}

	tokenUserID := claims.Subject

	refreshTokenInDB, err := i.redisClient.GetToken(ctx, tokenUserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to authorize")
	}
	// verify that the token is whitelisted
	if refreshToken != refreshTokenInDB {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token`")
	}

	// create new tokens
	intUserID, _ := strconv.ParseInt(tokenUserID, 10, 32)
	int32UserID := int32(intUserID)
	newAccessToken, err := i.jwtManager.GenerateAccessToken(int32UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to authorize")
	}
	newRefreshToken, err := i.jwtManager.GenerateRefreshToken(int32UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to authorize")
	}

	err = i.redisClient.StoreRefreshTokenStringID(ctx, tokenUserID, newRefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to authorize")
	}

	authCtx := context.WithValue(ctx, userIdKey, claims.Subject)
	authCtx = context.WithValue(authCtx, accTokKey, newAccessToken)
	authCtx = context.WithValue(authCtx, refrTokKey, newRefreshToken)

	return authCtx, nil
}
