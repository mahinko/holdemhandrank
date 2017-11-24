package holdemhandrank

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
)

// Hand types
var handtypes = []string{
	"Invalid hand",
	"High Card",
	"One Pair",
	"Two Pairs",
	"Three of a Kind",
	"Straight",
	"Flush",
	"Full House",
	"Four of a Kind",
	"Straight Flush",
}

// Deck
var Cards = map[string]uint32{
	"2c": 1,
	"2d": 2,
	"2h": 3,
	"2s": 4,
	"3c": 5,
	"3d": 6,
	"3h": 7,
	"3s": 8,
	"4c": 9,
	"4d": 10,
	"4h": 11,
	"4s": 12,
	"5c": 13,
	"5d": 14,
	"5h": 15,
	"5s": 16,
	"6c": 17,
	"6d": 18,
	"6h": 19,
	"6s": 20,
	"7c": 21,
	"7d": 22,
	"7h": 23,
	"7s": 24,
	"8c": 25,
	"8d": 26,
	"8h": 27,
	"8s": 28,
	"9c": 29,
	"9d": 30,
	"9h": 31,
	"9s": 32,
	"tc": 33,
	"td": 34,
	"th": 35,
	"ts": 36,
	"jc": 37,
	"jd": 38,
	"jh": 39,
	"js": 40,
	"qc": 41,
	"qd": 42,
	"qh": 43,
	"qs": 44,
	"kc": 45,
	"kd": 46,
	"kh": 47,
	"ks": 48,
	"ac": 49,
	"ad": 50,
	"ah": 51,
	"as": 52,
}

// Holds lookup table from HandsRanks.dat
var ranks []byte

// Holds evaluation results
type HandRank struct {
	HandType int
	HandRank int
	Value    int
	HandName string
}

// Read lookup table into memory
func init() {

	// Read the ranks file
	var err error
	ranks, err = ioutil.ReadFile("HandRanks.dat")
	if err != nil {
		log.Fatal(err)
	}
}

// Evaluate hand
func Eval(cards []string) HandRank {

	var p uint32 = 53
	for i := 0; i < len(cards); i++ {
		p = evalCard(p + Cards[cards[i]])
	}

	return HandRank{
		HandType: int(p >> 12),
		HandRank: int(p & 0x00000fff),
		Value:    int(p),
		HandName: handtypes[p>>12],
	}
}



// Evaluate a card
func evalCard(card uint32) (e uint32) {

	// Create buffer on ranks
	buf := bytes.NewBuffer(ranks)

	// Skip until the card
	buf.Next(int(card) * 4)

	// Read the card data
	binary.Read(buf, binary.LittleEndian, &e)

	// Return result
	return
}
