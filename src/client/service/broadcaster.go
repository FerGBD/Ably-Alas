package service

import (
	"ably-go-quickstart/client/models"
	"context"
	"fmt"

	"github.com/ably/ably-go/ably"
)

type FootballService struct {
	Channel *ably.RealtimeChannel
}

func (s *FootballService) PublishLance(tipo, desc string, minuto int) error {
	lance := models.Lance{
		Tipo:      tipo,
		Descricao: desc,
		Minuto:    minuto,
		Autor:     "Jornalista-Go",
	}

	return s.Channel.Publish(context.Background(), "evento_jogo", lance)
}

func (s *FootballService) SubscribeLances() {
	s.Channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		fmt.Printf("\n [%s] %s", msg.Name, msg.Data)
	})
}
