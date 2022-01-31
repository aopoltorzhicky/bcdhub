package helpers

import (
	"time"

	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// InitSentry -
func InitSentry(debug bool, environment, dsn string) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Environment:      environment,
		Debug:            debug,
		AttachStacktrace: true,
		BeforeSend:       beforeSend,
	}); err != nil {
		logger.Info().Msgf("Sentry initialization failed: %v\n", err)
	}
}

func beforeSend(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	logger.Info().Msgf("[Sentry message] %s", event.Message)
	return event
}

// SentryMiddleware -
func SentryMiddleware() gin.HandlerFunc {
	return sentrygin.New(sentrygin.Options{})
}

// SetTagSentry -
func SetTagSentry(key, value string) {
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag(key, value)
	})
}

// SetUserIDSentry -
func SetUserIDSentry(id string) {
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{
			ID: id,
		})
	})
}

// CatchPanicSentry -
func CatchPanicSentry() {
	if err := recover(); err != nil {
		sentry.CurrentHub().WithScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelError)
		})
		sentry.CurrentHub().Recover(err)
		sentry.Flush(time.Second * 5)
	}
}

// CatchErrorSentry -
func CatchErrorSentry(err error) {
	sentry.CaptureEvent(&sentry.Event{
		Message: err.Error(),
		Level:   sentry.LevelError,
		Exception: []sentry.Exception{
			{
				Value:      err.Error(),
				Stacktrace: sentry.ExtractStacktrace(errors.WithStack(err)),
			},
		},
	})
	sentry.Flush(time.Second * 5)
}

// GetLocalSentry -
func GetLocalSentry() *sentry.Hub {
	return sentry.CurrentHub().Clone()
}

// SetLocalTagSentry -
func SetLocalTagSentry(hub *sentry.Hub, key, value string) {
	hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag(key, value)
	})
}

// LocalCatchErrorSentry -
func LocalCatchErrorSentry(hub *sentry.Hub, err error) {
	hub.CaptureException(err)
	hub.Flush(time.Second * 5)
}

// LocalCatchPanicSentry -
func LocalCatchPanicSentry(hub *sentry.Hub) {
	if err := recover(); err != nil {
		hub.CaptureMessage(err.(string))
		hub.Flush(time.Second * 5)
	}
}
