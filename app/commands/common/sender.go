package common

import (
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rs/zerolog/log"
	"sync"
)

var (
	embedsChan = make(chan discord.Embed, 100)
)

func SendLogEmbed(embed discord.Embed) error {
	session := utils.GetSession()
	channel := GetLoggerChannel()
	_, err := session.SendMessage(channel, "", embed)
	return err
}

func SenderWorker(embeds <-chan discord.Embed, wg *sync.WaitGroup) {
	log.Debug().Msgf("Sender worker created")
	defer wg.Done()

	for {
		embed := <-embeds
		err := SendLogEmbed(embed)
		if err != nil {
			log.Error().Err(err).Msg("Error sending log embed")
		}
	}
}

func CreateSenderWorker() {
	var wg sync.WaitGroup

	defer close(embedsChan)
	wg.Add(1)

	go SenderWorker(embedsChan, &wg)

	wg.Wait()
}

func AddEmbedToQueue(embed discord.Embed) {
	embedsChan <- embed
}
