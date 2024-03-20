package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)


type GameData struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

func main() {
	godotenv.Load()
	
	file, err := os.Open(os.Getenv("FILE_PATH"))
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

    reader := bufio.NewReader(file)

	// todo make it a constructor
	gameData := createGame()
    
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
			gameData.addTotalKills()
			killedBy := lineContent[5]
			playerKilled := lineContent[5]
			gameData.addPlayer(killedBy, playerKilled)
			gameData.addKill(killedBy, playerKilled)
			
		case strings.Contains(action, "clientconnect"):
			continue
		case strings.Contains(action, "shutdowngame"):
			gameData = createGame()
		}

		// Convert the struct to JSON
		jsonData, err := json.MarshalIndent(gameData, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Print the JSON data
		fmt.Println(string(jsonData))
		


    }

}

func createGame() *GameData {
	return &GameData{
		TotalKills: 0,
		Kills: make(map[string]int),
		Players: []string{},
	}
}

func (g *GameData) addPlayer(killedBy string, player string) {
	if (killedBy == "<world>") {
		return
	}
	for _, p := range g.Players {
		if p == player {
			return
		}
	}
	g.Players = append(g.Players, player)
}

func (g *GameData) removePlayer(player string) {
	for i, p := range g.Players {
		if p == player {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}
}

func (g *GameData) addKill(killedBy string, player string) {
	if (killedBy == "<world>") {
		g.removeKill(player)
	}
	g.Kills[killedBy] = g.Kills[killedBy] + 1
}

func (g *GameData) removeKill(player string) {
	g.Kills[player] = g.Kills[player] - 1
	if g.Kills[player] < 0 {
		g.Kills[player] = 0
	
	}
}

func (g *GameData) addTotalKills() {
	g.TotalKills++
}


// Remove leading and trailing spaces from a string
func removeTrailingSpaces(text string) string {
	return strings.TrimSpace(text)

}
