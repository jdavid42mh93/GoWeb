package main

import (
	"fmt"
	"gopostgres/db"
	"gopostgres/models"
)

func main() {
	db.Connect()
	db.Ping()
	//fmt.Println(db.ExistsTable("users"))
	//user := models.CreateUser("David", "Suarez", "csjdavid42mh93", "@Admin2021", "csjdavid42mh93@gmail.com")
	//fmt.Println(user)
	//db.CreateTable(models.UserSchema, "users")
	//users := models.ListUsers()
	user := models.GetUser(2)
	/*user.FirstName = "Segundo"
	user.LastName = "Callataxi"
	user.Username = "scallataxi"
	user.Password = "@Admin2022"
	user.Email = "scallataxi@gmail.com"
	user.Save()*/
	user.Delete()
	fmt.Println(user)
	//fmt.Println(users)
	db.Close()
}
