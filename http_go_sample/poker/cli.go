package poker

import (
	"bufio"
	"http_go_sample/webserver"
	"io"
	"strings"
)

type CLI struct {
	PlayerStore webserver.PlayerStore
	In          *bufio.Scanner
}

func NewCLI(store webserver.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.PlayerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
