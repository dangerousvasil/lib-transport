package bgserver

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"lib-transport/interceptor"
	"lib-transport/jwtman"
	"time"
)

var (
	ErrMetadataIsNotProvided = errors.New("metadata is not provided")
	ErrAccessTokenIsInvalid  = errors.New("access token is invalid")
)

type JWTManager interface {
	Verify(accessToken string) (*jwtman.UserClaims, error)
	Generate(user *jwtman.UserClaims, tokenDuration time.Duration, otp bool) (string, error)
}

type BaseService struct {
	JWTManager JWTManager
}

func (s *BaseService) GetUserClaimsFromCTX(ctx context.Context) (*jwtman.UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMetadataIsNotProvided
	}

	accessToken := ""
	values := md["authorization"]
	if len(values) == 0 {
		gCookie := md["grpcgateway-cookie"]
		if len(gCookie) > 0 {
			gHeaders := interceptor.CookieHeader(gCookie[0])
			for _, header := range gHeaders {
				if header.Name == "X-TOKEN" {
					accessToken = header.Value
					break
				}
			}
		} else {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided md: %v", md)
		}
	} else {
		accessToken = values[0]
	}

	claims, err := s.JWTManager.Verify(accessToken)
	if err != nil {
		slog.Error("get claims from accessToken in GetUserClaimsFromCTX", "err", err, "md", md)
		return nil, errors.Wrap(ErrAccessTokenIsInvalid, err.Error())
	}
	return claims, nil
}
