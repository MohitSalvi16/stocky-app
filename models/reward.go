package models

import "time"

type Reward struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Stock     string    `json:"stock"`
    Shares    float64   `json:"shares"`
    Timestamp time.Time `json:"timestamp"`
}
