package models

import "time"

type Ledger struct {
    ID          int       `json:"id"`
    UserID      int       `json:"user_id"`
    StockSymbol string    `json:"stock_symbol"`
    Units       float64   `json:"units"`
    CashOutflow float64   `json:"cash_outflow"`
    Fees        float64   `json:"fees"`
    Timestamp   time.Time `json:"timestamp"`
}
