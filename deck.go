package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create type deck
//which is slice of string
type deck []string

//here deck is type of return
//here we r not adding receiver because we are creating new cards...no need existing things
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	//if u don't want to use index of for loop u can use _ to tell go to ignore
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//(d deck) ==> receiver ==> now any var of type [deck] gets access to print()
func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//(d deck) ==> receiver ==> now any var of type [deck] gets access to print()
//[d] is the instance of [cards] which is declared in main.go
// func (cards deck) Print() {
// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}
// }

//we need deck type
//we should call the function with 1st arugment of type [deck/string] , should be in order
//second () => we are returning two value type of deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// to convert type from deck to []string
func (d deck) toString() string {
	//to convert type of []string
	// []string(d)

	//	//to convert type []string to string
	return strings.Join([]string(d), ",")
}

//saving file to hard drive
func (d deck) saveToFile(filename string) error {
	//[0666] is a permission code, here is anyone can read and write
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//option 1 => log the error and return a call to newDeck()
		//option 2 => log the error and entirely quit the program

		fmt.Println("Error :: ", err)
		os.Exit(1)
	}

	//	//to convert type string to []string then type deck
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	//[time.Now().UnixNano()] =>every time runs the program, we r taking current time and returning int64

	//random num generation
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		//sedo random generation
		// newPosition := rand.Intn(len(d) - 1) ==> old code, which we didn't get exact suffle
		newPosition := r.Intn(len(d) - 1)

		//swapping / randomization
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
