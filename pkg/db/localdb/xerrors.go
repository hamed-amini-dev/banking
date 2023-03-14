package localdb

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound = status.Error(codes.NotFound, "data not found")
)
