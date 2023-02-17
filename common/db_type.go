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
