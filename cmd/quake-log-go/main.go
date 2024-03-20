package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fnunezzz/quake-log-go/internal/domain"
	helpers "github.com/fnunezzz/quake-log-go/internal/helpers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	
	file, err := os.Open(os.Getenv("FILE_PATH"))
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

    processGameData(file)

}

func processGameData(file *os.File) {
	reader := bufio.NewReader(file)

	gameData := domain.CreateGame()
    
	// Infinite loop
	// keeps iterating until an error occurs (end of file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }

		lineNormalized := removeTrailingSpaces(line)
        
		lineContent := strings.Split(lineNormalized, " ")
		action := strings.ToLower(lineContent[1])
		
		switch {
		case strings.Contains(action, "initgame"):
			continue
		case strings.Contains(action, "kill"):
			killedBy := lineContent[5]
			playerKilled := lineContent[5]
			gameData.CalculateKillPoints(killedBy, playerKilled)
			
		case strings.Contains(action, "clientconnect"):
			continue
		case strings.Contains(action, "shutdowngame"):
			gameData = domain.CreateGame()
		}

		// Convert the struct to JSON
		jsonData, err := helpers.ToJson(gameData)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Print the JSON data
		fmt.Println(string(jsonData))

    }
}






// Remove leading and trailing spaces from a string
func removeTrailingSpaces(text string) string {
	return strings.TrimSpace(text)

}
