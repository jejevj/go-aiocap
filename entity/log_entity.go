package entity

import (
	"time"

	"github.com/google/uuid"
)

type SystemLog struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Action    string    `json:"action" gorm:"type:varchar(255)"`
	Endpoint  string    `json:"endpoint" gorm:"type:varchar(255)"`
	Method    string    `json:"method" gorm:"type:varchar(50)"`
	UserID    *string   `json:"user_id" gorm:"type:uuid"`
	Request   string    `json:"request" gorm:"type:text"`
	Response  string    `json:"response" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
