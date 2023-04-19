package server_lib

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	context_lib "github.com/thesisK19/buildify/library/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Claims struct for JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Skip token verification for HealthCheck GKE and SignIn, SignUp api
	if strings.Contains(info.FullMethod, "HealthCheck") ||
		strings.Contains(info.FullMethod, "SignUp") || strings.Contains(info.FullMethod, "SignIn") {
		return handler(ctx, req)
	}

	// Extract the JWT token from the context or request headers
	token, err := extractTokenFromContextOrHeaders(ctx)
	if err != nil {
		return nil, err
	}

	// Verify and parse the JWT token to extract the claims
	claims, err := verifyAndParseJWT(token)
	if err != nil {
		return nil, err
	}

	// Add the extracted username to the context
	ctx = context.WithValue(ctx, context_lib.USERNAME, claims.Username)

	// Append the token to the outgoing metadata
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	// Call the next handler in the chain
	return handler(ctx, req)
}

func extractTokenFromContextOrHeaders(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "Missing metadata")
	}
	authorization := md.Get("authorization")
	if len(authorization) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "Missing token")
	}
	token := authorization[0]
	return token, nil
}

func verifyAndParseJWT(tokenString string) (*Claims, error) {
	// Retrieve the jwtSecret environment variable
	jwtSecret := os.Getenv("jwtSecret")

	if jwtSecret == "" {
		// Handle the case when jwtSecret is not set
		return nil, fmt.Errorf("jwtSecret environment variable is not set")
	}

	// Parse and validate the JWT token
	parsedToken, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil // Use your secret key for token validation
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
	}

	// Check if the token is valid and not expired
	if !parsedToken.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}

	// Extract the "username" claim from the token
	claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("failed to extract JWT claims")
	}

	return claims, nil
}
