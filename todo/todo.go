package todo

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var validate *validator.Validate

type Todo struct {
	Title string `json:"text"`
	gorm.Model
}

// กำหนดชื่อ table ด้วยตัวเอง
func (Todo) TableName() string {
	return "todos"
}

type ValidateHandler struct {
	db *gorm.DB
}

type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (t *TodoHandler) NewTask(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := t.db.Create(&todo)
	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": todo.Model.ID,
	})
}

func (t *TodoHandler) GetFirstTask(c *gin.Context) {
	var todo []Todo

	r := t.db.First(&todo)
	if err := r.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (t *TodoHandler) GetAllTask(c *gin.Context) {
	var todo []Todo

	r := t.db.Find(&todo)
	if err := r.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}
