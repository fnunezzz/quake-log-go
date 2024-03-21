package domain

import "testing"

func TestCreateGameReport(t *testing.T) {
	g := CreateGameReport()
	if g == nil {
		t.Errorf("CreateGame() = nil, want &GameData{}")
	}
}

// This test is to check if the game is able to handle one kill by a player
// Should add to the total kills and to the player
func TestCalculateKillPointsForOneKillByPlayer(t *testing.T) {
	g := CreateGameReport()
	g.CalculateKillPoints("player1", "player2")
	if g.TotalKills != 1 {
		t.Errorf("CalculateKillPoints().TotalKills = %d, want 1", g.TotalKills)
	}
	if g.Kills["player1"] != 1 {
		t.Errorf("CalculateKillPoints().Kills[player1] = %d, want 1", g.Kills["player1"])
	}
}

// This test is to check if the game is able to handle a kill by the world
// Should add to the total kills but not to the player
func TestCalculateKillPointsForOneKillByWorld(t *testing.T) {
	g := CreateGameReport()
	g.CalculateKillPoints("<world>", "player2")
	if g.TotalKills != 1 {
		t.Errorf("CalculateKillPoints().TotalKills = %d, want 1", g.TotalKills)
	}
	if g.Kills["player1"] != 0 {
		t.Errorf("CalculateKillPoints().Kills[player1] = %d, want 0", g.Kills["<world>"])
	}
}

// This test is to check if the game is able to handle a new player
// Should add the player to the list of players
func TestAddingNewPlayer(t *testing.T) {
	g := CreateGameReport()
	g.NewPlayer("player1")
	if len(g.Players) != 1 {
		t.Errorf("NewPlayer().Players.length = %d, want 1", len(g.Players))
	}
	if g.Players[0] != "player1" {
		t.Errorf("NewPlayer().Players[0] = %s, want player1", g.Players[0])
	}
}

// This test is to check if the game is able to handle the word "world" on the player name
// Should add the player to the list of players
func TestAddingPlayerWithWorldInName(t *testing.T) {
	g := CreateGameReport()
	g.NewPlayer("player_world")
	if len(g.Players) != 1 {
		t.Errorf("NewPlayer().Players.length = %d, want 1", len(g.Players))
	}
}

// This test is to check if the game is able to handle the same player name
// Should not add the same player twice to the list of players
func TestAddingPlayersWithSameName(t *testing.T) {
	g := CreateGameReport()
	g.NewPlayer("player1")
	g.NewPlayer("player1")
	if len(g.Players) != 1 {
		t.Errorf("NewPlayer().Players.length = %d, want 1", len(g.Players))
	}
}

func TestAddingMultiplePlayers(t *testing.T) {
	g := CreateGameReport()
	g.NewPlayer("player1")
	g.NewPlayer("player2")
	g.NewPlayer("player3")
	if len(g.Players) != 3 {
		t.Errorf("NewPlayer().Players.length = %d, want 3", len(g.Players))
	}
}

// This test is to check if the game is able to handle the world as the player
func TestAddingWorldPlayer(t *testing.T) {
	g := CreateGameReport()
	g.NewPlayer("<world>")
	if len(g.Players) != 0 {
		t.Errorf("NewPlayer().Players.length = %d, want 0", len(g.Players))
	}
}

// This test is to check if the game is able to handle multiple kills by several players and the world
func TestCalculateKillPointsForMultipleKillByPlayers(t *testing.T) {
	g := CreateGameReport()
	g.CalculateKillPoints("player1", "player2")
	g.CalculateKillPoints("player1", "player2")
	g.CalculateKillPoints("player1", "player2")
	g.CalculateKillPoints("player1", "player2")
	g.CalculateKillPoints("player2", "player1")
	g.CalculateKillPoints("<world>", "player3")

	if g.TotalKills != 6 {
		t.Errorf("CalculateKillPoints().TotalKills = %d, want 1", g.TotalKills)
	}
	if g.Kills["player1"] != 4 {
		t.Errorf("CalculateKillPoints().Kills[player1] = %d, want 1", g.Kills["player1"])
	}
	if g.Kills["player2"] != 1 {
		t.Errorf("CalculateKillPoints().Kills[player2] = %d, want 1", g.Kills["player2"])
	}
	if g.Kills["player3"] != -1 {
		t.Errorf("CalculateKillPoints().Kills[player3] = %d, want -1", g.Kills["player3"])
	}
	if g.Kills["<world>"] != 0 {
		t.Errorf("CalculateKillPoints().Kills[<world>] = %d, want 0", g.Kills["<world>"])
	}
}