package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "stocky-assignment/db"
)

type RewardRequest struct {
    UserID int     `json:"user_id"`
    Stock  string  `json:"stock"`
    Shares float64 `json:"shares"`
}

func PostReward(c *gin.Context) {
    var req RewardRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    t := time.Now()
    _, err := db.DB.Exec(`INSERT INTO rewards (user_id, stock_symbol, shares, timestamp) VALUES ($1,$2,$3,$4)`,
        req.UserID, req.Stock, req.Shares, t)
    if err != nil {
        logrus.Error("Insert reward failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
        return
    }
    price := GetRandomStockPrice(req.Stock)
    total := price * req.Shares
    fees := total * 0.01
    _, err = db.DB.Exec(`INSERT INTO ledger (user_id, stock_symbol, units, cash_outflow, fees, timestamp) VALUES ($1,$2,$3,$4,$5,$6)`,
        req.UserID, req.Stock, req.Shares, total, fees, t)
    if err != nil {
        logrus.Error("Insert ledger failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ledger error"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Reward recorded successfully"})
}
