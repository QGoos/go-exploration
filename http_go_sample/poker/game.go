package poker

import (
	"http_go_sample/webserver"
	"io"
	"os"
	"time"
)

type TexasHoldem struct {
	Alerter BlindAlerter
	Store   webserver.PlayerStore
}

func NewGame(alerter BlindAlerter, store webserver.PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		Alerter: alerter,
		Store:   store,
	}
}

func (p *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.Alerter.ScheduleAlertAt(blindTime, blind, os.Stdout)
		blindTime = blindTime + blindIncrement
	}
}

func (p *TexasHoldem) Finish(winner string) {
	p.Store.RecordWin(winner)
}
