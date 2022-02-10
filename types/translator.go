package types

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Identifier string `json:"identifier"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
}

type FeedUser struct {
	Username   string `json:"username"`
	Identifier string `json:"identifier"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}
