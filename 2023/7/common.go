package main

import (
	"bufio"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Solve(input string, joker bool) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	hands := make([]Hand, 0)
	for sc.Scan() {
		hand := NewHand(sc.Text(), joker)
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[j].Stronger(hands[i])
	})

	totalWinnings := 0
	for i := range hands {
		rank := i + 1
		bid := hands[i].Bid
		totalWinnings += rank * bid
	}

	return fmt.Sprint(totalWinnings)
}

type (
	Card     byte
	HandType int
)

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var PowerMap = map[byte]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

var PowerMapJoker = map[byte]int{
	'A': 14, 'K': 13, 'Q': 12, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, 'J': 1,
}

type Cards []Card

func (c Cards) AsMap() map[Card]int {
	m := make(map[Card]int)
	for _, card := range c {
		m[Card(card)] += 1
	}
	return m
}

func (c Cards) Stronger(other Cards, joker bool) bool {
	m := PowerMap
	if joker {
		m = PowerMapJoker
	}

	for i := 0; i < 5; i++ {
		left, right := m[byte(c[i])], m[byte(other[i])]
		if left > right {
			return true
		}
		if left < right {
			return false
		}
	}
	return false
}

type Hand struct {
	Cards Cards
	Type  HandType
	Bid   int
	Joker bool
}

func NewHand(s string, joker bool) Hand {
	split := strings.Split(s, " ") // "32T3K 765"

	// Parse cards
	cards := Cards{}
	for _, card := range split[0] {
		cards = append(cards, Card(byte(card)))
	}

	// Figure out the type
	handType := findHandType(cards.AsMap())

	// If the joker is in play, find the theoretical max power of the card by changing joker to every possible type of a card until we get the max power.
	if joker {
		maxPower := handType
		for k := range PowerMapJoker {
			theoreticalCards := Cards{}
			for _, card := range strings.ReplaceAll(split[0], "J", string(k)) {
				theoreticalCards = append(theoreticalCards, Card(byte(card)))
			}
			maxPower = max(maxPower, findHandType(theoreticalCards.AsMap()))
		}
		handType = maxPower
	}

	// Parse bid
	bid := MustParse(split[1])

	return Hand{
		Cards: cards,
		Type:  handType,
		Bid:   bid,
		Joker: joker,
	}
}

func findHandType(cards map[Card]int) HandType {
	unique := make([]Card, 0)
	maxRepeatingCards := 0
	for k, v := range cards {
		if !slices.Contains(unique, k) {
			unique = append(unique, k)
		}
		maxRepeatingCards = max(maxRepeatingCards, int(v))
	}

	switch {
	case maxRepeatingCards == 5 && len(unique) == 1: // Five of a kind, where all five cards have the same label: AAAAA
		return FiveOfAKind
	case maxRepeatingCards == 4 && len(unique) == 2: // Four of a kind, where four cards have the same label and one card has a different label
		return FourOfAKind
	case maxRepeatingCards == 3 && len(unique) == 2: // Full house, where three cards have the same label, and the remaining two cards share a different label
		return FullHouse
	case maxRepeatingCards == 3 && len(unique) == 3: // Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand
		return ThreeOfAKind
	case maxRepeatingCards == 2 && len(unique) == 3: // Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label
		return TwoPair
	case maxRepeatingCards == 2 && len(unique) == 4: // One pair, where two cards share one label, and the other three cards have a different label from the pair and each other
		return OnePair
	default: // High card, where all cards' labels are distinct
		return HighCard
	}
}

func (h Hand) Stronger(other Hand) bool {
	if h.Type > other.Type {
		return true
	}
	if h.Type < other.Type {
		return false
	}
	return h.Cards.Stronger(other.Cards, h.Joker)
}

func CopySlice[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
