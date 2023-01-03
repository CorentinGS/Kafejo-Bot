package services

import (
	"github.com/corentings/kafejo-bot/commands/karma"
	"github.com/corentings/kafejo-bot/data/cmdHandler"
	"github.com/corentings/kafejo-bot/data/infrastructures"
	"github.com/corentings/kafejo-bot/data/repositories"
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/corentings/kafejo-bot/models"
	"github.com/corentings/kafejo-bot/utils"
)

type KarmaService struct {
	interfaces.IKarmaRepository
}

func (k *kernel) InjectKarmaCommandHandler() karma.Command {
	DBConn := infrastructures.GetDBConn()
	karmaRepo := &repositories.KarmaRepository{DBConn: DBConn}
	karmaService := &KarmaService{IKarmaRepository: karmaRepo}
	karmaCommand := karma.Command{IKarmaService: karmaService, IHandler: cmdHandler.GetHandler()}
	return karmaCommand
}

func (KarmaService *KarmaService) CreateKarma(karma models.Karma) (models.Karma, error) {
	err := KarmaService.IKarmaRepository.CreateKarma(&karma)
	if err != nil {
		return models.Karma{}, err
	}
	return karma, nil
}

func (KarmaService *KarmaService) UpdateKarma(karma models.Karma) (models.Karma, error) {
	err := KarmaService.IKarmaRepository.UpdateKarma(&karma)
	if err != nil {
		return models.Karma{}, err
	}
	return karma, nil
}

func (KarmaService *KarmaService) GetKarma(userID, guildID string) (models.Karma, error) {
	userKarma, err := KarmaService.IKarmaRepository.GetKarmaByUserIDAndGuildID(userID, guildID)
	if err != nil || userKarma.ID == 0 {
		return models.Karma{}, utils.ErrGetKarma
	}

	return userKarma, nil
}

func (KarmaService *KarmaService) GetTopKarma(guildID string) ([]models.Karma, error) {
	topKarma, err := KarmaService.IKarmaRepository.GetTopKarmaByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return topKarma, nil
}

func (KarmaService *KarmaService) IncrementKarma(userID, guildID string) (models.Karma, error) {
	userKarma, err := KarmaService.GetKarma(userID, guildID)
	if err != nil {
		return models.Karma{}, err
	}

	if userKarma.ID == 0 {
		userKarma = models.Karma{UserID: userID, GuildID: guildID, Value: 1}
		_, err = KarmaService.CreateKarma(userKarma)
		if err != nil {
			return models.Karma{}, err
		}
	} else {
		userKarma.AddKarma(1)
		_, err = KarmaService.UpdateKarma(userKarma)
		if err != nil {
			return models.Karma{}, err
		}
	}

	return userKarma, nil
}
