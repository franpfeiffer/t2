package main

import (
	"fmt"
	"os"
	"time"
	"log"

	"github.com/charmbracelet/huh"
)

func waitForEnter() {
	_, _ = fmt.Scanln()
}

func saveTimeToFile(minutes int) {
	file, err := os.OpenFile("tracked-time.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Time tracked: %d minute(s)\n", minutes))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Done!")
	}

}

func main() {
	var confirm bool
	var startTime time.Time

	startForm := huh.NewConfirm().
		Title("Start the timer? (can use vim motions)").
		Affirmative("Yes").
		Negative("No").
		Value(&confirm)

	err := startForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	if confirm {
		startTime = time.Now()
	}

	Title("Press Enter to stop the timer.")
	waitForEnter()


	elapsedTime := time.Since(startTime)
	minutes := int(elapsedTime.Minutes())

	stopForm := huh.NewConfirm().
		Title(fmt.Sprintf("Timer stopped. Total time: %d minute(s).", minutes)).
		Value(&confirm)

	err = stopForm.Run()

	if err != nil {
		log.Fatal(err)
	}

	saveTimeToFile(minutes)
}

