package dto

import (
	"github.com/google/uuid"
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
)
