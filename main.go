package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
r := gin.Default()

// API v1
router := r.Group("/state")
{
	//router.POST("/create", postBlog)
	router.GET("/", getStates)
	//router.POST("/update/:id", updateBlog)
	//router.GET("/delete/:id", deleteBlog)
}

// By default it serves on :8080
r.Run()
  
}

func getStates(c *gin.Context) {
	c.JSON(200, gin.H{"message": "A new Record Created!"})
}

func readBlog(c *gin.Context) {
	c.JSON(200, gin.H{"message": "All Records"})
}

func updateBlog(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Updated!"})
}
func deleteBlog(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Deleted!"})
}

