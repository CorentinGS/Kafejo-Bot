package karma

import (
	"context"
	"github.com/corentings/kafejo-bot/domain"
)

// IUseCase is the user use case interface.
type IUseCase interface {
	CreateKarma(ctx context.Context, karma domain.Karma) (domain.Karma, error)
	UpdateKarma(ctx context.Context, karma domain.Karma) (domain.Karma, error)
	GetKarma(ctx context.Context, userID, guildID string) (domain.Karma, error)
	GetTopKarma(ctx context.Context, guildID string) ([]domain.Karma, error)
	IncrementKarma(ctx context.Context, userID, guildID string) (domain.Karma, error)
}

// IRepository is the user repository interface.
type IRepository interface {
	CreateKarma(ctx context.Context, karma *domain.Karma) error
	UpdateKarma(ctx context.Context, karma *domain.Karma) error
	GetKarmaByGuildID(ctx context.Context, guildID string) ([]domain.Karma, error)
	GetKarmaByUserIDAndGuildID(ctx context.Context, userID, guildID string) (domain.Karma, error)
	GetTopKarmaByGuildID(ctx context.Context, guildID string) ([]domain.Karma, error)
}
