package task

type CreateDto struct {
	Method  string            `json:"method" binding:"required"`
	URL     string            `json:"url" binding:"required,url"`
	Headers map[string]string `json:"headers"`
}
