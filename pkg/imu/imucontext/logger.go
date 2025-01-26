package imucontext

import (
	"context"

	"github.com/sirupsen/logrus"
)

//goland:noinspection GoUnusedExportedFunction
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

//goland:noinspection GoUnusedExportedFunction
func GetLogger(ctx context.Context) *logrus.Entry {
	return ctx.Value(LoggerKey).(*logrus.Entry)
}

//goland:noinspection GoUnusedExportedFunction
func WithField(ctx context.Context, key string, value string) {
	logger := ctx.Value(LoggerKey).(*logrus.Entry)
	*logger = *logger.WithField(key, value)
}

//goland:noinspection GoUnusedExportedFunction
func WithFields(ctx context.Context, fields map[string]any) {
	logger := ctx.Value(LoggerKey).(*logrus.Entry)
	*logger = *logger.WithFields(fields)
}
