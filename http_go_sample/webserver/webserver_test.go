package webserver

import (
	"io"
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

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Darius": 20,
			"Ivern":  10,
		},
		nil,
		nil,
	}
	server, _ := NewPlayerServer(&store, dummyGame)
	t.Run("returns Darius's score", func(t *testing.T) {
		request := GetScoreRequest("Darius")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Ivern's score", func(t *testing.T) {
		request := GetScoreRequest("Ivern")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("player does not exist: 404s", func(t *testing.T) {
		request := GetScoreRequest("DNE")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusNotFound

		AssertStatus(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server, _ := NewPlayerServer(&store, dummyGame)
	t.Run("returns accepted on POST", func(t *testing.T) {
		player := "Darius"

		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)
		AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("returns 200 for /league on GET", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tara", 14},
		}
		store := StubPlayerStore{nil, nil, wantedLeague}
		server, _ := NewPlayerServer(&store, dummyGame)

		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)

		AssertStatus(t, response.Code, http.StatusOK)
		AsserContentTypeJson(t, response)
		AssertLeague(t, got, wantedLeague)
	})
}
