package http

type CreateSpaceRequest struct {
    OrgID       string `json:"org_id" binding:"required"`
    Slug        string `json:"slug" binding:"required"`
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
    IsPrivate   bool   `json:"is_private"`
}
