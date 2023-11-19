package main

// import modules
import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// Create New Router With `gin` Engine
	router := gin.Default()
	// Specifies where the html files are stored
	router.LoadHTMLGlob("templates/*.html")
	// the routers gives two args: 1.path in url. 2.function runner
	router.GET("/hello", getHello)
	router.GET("/hello/:name", getHelloName)
	router.GET("/data", datas)
	router.GET("/form", getForm)
	router.POST("/form", postForm)
	// we execute router on localhost:8000
	err := router.Run("localhost:8080")
	// if the router hase error we shouting down that
	if err != nil {
		// exit from router and write router errors
		log.Fatal(err)
	}

}

// it just return hello world
func getHello(c *gin.Context) {
	// as string we write `hello world`. this page does not use html page
	c.String(http.StatusOK, "Hello World!")
}

// we get a parameter in url and write this to html file: `localhost:8080/hello/{Name}`
func getHelloName(c *gin.Context) {
	name := c.Param("name")
	// we passed the html page on this router
	c.HTML(http.StatusOK, "hello.html", name)
}

// this function return html page with many data with `gin.H`
func datas(c *gin.Context) {
	// originally we send data as map or dictionary with gin.H. in left side name of variable and right side value of variable
	c.HTML(http.StatusOK, "data.html", gin.H{
		"name":  "GodFather",
		"food":  "Kebab",
		"color": "Yellow",
	})
}

// this func just return html page with null objects ( it just a blank form for send data to post form )
func getForm(c *gin.Context) {
	// context or object is nil or null
	c.HTML(http.StatusOK, "form.html", nil)
}

// we send data from getForm and write on html form with method POST
func postForm(c *gin.Context) {
	// The information is originally taken from the post method and stored in our variables
	name := c.PostForm("name")
	color := c.PostForm("color")

	c.HTML(http.StatusOK, "post_form.html", gin.H{
		"name":  name,
		"color": color,
	})
}
