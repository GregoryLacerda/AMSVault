package request

type StoryRequestViewModel struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Kind        string `json:"kind"`
	Season      int64  `json:"season,omitempty"`
	Episode     int64  `json:"episode,omitempty"`
	Volume      int64  `json:"volume,omitempty"`
	Chapter     int64  `json:"chapter,omitempty"`
	Status      string `json:"status"`
	User        string `json:"user"`
}