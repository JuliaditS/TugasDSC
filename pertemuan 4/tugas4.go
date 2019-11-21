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

type Gebetan struct {
	UserID string `form:"UserID" json:"UserID" binding:"required`
	Nama string `form:"nama" json:"nama" binding:"required`
	Umur string `form:"umur" json:"umur" binding:"required`
	Alamat string `form:"alamat" json:"alamat" binding:"required`

}

var gebetans []Gebetan = []Gebetan{}

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
					"Nama": v.Nama,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Nama: claims["Nama"].(string),
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
							Password:  password,
							Nama: user.Nama,
						}, nil
					}
				}
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			var loginVals User
			if v, ok := data.(*User); ok && v.Username == loginVals.Username {
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

	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	usersGroup := router.Group("/Gebetan")
	{	
		usersGroup.Use(authMiddleware.MiddlewareFunc())
		{
			usersGroup.GET("/", getAllGebetan)
			usersGroup.POST("/", addGebetan)
			usersGroup.DELETE("/:nama", deleteGebetan)
		}
		
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

func addGebetan(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	var newGebetan Gebetan
	newGebetan.UserID = claims[identityKey].(string)
	if err := c.ShouldBindJSON(&newGebetan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Username": claims[identityKey],
			"Nama": user.(*User).Nama,
			"pesan": "Gagal menambahkan gebtan, pastikan anda tidak jomblo",
		})
		return
	}

	gebetans = append(gebetans, newGebetan)
	c.JSON(http.StatusCreated, gin.H{
		"Username": claims[identityKey],
		"Nama": user.(*User).Nama,
		"pesan":         "Berhasil menambahkan gebetan baru,selamat anda bucin!!!",
		"Gebetan_baru": newGebetan,
	})
}

func getAllGebetan(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	var listGebetan Gebetan 
	var myGebetans []Gebetan = []Gebetan{}
	for index, gebetan := range gebetans {
		if claims[identityKey].(string) == gebetans[index].UserID {
			listGebetan.UserID = gebetan.UserID
			listGebetan.Nama = gebetan.Nama
			listGebetan.Umur = gebetan.Umur
			listGebetan.Alamat = gebetan.Alamat
			myGebetans = append(myGebetans, listGebetan)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  myGebetans,
		"pesan": "Berhasil mendapatkan semua data Gebetan",
	})
}

func deleteGebetan(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	nama := c.Param("nama")
	ketemu := false
	for index, gebetan := range gebetans {
		if gebetan.Nama == nama && gebetan.UserID == claims[identityKey].(string) {
			ketemu = true
			gebetans = append(gebetans[:index], gebetans[index+1:]...)
			break
		}
	}

	var listGebetan Gebetan 
	var myGebetans []Gebetan = []Gebetan{}
	for index, gebetan := range gebetans {
		if claims[identityKey].(string) == gebetans[index].UserID {
			listGebetan.UserID = gebetan.UserID
			listGebetan.Nama = gebetan.Nama
			listGebetan.Umur = gebetan.Umur
			listGebetan.Alamat = gebetan.Alamat
			myGebetans = append(myGebetans, listGebetan)
		}
	}

	if ketemu {
		c.JSON(http.StatusOK, gin.H{
			"pesan":         "Berhasil moveon dari mantan dengan nama " + nama,
			"data": myGebetans,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan gebetan dengan nama " + nama + "Pastikan anda tidak jomblo!!!",
		})
	}
}