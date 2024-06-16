package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"http_go_sample/dummy"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := dummy.NewInMemoryPlayerStore[Player]()
	server := NewPlayerServer(store)
	player := "Darius"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, getScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Darius", 3},
		}
		assertLeague(t, got, want)
	})
}
