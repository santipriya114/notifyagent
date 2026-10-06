package grpc

import (
	"context"
	"os"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestValidateToken(t *testing.T) {

	os.Setenv("GRPC_TOKEN", "test-token")

	md := metadata.New(map[string]string{
		"authorization": "Bearer test-token",
	})

	ctx := metadata.NewIncomingContext(
		context.Background(),
		md,
	)

	err := validateToken(ctx)

	if err != nil {
		t.Fatal(err)
	}
}