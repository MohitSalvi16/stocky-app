package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "stocky-assignment/db"
    "stocky-assignment/routes"
)

func main() {
    logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
    db.Connect()
    r := gin.Default()
    routes.RegisterRoutes(r)
    logrus.Info("🚀 Stocky server started on :8080")
    r.Run(":8080")
}
