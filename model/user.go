package model

import "fmt"

const (
	TNUser = "users"
)

func GetUser(id int) *User {
	return nil
}

func LoginCheck(name, password string) (*User, error) {
	query := fmt.Sprintf(`select id,name from %s where name=? and password=?`, TNUser)
	row := mysql.QueryRow(query, name, password)

	u := &User{}
	return u, row.Scan(&u.ID, &u.Name)
}

type User struct {
	ID   int
	Name string
}
