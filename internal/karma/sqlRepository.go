package karma

import (
	"context"

	"github.com/corentings/kafejo-bot/domain"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SQLRepository struct {
	DBConn *gorm.DB
}

func NewSQLRepository(dbConn *gorm.DB) IRepository {
	return SQLRepository{DBConn: dbConn}
}

func (r SQLRepository) CreateKarma(ctx context.Context, karma *domain.Karma) error {
	// Create a new karma
	if err := r.DBConn.WithContext(ctx).Create(&karma).Error; err != nil {
		return errors.Wrap(err, "error creating karma")
	}

	return nil
}

func (r SQLRepository) UpdateKarma(ctx context.Context, karma *domain.Karma) error {
	// Update karma
	if err := r.DBConn.WithContext(ctx).Save(&karma).Error; err != nil {
		return errors.Wrap(err, "error updating karma")
	}

	return nil
}

func (r SQLRepository) GetKarmaByGuildID(ctx context.Context, guildID string) ([]domain.Karma, error) {
	// Get all karma from a guild
	var karma []domain.Karma
	if err := r.DBConn.WithContext(ctx).Where("guild_id = ?", guildID).Find(&karma).Error; err != nil {
		return nil, errors.Wrap(err, "error getting karma by guild id")
	}

	return karma, nil
}

func (r SQLRepository) GetKarmaByUserIDAndGuildID(ctx context.Context, userID, guildID string) (domain.Karma, error) {
	var karma domain.Karma
	if err := r.DBConn.WithContext(ctx).Where("user_id = ? AND guild_id = ?", userID, guildID).Find(&karma).Error; err != nil {
		return domain.Karma{}, errors.Wrap(err, "error getting karma by user id and guild id")
	}

	return karma, nil
}

func (r SQLRepository) GetTopKarmaByGuildID(ctx context.Context, guildID string) ([]domain.Karma, error) {
	// Get top karma from a guild (10)
	var karma []domain.Karma
	if err := r.DBConn.WithContext(ctx).Where("guild_id = ?", guildID).Order("value desc").Limit(10).Find(&karma).Error; err != nil {
		return nil, errors.Wrap(err, "error getting top karma by guild id")
	}

	return karma, nil
}
