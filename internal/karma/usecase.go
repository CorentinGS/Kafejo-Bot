package karma

import (
	"context"
	"github.com/corentings/kafejo-bot/domain"
	"github.com/corentings/kafejo-bot/utils"
)

type UseCase struct {
	IRepository
}

func NewUseCase(repo IRepository) IUseCase {
	return &UseCase{
		IRepository: repo,
	}
}

func (u UseCase) CreateKarma(ctx context.Context, karma domain.Karma) (domain.Karma, error) {
	err := u.IRepository.CreateKarma(ctx, &karma)
	if err != nil {
		return domain.Karma{}, err
	}
	return karma, nil
}

func (u UseCase) UpdateKarma(ctx context.Context, karma domain.Karma) (domain.Karma, error) {
	err := u.IRepository.UpdateKarma(ctx, &karma)
	if err != nil {
		return domain.Karma{}, err
	}
	return karma, nil
}

func (u UseCase) GetKarma(ctx context.Context, userID, guildID string) (domain.Karma, error) {
	userKarma, err := u.IRepository.GetKarmaByUserIDAndGuildID(ctx, userID, guildID)
	if err != nil || userKarma.ID == 0 {
		return domain.Karma{}, utils.ErrGetKarma
	}

	return userKarma, nil
}

func (u UseCase) GetTopKarma(ctx context.Context, guildID string) ([]domain.Karma, error) {
	topKarma, err := u.IRepository.GetTopKarmaByGuildID(ctx, guildID)
	if err != nil {
		return nil, err
	}

	return topKarma, nil
}

func (u UseCase) IncrementKarma(ctx context.Context, userID, guildID string) (domain.Karma, error) {
	userKarma, _ := u.GetKarma(ctx, userID, guildID)

	if userKarma.ID == 0 {
		userKarma = domain.Karma{UserID: userID, GuildID: guildID, Value: 1}
		_, err := u.CreateKarma(ctx, userKarma)
		if err != nil {
			return domain.Karma{}, err
		}
	} else {
		userKarma.AddKarma(1)
		_, err := u.UpdateKarma(ctx, userKarma)
		if err != nil {
			return domain.Karma{}, err
		}
	}

	return userKarma, nil
}
