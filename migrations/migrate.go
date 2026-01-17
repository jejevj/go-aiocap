package migrations

import (
	"fmt"

	"github.com/jejevj/go-aiocap/entity"
	"github.com/jejevj/go-aiocap/entity/course"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	queries := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	}

	for _, query := range queries {
		result := db.Exec(query)
		if result.Error != nil {
			fmt.Println("Error executing query:", result.Error)
		} else {
			fmt.Println("Executed query successfully:", query)
		}
	}

	if err := db.AutoMigrate(
		&entity.User{},
		&course.CourseCustomer{},
		&course.Course{},
		&course.CourseKehadiran{},
		&course.CoursePeserta{},
		&entity.SystemLog{},
	); err != nil {
		return err
	}

	return nil
}
