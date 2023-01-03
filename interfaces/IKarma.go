package interfaces

import "github.com/corentings/kafejo-bot/models"

type IKarmaRepository interface {
	CreateKarma(karma *models.Karma) error
	UpdateKarma(karma *models.Karma) error
	GetKarmaByGuildID(guildID string) ([]models.Karma, error)
	GetKarmaByUserIDAndGuildID(userID, guildID string) (models.Karma, error)
	GetTopKarmaByGuildID(guildID string) ([]models.Karma, error)
}

type IKarmaService interface {
	CreateKarma(karma models.Karma) (models.Karma, error)
	UpdateKarma(karma models.Karma) (models.Karma, error)
	GetKarma(userID, guildID string) (models.Karma, error)
	GetTopKarma(guildID string) ([]models.Karma, error)
	IncrementKarma(userID, guildID string) (models.Karma, error)
}
