package course

import (
	"github.com/google/uuid"
	"github.com/jejevj/go-aiocap/entity"
	"gorm.io/gorm"
)

type CourseKehadiran struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	PesertaID  uuid.UUID `json:"peserta_id"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	CourseID   uuid.UUID `gorm:"type:uuid;not null" json:"course_id"`

	Peserta  CoursePeserta  `gorm:"foreignKey:PesertaID;references:ID" json:"peserta,omitempty"`
	Course   Course         `gorm:"foreignKey:CourseID;references:ID" json:"course,omitempty"`
	Customer CourseCustomer `gorm:"foreignKey:CustomerID;references:ID" json:"customer,omitempty"`

	entity.Timestamp
}

func (u *CourseKehadiran) BeforeCreate(tx *gorm.DB) error {
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
