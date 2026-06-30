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
	auth.POST("/artifacts/:id/ai", h.GenerateAI)
	auth.POST("/artifacts", h.CreateArtifact)
	auth.GET("/spaces/:space_id/artifacts", h.ListArtifacts)
	auth.GET("/spaces/:space_id/feed", h.Feed)

	return r
}
