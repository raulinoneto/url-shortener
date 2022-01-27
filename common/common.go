// Package common is used in all app
package common

type Key string

func (k Key) String() string {
	return string(k)
}

const (
	TrackingIDKey Key = "x-tracking-id"
	ObjIDKey      Key = "objID"
	LoggerKey     Key = "logger"
	ErrorKey      Key = "error"
)

const (
	HealthCheckPath = "/healthcheck"
	AppDefaultPath  = "/v1/url-handlers"
)
