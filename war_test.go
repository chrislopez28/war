package main

import (
	"cards"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBattle(t *testing.T) {
	g := CreateGame()

	topCardBefore := g.players[0].hand[len(g.players[0].hand)-1]

	g.Battle()
	g.Battle()
	g.Battle()

	topCardAfter := g.players[0].hand[len(g.players[0].hand)-1]

	fmt.Println(topCardBefore, topCardAfter)
	if cmp.Equal(topCardBefore, topCardAfter) {
		t.Errorf("Top card the same after battle")
	}

}

func TestCreateGame(t *testing.T) {
	g := CreateGame()

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
	g := CreateGame()

	status := g.isGameFinished()
	expectedResult := false

	if expectedResult != status {
		t.Errorf("Expected game to be in progress")
	}

}
