package repositories

import (
	"github.com/corentings/kafejo-bot/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type KarmaRepository struct {
	DBConn *gorm.DB
}

func (KarmaRepository *KarmaRepository) CreateKarma(karma *models.Karma) error {
	// Create a new karma
	if err := KarmaRepository.DBConn.Create(&karma).Error; err != nil {
		return errors.Wrap(err, "error creating karma")
	}

	return nil
}

func (KarmaRepository *KarmaRepository) UpdateKarma(karma *models.Karma) error {
	// Update karma
	if err := KarmaRepository.DBConn.Save(&karma).Error; err != nil {
		return errors.Wrap(err, "error updating karma")
	}

	return nil
}

func (KarmaRepository *KarmaRepository) GetKarmaByGuildID(guildID string) ([]models.Karma, error) {
	// Get all karma from a guild
	var karma []models.Karma
	if err := KarmaRepository.DBConn.Where("guild_id = ?", guildID).Find(&karma).Error; err != nil {
		return nil, errors.Wrap(err, "error getting karma by guild id")
	}

	return karma, nil
}

func (KarmaRepository *KarmaRepository) GetKarmaByUserIDAndGuildID(userID, guildID string) (models.Karma, error) {
	var karma models.Karma
	if err := KarmaRepository.DBConn.Where("user_id = ? AND guild_id = ?", userID, guildID).Find(&karma).Error; err != nil {
		return models.Karma{}, errors.Wrap(err, "error getting karma by user id and guild id")
	}

	return karma, nil
}

func (KarmaRepository *KarmaRepository) GetTopKarmaByGuildID(guildID string) ([]models.Karma, error) {
	// Get top karma from a guild (10)
	var karma []models.Karma
	if err := KarmaRepository.DBConn.Where("guild_id = ?", guildID).Order("value desc").Limit(10).Find(&karma).Error; err != nil {
		return nil, errors.Wrap(err, "error getting top karma by guild id")
	}

	return karma, nil
}
