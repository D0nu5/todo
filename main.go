package main

import (
	_ "encoding/json"
	_ "fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/d0nu5/todo/auth"
	"github.com/d0nu5/todo/todo"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect databases")
	}

	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()

	r.GET("/tokenz", auth.AccessToken("MySignature"))

	protected := r.Group("", auth.Protect([]byte("MySignature")))

	handler := todo.NewTodoHandler(db)
	protected.GET("/todo", handler.GetFirstTask)
	protected.GET("/todos", handler.GetAllTask)
	protected.POST("/todos", handler.NewTask)

	r.Run()
}

// type UserHandler struct {
// 	db *gorm.DB
// }

// func (h *UserHandler) User(c *gin.Context) {
// 	var u User
// 	h.db.First(&u)
// 	c.JSON(200, u)
// }

// func pingPongHandler(c *gin.Context) {
// 	var u User
// 	db.First(&u)
// 	c.JSON(200, u)
// }
