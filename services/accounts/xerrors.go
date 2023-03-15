package accounts

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
	ErrEnoughBalance = status.Error(codes.FailedPrecondition, "doesn't have enough balance to transfer")
)
