package main

import (
	"fmt"
)

var cardValue map[string]int = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

func main() {
	fmt.Println(FirstTurn("five", "queen", "ace"))
}

func ParseCard(card string) int {
	return cardValue[card]
}

func FirstTurn(card1, card2, dealerCard string) string {
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	valueDealer := ParseCard(dealerCard)
	sumCards := value1 + value2

	if sumCards == 22 {
		return "P"
	} else if sumCards == 21 {
		if valueDealer > 9 {
			return "S"
		} else {
			return "W"
		}
	} else if sumCards > 16 {
		return "S"
	} else if sumCards > 11 {
		if valueDealer > 6 {
			return "H"
		} else {
			return "S"
		}
	} else {
		return "H"
	}

}
