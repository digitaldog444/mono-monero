package models

type User struct {
	UserId   uint   `gorm:"primaryKey,index"`
	Username string `gorm:"unique"`
	Password string
	Settings UserSettings
	PGPKey   Keys
}

type Wallet struct {
	WalletId uint `gorm:"primaryKey, index"`
	Address  string
	UserId   uint
}

type Keys struct {
	UserId    uint `gorm:"primaryKey,index"`
	PublicKey string
}

type UserSettings struct {
	UserId     uint `gorm:"primaryKey,index"`
	Currency   string
	TFAEnabled bool
}