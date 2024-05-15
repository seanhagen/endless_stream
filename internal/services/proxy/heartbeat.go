package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability"
)

// HealthReporter ...
type HealthReporter interface {
	Health(context.Context) error
}

// HealthConfig ...
type HealthConfig struct {
	Reporter          HealthReporter
	MaxFailed         int
	TimeBetweenChecks time.Duration
}

type healthManager struct {
	reporter          HealthReporter
	failedChecks      int
	maxFailedChecks   int
	lastCheck         time.Time
	timeBetweenChecks time.Duration
	logger            observability.Logger
}

const (
	maxFailedMinimum = 1
	minTimeBetween   = time.Millisecond * 200
)

var (
	// ErrRequireHealthReporter ...
	ErrRequireHealthReporter = errors.New("health config requires HealthReporter interface")
	// ErrMaxFailedBelowMinimum ...
	ErrMaxFailedBelowMinimum = fmt.Errorf(
		"health config value 'MaxFailed' below minimum value of %v",
		maxFailedMinimum,
	)
	// ErrTimeBetweenChecksBelowMinimum ...
	ErrTimeBetweenChecksBelowMinimum = fmt.Errorf(
		"health config value 'TimeBetweenChecks' below minimum value of %s",
		minTimeBetween,
	)
)

func newHealthManager(hc HealthConfig, logger observability.Logger) (*healthManager, error) {
	if isNil(hc.Reporter) {
		return nil, ErrRequireHealthReporter
	}

	if hc.MaxFailed < maxFailedMinimum {
		return nil, ErrMaxFailedBelowMinimum
	}

	if hc.TimeBetweenChecks < minTimeBetween {
		return nil, ErrTimeBetweenChecksBelowMinimum
	}

	hm := healthManager{
		reporter:          hc.Reporter,
		maxFailedChecks:   hc.MaxFailed,
		timeBetweenChecks: hc.TimeBetweenChecks,
		logger:            logger,
	}

	return &hm, nil
}

// shouldCheck ...
func (hm *healthManager) shouldCheck(beatTime time.Time) bool {
	if hm.lastCheck.IsZero() {
		return true
	}
	diff := beatTime.Sub(hm.lastCheck)
	return diff > hm.timeBetweenChecks
}

// heartbeat ...
func (hm *healthManager) heartbeat(ctx context.Context, beatTime time.Time) error {
	doCheck := hm.shouldCheck(beatTime)

	var err error
	if doCheck {
		hm.lastCheck = beatTime
		err = hm.report(ctx)
	}

	if err != nil {
		return fmt.Errorf("heartbeat health check failed: %w", err)
	}

	return nil
}

// shutdown ...
func (hm *healthManager) shutdown(ctx context.Context) {
	hm.logger.Debug(ctx, "health check manager shutting down")
}

// report ...
func (hm *healthManager) report(ctx context.Context) error {
	hm.logger.Debug(ctx, "sending health ping")
	err := hm.reporter.Health(ctx)
	if err != nil {
		hm.logger.Error(ctx, "health ping returned error", observability.ErrorAttr(err))
		hm.failedChecks++
	} else {
		hm.failedChecks = 0
	}

	if err != nil && hm.failedChecks >= hm.maxFailedChecks {
		hm.logger.Error(ctx, "too many failed checks")
		return fmt.Errorf("max failed checks, last error: %w", err)
	}

	return nil
}
