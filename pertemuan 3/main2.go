package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
	Nim      string `json:"nim"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	Umur     int    `json:"umur"`
}

var users []User = []User{
	User{Username: "james123", Nim: "1", Password: "123", Nama: "James Catalunya", Umur: 12},
	User{Username: "catalunya321", Nim: "2", Password: "123", Nama: "Rodrigo Catalunya", Umur: 15},
	User{Username: "aderay", Nim: "3", Password: "123", Nama: "Ade Ray", Umur: 19},
	User{Username: "bigo", Nim: "4", Password: "123", Nama: "Bigo - The Most Popular Live Video Chat", Umur: 18},
	User{Username: "michat", Nim: "5", Password: "123", Nama: "Mi Chat", Umur: 19},
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/", getAllUsers)
		usersGroup.POST("/", addNewUser)
		usersGroup.PUT("/:username", updateUserData)
		usersGroup.DELETE("/:username", deleteUser)
		usersGroup.POST("/login", login)
	}

	indexGroup := router.Group("/index")
	{
		indexGroup.GET("/", getHandler)
		indexGroup.GET("/query", getQueryHandler)
		indexGroup.PUT("/:username", putHandler)
		indexGroup.POST("/", postHandler)
		indexGroup.DELETE("/:username", deleteHandler)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router.Run(":" + port)
}

func getHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pesan": "Ini adalah pesan dari GET",
	})
}

func getQueryHandler(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{
		"pesan": "Ini adalah pesan dari GET",
		"query": "Query yang diberikan adalah " + username,
	})
}

func putHandler(c *gin.Context) {
	var body gin.H

	parameter := c.Param("username")
	c.BindJSON(&body)
	c.JSON(http.StatusOK, gin.H{
		"pesan":     "Ini adalah pesan dari PUT",
		"parameter": "Parameter yang dikirimkan adalah " + parameter,
		"body":      body,
	})
}

func postHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Parameter yang diberikan salah",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan":    "Ini adalah pesan dari POST",
		"username": user.Username,
		"nama":     user.Nama,
		"nim":      user.Nim,
		"umur":     user.Umur,
	})
}

func deleteHandler(c *gin.Context) {
	parameter := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"pesan":     "Ini adalah pesan dari DELETE",
		"parameter": "Parameter yang dikirimkan adalah " + parameter,
	})
}

func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"pesan": "Berhasil mendapatkan semua pengguna",
	})
}

func addNewUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal menambahkan pengguna baru, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{
		"pesan":         "Berhasil menambahkan pengguna baru",
		"pengguna_baru": newUser,
	})
}

func updateUserData(c *gin.Context) {
	var data gin.H
	c.BindJSON(&data)

	username := c.Param("username")
	ketemu := false
	for index, user := range users {
		if user.Username == username {
			ketemu = true
			if data["nim"] != nil {
				users[index].Nim = data["nim"].(string)
			}

			if data["nama"] != nil {
				users[index].Nama = data["nama"].(string)
			}

			if data["password"] != nil {
				users[index].Password = data["password"].(string)
			}

			if data["umur"] != nil {
				users[index].Umur = data["umur"].(int)
			}
			break
		}
	}

	if ketemu {
		c.JSON(http.StatusOK, gin.H{
			"pesan":         "Berhasil melakukan perubahan pada data pengguna dengan username " + username,
			"list_pengguna": users,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan data pengguna dengan username " + username,
		})
	}
}

func deleteUser(c *gin.Context) {
	username := c.Param("username")
	ketemu := false
	for index, user := range users {
		if user.Username == username {
			ketemu = true
			users = append(users[:index], users[index+1:]...)
			break
		}
	}

	if ketemu {
		c.JSON(http.StatusOK, gin.H{
			"pesan":         "Berhasil menghapus data pengguna dengan username " + username,
			"list_pengguna": users,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan data pengguna dengan username " + username,
		})
	}
}

func login(c *gin.Context) {
	var data gin.H
	c.Bind(&data)

	for _, user := range users {
		if user.Username == data["username"] {
			if user.Password == data["password"] {
				c.JSON(http.StatusOK, gin.H{
					"pesan":    "Berhasil melakukan login",
					"pengguna": user,
				})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"pesan": "Password yang dimasukkan salah?",
				})
			}
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"pesan": "Tidak ditemukan data pengguna dengan username " + data["username"].(string),
	})
}