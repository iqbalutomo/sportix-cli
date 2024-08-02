package entity

type Wallet struct {
	WalletID  uint
	UserID    uint
	Balance   float64
	UpdatedAt string
	CreatedAt string
}

type WalletCurrentUser struct {
	WalletID    uint
	OwnerWallet uint
	Balance     float64
}
