package http

import (
	"database/sql"

	"github.com/aaronaustintx/sovereign-backend/internal/config"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()

	h := NewHandlers(db, cfg)

	r.GET("/health", h.Health)

	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)

	auth := r.Group("/")
	auth.Use(h.AuthMiddleware())
	auth.GET("/me", h.Me)
	auth.POST("/orgs", h.CreateOrg) // you can protect this if you want

	return r
}
