package binding

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

const requestTimeout = 30 * time.Second

type App struct {
	mu         sync.RWMutex
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewApp() *App {
	return &App{}
}

// Startup is called by Wails when the app starts. The context is kept
// as the parent of every frontend request context.
func (a *App) Startup(ctx context.Context) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.ctx, a.cancelFunc = context.WithCancel(ctx)
}

func (a *App) Shutdown(ctx context.Context) {
	slog.Info("app shutting down")

	a.mu.RLock()
	cancel := a.cancelFunc
	a.mu.RUnlock()

	if cancel != nil {
		cancel()
	}
}

// requestContext derives a context for a single frontend call. It is
// cancelled on app shutdown and times out after requestTimeout.
func (a *App) requestContext() (context.Context, context.CancelFunc) {
	a.mu.RLock()
	base := a.ctx
	a.mu.RUnlock()

	if base == nil {
		base = context.Background()
	}
	return context.WithTimeout(base, requestTimeout)
}
