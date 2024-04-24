package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer_Constructor_RequiredOptions(t *testing.T) {
	tests := []struct {
		name  string
		conf  Config
		valid bool
	}{
		{
			name: "empty config is not valid",
		},
		{
			name: "config with only GameSDK set is invalid",
			conf: Config{GameSDK: &testAgonesSDK{}},
		},
		{
			name: "config with only GameSDK & Health set is valid",
			conf: Config{
				GameSDK: &testAgonesSDK{}, //nolint:exhaustruct
				Health: HealthConfig{
					Reporter:          nil,
					MaxFailed:         3,
					TimeBetweenChecks: time.Second,
				},
				Logger: logs.NewTestLog(t, &logs.Config{}),
			},
			valid: true,
		},
	}

	ctx := context.Background()

	for i, x := range tests {
		tt := x
		t.Run(
			fmt.Sprintf("test %d %s", i+1, tt.name),
			func(t *testing.T) {
				t.Parallel()
				handler, err := Create(ctx, tt.conf)

				if tt.valid {
					assert.NotNil(t, handler)
					assert.NoError(t, err)
					return
				}

				assert.Nil(t, handler)
				assert.Error(t, err)
			},
		)
	}
}

func TestServer_StartsAndLifecycleProgression(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeTicks := make(chan time.Time, 10)
	expectDuration := time.Second
	testSDK := &testAgonesSDK{} //nolint:exhaustruct

	buf := bytes.NewBuffer(nil)
	logConf := &logs.Config{
		Out:     buf,
		LogType: logs.LogTypeText,
		Level:   slog.LevelDebug,
	}
	logger := logs.NewTestLog(t, logConf)

	var gotDuration time.Duration
	tickerFn := func(d time.Duration) <-chan time.Time { //nolint:wsl
		gotDuration = d

		return fakeTicks
	}

	conf := Config{
		GameSDK: testSDK,
		Health: HealthConfig{
			Reporter:          nil,
			MaxFailed:         2,
			TimeBetweenChecks: time.Second,
		},
		Logger:      logger,
		buildTicker: tickerFn,
	}

	hdlr, err := Create(ctx, conf)
	require.NoError(t, err)
	require.NotNil(t, hdlr)

	err = hdlr.Start(ctx)
	assert.NoError(t, err, "unable to call Start on server")

	fakeTicks <- time.Date(2024, 1, 1, 0, 30, 0, 500, time.UTC)

	time.Sleep(time.Second * 2)

	err = hdlr.Stop(ctx)
	assert.NoError(t, err, "unable to call Stop on server")

	time.Sleep(time.Second)

	assert.Equal(t, 1, testSDK.readyCalls, "expected at least one call to SDK Ready method")
	assert.Equal(t, 1, testSDK.healthCalls, "expected at least one call to SDK Health method")
	assert.Equal(
		t,
		expectDuration,
		gotDuration,
		"expected ticker to get created with specific duration",
	)

	assert.NotEmpty(t, buf.String())
	assert.Contains(t, buf.String(), "heartbeat notify")
	assert.Contains(t, buf.String(), "health check manager shutting down")
	assert.Contains(t, buf.String(), "heartbeat loop")

	t.Logf("logger output: \n%s", buf.String())
}

var errNotImplementedYet = errors.New("not implemented yet")

type testAgonesSDK struct {
	allocateCalls int
	readyCalls    int
	healthCalls   int
}

// Allocate ...
func (sdk *testAgonesSDK) Allocate(_ context.Context) error {
	sdk.allocateCalls++

	return nil
}

// GameServer ...
func (sdk *testAgonesSDK) GameServer(_ context.Context) (*GameServer, error) {
	return nil, errNotImplementedYet
}

// Health ...
func (sdk *testAgonesSDK) Health(_ context.Context) error {
	sdk.healthCalls++

	return nil
}

// Ready ...
func (sdk *testAgonesSDK) Ready(_ context.Context) error {
	sdk.readyCalls++

	return nil
}

// Reserve ...
func (sdk *testAgonesSDK) Reserve(_ context.Context, _ time.Duration) error {
	return nil
}

// SetAnnotation ...
func (sdk *testAgonesSDK) SetAnnotation(_, _ string) error {
	return nil
}

// SetLabel ...
func (sdk *testAgonesSDK) SetLabel(_, _ string) error {
	return nil
}

// Shutdown ...
func (sdk *testAgonesSDK) Shutdown(_ context.Context) error {
	return nil
}
