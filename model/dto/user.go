package dto

type User struct {
	ID        int64  `json:"id"`
	Token     string `json:"token"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      int32  `json:"role"`
}
