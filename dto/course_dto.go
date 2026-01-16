package dto

import (
	"github.com/google/uuid"
)

type (
	CourseCreateRequest struct {
		CourseName        string    `json:"course_name"`
		CourseDescription string    `json:"course_description"`
		CourseLocatoin    string    `json:"course_location"`
		CourseClient      uuid.UUID `json:"course_client"`
		IsVerified        bool      `json:"is_verivied"`
		CreatedByID       uuid.UUID `json:"created_by_id"`
		ChangedByID       uuid.UUID `json:"changed_by_id"`
	}

	CourseResponse struct {
		ID                string                 `json:"id"`
		CourseName        string                 `json:"course_name"`
		CourseDescription string                 `json:"course_description"`
		CourseLocatoin    string                 `json:"course_location"`
		CourseClient      uuid.UUID              `json:"course_client"`
		CourseCustomer    CourseCustomerResponse `json:"client_data"`
		IsVerified        bool                   `json:"is_verivied"`
		CreatedByID       uuid.UUID              `json:"created_by_id"`
		ChangedByID       uuid.UUID              `json:"changed_by_id"`
	}

	CoursePaginationResponse struct {
		Data []CourseResponse `json:"data"`
		PaginationResponse
	}
)
