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

func (g *Game) Battle() error {
	var c1, c2, f1, f2 cards.Card
	var err error
	var playerIndex int

	fmt.Printf("-- Battle #%v --\n", g.numBattles)

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
			fmt.Println("War!")

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
	g.printScores()
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

	fmt.Printf("%s: (Player1: %v vs. Player2: %v)\n", result, c1.String(), c2.String())

	return result
}

// Main -----------------------------------------------------------------------

func main() {
	p1 := Player{name: "P1", score: 0, hand: []cards.Card{}}
	p2 := Player{name: "P2", score: 0, hand: []cards.Card{}}

	game := CreateGame([]Player{p1, p2})

	for _, player := range game.players {
		fmt.Println(player.hand)
	}

	game.printScores()

	fmt.Println(game.numBattles)

	for !game.isGameFinished() {
		game.Battle()

		if game.numBattles > 50000 {
			fmt.Println("Game aborted. No winner after 50000 battles.")
			break
		}
	}

}
