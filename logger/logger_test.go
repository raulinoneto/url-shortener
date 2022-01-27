package logger

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"tools/common"
)

func Test(t *testing.T) {
	t.Run("should create a new logger", func(t *testing.T) {
		ctx := context.Background()
		ctx, logger, _ := FromContext(ctx, "")
		assert.NotNil(t, logger)
		assert.NotNil(t, ctx)
		assert.Equal(t, logger, ctx.Value(common.LoggerKey))
	})
	t.Run("should return existing logger", func(t *testing.T) {
		loggerExpected := logrus.New()
		ctx := context.WithValue(context.Background(), common.LoggerKey, loggerExpected)
		ctx, logger, _ := FromContext(ctx, "")
		assert.NotNil(t, logger)
		assert.NotNil(t, ctx)
		assert.Equal(t, loggerExpected, ctx.Value(common.LoggerKey))
		assert.Equal(t, loggerExpected, logger)
	})
}
