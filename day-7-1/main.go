package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

var CARD_STRENGTH = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

var HAND_STRENGTH = map[string]int{
	"11111": 1,
	"2111":  2,
	"221":   3,
	"311":   4,
	"32":    5,
	"41":    6,
	"5":     7,
}

type Card struct {
	Key   string
	Value int
}

type CardList []Card

func (p CardList) Len() int { return len(p) }
func (p CardList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return CARD_STRENGTH[p[i].Key] < CARD_STRENGTH[p[j].Key]
	}

	return p[i].Value < p[j].Value
}
func (p CardList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p CardList) Type() string {
	var str string
	for _, card := range p {
		stringValue := strconv.Itoa(card.Value)
		str += stringValue
	}
	return str
}
func (p CardList) KeyType() string {
	var str string
	for _, card := range p {
		str += card.Key
	}
	return str
}

func rankByCount(cardFrequencies map[string]int) CardList {
	pl := make(CardList, len(cardFrequencies))
	i := 0
	for k, v := range cardFrequencies {
		pl[i] = Card{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Hand struct {
	Cards  CardList
	Bid    int
    String string
}

type HandList []Hand

func (hands HandList) Len() int { return len(hands) }
func (hands HandList) Less(i, j int) bool {
	if HAND_STRENGTH[hands[i].Cards.Type()] == HAND_STRENGTH[hands[j].Cards.Type()] {
		for k := 0; k < 5; k++ {
			if CARD_STRENGTH[string(hands[i].String[k])] != CARD_STRENGTH[string(hands[j].String[k])] {
				return CARD_STRENGTH[string(hands[i].String[k])] < CARD_STRENGTH[string(hands[j].String[k])]
			}
		}
		return HAND_STRENGTH[hands[i].Cards.Type()] < HAND_STRENGTH[hands[j].Cards.Type()]
	}

	return HAND_STRENGTH[hands[i].Cards.Type()] < HAND_STRENGTH[hands[j].Cards.Type()]
}
func (hands HandList) Swap(i, j int) { hands[i], hands[j] = hands[j], hands[i] }

func rankByType(hands HandList) HandList {
	sort.Sort(hands)
	return hands
}

func main() {
	text := utils.ReadFile("./input.txt")

	hands := make(HandList, len(text))

	for i, textLine := range text {
		data := strings.Split(textLine, " ")
		hand := data[0]
		bid, _ := strconv.Atoi(data[1])

		cardCount := make(map[string]int)
		for _, card := range hand {
			_, exists := cardCount[string(card)]
			if exists {
				cardCount[string(card)] += 1
			} else {
				cardCount[string(card)] = 1
			}
		}
		rankedCardCound := rankByCount(cardCount)
		rankedHand := Hand{Cards: rankedCardCound, Bid: bid, String: hand}
		hands[i] = rankedHand
	}
	rankedType := rankByType(hands)

	result := 0

	for count, hand := range rankedType {
        result += hand.Bid * (count + 1)
		fmt.Println(hand.String, "\t", count+1, "\tX\t", hand.Bid, "\t", hand.Cards.Type(), "\t", result)
	}

	fmt.Println(result)
}
