package localdb

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/****
 * Constant user error for local error handling
 *
 *
 */

var (
	ErrNotFound = status.Error(codes.NotFound, "data not found")
)
