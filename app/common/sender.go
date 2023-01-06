package common

import (
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rs/zerolog/log"
	"sync"
)

type MessageItem struct {
	Embed   discord.Embed
	Channel discord.ChannelID
	Content string
}

var (
	embedsChan = make(chan MessageItem, 100)
)

func SendLogEmbed(embed MessageItem) error {
	session := utils.GetSession()
	_, err := session.SendMessage(embed.Channel, embed.Content, embed.Embed)
	return err
}

func SenderWorker(embeds <-chan MessageItem, wg *sync.WaitGroup) {
	log.Debug().Msgf("Sender worker created")
	defer wg.Done()

	for {
		embed := <-embeds
		err := SendLogEmbed(embed)
		if err != nil {
			log.Error().Err(err).Msg("Error sending embed")
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

func AddEmbedToQueue(embed MessageItem) {
	embedsChan <- embed
}
