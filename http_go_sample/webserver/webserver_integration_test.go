package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"http_go_sample/dummy"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := dummy.NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Darius"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, getScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
