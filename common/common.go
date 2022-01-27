// Package common is used in all app
package common

type Key string

func (k Key) String() string {
	return string(k)
}

const (
	TrackingIDKey Key = "x-tracking-id"
	TenantIDKey   Key = "x-tenant-id"
	ObjIDKey      Key = "objID"
	LoggerKey     Key = "logger"
	ErrorKey      Key = "error"
)

const (
	HealthCheckPath = "/healthcheck"
	AppDefaultPath  = "/v1/app_name"
)
