package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {

	// Creates a gin Engine without any middleware
	r := gin.New()

	gin.SetMode(gin.ReleaseMode)

	// Add Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Routing
	routes(r)

	//Start listeners
	httpListeners(r)

}

func routes(r *gin.Engine) {

	// Static assets
	r.StaticFS("webstatic", http.Dir("./webstatic"))

	// Views
	r.LoadHTMLGlob("webviews/*")

	// Other routes
	r.GET("/", home)
	r.GET("/checkout", checkout)
	//r.POST("/intent", intent)
	r.GET("/payment_finished/:pi", paymentFinished)
	r.POST("/webhook", webhook)
}

func httpListeners(r *gin.Engine) {

	////////////
	// go rutine that listens in port 80 and redirects traffic to port 443
	srv := &http.Server{
		Addr:           ":80",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    5 * time.Second,
		MaxHeaderBytes: 1 << 20, //1MB
		Handler:        r,
	}

	srv.ListenAndServe()

}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Main website",
	})
}
