package ai_outputs

import "time"

type AIOutput struct {
	ID         string    `json:"id"`
	ArtifactID string    `json:"artifact_id"`
	Model      string    `json:"model"`
	Prompt     string    `json:"prompt"`
	Output     string    `json:"output"`
	CreatedAt  time.Time `json:"created_at"`
}
