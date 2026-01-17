package service

import (
	"context"

	"github.com/jejevj/go-aiocap/dto"
	"github.com/jejevj/go-aiocap/entity/course"
	"github.com/jejevj/go-aiocap/repository"
)

type (
	CourseService interface {
		AddCourse(ctx context.Context, req dto.CourseCreateRequest) (dto.CourseResponse, error)
	}

	courseService struct {
		customerRepo repository.CourseRepository
		cstRepo      repository.CourseCustomerRepository // Fixed: renamed to avoid confusion
		jwtService   JWTService
	}
)

func NewCourseService(customerRepo repository.CourseRepository, cstRepo repository.CourseCustomerRepository, jwtService JWTService) CourseService {
	return &courseService{
		customerRepo: customerRepo,
		cstRepo:      cstRepo,
		jwtService:   jwtService,
	}
}

func (s *courseService) AddCourse(ctx context.Context, req dto.CourseCreateRequest) (dto.CourseResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	customer := course.Course{
		CourseName:        req.CourseName,
		CourseDescription: req.CourseDescription,
		CourseLocation:    req.CourseLocatoin,
		CourseClient:      req.CourseClient,
		IsVerified:        req.IsVerified,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		CreatedByID:       req.CreatedByID,
		ChangedByID:       req.ChangedByID,
	}
	// Fixed: Convert uuid.UUID to string for the ID parameter
	clientDetails, err := s.cstRepo.GetCourseCustomerById(ctx, req.CourseClient.String()) // Fixed: removed string() conversion
	if err != nil {
		return dto.CourseResponse{}, dto.ErrCreateCourseClientNotFound
	}

	userReg, err := s.customerRepo.AddCourse(ctx, customer)
	if err != nil {
		return dto.CourseResponse{}, dto.ErrCreateUser
	}

	// Fixed: Convert course.CourseCustomer to dto.CourseCustomerResponse
	return dto.CourseResponse{
		ID:                userReg.ID.String(),
		CourseName:        userReg.CourseName,
		CourseDescription: userReg.CourseDescription,
		CourseLocatoin:    userReg.CourseLocation,
		CourseClient:      userReg.CourseClient, // Fixed: convert UUID to string for response
		CourseCustomer: dto.CourseCustomerResponse{ // Fixed: convert to DTO type
			ID:              clientDetails.ID.String(),
			CustomerName:    clientDetails.CustomerName,
			CustomerEmail:   clientDetails.CustomerEmail,
			ContactName:     clientDetails.ContactName,
			PhoneNumber:     clientDetails.PhoneNumber,
			CustomerAddress: clientDetails.CustomerAddress,
			CreatedByID:     clientDetails.CreatedByID,
			ChangedByID:     clientDetails.ChangedByID,
		},
		IsVerified:  userReg.IsVerified,
		StartDate:   userReg.StartDate,
		EndDate:     userReg.EndDate,
		CreatedByID: userReg.CreatedByID,
		ChangedByID: userReg.ChangedByID,
	}, nil
}
