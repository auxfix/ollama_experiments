package main

import (
    "net/http"
    "context"

    "github.com/gin-gonic/gin"
    lama "github.com/ollama/ollama/api"
)

func main() {
    router := gin.Default()
    router.SetTrustedProxies([]string{}) // Silences the "trusted all proxies" warning

    // CORS and Cache Middleware
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Disable caching completely
        c.Writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
        c.Writer.Header().Set("Pragma", "no-cache")
        c.Writer.Header().Set("Expires", "0")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })

    router.POST("/ask_mii", askMii)

    router.Run("localhost:6767")
}

type RequestBody struct {
    AskMsg string `json:"ask_msk" binding:"required"`
}


func askMii(c *gin.Context) {
    var request RequestBody

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status": "error",
            "message": "Invalid request body. Must include a JSON field 'ask_msk' containing the text message.",
        })
        return
    }

    ask := request.AskMsg

    ctx := context.Background()
    client, err_cl := lama.ClientFromEnvironment()

    if err_cl != nil {
      panic(err_cl)
    }
  
    req := &lama.GenerateRequest{
      Model:  "gemma4:12b-it-q4_K_M",
      Prompt: ask,
    }
  
    var fullResponse string
    err := client.Generate(ctx, req, func(resp lama.GenerateResponse) error {
      fullResponse += resp.Response
      return nil
    })

    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "status":  "error",
        "message": err.Error(),
      })
      return
    }

    c.JSON(http.StatusOK, gin.H{
      "status":  "success",
      "message": fullResponse,
    })
}