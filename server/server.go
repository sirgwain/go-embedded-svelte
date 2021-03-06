package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Start() {

	// create the data dir
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("data/session.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	store := gormsessions.NewStore(db, true, []byte("secret"))

	r := gin.Default()
	r.Use(sessions.Sessions("session", store))

	// create a static folder for embedded assets
	embeddedFolder := EmbedFolder(GetAssets(), "frontend/build")
	staticServer := static.Serve("/", embeddedFolder)
	r.Use(staticServer)

	// weird workaround for embedded static assets (from gin-contrib/static github issue)
	r.NoRoute(func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		log.Println("count", count)
		session.Save()
		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})

	r.POST("/api/login", Login)
	r.GET("/api/logout", Logout)

	// authorized routes
	ar := r.Group("/api")
	ar.Use(AuthRequired)
	ar.GET("/me", Me)

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
