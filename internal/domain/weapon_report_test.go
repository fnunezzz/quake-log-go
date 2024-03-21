package domain

import "testing"

func TestCreateWeaponReport(t *testing.T) {
	g := CreateWeponReport()
	if g == nil {
		t.Errorf("CreateGame() = nil, want &GameData{}")
	}
}

// Test adding a single kill/weapon to the report
func TestAddSingleKill(t *testing.T) {
	g := CreateWeponReport()
	g.AddKill("MOD_SHOTGUN")
	if g.KillsByMeans["MOD_SHOTGUN"] != 1 {
		t.Errorf("AddKill().KillsByMeans[MOD_SHOTGUN] = %d, want 1", g.KillsByMeans["MOD_SHOTGUN"])
	}
}

// Test adding a non existing weapon to the report
func TestAddNonExistingKill(t *testing.T) {
	g := CreateWeponReport()
	g.AddKill("TEST_INVALID_WEAPON")
	if g.KillsByMeans["TEST_INVALID_WEAPON"] != 0 {
		t.Errorf("AddKill().KillsByMeans[TEST_INVALID_WEAPON] = %d, want 0", g.KillsByMeans["TEST_INVALID_WEAPON"])
	}
}

// Test adding multiple valid weapons to the report
func TestAddMultipleKills(t *testing.T) {
	g := CreateWeponReport()
	g.AddKill("MOD_SHOTGUN")
	g.AddKill("MOD_SHOTGUN")
	g.AddKill("MOD_SHOTGUN")
	g.AddKill("MOD_PROXIMITY_MINE")
	if g.KillsByMeans["MOD_SHOTGUN"] != 3 {
		t.Errorf("AddKill().KillsByMeans[MOD_SHOTGUN] = %d, want 3", g.KillsByMeans["MOD_SHOTGUN"])
	}
	if g.KillsByMeans["MOD_PROXIMITY_MINE"] != 1 {
		t.Errorf("AddKill().KillsByMeans[MOD_PROXIMITY_MINE] = %d, want 1", g.KillsByMeans["MOD_PROXIMITY_MINE"])
	}
}

// Test adding multiple valid and invalid weapons to the report
func TestAddMultipleKillsWithInvalid(t *testing.T) {
	g := CreateWeponReport()
	g.AddKill("MOD_SHOTGUN")
	g.AddKill("MOD_SHOTGUN")
	g.AddKill("MOD_SHOTGUN")
	g.AddKill("MOD_PROXIMITY_MINE")
	g.AddKill("TEST_INVALID_WEAPON")
	if g.KillsByMeans["MOD_SHOTGUN"] != 3 {
		t.Errorf("AddKill().KillsByMeans[MOD_SHOTGUN] = %d, want 3", g.KillsByMeans["MOD_SHOTGUN"])
	}
	if g.KillsByMeans["MOD_PROXIMITY_MINE"] != 1 {
		t.Errorf("AddKill().KillsByMeans[MOD_PROXIMITY_MINE] = %d, want 1", g.KillsByMeans["MOD_PROXIMITY_MINE"])
	}
	
	if g.KillsByMeans["TEST_INVALID_WEAPON"] != 0 {
		t.Errorf("AddKill().KillsByMeans[TEST_INVALID_WEAPON] = %d, want 0", g.KillsByMeans["TEST_INVALID_WEAPON"])
	}

}