package http

type AIRequest struct {
    Model  string `json:"model" binding:"required"`
    Prompt string `json:"prompt" binding:"required"`
}
