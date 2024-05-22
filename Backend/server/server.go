package server

import (
	"attendance-app/handlers"
	"attendance-app/middleware"
	"attendance-app/store"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Performserver(m *store.MongoStore) {

	router := gin.Default()
	//router.Use(cors.Default())
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, token") // Include 'token' header
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.POST("/users/signup", func(c *gin.Context) {
		handlers.SignUp(c, m)

	})
	// router.POST("/admin/signup", func(c *gin.Context) {
	// 	handlers.SignUpAdmin(c, m)

	// })
	router.POST("/admin", func(c *gin.Context) {
		handlers.LoginAdmin(c, m)

	})
	router.POST("/user/login", func(c *gin.Context) {
		handlers.Login(c, m)

	})
	router.POST("/markattendance", func(c *gin.Context) {
		handlers.InsertAttendance(c, m)
	})
	router.Use(middleware.Authentication())
	router.POST("/admin/userslist", func(c *gin.Context) {
		handlers.GetUsers(c, m)

	})
	router.POST("/user/dashboard", func(c *gin.Context) {
		handlers.GetUsersAttendance(c, m)

	})

	// Open connection with MongoDB
	if err := m.OpenConnectionWithMongoDB("mongodb://localhost:27017", "Attendance-app"); err != nil {
		log.Fatalf("Failed to open connection with MongoDB: %v", err)
	}

	//runs the server with localhost
	if err := router.Run(":9000"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)

	}

}
