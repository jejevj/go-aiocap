package course

import (
	"github.com/google/uuid"
	"github.com/jejevj/go-aiocap/entity"
	"gorm.io/gorm"
)

type CourseCustomer struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CustomerName    string    `json:"customer_name"`
	ContactName     string    `json:"contact_name"`
	PhoneNumber     string    `json:"phone_number"`
	CustomerAddress string    `json:"customer_address"`
	CreatedByID     uuid.UUID `gorm:"type:uuid;not null" json:"created_by_id"`
	ChangedByID     uuid.UUID `gorm:"type:uuid;not null" json:"changed_by_id"`
	// Relationship fields
	CreatedByUser entity.User `gorm:"foreignKey:CreatedByID;references:ID" json:"created_by_user,omitempty"`
	ChangedByUser entity.User `gorm:"foreignKey:ChangedByID;references:ID" json:"changed_by_user,omitempty"`

	entity.Timestamp
}

func (u *CourseCustomer) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// var err error
	// // u.ID = uuid.New()
	// u.Password, err = helpers.HashPassword(u.Password)
	// if err != nil {
	// 	return err
	// }
	return nil
}
