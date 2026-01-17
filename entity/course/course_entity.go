package course

import (
	"time"

	"github.com/google/uuid"
	"github.com/jejevj/go-aiocap/entity"
	"gorm.io/gorm"
)

type Course struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CourseName        string    `json:"course_name"`
	CourseDescription string    `json:"course_description"`
	CourseLocation    string    `json:"course_location"`
	CourseClient      uuid.UUID `gorm:"type:uuid;not null" json:"course_client_id"`
	CreatedByID       uuid.UUID `gorm:"type:uuid;not null" json:"created_by_id"`
	ChangedByID       uuid.UUID `gorm:"type:uuid;not null" json:"changed_by_id"`
	IsVerified        bool      `json:"is_verified"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	// Relationship fields
	CourseClientCustomer CourseCustomer `gorm:"foreignKey:CourseClient;references:ID" json:"course_client_customer,omitempty"`
	CreatedByUser        entity.User    `gorm:"foreignKey:CreatedByID;references:ID" json:"created_by_user,omitempty"`
	ChangedByUser        entity.User    `gorm:"foreignKey:ChangedByID;references:ID" json:"changed_by_user,omitempty"`

	entity.Timestamp
}

func (u *Course) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return nil
}
