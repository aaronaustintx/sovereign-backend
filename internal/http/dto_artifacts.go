package http

type CreateArtifactRequest struct {
	SpaceID string `json:"space_id" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Title   string `json:"title" binding:"required"`
}
