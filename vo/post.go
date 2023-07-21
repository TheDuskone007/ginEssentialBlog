package vo

type CreatePostRequest struct {
	CategoryId uint   `json:"categoryId" gorm:"not null"`
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
}
