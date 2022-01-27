// Package logger implements functions to retrieve logger by context
package logger

import (
	"context"
	"os"

	"github.com/raulinoneto/url-shortener/common"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type SegmentCloser interface {
	Close(err error)
}

func Init(ctx context.Context) (context.Context, logrus.FieldLogger) {

	tid, ok := ctx.Value(common.TrackingIDKey).(string)
	if !ok || len(tid) == 0 {
		tid = uuid.NewString()
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	loggerEntry := logger.WithField(common.TrackingIDKey.String(), tid)

	return context.WithValue(ctx, common.LoggerKey, loggerEntry), loggerEntry
}

func FromContext(ctx context.Context) (context.Context, logrus.FieldLogger) {

	logger, ok := ctx.Value(common.LoggerKey).(logrus.FieldLogger)
	if ok {
		return ctx, logger
	}

	ctx, logger = Init(ctx)

	return ctx, logger
}
