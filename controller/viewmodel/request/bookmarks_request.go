package request

type BookmarksRequestViewModel struct {
	ID     string                `json:"id"`
	UserID int64                 `json:"user_id"`
	Story  StoryRequestViewModel `json:"story"`
}
