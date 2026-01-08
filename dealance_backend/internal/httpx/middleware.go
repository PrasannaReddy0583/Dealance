package httpx

import "net/http"

type Middleware func(http.Handler) http.Handler

func applyMiddleware(h http.Handler) http.Handler {
	middlewares := []Middleware{
		recoverMiddleware,
		// authMiddleware (later)
		// rateLimitMiddleware (later)
		// tracingMiddleware (later)
		// jwt_middleware.go (later needs to wiredin)
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
