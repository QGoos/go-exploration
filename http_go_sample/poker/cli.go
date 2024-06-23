package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadPlayerWinnerErrMsg = "bad value recieved for winner, please enter {Name} wins"

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type CLI struct {
	In   *bufio.Scanner
	Out  io.Writer
	Game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		In:   bufio.NewScanner(in),
		Out:  out,
		Game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.Out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		fmt.Fprint(cli.Out, BadPlayerInputErrMsg)
		return
	}

	cli.Game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(cli.Out, BadPlayerWinnerErrMsg)
		return
	}

	cli.Game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadPlayerWinnerErrMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}
