package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// Create New Router With `gin` Engine
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/hello", getHello)
	router.GET("/hello/:name", getHelloName)
	router.GET("/data", datas)
	router.GET("/form", getForm)
	router.POST("/form", postForm)
	err := router.Run("localhost:8080")
	// if the router hase a error we shouting down that
	if err != nil {
		log.Fatal(err)
	}

}

// it just return hello world
func getHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

// we get a parameter in url and write this to html file: `localhost:8080/hello/{Name}`
func getHelloName(c *gin.Context) {
	name := c.Param("name")
	c.HTML(http.StatusOK, "hello.html", name)
}

// this function return html page with many data with `gin.H`
func datas(c *gin.Context) {
	c.HTML(http.StatusOK, "data.html", gin.H{
		"name":  "Mahdi",
		"food":  "Kebab",
		"color": "Yellow",
	})
}

// this func just return html page with null objects ( it just a blank form for send data to post form )
func getForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

// we send data from getForm and write on html form with method POST
func postForm(c *gin.Context) {
	name := c.PostForm("name")
	color := c.PostForm("color")

	c.HTML(http.StatusOK, "post_form.html", gin.H{
		"name":  name,
		"color": color,
	})
}
