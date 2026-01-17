package repository

import (
	"context"
	"math"

	"github.com/jejevj/go-aiocap/dto"
	"github.com/jejevj/go-aiocap/entity/course"
	"gorm.io/gorm"
)

type (
	CourseCustomerRepository interface {
		AddCourseCustomer(ctx context.Context, customer course.CourseCustomer) (course.CourseCustomer, error)
		CheckEmail(ctx context.Context, email string) (course.CourseCustomer, bool, error)
		CheckName(ctx context.Context, name string) (course.CourseCustomer, bool, error)
		CheckPhone(ctx context.Context, phone string) (course.CourseCustomer, bool, error)
		GetAllCourseCustomer(ctx context.Context, req dto.PaginationRequest) (dto.GetAllCCResponse, error)
		GetCourseCustomerById(ctx context.Context, id string) (course.CourseCustomer, error)
		UpdateCourseCustomer(ctx context.Context, customer course.CourseCustomer) (course.CourseCustomer, error)
		GetAllCourseCustomerForExport(ctx context.Context, req dto.PaginationRequest) ([]dto.CourseCustomerExport, error)
		DeleteCourseCustomer(ctx context.Context, id string) error
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
func (r *courseCustomerRepository) GetAllCourseCustomer(ctx context.Context, req dto.PaginationRequest) (dto.GetAllCCResponse, error) {
	tx := r.db

	var customers []course.CourseCustomer
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	// Build the query with search if provided
	query := tx.WithContext(ctx).Model(&course.CourseCustomer{})

	// Apply search filter if search term is provided
	if req.Search != "" {
		query = query.Where("customer_name LIKE ? OR customer_email LIKE ? OR contact_name LIKE ? OR phone_number LIKE ?",
			"%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	// Count the filtered results
	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllCCResponse{}, err
	}

	// Apply pagination to the query
	offset := (req.Page - 1) * req.PerPage
	if err := query.Offset(offset).Limit(req.PerPage).Find(&customers).Error; err != nil {
		return dto.GetAllCCResponse{}, err
	}
	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllCCResponse{
		CourseCustomer: customers,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *courseCustomerRepository) GetCourseCustomerById(ctx context.Context, id string) (course.CourseCustomer, error) {
	tx := r.db

	var customer course.CourseCustomer
	if err := tx.WithContext(ctx).Where("id = ?", id).First(&customer).Error; err != nil {
		return course.CourseCustomer{}, err
	}

	return customer, nil
}

func (r *courseCustomerRepository) UpdateCourseCustomer(ctx context.Context, customer course.CourseCustomer) (course.CourseCustomer, error) {
	tx := r.db

	if err := tx.WithContext(ctx).Updates(&customer).Error; err != nil {
		return course.CourseCustomer{}, err
	}

	return customer, nil
}

func (r *courseCustomerRepository) DeleteCourseCustomer(ctx context.Context, id string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&course.CourseCustomer{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (r *courseCustomerRepository) GetAllCourseCustomerForExport(ctx context.Context, req dto.PaginationRequest) ([]dto.CourseCustomerExport, error) {
	tx := r.db

	var customers []course.CourseCustomer
	var exportData []dto.CourseCustomerExport

	// Build the query with search if provided
	query := tx.WithContext(ctx).Model(&course.CourseCustomer{})

	// Apply search filter if search term is provided
	if req.Search != "" {
		query = query.Where("customer_name LIKE ? OR customer_email LIKE ? OR contact_name LIKE ? OR phone_number LIKE ?",
			"%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	// Find all records (no pagination for export)
	if err := query.Find(&customers).Error; err != nil {
		return []dto.CourseCustomerExport{}, err
	}

	// Convert to export format
	for _, customer := range customers {
		exportData = append(exportData, dto.CourseCustomerExport{
			CustomerName:    customer.CustomerName,
			CustomerEmail:   customer.CustomerEmail,
			ContactName:     customer.ContactName,
			PhoneNumber:     customer.PhoneNumber,
			CustomerAddress: customer.CustomerAddress,
			CreatedAt:       customer.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       customer.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return exportData, nil
}
