package common

type User struct {
	ID       string
	Name     string
	Password string
	Email    string
	Mobile   string
	Age      int
	Gender   int
}

type Diner struct {
	ID          int32  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Category    string `json:"category,omitempty"`
	Address     string `json:"address,omitempty"`
	URL         string `gorm:"url" json:"url,omitempty"`
	Ratings     string `json:"ratings,omitempty"`
	Reviews     string `json:"reviews,omitempty"`
	Reviewers   string `json:"reviewers,omitempty"`
	ReviewTimes string `json:"review_times,omitempty"`
	ImageURL    string `gorm:"image_url" json:"image_url,omitempty"`
}

type Comment struct {
	ID      int32  `json:"id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
	DinerID int32  `json:"diner_id,omitempty"`
	Msg     string `json:"msg"`
}

type Dialog struct {
	ID       int32  `json:"id,omitempty"`
	FatherID int32  `json:"father_id,omitempty" gorm:"father_id"`
	UserID   string `json:"user_id,omitempty" gorm:"user_id"`
	DinerID  int32  `json:"diner_id,omitempty" gorm:"diner_id"`
	Msg      string `json:"msg,omitempty"`
}

type Ad struct {
	ID      int32 `json:"id"`
	DinerID int32 `json:"diner_id" gorm:"diner_id"`
	Vip     int32 `json:"vip"`
	EndDate int32 `json:"end_date" gorm:"end_date"`
}
