	package users

	import "time"

	type Role string

	const (
		RoleCustomer Role = "customer"
		RoleAdmin    Role = "admin"
		RoleSeller   Role = "seller"
	)

	type User struct {
		ID        string    `json:"id" db:"id"`
		Name      string    `json:"name" db:"name"`
		Email     string    `json:"email" db:"email"`
		Password  string    `json:"password" db:"password"`
		Role      Role      `json:"role" db:"role"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	}

	type UserResponse struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Role      Role      `json:"role"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	}
