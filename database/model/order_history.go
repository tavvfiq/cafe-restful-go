package model

// OrderHistory pivot tabel to menu, user, and history
type OrderHistory struct {
	MenuID    uint
	HistoryID uint
	UserID    uint
	Quantity  int
}
