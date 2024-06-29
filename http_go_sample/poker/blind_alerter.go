package poker

import (
	"fmt"
	"io"
	"os"
	"time"
)

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
