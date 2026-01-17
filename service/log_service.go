package service

import (
	"context"

	"github.com/jejevj/go-aiocap/entity"
	"github.com/jejevj/go-aiocap/repository"
)

type LogService interface {
	LogAction(ctx context.Context, action, endpoint, method, userID, request, response string) error
	GetLogs(ctx context.Context, page, limit int) ([]entity.SystemLog, error)
	GetLogByID(ctx context.Context, id string) (*entity.SystemLog, error)
}

type logService struct {
	logRepository repository.LogRepository
}

func NewLogService(lr repository.LogRepository) LogService {
	return &logService{
		logRepository: lr,
	}
}

func (s *logService) LogAction(ctx context.Context, action, endpoint, method, userID, request, response string) error {
	log := &entity.SystemLog{
		Action:   action,
		Endpoint: endpoint,
		Method:   method,
		UserID:   &userID,
		Request:  request,
		Response: response,
	}
	return s.logRepository.CreateLog(ctx, log)
}

func (s *logService) GetLogs(ctx context.Context, page, limit int) ([]entity.SystemLog, error) {
	return s.logRepository.GetLogs(ctx, page, limit)
}

func (s *logService) GetLogByID(ctx context.Context, id string) (*entity.SystemLog, error) {
	return s.logRepository.GetLogByID(ctx, id)
}
