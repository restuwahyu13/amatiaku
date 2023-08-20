package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func StartScreenTime(nextRequestWatchDogs int) {
	startRequest := time.Now().Format("2006-01-02 15:04:05")
	nextRequest := time.Now().Add(time.Second * time.Duration(nextRequestWatchDogs)).Format("2006-01-02 15:04:05")

	if nextRequestWatchDogs == 0 {
		nextRequest = "unknown"
	}

	fmt.Printf("======================================================================================\n")
	fmt.Printf("\n")
	fmt.Printf("====================== START REQUEST TIME: %s =======================\n", startRequest)
	fmt.Printf("====================== NEXT REQUEST TIME: %s  =======================\n", nextRequest)
	fmt.Printf("\n")
	fmt.Printf("======================================================================================\n")
}

func EndScreenTime(nextRequestWatchDogs int) {
	startRequest := time.Now().Add(time.Second * time.Duration(nextRequestWatchDogs)).Format("2006-01-02 15:04:05")
	nextRequest := time.Now().Add(time.Second * time.Duration(nextRequestWatchDogs+nextRequestWatchDogs)).Format("2006-01-02 15:04:05")

	if nextRequestWatchDogs == 0 {
		nextRequest = "unknown"
	}

	fmt.Printf("======================================================================================\n")
	fmt.Printf("\n")
	fmt.Printf("====================== START REQUEST TIME: %s =======================\n", startRequest)
	fmt.Printf("====================== NEXT REQUEST TIME: %s  =======================\n", nextRequest)
	fmt.Printf("\n")
	fmt.Printf("======================================================================================\n")
}
