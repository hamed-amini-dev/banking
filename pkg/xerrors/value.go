package xerrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/****
 * Constant user error for global error handling
 *
 *
 */

var (
	ErrServiceUnavailableMessage = status.Error(codes.Unavailable, "service unavailable message")
)
