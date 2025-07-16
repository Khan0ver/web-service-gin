package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type userAuthDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userRegDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authRespond struct {
	Token string `json:"token"`
}

var IdGlobal = 0 //Delete -> DB autoInc

var users = []UserInfo{
	{ID: strconv.Itoa(IdGlobal), Name: "Qq Name", Email: "qqqq@qq.qq", Password: "qqqqwwww"},
}

// FOR DEBUG POSTMAN
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)

}

func performLogin(c *gin.Context) {
	var newUser userAuthDTO

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	var user *UserInfo //nil

	for i := 0; i < len(users); i++ {
		if users[i].Email == newUser.Email && users[i].Password == newUser.Password {
			user = &users[i]
		}
	}

	if user != nil {
		c.IndentedJSON(http.StatusOK, authRespond{Token: "Temp Placeholder"}) //TODO JWT TOKEN
	} else {
		c.IndentedJSON(http.StatusBadRequest, authRespond{Token: ""})
	}
}

func createUser(c *gin.Context) {
	var newUser userRegDTO

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	var hasUser = false

	for i := 0; i < len(users); i++ {
		if users[i].Email == newUser.Email {
			hasUser = true
		}
	}

	if hasUser {
		IdGlobal++
		users = append(users, UserInfo{
			ID:       strconv.Itoa(IdGlobal),
			Name:     newUser.Name,
			Email:    newUser.Email,
			Password: newUser.Password, //ADD BCRYPT
		})
		c.IndentedJSON(http.StatusCreated, authRespond{Token: "Successfuly Reg"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, authRespond{Token: ""})
	}
}

// func getUserById(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range users {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
// }

func main() {
	router := gin.Default()
	router.POST("/auth/login", performLogin)
	router.POST("/auth/reg", createUser)

	router.Run("192.168.0.104:8080")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
