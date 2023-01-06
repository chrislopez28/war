package main

import (
	"fmt"

	"github.com/chrislopez28/cards"
)

var cardValue = map[cards.CardValue]int{
	cards.Two:   2,
	cards.Three: 3,
	cards.Four:  4,
	cards.Five:  5,
	cards.Six:   6,
	cards.Seven: 7,
	cards.Eight: 8,
	cards.Nine:  9,
	cards.Ten:   10,
	cards.Jack:  11,
	cards.Queen: 12,
	cards.King:  13,
	cards.Ace:   14,
}

// Player ---------------------------------------------------------------------

type Player struct {
	name  string
	score int
	hand  []cards.Card
}

func (p Player) Name() string {
	return p.name
}

// Game -----------------------------------------------------------------------

type Game struct {
	players    []Player
	numBattles int64
}

func CreateGame(players []Player) Game {
	d := cards.LoadDeck()
	d.Shuffle()

	for i := range players {
		players[i].hand = d.DealCards(26)
	}

	g := Game{players: players}
	g.updateScores()
	g.numBattles = 0

	return g
}

func (g *Game) updateScores() {
	for i, player := range g.players {
		g.players[i].score = len(player.hand)
	}
}

func (g *Game) isGameFinished() bool {
	game := *g

	count := 0

	for _, player := range game.players {
		if !cards.IsCardStackEmpty(player.hand) {
			count++
		}
	}

	return count <= 1
}

func (g *Game) Battle(printMessages bool) error {
	var c1, c2, f1, f2 cards.Card
	var err error
	var playerIndex int

	if printMessages {
		fmt.Printf("-- Battle #%v --\n", g.numBattles)
	}

	c1, g.players[0].hand, err = cards.TakeCard(g.players[0].hand)

	if err != nil {
		return err
	}

	c2, g.players[1].hand, err = cards.TakeCard(g.players[1].hand)

	if err != nil {
		return err
	}

	result := CompareCards(c1, c2)

	usedCards := []cards.Card{c1, c2}

	switch result {
	case "P1 Wins":
		playerIndex = 0
	case "P2 Wins":
		playerIndex = 1
	case "Tie":
		for result == "Tie" {
			if printMessages {
				fmt.Println("War!")
			}

			f1, g.players[0].hand, err = cards.TakeCard(g.players[0].hand)

			if err != nil {
				return err
			}

			f2, g.players[1].hand, err = cards.TakeCard(g.players[1].hand)

			if err != nil {
				return err
			}

			usedCards = append(usedCards, f1, f2)

			c1, g.players[0].hand, err = cards.TakeCard(g.players[0].hand)

			if err != nil {
				return err
			}

			c2, g.players[1].hand, err = cards.TakeCard(g.players[1].hand)

			if err != nil {
				return err
			}

			result = CompareCards(c1, c2)
			usedCards = append(usedCards, c1, c2)
		}

		if result == "P1 Wins" {
			playerIndex = 0
		}
		if result == "P2 Wins" {
			playerIndex = 1
		}
	}

	usedCards = cards.Shuffle(usedCards)
	for _, card := range usedCards {
		g.players[playerIndex].hand, err = cards.InsertCardBottom(card, g.players[playerIndex].hand)
		if err != nil {
			return err
		}
	}

	g.updateScores()
	if printMessages {
		g.printScores()
	}
	g.incrementBattleCount()

	return nil
}

func (g *Game) incrementBattleCount() {
	g.numBattles = g.numBattles + 1
}

func (g *Game) printScores() {
	fmt.Println("*** Scores ***")
	for _, player := range g.players {
		fmt.Printf("%v:%v  ", player.Name(), player.score)
	}
	fmt.Println()
}

// Helper Functions -----------------------------------------------------------

func CompareCards(c1 cards.Card, c2 cards.Card) string {
	val1 := cardValue[c1.Value]
	val2 := cardValue[c2.Value]
	result := ""

	if val1 > val2 {
		result = "P1 Wins"
	}
	if val1 < val2 {
		result = "P2 Wins"
	}
	if val1 == val2 {
		result = "Tie"
	}

	if false {
		fmt.Printf("%s: (Player1: %v vs. Player2: %v)\n", result, c1.String(), c2.String())
	}

	return result
}

func runAutomatedTwoPlayerGame() string {
	p1 := Player{name: "P1", score: 0, hand: []cards.Card{}}
	p2 := Player{name: "P2", score: 0, hand: []cards.Card{}}

	game := CreateGame([]Player{p1, p2})

	for !game.isGameFinished() {
		game.Battle(false)

		if game.numBattles > 10000 {
			fmt.Println("Game aborted. No winner after 10000 battles.")
			return "Aborted"
		}
	}

	// fmt.Printf("P1 %v, P2 %v\n", game.players[0].score, game.players[1].score)

	if game.players[0].score == game.players[1].score {
		return "Tie"
	}

	if game.players[0].score > game.players[1].score {
		return "P1 wins"
	}

	return "P2 wins"
}

func runSimulations(n int) {
	unfinishedCount, p1wins, p2wins := 0, 0, 0
	var result string

	fmt.Println("Running...")

	for i := 0; i < n; i++ {
		// fmt.Printf("Running game %v\n", i+1)
		result = runAutomatedTwoPlayerGame()
		switch result {
		case "Aborted":
			unfinishedCount++
		case "P1 wins":
			p1wins++
		case "P2 wins":
			p2wins++
		default:
		}
	}

	fmt.Printf("Ran %v games. P1 won %v games, P2 won %v games, %v games did not complete.\n", n, p1wins, p2wins, unfinishedCount)
}

// Main -----------------------------------------------------------------------

func main() {
	runSimulations(1000)
}
