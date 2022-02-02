package models

type Comment struct {
	ID      int    `json:"id" form:"id"`
	UserID  int    `json:"user_id" form:"user_id"`
	EventID int    `json:"event_id" form:"event_id"`
	Content string `json:"content" form:"content"`
}
