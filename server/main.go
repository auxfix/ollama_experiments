package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/ask_mii", askMii)

    router.Run("localhost:6767")
}

// getAlbums responds with the list of all albums as JSON.
func askMii(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "we all gonna lie. sooner or later")
}