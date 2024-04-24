package logs

import (
	"bytes"
	"context"
	"testing"

	"github.com/seanhagen/endless_stream/internal/observability"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogs_NewTestLog(t *testing.T) {
	logger := NewTestLog(t, nil)
	assert.Implements(t, (*observability.Logger)(nil), logger)
}

func TestLogs_TestLog_UsesWriter(t *testing.T) {
	ctx := context.TODO()

	t.Run("when provided an io.Writer", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		conf := Config{
			Out: buf,
		}

		tl := NewTestLog(t, &conf)
		tl.Info(ctx, "test log")

		writerTmp := tl.GetOutput()
		require.IsType(t, (*bytes.Buffer)(nil), writerTmp)

		assert.NotEmpty(t, buf.String())
	})

	t.Run("when provided a nil configuration", func(t *testing.T) {
		tl := NewTestLog(t, nil)
		tl.Info(ctx, "test log")

		writerTmp := tl.GetOutput()
		require.IsType(t, (*testLogOutput)(nil), writerTmp)

		writer := writerTmp.(*testLogOutput)
		assert.Contains(t, writer.logs[0], "test log")
	})
}
