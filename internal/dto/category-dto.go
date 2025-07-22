package dto

type CategoryInput struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type UpdateCategoryInput struct {
	Name *string `json:"name"`
	Icon *string `json:"icon"`
}
