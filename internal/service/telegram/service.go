package telegram

import (
	"context"
	"fmt"
	"time"

	tele "gopkg.in/telebot.v3"

	"github.com/BigDwarf/sahtian/internal/service/users"
)

type Service struct {
	Bot          *tele.Bot
	appUrl       string
	usersService *users.Service
}

type Option func(*Service)

func NewService(token, appUrl string, options ...Option) (*Service, error) {
	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	err = b.SetMyDescription(`Sahtian`, "en")
	if err != nil {
		return nil, err
	}

	startCmd := tele.Command{
		Text:        "start",
		Description: "Start SahtianShop",
	}
	err = b.SetCommands([]tele.Command{startCmd})
	if err != nil {
		return nil, err
	}

	s := &Service{
		Bot:    b,
		appUrl: appUrl,
	}

	for _, option := range options {
		option(s)
	}

	return s, nil
}

func (s *Service) Start() {
	s.Bot.Start()
}

func (s *Service) Stop() {
	s.Bot.Stop()
}

func (s *Service) SendMessage(ctx context.Context, userId int64, text string) error {
	recipient, err := s.Bot.ChatByID(userId)
	if err != nil {
		return err
	}

	_, err = s.Bot.Send(recipient, text)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
