package domain

type GameData struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

func CreateGame() *GameData {
	return &GameData{
		TotalKills: 0,
		Kills:      make(map[string]int),
		Players:    []string{},
	}
}

func (g *GameData) NewPlayer(player string) {
	// <world> is a reserverd word for the game. Any other combination is accepted
	if player == "<world>" {
		return
	}
	g.addPlayer(player)
}

func (g *GameData) addPlayer(player string) {
	for _, p := range g.Players {
		if p == player {
			return
		}
	}
	g.Players = append(g.Players, player)
}

func (g *GameData) addKill(player string) {
	g.Kills[player] = g.Kills[player] + 1
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

func (g *GameData) CalculateKillPoints(killedBy string, playerKilled string) {
	g.addTotalKills()
	if killedBy != "<world>" {
		g.addPlayer(killedBy)
		g.addKill(killedBy)
	} else {
		g.removeKill(playerKilled)
	}
}

func (g *GameData) ResetGame() {
	g.TotalKills = 0
	g.Players = []string{}
	g.Kills = make(map[string]int)
}
