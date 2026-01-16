package course

import (
	"github.com/google/uuid"
	"github.com/jejevj/go-aiocap/entity"
	"gorm.io/gorm"
)

type CoursePeserta struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	PesertaName    string    `json:"peserta_name"`
	PhoneNumber    string    `json:"phone_number"`
	PesertaAddress string    `json:"peserta_address"`
	CourseID       uuid.UUID `gorm:"type:uuid;not null" json:"course_id"`
	CreatedByID    uuid.UUID `gorm:"type:uuid;not null" json:"created_by_id"`
	ChangedByID    uuid.UUID `gorm:"type:uuid;not null" json:"changed_by_id"`
	// Relationship fields
	Course        Course      `gorm:"foreignKey:CourseID;references:ID" json:"course,omitempty"`
	CreatedByUser entity.User `gorm:"foreignKey:CreatedByID;references:ID" json:"created_by_user,omitempty"`
	ChangedByUser entity.User `gorm:"foreignKey:ChangedByID;references:ID" json:"changed_by_user,omitempty"`

	entity.Timestamp
}

func (u *CoursePeserta) BeforeCreate(tx *gorm.DB) error {
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
