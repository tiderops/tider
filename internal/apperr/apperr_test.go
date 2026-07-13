package apperr

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestClassify(t *testing.T) {
	gr := schema.GroupResource{Group: "", Resource: "pods"}

	tests := []struct {
		name string
		err  error
		want Kind
	}{
		{name: "k8s not found", err: apierrors.NewNotFound(gr, "web-1"), want: KindNotFound},
		{name: "k8s forbidden", err: apierrors.NewForbidden(gr, "web-1", errors.New("rbac")), want: KindForbidden},
		{name: "k8s unauthorized", err: apierrors.NewUnauthorized("token expired"), want: KindForbidden},
		{name: "k8s conflict", err: apierrors.NewConflict(gr, "web-1", errors.New("stale")), want: KindConflict},
		{name: "k8s invalid", err: apierrors.NewInvalid(schema.GroupKind{Kind: "Pod"}, "web-1", nil), want: KindValidation},
		{name: "k8s timeout", err: apierrors.NewTimeoutError("too slow", 1), want: KindTimeout},
		{name: "context deadline", err: context.DeadlineExceeded, want: KindTimeout},
		{name: "network error", err: &net.OpError{Op: "dial", Err: errors.New("connection refused")}, want: KindUnreachable},
		{name: "wrapped not found", err: fmt.Errorf("getting pod: %w", apierrors.NewNotFound(gr, "web-1")), want: KindNotFound},
		{name: "plain error", err: errors.New("boom"), want: KindInternal},
		{name: "already classified", err: New(KindUnreachable, "not registered"), want: KindUnreachable},
		{name: "wrapped classified", err: fmt.Errorf("ctx: %w", New(KindNotFound, "gone")), want: KindNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Classify(tt.err); got != tt.want {
				t.Errorf("Classify(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	if Normalize(nil) != nil {
		t.Fatal("Normalize(nil) should be nil")
	}

	err := Normalize(apierrors.NewNotFound(schema.GroupResource{Resource: "pods"}, "web-1"))
	if got, want := err.Error()[:len(KindNotFound)], string(KindNotFound); got != want {
		t.Errorf("Normalize() message prefix = %q, want %q", got, want)
	}

	classified := New(KindTimeout, "slow")
	if Normalize(classified) != classified {
		t.Error("Normalize() should pass through an already classified error")
	}
}
