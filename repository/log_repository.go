package repository

import (
	"context"

	"github.com/jejevj/go-aiocap/entity"
	"gorm.io/gorm"
)

type LogRepository interface {
	CreateLog(ctx context.Context, log *entity.SystemLog) error
	GetLogs(ctx context.Context, page, limit int) ([]entity.SystemLog, error)
	GetLogByID(ctx context.Context, id string) (*entity.SystemLog, error)
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{
		db: db,
	}
}

func (r *logRepository) CreateLog(ctx context.Context, log *entity.SystemLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *logRepository) GetLogs(ctx context.Context, page, limit int) ([]entity.SystemLog, error) {
	var logs []entity.SystemLog
	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&logs).Error
	return logs, err
}

func (r *logRepository) GetLogByID(ctx context.Context, id string) (*entity.SystemLog, error) {
	var log entity.SystemLog
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&log).Error
	return &log, err
}
