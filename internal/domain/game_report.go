package domain

type GameReportData struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

func CreateGameReport() *GameReportData {
	return &GameReportData{
		TotalKills: 0,
		Kills:      make(map[string]int),
		Players:    []string{},
	}
}

func (g *GameReportData) NewPlayer(player string) {
	// <world> is a reserverd word for the game. Any other combination is accepted
	if player == "<world>" {
		return
	}
	g.addPlayer(player)
}

func (g *GameReportData) addPlayer(player string) {
	for _, p := range g.Players {
		if p == player {
			return
		}
	}
	g.Players = append(g.Players, player)
}

func (g *GameReportData) addKill(player string) {
	g.Kills[player] = g.Kills[player] + 1
}

func (g *GameReportData) removeKill(player string) {
	g.Kills[player] = g.Kills[player] - 1
}

func (g *GameReportData) addTotalKills() {
	g.TotalKills++
}

func (g *GameReportData) CalculateKillPoints(killedBy string, playerKilled string) {
	g.addTotalKills()
	if killedBy != "<world>" {
		g.addPlayer(killedBy)
		g.addKill(killedBy)
	} else {
		g.removeKill(playerKilled)
	}
}

func (g *GameReportData) ResetGame() {
	g.TotalKills = 0
	g.Players = []string{}
	g.Kills = make(map[string]int)
}
