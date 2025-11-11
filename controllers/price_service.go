package controllers

import (
    "math/rand"
    "time"
)

func GetRandomStockPrice(symbol string) float64 {
    rand.Seed(time.Now().UnixNano())
    min, max := 2000.0, 4000.0
    return min + rand.Float64()*(max-min)
}
