package http

import (
	"database/sql"
	"net/http"

	"github.com/aaronaustintx/sovereign-backend/internal/ai"
	"github.com/aaronaustintx/sovereign-backend/internal/ai_outputs"
	"github.com/aaronaustintx/sovereign-backend/internal/artifacts"
	"github.com/aaronaustintx/sovereign-backend/internal/config"
	"github.com/aaronaustintx/sovereign-backend/internal/events"
	"github.com/aaronaustintx/sovereign-backend/internal/orgs"
	"github.com/aaronaustintx/sovereign-backend/internal/spaces"
	"github.com/aaronaustintx/sovereign-backend/internal/users"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	orgService      *orgs.Service
	userService     *users.Service
	spaceService    *spaces.Service
	artifactService *artifacts.Service
	aiOutputService *ai_outputs.Service
	jwtManager      *JWTManager
	aiClient        *ai.Client
	eventService    *events.Service
	userRepo        *users.Repository
	db              *sql.DB
	cfg             config.Config
}

func NewHandlers(db *sql.DB, cfg config.Config) *Handlers {
	return &Handlers{
		orgService:      orgs.NewService(db),
		userService:     users.NewService(db),
		spaceService:    spaces.NewService(db),
		artifactService: artifacts.NewService(db),
		aiOutputService: ai_outputs.NewService(db),
		eventService:    events.NewService(db),
		jwtManager:      NewJWTManager(cfg),
		aiClient:        ai.New(cfg.OpenAIKey),
		db:              db,
		cfg:             cfg,
		userRepo:        users.NewRepository(db),
	}
}

func (h *Handlers) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handlers) CreateOrg(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	org, err := h.orgService.CreateOrg(req.Name, req.Slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, org)
}
func (h *Handlers) Signup(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.userService.CreateUser(req.Email, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// func (h *Handlers) Login(c *gin.Context) {
// 	var req struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
// 		return
// 	}

// 	user, err := h.userService.Authenticate(req.Email, req.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
// 		return
// 	}

// 	token, err := h.jwtManager.GenerateToken(user.ID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"token": token,
// 		"user":  user,
// 	})
// }
