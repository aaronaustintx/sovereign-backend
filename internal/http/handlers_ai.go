package http

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func (h *Handlers) GenerateAI(c *gin.Context) {
    artifactID := c.Param("id")
    userID := c.GetString("user_id")

    var req AIRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    // 1. Load artifact
    artifact, err := h.artifactService.GetArtifact(artifactID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "artifact not found"})
        return
    }

    // 2. Call AI model
    output, err := h.aiClient.Generate(req.Model, req.Prompt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "AI generation failed"})
        return
    }

    // 3. Store AI output
    aiOut, err := h.aiOutputService.CreateOutput(
        artifact.ID,
        req.Model,
        req.Prompt,
        output,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save AI output"})
        return
    }

    // 4. Log event
    _, _ = h.eventService.LogEvent(
        artifact.SpaceID,
        userID,
        artifact.ID,
        "ai_generated",
        `{"model":"`+req.Model+`"}`,
    )

    // 5. Return result
    c.JSON(http.StatusOK, gin.H{
        "artifact_id": artifact.ID,
        "model":       req.Model,
        "output":      output,
        "record":      aiOut,
    })
}
