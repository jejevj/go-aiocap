package repository

import (
	"context"

	"github.com/jejevj/go-aiocap/entity/course"
	"gorm.io/gorm"
)

type (
	CourseCustomerRepository interface {
		AddCourseCustomer(ctx context.Context, customer course.CourseCustomer) (course.CourseCustomer, error)
		CheckEmail(ctx context.Context, email string) (course.CourseCustomer, bool, error)
		CheckName(ctx context.Context, name string) (course.CourseCustomer, bool, error)
		CheckPhone(ctx context.Context, phone string) (course.CourseCustomer, bool, error)
		GetAllCourseCustomer(ctx context.Context) ([]course.CourseCustomer, error)
		GetCourseCustomerById(ctx context.Context, id string) (course.CourseCustomer, error)
	}
	courseCustomerRepository struct {
		db *gorm.DB
	}
)

func NewCourseCustomerRepository(db *gorm.DB) CourseCustomerRepository {
	return &courseCustomerRepository{
		db: db,
	}
}

func (r *courseCustomerRepository) AddCourseCustomer(ctx context.Context, customer course.CourseCustomer) (course.CourseCustomer, error) {
	tx := r.db

	if err := tx.WithContext(ctx).Create(&customer).Error; err != nil {
		return course.CourseCustomer{}, err
	}

	return customer, nil
}

func (r *courseCustomerRepository) CheckEmail(ctx context.Context, email string) (course.CourseCustomer, bool, error) {
	tx := r.db

	var user course.CourseCustomer
	if err := tx.WithContext(ctx).Where("customer_email = ?", email).Take(&user).Error; err != nil {
		return course.CourseCustomer{}, false, err
	}

	return user, true, nil
}
func (r *courseCustomerRepository) CheckName(ctx context.Context, name string) (course.CourseCustomer, bool, error) {
	tx := r.db

	var user course.CourseCustomer
	if err := tx.WithContext(ctx).Where("customer_name = ?", name).Take(&user).Error; err != nil {
		return course.CourseCustomer{}, false, err
	}

	return user, true, nil
}
func (r *courseCustomerRepository) CheckPhone(ctx context.Context, phone string) (course.CourseCustomer, bool, error) {
	tx := r.db

	var user course.CourseCustomer
	if err := tx.WithContext(ctx).Where("phone_number = ?", phone).Take(&user).Error; err != nil {
		return course.CourseCustomer{}, false, err
	}

	return user, true, nil
}

func (r *courseCustomerRepository) GetAllCourseCustomer(ctx context.Context) ([]course.CourseCustomer, error) {
	tx := r.db

	var customers []course.CourseCustomer
	if err := tx.WithContext(ctx).Find(&customers).Error; err != nil {
		return []course.CourseCustomer{}, err
	}

	return customers, nil
}

func (r *courseCustomerRepository) GetCourseCustomerById(ctx context.Context, id string) (course.CourseCustomer, error) {
	tx := r.db

	var customer course.CourseCustomer
	if err := tx.WithContext(ctx).Where("id = ?", id).First(&customer).Error; err != nil {
		return course.CourseCustomer{}, err
	}

	return customer, nil
}
