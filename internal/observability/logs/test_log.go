package logs

import (
	"io"
	"log/slog"
	"testing"

	"github.com/seanhagen/endless_stream/internal/observability"
)

// NewTestLog ...
func NewTestLog(t *testing.T, conf *Config) observability.TestLogger {
	t.Helper()

	if conf == nil {
		conf = &Config{
			Out:       NewTestLogOutput(t, false),
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

func NewTestLogOutput(t *testing.T, outputTestLogs bool) io.Writer {
	t.Helper()

	return &testLogOutput{t: t, outputTestLogs: outputTestLogs, logs: []string{}}
}

type testLogOutput struct {
	t *testing.T

	outputTestLogs bool
	logs           []string
}

// Write ...
func (tlo *testLogOutput) Write(data []byte) (int, error) {
	tlo.t.Helper()
	tlo.logs = append(tlo.logs, string(data))
	if tlo.outputTestLogs {
		tlo.t.Log(string(data))
	}
	return len(data), nil
}
