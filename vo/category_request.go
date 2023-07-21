package vo

// CreateCategoryRequest 定义一个验证数据validData
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
