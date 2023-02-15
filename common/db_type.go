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
	ID          int32
	Name        string
	Category    string
	Address     string
	URL         string `gorm:"url"`
	Ratings     string
	Reviews     string
	Reviewers   string
	ReviewTimes string
	ImageURL    string `gorm:"image_url"`
}
