package routes

import (
    "github.com/gin-gonic/gin"
    "stocky-assignment/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/reward", controllers.PostReward)
    r.GET("/today-stocks/:userId", controllers.GetTodayStocks)
    r.GET("/stats/:userId", controllers.GetStats)
    r.GET("/historical-inr/:userId", controllers.GetHistoricalINR)
    r.GET("/portfolio/:userId", controllers.GetPortfolio)
}
