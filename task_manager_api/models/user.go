package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" required:"true" unique:"true"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}
