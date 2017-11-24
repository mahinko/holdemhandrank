package holdemhandrank

import (
	"testing"
)

func TestHighCard(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "2d", "4h", "js", "3c", "7h", "9s"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "High Card" {
		t.Error("High Card isn't recognized.")
	}
}


func TestOnePair(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "2d", "2h", "js", "3c", "5h", "8s"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "One Pair" {
		t.Error("One Pair isn't recognized.")
	}
}

func TestTwoPairs(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "2d", "2h", "js", "3c", "5h", "as"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "Two Pairs" {
		t.Error("Two Pairs isn't recognized.")
	}
}

func TestThreeOfKind(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "as", "ah", "2s", "3c", "5h", "ts"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "Three of a Kind" {
		t.Error("Three of Kind isn't recognized.")
	}
}

func TestStraight(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "ks", "qh", "js", "3c", "5h", "ts"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "Straight" {
		t.Error("Straight isn't recognized.")
	}
}

func TestFlush(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "2d", "4d", "jd", "3d", "7h", "9s"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "Flush" {
		t.Error("Flush isn't recognized.")
	}
}

func TestFullHouse(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "2d", "2h", "2s", "3c", "5h", "as"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "Full House" {
		t.Error("Full House isn't recognized.")
	}
}

func TestFourOfKind(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "as", "ah", "ac", "3c", "5h", "ts"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "Four of a Kind" {
		t.Error("Four of Kind isn't recognized.")
	}
}

func TestStraightFlush(t *testing.T) {

	// Set a hand
	hand := []string{"ad", "kd", "qd", "jd", "td", "7h", "9s"}

	// Evaluate the hand
	handRank := Eval(hand)

	// Verify result
	if handRank.HandName != "Straight Flush" {
		t.Error("Straight isn't recognized.")
	}
}