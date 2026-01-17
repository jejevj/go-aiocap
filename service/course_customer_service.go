package service

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jejevj/go-aiocap/dto"
	"github.com/jejevj/go-aiocap/entity/course"
	"github.com/jejevj/go-aiocap/repository"
	"github.com/xuri/excelize/v2"
)

type (
	CourseCustomerService interface {
		AddCourseCustomer(ctx context.Context, req dto.CourseCustomerCreateRequest) (dto.CourseCustomerResponse, error)
		GetAllCourseCustomer(ctx context.Context, req dto.PaginationRequest) (dto.CourseCustomerPaginationResponse, error)
		GetCourseCustomerById(ctx context.Context, id string) (dto.CourseCustomerResponse, error)
		UpdateCourseCustomer(ctx context.Context, req dto.CourseCustomerUpdateRequest, userId string) (dto.CourseCustomerResponse, error)
		ExportCourseCustomerToExcel(ctx context.Context, req dto.PaginationRequest) ([]byte, error)
	}

	courseCustomerService struct {
		customerRepo repository.CourseCustomerRepository
		jwtService   JWTService
	}
)

func NewCourseCustomerService(customerRepo repository.CourseCustomerRepository, jwtService JWTService) CourseCustomerService {
	return &courseCustomerService{
		customerRepo: customerRepo,
		jwtService:   jwtService,
	}
}

func (s *courseCustomerService) AddCourseCustomer(ctx context.Context, req dto.CourseCustomerCreateRequest) (dto.CourseCustomerResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	_, flag2, _ := s.customerRepo.CheckName(ctx, req.CustomerName)
	if flag2 {
		return dto.CourseCustomerResponse{}, dto.ErrNameAlreadyExists
	}

	_, flag, _ := s.customerRepo.CheckEmail(ctx, req.CustomerEmail)
	if flag {
		return dto.CourseCustomerResponse{}, dto.ErrEmailAlreadyExists
	}

	_, flag3, _ := s.customerRepo.CheckPhone(ctx, req.PhoneNumber)
	if flag3 {
		return dto.CourseCustomerResponse{}, dto.ErrPhoneAlreadyExists
	}
	customer := course.CourseCustomer{
		CustomerName:    req.CustomerName,
		CustomerEmail:   req.CustomerEmail,
		ContactName:     req.ContactName,
		PhoneNumber:     req.PhoneNumber,
		CustomerAddress: req.CustomerAddress,
		CreatedByID:     req.CreatedByID,
		ChangedByID:     req.ChangedByID,
	}

	userReg, err := s.customerRepo.AddCourseCustomer(ctx, customer)
	if err != nil {
		return dto.CourseCustomerResponse{}, dto.ErrCreateUser
	}

	return dto.CourseCustomerResponse{
		ID:              userReg.ID.String(),
		CustomerName:    userReg.CustomerName,
		CustomerEmail:   userReg.CustomerEmail,
		ContactName:     userReg.ContactName,
		PhoneNumber:     userReg.PhoneNumber,
		CustomerAddress: userReg.CustomerAddress,
		CreatedByID:     userReg.CreatedByID,
		ChangedByID:     userReg.ChangedByID,
	}, nil
}

func (s *courseCustomerService) GetAllCourseCustomer(ctx context.Context, req dto.PaginationRequest) (dto.CourseCustomerPaginationResponse, error) {
	dataWithPaginate, err := s.customerRepo.GetAllCourseCustomer(ctx, req)

	if err != nil {
		return dto.CourseCustomerPaginationResponse{}, err
	}

	var responses []dto.CourseCustomerResponse
	for _, customer := range dataWithPaginate.CourseCustomer {
		response := dto.CourseCustomerResponse{
			ID:              customer.ID.String(),
			CustomerName:    customer.CustomerName,
			CustomerEmail:   customer.CustomerEmail,
			ContactName:     customer.ContactName,
			PhoneNumber:     customer.PhoneNumber,
			CustomerAddress: customer.CustomerAddress,
			CreatedByID:     customer.CreatedByID,
			ChangedByID:     customer.ChangedByID,
		}
		responses = append(responses, response)
	}
	return dto.CourseCustomerPaginationResponse{
		Data: responses,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *courseCustomerService) GetCourseCustomerById(ctx context.Context, id string) (dto.CourseCustomerResponse, error) {
	customer, err := s.customerRepo.GetCourseCustomerById(ctx, id)
	if err != nil {
		return dto.CourseCustomerResponse{}, err
	}

	return dto.CourseCustomerResponse{
		ID:              customer.ID.String(),
		CustomerName:    customer.CustomerName,
		CustomerEmail:   customer.CustomerEmail,
		ContactName:     customer.ContactName,
		PhoneNumber:     customer.PhoneNumber,
		CustomerAddress: customer.CustomerAddress,
		CreatedByID:     customer.CreatedByID,
		ChangedByID:     customer.ChangedByID,
	}, nil
}
func (s *courseCustomerService) UpdateCourseCustomer(ctx context.Context, req dto.CourseCustomerUpdateRequest, userId string) (dto.CourseCustomerResponse, error) {
	user, err := s.customerRepo.GetCourseCustomerById(ctx, userId)
	if err != nil {
		return dto.CourseCustomerResponse{}, dto.ErrUserNotFound
	}

	data := course.CourseCustomer{
		ID:              user.ID,
		CustomerName:    req.CustomerName,
		CustomerEmail:   req.CustomerEmail,
		ContactName:     req.ContactName,
		PhoneNumber:     req.PhoneNumber,
		CustomerAddress: req.CustomerAddress,
		CreatedByID:     req.CreatedByID,
		ChangedByID:     req.ChangedByID,
	}

	userUpdate, err := s.customerRepo.UpdateCourseCustomer(ctx, data)
	if err != nil {
		return dto.CourseCustomerResponse{}, dto.ErrUpdateUser
	}

	return dto.CourseCustomerResponse{
		ID:              userUpdate.ID.String(),
		CustomerName:    userUpdate.CustomerName,
		CustomerEmail:   userUpdate.CustomerEmail,
		ContactName:     userUpdate.ContactName,
		PhoneNumber:     userUpdate.PhoneNumber,
		CustomerAddress: userUpdate.CustomerAddress,
		CreatedByID:     userUpdate.CreatedByID,
		ChangedByID:     userUpdate.ChangedByID,
	}, nil
}

// Export to Excel functionality
func (s *courseCustomerService) ExportCourseCustomerToExcel(ctx context.Context, req dto.PaginationRequest) ([]byte, error) {
	// Get all course customers for export
	exportData, err := s.customerRepo.GetAllCourseCustomerForExport(ctx, req)
	if err != nil {
		return nil, err
	}

	// Create Excel file
	return s.createExcelFile(exportData)
}

// Helper method to create Excel file
func (s *courseCustomerService) createExcelFile(data []dto.CourseCustomerExport) ([]byte, error) {
	// This would use a library like "github.com/tealeg/xlsx" or "github.com/xuri/excelize/v2"
	// For now, we'll return a placeholder - you'll need to implement this with actual Excel library

	// Example using github.com/xuri/excelize/v2:

	f := excelize.NewFile()
	f.SetSheetName("Sheet1", "Sheet1")

	// Set headers
	headers := []string{"Customer Name", "Email", "Contact Name", "Phone Number", "Address", "Created At", "Updated At"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue("Sheet1", cell, header)
	}

	// Set data
	for i, record := range data {
		row := i + 2
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), record.CustomerName)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), record.CustomerEmail)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), record.ContactName)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), record.PhoneNumber)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), record.CustomerAddress)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), record.CreatedAt)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), record.UpdatedAt)
	}

	// Save to buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
