package imucontext

type ContextKey string

const (
	TenantID       = ContextKey("tenant_id")
	UserID         = ContextKey("user_id")
	CorrelationID  = ContextKey("correlation_id")
	UserPermission = ContextKey("user_permission")
	LoggerKey      = ContextKey("logger")
)
