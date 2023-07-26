package init

type BaseModel struct {
}

type ID struct {
	ID int `gorm:"primary_key" json:"id"`
}

type Timestamps struct {
	CreatedAt int64 `gorm:"autoUpdateTime:milli;column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli;column:updated_at" db:"updated_at" json:"updated_at"`
}
