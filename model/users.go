package model

type Users struct {
	ID            int
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
	JwtToken      string  `json:"jwt_token"`
}
