package main

import(
	"log"
	"net/http"
	"os"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required`
	Password string `form:"password" json:"password" binding:"required`
	Nama string `form:"nama" json:"nama" binding:"required`

}

var users []User = []User{
	User{Username: "james123", Password: "123", Nama: "James Catalunya"},
	User{Username: "catalunya321", Password: "123", Nama: "Rodrigo Catalunya"},
	User{Username: "aderay", Password: "123", Nama: "Ade Ray"},
	User{Username: "capek", Password: "123", Nama: "capek"},
	User{Username: "micak", Password: "123", Nama: "Mi Chat"},
}

var identityKey = "id"

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals User
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			for _, user := range users {
				if user.Username == userID {
					if user.Password == password {
						return &User{
							Username:  userID,
							Password:  "Bo-Yi",
							Nama: "Wu",
						}, nil
					}
				}
			}

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.Username == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/", getAllUsers)
		usersGroup.POST("/", addNewUser)
		usersGroup.PUT("/:username", updateUserData)
		usersGroup.DELETE("/:username", deleteUser)
		usersGroup.POST("/login", authMiddleware.LoginHandler)
	}

	// usersGroup := router.Group("/todos")
	// {
	// 	usersGroup.GET("/", getAllUsers)
	// 	usersGroup.POST("/", addNewUser)
	// 	usersGroup.PUT("/:id_todo", updateUserData)
	// 	usersGroup.DELETE("/:id_todo", deleteUser)
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router.Run(":" + port)
}

func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"pesan": "Berhasil mendapatkan semua data user",
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
	c.ShouldBindJSON(&data)

	username := c.Param("username")
	ketemu := false
	for index, user := range users {
		if user.Username == username {
			ketemu = true
			if data["nama"] != nil {
				users[index].Nama = data["nama"].(string)
			}

			if data["password"] != nil {
				users[index].Password = data["password"].(string)
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

// func login(c *gin.Context) {
// 	var data gin.H
// 	c.Bind(&data)

// 	for _, user := range users {
// 		if user.Username == data["username"] {
// 			if user.Password == data["password"] {
// 				sign := jwt.New(jwt.GetSigningMethod("HS256"))
// 				token, err := sign.SignedString([]byte("secret"))
// 				if err != nil {
// 					c.JSON(http.StatusInternalServerError, gin.H{
// 						"message": err.Error(),
// 					})
// 					c.Abort()
// 				}

// 				c.JSON(http.StatusOK, gin.H{
// 					"pesan":    "Berhasil melakukan login",
// 					"pengguna": user,
// 				})
// 			} else {
// 				c.JSON(http.StatusUnauthorized, gin.H{
// 					"pesan": "Password yang dimasukkan salah?",
// 				})
// 			}
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{
// 		"pesan": "Tidak ditemukan data pengguna dengan username " + data["username"].(string),
// 	})
// }
