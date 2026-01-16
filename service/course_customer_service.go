package service

import (
	"context"

	"github.com/jejevj/go-aiocap/dto"
	"github.com/jejevj/go-aiocap/entity/course"
	"github.com/jejevj/go-aiocap/repository"
)

type (
	CourseCustomerService interface {
		AddCourseCustomer(ctx context.Context, req dto.CourseCustomerCreateRequest) (dto.CourseCustomerResponse, error)
		GetAllCourseCustomer(ctx context.Context) ([]dto.CourseCustomerResponse, error)
		GetCourseCustomerById(ctx context.Context, id string) (dto.CourseCustomerResponse, error)
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

func (s *courseCustomerService) GetAllCourseCustomer(ctx context.Context) ([]dto.CourseCustomerResponse, error) {
	customers, err := s.customerRepo.GetAllCourseCustomer(ctx)
	if err != nil {
		return []dto.CourseCustomerResponse{}, err
	}

	var responses []dto.CourseCustomerResponse
	for _, customer := range customers {
		responses = append(responses, dto.CourseCustomerResponse{
			ID:              customer.ID.String(),
			CustomerName:    customer.CustomerName,
			CustomerEmail:   customer.CustomerEmail,
			ContactName:     customer.ContactName,
			PhoneNumber:     customer.PhoneNumber,
			CustomerAddress: customer.CustomerAddress,
			CreatedByID:     customer.CreatedByID,
			ChangedByID:     customer.ChangedByID,
		})
	}

	return responses, nil
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
