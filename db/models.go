// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
