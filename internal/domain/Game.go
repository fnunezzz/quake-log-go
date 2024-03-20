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

func (g *GameData) addPlayer(killedBy string, player string) {
	if killedBy == "<world>" {
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
	if killedBy == "<world>" {
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

func (g *GameData) CalculateKillPoints(killedBy string, playerKilled string) {
	g.addTotalKills()
	g.addPlayer(killedBy, playerKilled)
	g.addKill(killedBy, playerKilled)
}

func (g *GameData) ResetGame() {
	g.TotalKills = 0
	g.Players = []string{}
	g.Kills = make(map[string]int)
}
