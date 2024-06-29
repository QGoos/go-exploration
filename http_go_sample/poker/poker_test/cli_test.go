package poker

import (
	"bytes"
	"http_go_sample/poker"
	"http_go_sample/webserver"
	"io"
	"strings"
	"testing"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &webserver.StubPlayerStore{}

// var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int, alertsDestination io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		game := &GameSpy{}
		in := userSends("7", "Chris wins")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		wantPrompt := poker.PlayerPrompt

		assertStartedWith(t, game, 7)
		assertMessageSentToUser(t, stdout, wantPrompt)
	})
	t.Run("finish game with 'Chris' as winner", func(t *testing.T) {
		in := userSends("7", "Chris wins")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertStartedWith(t, game, 7)
		assertFinishedWith(t, game.FinishedWith, "Chris")
	})
	t.Run("record 'Cleo' win from user input", func(t *testing.T) {
		in := userSends("7", "Cleo wins")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertStartedWith(t, game, 7)
		assertFinishedWith(t, game.FinishedWith, "Cleo")
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertNotStarted(t, game)

		wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrMsg
		assertMessageSentToUser(t, stdout, wantPrompt)
	})
	t.Run("it handles the case where the user is snarky and puts something other than wins", func(t *testing.T) {
		in := userSends("7", "Chris says something snarky")
		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		cli := poker.NewCLI(in, stdout, game)

		expectedError := poker.BadPlayerWinnerErrMsg

		cli.PlayPoker()
		assertStarted(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, expectedError)
	})
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertNotStarted(t testing.TB, game *GameSpy) {
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertStarted(t testing.TB, game *GameSpy) {
	if !game.StartCalled {
		t.Errorf("game should have started")
	}
}

func assertFinishedWith(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected finish called with %q but got %q", want, got)
	}
}

func assertStartedWith(t testing.TB, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()
	if game.StartedWith != numberOfPlayersWanted {
		t.Errorf("expected finish called with %q but got %q", numberOfPlayersWanted, game.StartedWith)
	}
}
