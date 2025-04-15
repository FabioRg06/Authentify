package domain

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
}
