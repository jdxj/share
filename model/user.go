package model

import (
	"errors"
	"fmt"
	"strings"
)

const (
	TNUser = "users"
)

var (
	ErrDuplicateName = errors.New("duplicate name")
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
	ID       int64
	Name     string
	Password string
}

func (u *User) Insert() error {
	query := fmt.Sprintf(`insert into %s (name,password)
values (?,?)`, TNUser)
	res, err := mysql.Exec(query, u.Name, u.Password)
	if err == nil {
		u.ID, err = res.LastInsertId()
		return err
	}

	if strings.Index(err.Error(), "Duplicate entry") >= 0 {
		return ErrDuplicateName
	}
	return err
}
