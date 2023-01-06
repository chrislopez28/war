package main

import (
	"fmt"
	"testing"

	"github.com/chrislopez28/cards"

	"github.com/google/go-cmp/cmp"
)

func TestBattle(t *testing.T) {
	p1 := Player{name: "P1", score: 0}
	p2 := Player{name: "P2", score: 0}

	g := CreateGame([]Player{p1, p2})

	topCardBefore := g.players[0].hand[len(g.players[0].hand)-1]

	g.Battle(false)
	g.Battle(false)
	g.Battle(false)

	topCardAfter := g.players[0].hand[len(g.players[0].hand)-1]

	fmt.Println(topCardBefore, topCardAfter)
	if cmp.Equal(topCardBefore, topCardAfter) {
		t.Errorf("Top card the same after battle")
	}

}

func TestCreateGame(t *testing.T) {
	p1 := Player{name: "P1", score: 0}
	p2 := Player{name: "P2", score: 0}

	g := CreateGame([]Player{p1, p2})

	for _, player := range g.players {
		fmt.Println(player.hand)
	}
}

func TestCompareCards(t *testing.T) {
	aceHearts := cards.Card{Suit: cards.Heart, Value: cards.Ace}
	twoHearts := cards.Card{Suit: cards.Heart, Value: cards.Two}
	aceSpades := cards.Card{Suit: cards.Spade, Value: cards.Ace}
	expectedResultOne := "P1 Wins"
	expectedResultTwo := "Tie"
	expectedResultThree := "P2 Wins"

	resultOne := CompareCards(aceHearts, twoHearts)
	resultTwo := CompareCards(aceHearts, aceSpades)
	resultThree := CompareCards(twoHearts, aceSpades)

	if expectedResultOne != resultOne {
		t.Errorf("Expected Player 1 to Win")
	}

	if expectedResultTwo != resultTwo {
		t.Errorf("Expected a Tie")
	}

	if expectedResultThree != resultThree {
		t.Errorf("Expected Player 2 to Win")
	}

}

func TestIsGameFinished(t *testing.T) {
	p1 := Player{name: "P1", score: 0}
	p2 := Player{name: "P2", score: 0}

	g := CreateGame([]Player{p1, p2})

	status := g.isGameFinished()
	expectedResult := false

	if expectedResult != status {
		t.Errorf("Expected game to be in progress")
	}

}
