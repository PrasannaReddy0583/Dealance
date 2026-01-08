package httpx

type contextKey string

const (
	ContextUserID   contextKey = "user_id"
	ContextUserRole contextKey = "user_role"
	ContextKYCLevel contextKey = "kyc_level"
)
