package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fnunezzz/quake-log-go/internal/domain"
	"github.com/fnunezzz/quake-log-go/internal/helpers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	
	file, err := os.Open(os.Getenv("FILE_PATH"))
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

    gameDataReport(file)

}



func gameDataReport(file *os.File) {
	reader := bufio.NewReader(file)

	var gameData *domain.GameData

    
	// Infinite loop
	// keeps iterating until an error occurs (end of file)
	gameNumber := 0
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
			gameData = domain.CreateGame()
			gameNumber++
			continue
		case strings.Contains(action, "clientuserinfochanged"):
			// not every player will have a kill. Some can just join the game and be AFK
			player := strings.Split(lineContent[3], "\\")[1]
			gameData.NewPlayer(player)
			continue
		case strings.Contains(action, "kill"):
			killedBy := lineContent[5]
			playerKilled := lineContent[7]
			gameData.CalculateKillPoints(killedBy, playerKilled)
		case strings.Contains(action, "shutdowngame"):
			jsonData, err := helpers.ToJson(gameData)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
	
			message := fmt.Sprintf("Game_%d: %s", gameNumber, string(jsonData))
			fmt.Println(message)
		default:
			continue
		}
		


		
    }
}






// Remove leading and trailing spaces from a string
func removeTrailingSpaces(text string) string {
	return strings.TrimSpace(text)

}
