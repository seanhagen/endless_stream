package logs

import (
	"log/slog"
	"testing"

	"github.com/seanhagen/endless_stream/internal/observability"
)

// NewTestLog ...
func NewTestLog(t *testing.T, conf *Config) observability.TestLogger {
	t.Helper()

	if conf == nil {
		conf = &Config{
			Out:       &testLogOutput{t: t, logs: []string{}},
			LogType:   LogTypeText,
			Level:     slog.LevelInfo,
			AddSource: true,
		}
	}

	tl, ok := New(conf).(observability.TestLogger)
	if !ok {
		panic("logger isn't a TestLogger?")
	}

	return tl
}

type testLogOutput struct {
	t *testing.T

	logs []string
}

// Write ...
func (tlo *testLogOutput) Write(b []byte) (int, error) {
	tlo.logs = append(tlo.logs, string(b))
	return len(b), nil
}
