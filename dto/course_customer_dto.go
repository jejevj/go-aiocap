package dto

import (
	"github.com/google/uuid"
	"github.com/jejevj/go-aiocap/entity/course"
)

type (
	CourseCustomerCreateRequest struct {
		CustomerName    string    `json:"customer_name"`
		CustomerEmail   string    `json:"customer_email"`
		ContactName     string    `json:"contact_name"`
		PhoneNumber     string    `json:"phone_number"`
		CustomerAddress string    `json:"customer_address"`
		CreatedByID     uuid.UUID `json:"created_by_id"`
		ChangedByID     uuid.UUID `json:"changed_by_id"`
	}
	CourseCustomerUpdateRequest struct {
		ID              uuid.UUID `json:"id"`
		CustomerName    string    `json:"customer_name"`
		CustomerEmail   string    `json:"customer_email"`
		ContactName     string    `json:"contact_name"`
		PhoneNumber     string    `json:"phone_number"`
		CustomerAddress string    `json:"customer_address"`
		CreatedByID     uuid.UUID `json:"created_by_id"`
		ChangedByID     uuid.UUID `json:"changed_by_id"`
	}

	CourseCustomerGetDetailsRequest struct {
		ID uuid.UUID `json:"id"`
	}

	CourseCustomerResponse struct {
		ID              string    `json:"id"`
		CustomerName    string    `json:"customer_name"`
		CustomerEmail   string    `json:"customer_email"`
		ContactName     string    `json:"contact_name"`
		PhoneNumber     string    `json:"phone_number"`
		CustomerAddress string    `json:"customer_address"`
		CreatedByID     uuid.UUID `json:"created_by_id"`
		ChangedByID     uuid.UUID `json:"changed_by_id"`
	}

	CourseCustomerPaginationResponse struct {
		Data []CourseCustomerResponse `json:"data"`
		PaginationResponse
	}

	GetAllCCResponse struct {
		CourseCustomer []course.CourseCustomer
		PaginationResponse
	}

	CourseCustomerExport struct {
		CustomerName    string `json:"customer_name"`
		CustomerEmail   string `json:"customer_email"`
		ContactName     string `json:"contact_name"`
		PhoneNumber     string `json:"phone_number"`
		CustomerAddress string `json:"customer_address"`
	}
)
