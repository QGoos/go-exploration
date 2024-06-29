package dummy

import (
	"http_go_sample/webserver"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

var (
	dummyGame = &GameSpy{}
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, closer := webserver.CreateTempFile(t, `[]`)
	defer closer()
	store, err := webserver.NewFileSystemPlayerStore(database) //NewInMemoryPlayerStore[webserver.League]()

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server, _ := webserver.NewPlayerServer(store, dummyGame)
	player := "Darius"

	server.ServeHTTP(httptest.NewRecorder(), webserver.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), webserver.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), webserver.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, webserver.GetScoreRequest(player))
		webserver.AssertStatus(t, response.Code, http.StatusOK)

		webserver.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, webserver.NewLeagueRequest())
		webserver.AssertStatus(t, response.Code, http.StatusOK)

		got := webserver.GetLeagueFromResponse(t, response.Body)
		want := []webserver.Player{
			{Name: "Darius", Wins: 3},
		}
		webserver.AssertLeague(t, got, want)
	})
}
