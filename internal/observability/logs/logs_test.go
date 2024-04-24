package logs

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/seanhagen/endless_stream/internal/observability"
	"github.com/stretchr/testify/assert"
)

func TestLogs_Logs(t *testing.T) {
	logger := New(nil)
	assert.Implements(t, (*observability.Logger)(nil), logger)
}

func TestLogs_OutputToBuffer(t *testing.T) {
	ctx := context.TODO()

	t.Run("debug", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		conf := Config{
			Out:     buf,
			LogType: LogTypeText,
			Level:   slog.LevelDebug,
		}

		expectStr := `test debug message`

		log := New(&conf)
		log.Debug(ctx, expectStr)

		assert.Contains(t, buf.String(), expectStr)
		assert.Contains(t, buf.String(), "level=DEBUG")
	})

	t.Run("info", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		conf := Config{
			Out:     buf,
			LogType: LogTypeText,
			Level:   slog.LevelDebug,
		}

		expectStr := "test info message"

		log := New(&conf)
		log.Info(ctx, expectStr)

		assert.Contains(t, buf.String(), expectStr)
		assert.Contains(t, buf.String(), "level=INFO")
	})

	t.Run("error", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		conf := Config{
			Out:     buf,
			LogType: LogTypeText,
			Level:   slog.LevelDebug,
		}

		expectStr := "test error message"

		log := New(&conf)
		log.Error(ctx, expectStr)

		assert.Contains(t, buf.String(), expectStr)
		assert.Contains(t, buf.String(), "level=ERROR")
	})

	t.Run("warn", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		conf := Config{
			Out:     buf,
			LogType: LogTypeText,
			Level:   slog.LevelDebug,
		}

		expectStr := "test warn message"

		log := New(&conf)
		log.Warn(ctx, expectStr)

		assert.Contains(t, buf.String(), expectStr)
		assert.Contains(t, buf.String(), "level=WARN")
	})
}
