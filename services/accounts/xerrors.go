package accounts

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrEnoughBalance = status.Error(codes.FailedPrecondition, "doesn't have enough balance to transfer")
)
