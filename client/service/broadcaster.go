package service

import (
	"ably-go-quickstart/client/models"
	"context"
	"encoding/json"
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

	// Serializar para JSON antes de enviar
	jsonData, err := json.Marshal(lance)
	if err != nil {
		return fmt.Errorf("erro ao serializar lance: %v", err)
	}

	return s.Channel.Publish(context.Background(), "lance", string(jsonData))
}

func (s *FootballService) SubscribeLances() {
	s.Channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		fmt.Printf("\n✅ [%s] %s\n", msg.Name, msg.Data)
	})
}
