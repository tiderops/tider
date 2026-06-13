package apperr

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

type Kind string

const (
	KindNotFound    Kind = "NOT_FOUND"
	KindForbidden   Kind = "FORBIDDEN"
	KindConflict    Kind = "CONFLICT"
	KindValidation  Kind = "VALIDATION"
	KindTimeout     Kind = "TIMEOUT"
	KindUnreachable Kind = "CLUSTER_UNREACHABLE"
	KindInternal    Kind = "INTERNAL"
)

type Error struct {
	Kind Kind
	Msg  string
	Err  error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Kind, e.Msg)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func New(kind Kind, msg string) *Error {
	return &Error{Kind: kind, Msg: msg}
}

func Normalize(err error) error {
	if err == nil {
		return nil
	}

	var ae *Error
	if errors.As(err, &ae) {
		return ae
	}

	return &Error{Kind: Classify(err), Msg: err.Error(), Err: err}
}

func Classify(err error) Kind {
	var ae *Error
	if errors.As(err, &ae) {
		return ae.Kind
	}

	switch {
	case apierrors.IsNotFound(err):
		return KindNotFound
	case apierrors.IsForbidden(err) || apierrors.IsUnauthorized(err):
		return KindForbidden
	case apierrors.IsConflict(err) || apierrors.IsAlreadyExists(err):
		return KindConflict
	case apierrors.IsInvalid(err) || apierrors.IsBadRequest(err):
		return KindValidation
	case apierrors.IsTimeout(err) || apierrors.IsServerTimeout(err),
		errors.Is(err, context.DeadlineExceeded):
		return KindTimeout
	case isConnectionError(err):
		return KindUnreachable
	default:
		return KindInternal
	}
}

func isConnectionError(err error) bool {
	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}

	var urlErr *url.Error
	return errors.As(err, &urlErr)
}
