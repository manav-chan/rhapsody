package models

type Blog struct {
	Id    uint   `json:"id"`
	Title string `json:title`
	Desc  string `json:desc`
	Image string `json:image`
	UserID string `json:userid`

	// Foreign key associated with user
	User User `json:"user";gorm:"foreignKey:UserId`

}