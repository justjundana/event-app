package models

type Participant struct {
	ID      int `json:"id" form:"id"`
	EventID int `json:"event_id" form:"event_id"`
	UserID  int `json:"user_id" form:"user_id"`
}
