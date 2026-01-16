package repository

import (
	"context"

	"github.com/jejevj/go-aiocap/entity/course"
	"gorm.io/gorm"
)

type (
	CourseRepository interface {
		AddCourse(ctx context.Context, customer course.Course) (course.Course, error)
		// CheckEmail(ctx context.Context, email string) (course.Course, bool, error)
		// CheckName(ctx context.Context, name string) (course.Course, bool, error)
		// CheckPhone(ctx context.Context, phone string) (course.Course, bool, error)
	}
	courseRepository struct {
		db *gorm.DB
	}
)

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{
		db: db,
	}
}

func (r *courseRepository) AddCourse(ctx context.Context, customer course.Course) (course.Course, error) {
	tx := r.db

	if err := tx.WithContext(ctx).Create(&customer).Error; err != nil {
		return course.Course{}, err
	}

	return customer, nil
}

// func (r *courseRepository) CheckEmail(ctx context.Context, email string) (course.Course, bool, error) {
// 	tx := r.db

// 	var user course.Course
// 	if err := tx.WithContext(ctx).Where("customer_email = ?", email).Take(&user).Error; err != nil {
// 		return course.Course{}, false, err
// 	}

// 	return user, true, nil
// }
// func (r *courseRepository) CheckName(ctx context.Context, name string) (course.Course, bool, error) {
// 	tx := r.db

// 	var user course.Course
// 	if err := tx.WithContext(ctx).Where("customer_name = ?", name).Take(&user).Error; err != nil {
// 		return course.Course{}, false, err
// 	}

// 	return user, true, nil
// }
// func (r *courseRepository) CheckPhone(ctx context.Context, phone string) (course.Course, bool, error) {
// 	tx := r.db

// 	var user course.Course
// 	if err := tx.WithContext(ctx).Where("phone_number = ?", phone).Take(&user).Error; err != nil {
// 		return course.Course{}, false, err
// 	}

// 	return user, true, nil
// }
