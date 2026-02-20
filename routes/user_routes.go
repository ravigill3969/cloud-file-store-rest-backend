package routes

import (
	"net/http"

	"backend/handlers"
	middleware "backend/middlewares"
	"github.com/redis/go-redis/v9"
)

func RegisterUserRoutes(mux *http.ServeMux, uh *handlers.UserHandler, redis *redis.Client) {
	authMw := &middleware.RedisStruct{
		RedisClient: redis,
	}

	mux.Handle("GET /api/file-store/users/get-user", authMw.AuthMiddleware((http.HandlerFunc(uh.GetUserInfo))))
	mux.Handle("POST /api/file-store/users/get-secret-key", authMw.AuthMiddleware((http.HandlerFunc(uh.GetSecretKey))))
	mux.Handle("PATCH /api/file-store/users/update-secret-key", authMw.AuthMiddleware((http.HandlerFunc(uh.UpdateSecretKey))))
	mux.Handle("PUT /api/file-store/users/update-password", authMw.AuthMiddleware((http.HandlerFunc(uh.UpdatePassword))))
	mux.Handle("PUT /api/file-store/users/update-user-info", authMw.AuthMiddleware((http.HandlerFunc(uh.UpdateUserInfo))))

	mux.HandleFunc("GET /api/file-store/users/logout", uh.Logout)
	mux.HandleFunc("POST /api/file-store/users/register", uh.Register)
	mux.HandleFunc("POST /api/file-store/users/login", uh.Login)
	mux.HandleFunc("GET /api/file-store/users/refresh-token", uh.RefreshTokenVerify)
}
