package models

import (
	"fmt"
	"gopostgres/db"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
}

type Users []User

const UserSchema string = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
    CONSTRAINT pk_users PRIMARY KEY(id)
);`

func NewUser(first_name, last_name, username, password, email string) *User {
	user := &User{FirstName: first_name, LastName: last_name, Username: username, Password: password, Email: email}
	return user
}

func (user *User) insert() {
	sql := "INSERT INTO users (first_name, last_name, username, password, email) VALUES ($1, $2, $3, $4, $5)"
	fmt.Println(sql)
	result, _ := db.Exec(sql, user.FirstName, user.LastName, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

func CreateUser(first_name, last_name, username, password, email string) *User {
	user := NewUser(first_name, last_name, username, password, email)
	user.Save()
	return user
}

func ListUsers() Users {
	sql := "SELECT id,first_name,last_name,username,password,email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

func GetUser(id int) *User {
	user := NewUser("", "", "", "", "")
	sql := "SELECT id,first_name,last_name,username,password,email FROM users WHERE id = $1"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Email)
	}
	return user
}

func (user *User) update() {
	sql := "UPDATE users SET first_name = $1, last_name = $2, username = $3, password = $4, email = $5 WHERE id = $6"
	db.Exec(sql, user.FirstName, user.LastName, user.Username, user.Password, user.Email, user.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id = $1"
	db.Exec(sql, user.Id)
}
