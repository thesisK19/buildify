package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/greet", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "greeting": "world",
        })
    })
    err := r.Run(":8882")
    if err != nil {
        log.Fatalf("failed to run: %v", err)
    }
}

// bazel run //:gazelle -- update-repos 