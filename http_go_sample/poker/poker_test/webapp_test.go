package poker

import (
	"http_go_sample/webserver"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var (
	dummyGame = &GameSpy{}
)

func TestGame(t *testing.T) {
	t.Run("Get /game returns 200", func(t *testing.T) {
		server, _ := webserver.NewPlayerServer(&webserver.StubPlayerStore{}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		webserver.AssertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		winner := "Ruth"
		passalong, _ := webserver.NewPlayerServer(dummyPlayerStore, dummyGame)
		server := httptest.NewServer(passalong)
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open a ws connection on %s %v", wsURL, err)
		}
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		assertStartedWith(t, dummyGame, 3)
		assertFinishedWith(t, dummyGame.FinishedWith, winner)
	})
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}
