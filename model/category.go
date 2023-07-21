package model

type Category struct {
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	//格式化时间格式，修改time.Time中的方法
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"`
}
