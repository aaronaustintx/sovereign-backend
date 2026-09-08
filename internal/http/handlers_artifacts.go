package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateArtifact(c *gin.Context) {
	var req CreateArtifactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID := c.GetString("user_id")

	artifact, err := h.artifactService.CreateArtifact(
		req.SpaceID,
		userID,
		req.Type,
		req.Title,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log event
	_, _ = h.eventService.LogEvent(
		req.SpaceID,
		userID,
		artifact.ID,
		"artifact_created",
		`{"type":"`+req.Type+`"}`,
	)

	c.JSON(http.StatusCreated, artifact)
}

func (h *Handlers) ListArtifacts(c *gin.Context) {
	spaceID := c.Param("space_id")

	artifacts, err := h.artifactService.ListArtifacts(spaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, artifacts)
}

func (h *Handlers) Feed(c *gin.Context) {
	spaceID := c.Param("space_id")

	feed, err := h.artifactService.Feed(spaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feed)
}
