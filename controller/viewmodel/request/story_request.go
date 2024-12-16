package request

type StoryRequestViewModel struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Season      int64  `json:"season,omitempty"`
	Episode     int64  `json:"episode,omitempty"`
	Volume      int64  `json:"volume,omitempty"`
	Chapter     int64  `json:"chapter,omitempty"`
	Status      string `json:"status"`
}
