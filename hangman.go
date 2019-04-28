package main

import (
	"bufio"
	"fmt"
	"github.com/tjarratt/babble"
	"os"
	"strings"
	"unicode"
)

var word string
var guessedWord string
var correctGuesses []string
var wrongGuesses []string

var reader = bufio.NewReader(os.Stdin)

const HiddenLetter string = "_"

func main() {
	InitializeGame()
	displayIntroduction()

	for done := false; done == false; done = TakeTurn() {}

	os.Exit(0)
}

func InitializeGame() {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	word = babbler.Babble()
	regenerateGuessedWord()
}

func TakeTurn() bool {
	var guess string
	for {
		guess = readInGuess()
		if len(guess) < 1 {
			fmt.Println("Did you enter any letters? Please enter a letter as guess!")
		} else if len(guess) > 1 {
			fmt.Println("Too Many letters! Please only enter a single letter guess!")
		} else if !unicode.IsLetter(rune(guess[0])) {
			fmt.Println("No special characters! Please enter a valid letter as a guess!")
		} else if arrayContainsGuess(correctGuesses, guess) || arrayContainsGuess(wrongGuesses, guess){
			fmt.Println("You've already guessed that letter! Guess again!")
		} else {
			break
		}
	}
	return processGuess(guess)
}

func readInGuess() string {
	fmt.Println("Please enter a letter: ")
	guess, _ := reader.ReadString('\n')
	guess = strings.TrimSpace(guess)
	guess = strings.ToLower(guess)
	fmt.Println("You guessed : ", guess)
	return guess
}

func processGuess(guess string) bool {
	if strings.Contains(word, guess) {
		fmt.Println("Sweet! The word has your guess!")
		correctGuesses = append(correctGuesses, guess)
		regenerateGuessedWord()
	} else {
		fmt.Println("Good Try! However the letter was not in the word.")
		wrongGuesses = append(wrongGuesses, guess)
	}
	displayGameStatus()
	if len(wrongGuesses) > 5 {
		displayLoseMessage()
		return true
	}
	if !strings.Contains(guessedWord, HiddenLetter) {
		displayWinMessage()
		return true
	}
	return false
}

func displayIntroduction() {
	fmt.Println("Welcome to Hangman!")
	fmt.Println()
	fmt.Println("How to play: ")
	fmt.Println("1. We'll choose a random word.")
	fmt.Println("2. You guess a letter you think may be in the word.")
	fmt.Println("3. You get 6 wrong guesses to save your friend Mr. Hangman.")
	fmt.Println("4. If you can guess all the letters and proclaim the word you WIN!")
	fmt.Println("5. If not... well Mr Hangman's wife will not be happy with you...")
	fmt.Println()
	fmt.Println("GOOD LUCK!")
}

func displayWinMessage() {
	fmt.Println()
	fmt.Println("You discovered the word! It took you ", len(correctGuesses) + len(wrongGuesses), " guesses")
	fmt.Println( "Your word was: ", word)
	fmt.Println("Congratulations! Mr Hangman is officially free to go!")
}

func displayLoseMessage()  {
	fmt.Println()
	fmt.Println("OH NO! You were unable to guess the word!")
	fmt.Println("Your word was: ", word)
	fmt.Println("Tough luck for Mrs. Hangman and their children. Better luck next time.")
}

func displayGameStatus() {
	fmt.Println()
	fmt.Println("Current word with guessed letters: ", guessedWord)
	fmt.Println("Guessed Correct Letters: ", correctGuesses)
	fmt.Println("Guessed Incorrect Letters: ", wrongGuesses)
	fmt.Println()
}

func regenerateGuessedWord() {
	guessedWord = ""
	for i := 0; i < len(word); i++ {
		if arrayContainsGuess(correctGuesses, string(word[i])) {
			guessedWord += string(word[i])
		} else {
			guessedWord += HiddenLetter
		}
		guessedWord += " "
	}
}

func arrayContainsGuess(list []string, guess string) bool {
	for _, listEntry := range list {
		if listEntry == guess {
			return true
		}
	}
	return false
}
