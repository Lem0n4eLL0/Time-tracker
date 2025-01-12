package structures

import "fmt"

type User struct {
	UserID   int    `db:"user_id" json:"id"`
	Name     string `db:"username" json:"username"`
	Email    any    `db:"email" json:"email"`
	Password string `db:"password_hash" json:"password"`
	Created  any    `db:"created_at" json:"created_at"`
	Updated  any    `db:"updated_at" json:"updated_at"`
	Role     string `db:"name" json:"role_id"`
}

func (u *User) DisplayUser() {
	fmt.Println("UserID: ", u.UserID,
		"\nName: ", u.Name,
		"\nEmail: ", u.Email,
		"\nPassword: ", u.Password,
		"\nCreated: ", u.Created,
		"\nUpdated: ", u.Updated,
		"\nRole: ", u.Role)
}

func (u *User) ResetNonAdminFild() {
	u.Email = ""
	u.Password = ""
	u.Created = ""
	u.Updated = ""
}
