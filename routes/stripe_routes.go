package routes

import (
	"net/http"

	"backend/handlers"
	middleware "backend/middlewares"
	"github.com/redis/go-redis/v9"
)

func StripeRoutes(mux *http.ServeMux, s *handlers.Stripe, redis *redis.Client) {
	authMw := &middleware.RedisStruct{
		RedisClient: redis,
	}
	mux.Handle("POST /api/file-store/stripe/create-session", authMw.AuthMiddleware(http.HandlerFunc(s.CreateCheckoutSession)))
	// mux.Handle("POST /api/file-store/stripe/verify-session", authMw.AuthMiddleware(http.HandlerFunc(s.VerifyCheckoutSession)))
	mux.Handle("POST /api/file-store/stripe/cancel-subscription", authMw.AuthMiddleware(http.HandlerFunc(s.CancelSubscription)))
}
