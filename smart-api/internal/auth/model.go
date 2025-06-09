package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int
	Username string
	Password string
}

type Charge struct {
	ID int `json:"id"`
	UserId int `json:"user_id"`
	Amount float64 `json:"amount"`
	Date string `json:"date"`
	Category string `json:"category"`
}
