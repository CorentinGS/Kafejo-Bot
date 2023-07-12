package interfaces

import "github.com/corentings/kafejo-bot/domain"

type IKarmaRepository interface {
	CreateKarma(karma *domain.Karma) error
	UpdateKarma(karma *domain.Karma) error
	GetKarmaByGuildID(guildID string) ([]domain.Karma, error)
	GetKarmaByUserIDAndGuildID(userID, guildID string) (domain.Karma, error)
	GetTopKarmaByGuildID(guildID string) ([]domain.Karma, error)
}

type IKarmaService interface {
	CreateKarma(karma domain.Karma) (domain.Karma, error)
	UpdateKarma(karma domain.Karma) (domain.Karma, error)
	GetKarma(userID, guildID string) (domain.Karma, error)
	GetTopKarma(guildID string) ([]domain.Karma, error)
	IncrementKarma(userID, guildID string) (domain.Karma, error)
}
