package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fnunezzz/quake-log-go/internal/domain"
	"github.com/fnunezzz/quake-log-go/internal/helpers"
)

func main() {
	
	var filePath string

	flag.StringVar(&filePath, "filepath", "", "Absolute path to the game log including the file name and extension")
	flag.Parse()

	if (filePath == "") {
		log.Fatalf("filePath is required and cannot be empty")
	}
	
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var gameReport *domain.GameReportData

    
	// Infinite loop
	// keeps iterating until an error occurs (end of file)
	gameNumber := 0
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }

		lineNormalized := helpers.RemoveTrailingSpaces(line)
        
		lineContent := strings.Split(lineNormalized, " ")
		action := strings.ToLower(lineContent[1])
		switch {
		case strings.Contains(action, "initgame"):
			gameReport = domain.CreateGameReport()
			gameNumber++
			continue
		case strings.Contains(action, "clientuserinfochanged"):
			// not every player will have a kill. Some can just join the game and be AFK
			player := strings.Split(lineContent[3], "\\")[1]
			gameReport.NewPlayer(player)
			continue
		case strings.Contains(action, "kill"):
			killedBy := lineContent[5]
			playerKilled := lineContent[7]
			gameReport.CalculateKillPoints(killedBy, playerKilled)
			continue
		case strings.Contains(action, "shutdowngame"):
			jsonData, err := helpers.ToJson(gameReport)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
	
			message := fmt.Sprintf("Game_%d: %s", gameNumber, string(jsonData))
			fmt.Println(message)
			continue
		default:
			continue
		}
	}
}



