package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
	"net/http"
)	

type User struct {
	Username string `gorm:"primary_key"`
	Password string
	Nama 	 string
	Todo []Todo `gorm:"foreignkey:Username;association_foreignkey:Username"`
}

type Todo struct {
	IdTodo    int `gorm:"primary_key;AUTO_INCREMENT"`
	Username  string
	Tugas     string
	Deskripsi string
	Deadline  string
	Status    bool `gorm:"default:false"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:@/golangdsc?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Errorf("Terjadi kesalahan saat membuka koneksi ke server MySQL...")
		return
	}
	fmt.Println("Berhasil melakukan koneksi ke server MySQL...")

	db.AutoMigrate(&User{}, &Todo{})
	db.Model(&Todo{}).AddForeignKey("Username", "Users(Username)", "CASCADE", "CASCADE")

	router := gin.Default()

	userGroup := router.Group("/users")
	{
		userGroup.POST("/", addUsers)
		userGroup.POST("/login", login)
		userGroup.PUT("/:username", updateUser)
	}

	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("/", getAllTodos)
		todoGroup.POST("/", addTodos)
		todoGroup.PUT("/:idTodo", updateTodos)
		todoGroup.DELETE("/:idTodo", deleteTodos)
	}

	router.Run(":5000")
}

func getAllTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusCreated, gin.H{
		"pesan": "Berhasil mendapatkan semua todo",
		"todo":  todos,
	})
}

func addTodos(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal menambahkan todo baru, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	db.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"pesan":     "Berhasil menambahkan todo baru",
		"todo_baru": todo,
	})
}

func updateTodos(c *gin.Context) {
	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal update todo, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	idTodo := c.Param("idTodo")

	var todo Todo
	if db.Where("id_todo = ?", idTodo).First(&todo).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan todo dengan id " + idTodo,
		})
		return
	}

	db.Model(&todo).Where("id_todo = ?", idTodo).Updates(updatedTodo)
	c.JSON(http.StatusOK, gin.H{
		"pesan":     "Berhasil melakukan update user",
		"todo_baru": todo,
	})
}

func deleteTodos(c *gin.Context) {
	idTodo := c.Param("idTodo")

	var todo Todo
	if db.Where("id_todo = ?", idTodo).First(&todo).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan data todo dengan id " + idTodo,
		})
		return
	}

	db.Model(&todo).Where("id_todo = ?", idTodo).Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"pesan":         "Berhasil melakukan update todo",
		"pengguna_baru": todo,
	})
}

func addUsers(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal menambahkan pengguna baru, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"pesan":         "Berhasil menambahkan pengguna baru",
		"pengguna_baru": user,
	})
}

func login(c *gin.Context) {
	var body gin.H
	c.BindJSON(&body)

	var user User
	if db.Where("username = ?", body["username"]).First(&user).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan data pengguna dengan username " + body["username"].(string),
		})
		return
	}

	if user.Password != body["password"] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"pesan": "Password yang dimasukkan salah",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan":         "Berhasil melakukan login",
		"pengguna_baru": user,
	})
}

func updateUser(c *gin.Context) {
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal update pengguna baru, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	username := c.Param("username")

	var user User
	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan data pengguna dengan username " + username,
		})
		return
	}

	db.Model(&user).Where("username = ?", username).Updates(updatedUser)
	c.JSON(http.StatusOK, gin.H{
		"pesan":         "Berhasil melakukan update user",
		"pengguna_baru": updatedUser,
	})
}