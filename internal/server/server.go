package server

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"time"

	"agones.dev/agones/pkg/sdk"
)

const (
	heartbeatTime         = time.Second
	defaultHealthCheckGap = time.Second
)

// GameServer ...
type GameServer = sdk.GameServer

// GameSDK ...
type GameSDK interface {
	Allocate(context.Context) error
	GameServer(context.Context) (*GameServer, error)
	Health(context.Context) error
	Ready(context.Context) error
	Reserve(context.Context, time.Duration) error
	SetAnnotation(string, string) error
	SetLabel(string, string) error
	Shutdown(context.Context) error
}

type heartbeatListener interface {
	heartbeat(context.Context, time.Time) error
	shutdown(context.Context)
}

type tickFn func(time.Duration) <-chan time.Time

// Config ...
type Config struct {
	// GameSDK is the interface with the Agones SDK
	GameSDK GameSDK

	// Health contains the configuration for the health checker that is
	// run by the server.
	Health HealthConfig

	// Logger ...
	Logger *slog.Logger

	buildTicker tickFn
}

// Handler ...
type Handler struct {
	sdk                GameSDK
	heartbeatCh        chan time.Time
	heartbeatListeners []heartbeatListener
	heartbeatCtx       context.Context
	heartbeatCtxCancel func()
	buildTicker        tickFn
	logger             *slog.Logger
}

// Create ...
func Create(ctx context.Context, conf Config) (*Handler, error) {
	if isNil(conf.Logger) {
		conf.Logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	} else {
		conf.Logger = conf.Logger.With("type", "handler")
	}

	if isNil(conf.GameSDK) {
		return nil, fmt.Errorf("missing required field GameSDK in configuration")
	}

	if isNil(conf.buildTicker) {
		conf.buildTicker = time.Tick
	}

	if isNil(conf.Health.Reporter) {
		conf.Health.Reporter = conf.GameSDK
	}

	if isNil(conf.Health.Logger) {
		conf.Health.Logger = conf.Logger.With("type", "heartbeat")
	}

	healthChecker, err := newHealthManager(conf.Health)
	if err != nil {
		return nil, fmt.Errorf("unable to configure health checker: %w", err)
	}

	hdlr := &Handler{
		sdk:         conf.GameSDK,
		buildTicker: conf.buildTicker,
		heartbeatCh: make(chan time.Time),
		heartbeatListeners: []heartbeatListener{
			healthChecker,
		},
		logger: conf.Logger,
	}

	return hdlr, nil
}

// Start ...
func (h *Handler) Start(ctx context.Context) error {
	h.logger.Log(ctx, slog.LevelDebug, "starting server handler")

	ctx, cancel := context.WithCancel(ctx)

	go h.heartbeat(ctx)

	h.heartbeatCtx = ctx
	h.heartbeatCtxCancel = cancel

	if err := h.sdk.Ready(ctx); err != nil {
		cancel()
		h.logger.ErrorContext(ctx, "sdk Ready method returned error", "err", err)
		return fmt.Errorf("unable to ready: %w", err)
	}

	return nil
}

// heartbeat ...
func (h *Handler) heartbeat(ctx context.Context) {
	tick := h.buildTicker(heartbeatTime)
	for {
		select {
		case <-ctx.Done():
			h.logger.DebugContext(ctx, "heartbeat loop, context done")
			return
		case t := <-tick:
			h.notifyHeartbeat(ctx, t)
		}
	}
}

// notifyHeartbeat ...
func (h *Handler) notifyHeartbeat(ctx context.Context, t time.Time) {
	h.logger.DebugContext(ctx, "heartbeat notify")
	for _, listener := range h.heartbeatListeners {
		if err := listener.heartbeat(ctx, t); err != nil {
			h.logger.ErrorContext(ctx, "heartbeat listener returned erorr", "err", err)
		}
	}
}

// Stop ...
func (h *Handler) Stop(ctx context.Context) error {
	h.logger.DebugContext(ctx, "handler stop")
	h.heartbeatCtxCancel()

	for _, listener := range h.heartbeatListeners {
		listener.shutdown(ctx)
	}

	return nil
}

func isNil(i any) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
