package interceptor

import (
	"context"
	"golang.org/x/exp/slog"
	"lib-transport/jwtman"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor is a server interceptor for authentication and authorization
type AuthInterceptor struct {
	jwtManager      *jwtman.JWTManager
	accessibleRoles map[string][]string
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwtManager *jwtman.JWTManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleRoles}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// Stream returns a server interceptor function to authenticate and authorize stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	accessToken := ""
	values := md["authorization"]
	if len(values) == 0 {
		gCookie := md["grpcgateway-cookie"]
		if len(gCookie) > 0 {
			gHeaders := CookieHeader(gCookie[0])
			for _, header := range gHeaders {
				if header.Name == "X-VALUN-TOKEN" {
					accessToken = header.Value
					break
				}
			}
		} else {
			return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}
	} else {
		accessToken = values[0]
	}

	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		slog.Error("get claims from accessToken in authorize", "err", err)
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		return status.Errorf(codes.Unauthenticated, "session time is expire")
	}

	return correspondClaims(claims.BotID != 0, claims.Roles, accessibleRoles)
}

func CookieHeader(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}
	return req.Cookies()
}

// correspondClaims проверка доступа польователя к ручке.
// боту доступны только ручки с ролью bot и должна быть ещё одна общая роль.
func correspondClaims(isBot bool, claimsRoles, accessibleRoles []string) error {
	var findedRole bool
	var hasBotRole = !isBot
L1:
	for i := range accessibleRoles {
		if !hasBotRole && accessibleRoles[i] == "bot" {
			hasBotRole = true
		}
		for j := range claimsRoles {
			if accessibleRoles[i] == claimsRoles[j] {
				findedRole = true
				continue L1
			}
		}
	}

	if hasBotRole && findedRole {
		return nil
	}
	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
