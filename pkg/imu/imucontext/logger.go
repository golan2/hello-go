package imucontext

import (
	"context"

	"github.com/sirupsen/logrus"
)

func NewEntry(ctx context.Context) (context.Context, *logrus.Entry) {
	logger, ok := ctx.Value(LoggerKey).(*logrus.Entry)
	if logger == nil || !ok {
		logger = logrus.NewEntry(logrus.StandardLogger())
	} else {
		logger = logger.Dup()
	}

	ctx = context.WithValue(ctx, LoggerKey, logger)
	*logger = *logger.WithContext(ctx)
	return ctx, logger
}

func GetLogger(ctx context.Context) *logrus.Entry {
	return ctx.Value(LoggerKey).(*logrus.Entry)
}

func WithField(ctx context.Context, key string, value string) {
	logger := ctx.Value(LoggerKey).(*logrus.Entry)
	*logger = *logger.WithField(key, value)
}

func WithFields(ctx context.Context, fields map[string]any) {
	logger := ctx.Value(LoggerKey).(*logrus.Entry)
	*logger = *logger.WithFields(fields)
}
