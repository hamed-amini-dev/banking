package xerrors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrServiceUnavailableMessage       = status.Error(codes.Unavailable, "service unavailable message")
	ErrInvalidSession                  = status.Error(codes.Unauthenticated, "invalid session")
	ErrUnauthorizedRequest             = status.Error(codes.Unauthenticated, "unauthorized request")
	ErrorUnImplemented                 = status.Error(codes.Unimplemented, "unimplemented")
	ErrNotFound                        = status.Error(codes.NotFound, "data not found")
	ErrUserNotFound                    = status.Error(codes.NotFound, "user not found")
	ErrUserInviteWithRepeatedRole      = status.Error(codes.AlreadyExists, "this user is already exist")
	ErrUserInviteUserHasOrgan          = status.Error(codes.FailedPrecondition, "this user is already has organization")
	ErrorExceededMaximumAllowedRecords = status.Error(codes.OutOfRange, "exceeded maximum number of allowed records")

	ErrFileIsADirectory               = status.Error(codes.InvalidArgument, "this file is a directory")
	ErrDirectoryNotFound              = status.Error(codes.NotFound, "this directory not found")
	ErrorNilArgs                      = status.Error(codes.InvalidArgument, "nil args")
	ErrParentFolderNotFound           = status.Error(codes.NotFound, "parent folder not found")
	ErrorInvalidXid                   = status.Error(codes.InvalidArgument, "invalid xid")
	ErrMethodNotSupported             = status.Error(codes.FailedPrecondition, "method not supported")
	ErrProblemTokenInvalid            = status.Error(codes.Unauthenticated, "problem token invalid")
	ErrCanNotDeleteYourself           = status.Error(codes.PermissionDenied, "you can't delete your account")
	ErrPermissionWasGrantedBefore     = status.Error(codes.AlreadyExists, "permission to this manager was granted before")
	ErrDeleteNonEmptyOrganization     = status.Error(codes.FailedPrecondition, "the organization is not empty")
	ErrNameLengthMustBeMoreThan2Chars = status.Error(codes.InvalidArgument, "name must be more than 2 characters")
	ErrManagerOfNoOrganization        = status.Error(codes.PermissionDenied, "you are a manager of no organization")
	ErrOrganizationNameEmpty          = status.Error(codes.InvalidArgument, "organization name is empty")
	ErrPaymentOrderNotFound           = status.Error(codes.NotFound, "order payment not found")
	ErrUserDoesNotHaveAPaymentAccount = status.Error(codes.FailedPrecondition, "the user doesn't have a payment account")
	ErrUserDidNotSaveTheirCard        = status.Error(codes.FailedPrecondition, "the user didn't have been saved their card")
	ErrInvalidToken                   = status.Error(codes.Unauthenticated, "invalid token")
	ErrCurrentSubscriptionIsOver      = status.Error(codes.FailedPrecondition, "current subscription is over")
	ErrAlreadyHasAnOrganization       = status.Error(codes.AlreadyExists, "already you have an organization")
	ErrInvalidModuleInitiation        = status.Error(codes.FailedPrecondition, "invalid module initiation")
	ErrPersonHasThisProfile           = status.Error(codes.AlreadyExists, "already you have this profile")
	ErrInvalidSharedStructureLink     = status.Error(codes.PermissionDenied, "invalid shared structure link")
	ErrManagerHasNoOrganization       = status.Error(codes.NotFound, "manager don't have organization")
)

var (
	ErrInvalidAccess error = status.Error(codes.PermissionDenied, "invalid role access")
	New                    = errors.New
	ErrNilArgsWith         = func(args ...string) error {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Nil args for %+v", args))
	}
	ErrNilParamsInDB = func(args ...string) error {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("DB: Nil params for %+v", args))
	}
)

var (
	ErrUnableToConnectGRPCService = func(addr string, err error) error {
		return status.Errorf(codes.Unavailable, "unable to connect gRPC server (%s): %v", addr, err)
	}
)
