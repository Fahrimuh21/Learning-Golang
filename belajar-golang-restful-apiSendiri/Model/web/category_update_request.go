package web

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `validate:"required, max=200, min=1"`
}
