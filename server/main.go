package main

import (
    "net/http"
    "context"
    "fmt"

    "github.com/gin-gonic/gin"
    lama "github.com/ollama/ollama/api"
)

func main() {
    router := gin.Default()
    router.SetTrustedProxies([]string{}) // Silences the "trusted all proxies" warning
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
            "message": "Invalid request body. Must include a JSON field 'ask_msg' containing the text message.",
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
  
    // Set headers for streaming
    c.Header("Content-Type", "text/event-stream")
    c.Header("Cache-Control", "no-cache")
    c.Header("Connection", "keep-alive")

    err := client.Generate(ctx, req, func(resp lama.GenerateResponse) error {
      // Write the chunk directly to the response writer
      fmt.Fprint(c.Writer, resp.Response)
      
      // Flush the writer to ensure the client receives the chunk immediately
      if c.Writer != nil {
        if flusher, ok := c.Writer.(http.Flusher); ok {
            flusher.Flush()
        }
      }

      if resp.Done {
        fmt.Fprint(c.Writer, "\n")
        if c.Writer != nil {
            if flusher, ok := c.Writer.(http.Flusher); ok {
                flusher.Flush()
            }
        }
      }
      return nil
    })
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "status":  "error",
        "message": err.Error(),
      })
      return
    }
}