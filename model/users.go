package model

type Users struct {
	ID            int
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
	Status        string  `json:"status"`
	JwtToken      string  `json:"jwt_token"`
}

type UserVerifications struct {
	ID         int
	UserID     int `json:"user_id"`
	User       Users
	VerifyCode string `json:"verify_code"`
}