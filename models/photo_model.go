package models

type Photo struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	Title      string `gorm:"type:varchar(255)" json:"title"`
	Caption    string `gorm:"type:varchar(255)" json:"caption"`
	PhotoUrl   string `gorm:"type:varchar(255)" json:"photoUrl"`
	UserID     int    `gorm:"index;not null" json:"user_id"`
	Created_At string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP()" json:"created_at"`
	Updated_At string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP()" json:"updated_at"`
}
