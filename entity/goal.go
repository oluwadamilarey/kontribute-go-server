package entity

//Goal represents goals table in database
type Goal struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name         string `gorm:"type:varchar(255)" json:"name"`
	Description  string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	UserID       string `gorm:"-" json:"token,omitempty"`
	TargetAmount string `gorm:"-" json:"token,omitempty"`
}
