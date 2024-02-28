package main

import (
	"fmt"
)

func main() {
	tubesVolume(2, 3)
	convertPoint(1000)
	oneToHundred()
	cetakPola(35)
	projectPriority(40, 25, 5)
	anagram("bowlehh", "bowleh")
	anagram("bowlehh", "bowlehh")

}

func tubesVolume(r, t float32) {
	const pi float32 = 3.14
	volume := pi * r * r * t
	fmt.Println(volume)
}

func convertPoint(point float32) {
	var convertedPoint string
	if point <= 100 && point >= 85 {
		convertedPoint = "A"
	} else if point <= 84 && point >= 70 {
		convertedPoint = "B"
	} else if point <= 69 && point >= 55 {
		convertedPoint = "C"
	} else if point <= 54 && point >= 40 {
		convertedPoint = "D"
	} else if point <= 39 && point >= 0 {
		convertedPoint = "E"
	} else {
		convertedPoint = "Invalid"
	}

	fmt.Println(convertedPoint)
}

func oneToHundred() {
	for i := 0; i <= 100; i++ {
		if i%4 == 0 && i%7 == 0 {
			fmt.Println("buzz")
			continue
		}
		if i%4 == 0 {
			fmt.Println(i * i)
			continue
		}
		if i%7 == 0 {
			fmt.Println(i * i * i)
			continue
		}
		fmt.Println(i)
	}
}

func cetakPola(angka int) {
	pola := ""

	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			if i%2 == 0 {
				pola += "I"
			} else {
				pola += "O"
			}
		}
	}
	fmt.Println(pola)
}

func projectPriority(budget, workTime, difficulty int) {
	var budgetScore, workTimeScore, difficultyScore int
	var result string
	switch {
	case budget <= 50 && budget >= 0:
		budgetScore = 4
	case budget <= 80 && budget >= 51:
		budgetScore = 3
	case budget <= 100 && budget >= 81:
		budgetScore = 2
	case budget > 100:
		budgetScore = 1
	}
	switch {
	case workTime <= 20 && workTime >= 0:
		workTimeScore = 10
	case workTime <= 30 && workTime >= 21:
		workTimeScore = 5
	case workTime <= 50 && workTime >= 31:
		workTimeScore = 2
	case workTime > 50:
		workTimeScore = 1
	}
	switch {
	case difficulty <= 3 && difficulty >= 0:
		difficultyScore = 10
	case difficulty <= 6 && difficulty >= 4:
		difficultyScore = 5
	case difficulty <= 10 && difficulty >= 7:
		difficultyScore = 1
	case difficulty > 10:
		difficultyScore = 0
	}
	
	totalScore := budgetScore + workTimeScore + difficultyScore
	switch {
	case totalScore <= 24 && totalScore >= 17:
		result = "High"
	case totalScore <= 16 && totalScore >= 10:
		result = "Medium"
	case totalScore <= 9 && totalScore >= 3:
		result = "Low"
	case totalScore <= 2:
		result = "Impossible"
	default:
		result = "Invalid"
	}

	fmt.Println("Butget :", budget, "- Budget Score:", budgetScore)
	fmt.Println("Work Time :", workTime, "- Work Time Score:", workTimeScore)
	fmt.Println("Difficulty :", difficulty, "- Difficulty Score:", difficultyScore)
	fmt.Println("Total Score :", totalScore,"- Result :", result)
}

func anagram(firstWord, secondWord string){
	firstWordLength := len(firstWord)
	secondWordLength := len(secondWord)
	result := ""

	if firstWordLength == secondWordLength {
		result = "anagram"
	} else {
		result = "not anagram"
	}
	fmt.Println("first word :", firstWord)
	fmt.Println("second word :", secondWord)
	fmt.Println(result)
}