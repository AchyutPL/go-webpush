package types

type Notification struct {
	Title string `json:"title" bson:"title"`
	Body  string `json:"body" bson:"body"`
}
