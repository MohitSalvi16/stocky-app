package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "stocky-assignment/db"
)

func GetTodayStocks(c *gin.Context) {
    userID := c.Param("userId")
    rows, _ := db.DB.Query(`SELECT stock_symbol, shares, timestamp FROM rewards WHERE user_id=$1 AND DATE(timestamp)=CURRENT_DATE`, userID)
    defer rows.Close()
    var res []map[string]interface{}
    for rows.Next() {
        var s string
        var sh float64
        var t time.Time
        rows.Scan(&s, &sh, &t)
        res = append(res, gin.H{"stock": s, "shares": sh, "timestamp": t})
    }
    c.JSON(http.StatusOK, res)
}

func GetStats(c *gin.Context) {
    userID := c.Param("userId")
    rows, _ := db.DB.Query(`SELECT stock_symbol,SUM(shares) FROM rewards WHERE user_id=$1 AND DATE(timestamp)=CURRENT_DATE GROUP BY stock_symbol`, userID)
    defer rows.Close()
    total := 0.0
    stocks := []map[string]interface{}{}
    for rows.Next() {
        var s string
        var sh float64
        rows.Scan(&s, &sh)
        p := GetRandomStockPrice(s)
        v := p * sh
        total += v
        stocks = append(stocks, gin.H{"stock": s, "total_shares": sh, "inr_value": v})
    }
    c.JSON(http.StatusOK, gin.H{"total_today": stocks, "portfolio_value": gin.H{"total_inr": total}})
}

func GetHistoricalINR(c *gin.Context) {
    userID := c.Param("userId")
    rows, err := db.DB.Query(`SELECT DATE(timestamp) AS d,stock_symbol,SUM(shares) FROM rewards WHERE user_id=$1 AND DATE(timestamp)<CURRENT_DATE GROUP BY d,stock_symbol ORDER BY d DESC`, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
        return
    }
    defer rows.Close()
    data := map[string]float64{}
    for rows.Next() {
        var d time.Time
        var s string
        var sh float64
        rows.Scan(&d, &s, &sh)
        p := GetRandomStockPrice(s)
        data[d.Format("2006-01-02")] += p * sh
    }
    c.JSON(http.StatusOK, gin.H{"historical_inr": data})
}

func GetPortfolio(c *gin.Context) {
    userID := c.Param("userId")
    rows, err := db.DB.Query(`SELECT stock_symbol,SUM(shares) FROM rewards WHERE user_id=$1 GROUP BY stock_symbol`, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
        return
    }
    defer rows.Close()
    total := 0.0
    port := []map[string]interface{}{}
    for rows.Next() {
        var s string
        var sh float64
        rows.Scan(&s, &sh)
        p := GetRandomStockPrice(s)
        v := p * sh
        total += v
        port = append(port, gin.H{"stock": s, "shares": sh, "current_inr": v})
    }
    c.JSON(http.StatusOK, gin.H{"portfolio": port, "total_inr": total})
}
