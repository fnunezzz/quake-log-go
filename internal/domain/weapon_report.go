package domain

var MEANS_OF_DEATH = map[string]string{
	"MOD_UNKNOWN":        "MOD_UNKNOWN",
	"MOD_SHOTGUN":        "MOD_SHOTGUN",
	"MOD_GAUNTLET":       "MOD_GAUNTLET",
	"MOD_MACHINEGUN":     "MOD_MACHINEGUN",
	"MOD_GRENADE":        "MOD_GRENADE",
	"MOD_GRENADE_SPLASH": "MOD_GRENADE_SPLASH",
	"MOD_ROCKET":         "MOD_ROCKET",
	"MOD_ROCKET_SPLASH":  "MOD_ROCKET_SPLASH",
	"MOD_PLASMA":         "MOD_PLASMA",
	"MOD_PLASMA_SPLASH":  "MOD_PLASMA_SPLASH",
	"MOD_RAILGUN":        "MOD_RAILGUN",
	"MOD_LIGHTNING":      "MOD_LIGHTNING",
	"MOD_BFG":            "MOD_BFG",
	"MOD_BFG_SPLASH":     "MOD_BFG_SPLASH",
	"MOD_WATER":          "MOD_WATER",
	"MOD_SLIME":          "MOD_SLIME",
	"MOD_LAVA":           "MOD_LAVA",
	"MOD_CRUSH":          "MOD_CRUSH",
	"MOD_TELEFRAG":       "MOD_TELEFRAG",
	"MOD_FALLING":        "MOD_FALLING",
	"MOD_SUICIDE":        "MOD_SUICIDE",
	"MOD_TARGET_LASER":   "MOD_TARGET_LASER",
	"MOD_TRIGGER_HURT":   "MOD_TRIGGER_HURT",
	"MOD_NAIL":           "MOD_NAIL",
	"MOD_CHAINGUN":       "MOD_CHAINGUN",
	"MOD_PROXIMITY_MINE": "MOD_PROXIMITY_MINE",
	"MOD_KAMIKAZE":       "MOD_KAMIKAZE",
	"MOD_JUICED":         "MOD_JUICED",
	"MOD_GRAPPLE":        "MOD_GRAPPLE",
}

type WeaponScoreReport struct {
	KillsByMeans map[string]int `json:"kills_by_means"`
}

func CreateWeponReport() *WeaponScoreReport {
	return &WeaponScoreReport{
		KillsByMeans: make(map[string]int),
	}
}

// Check if weapon part of source code
func (w *WeaponScoreReport) verifyWeapon(weapon string) bool {
	for _, c := range MEANS_OF_DEATH {
		if weapon == c {
			return true
		}
	}
	return false
}

// Add a weapon kill to the report
// If weapon is not part of the source code id does not add it to the report
func (w *WeaponScoreReport) AddKill(weapon string) {
	if !w.verifyWeapon(weapon) {
		return
	}
	w.KillsByMeans[weapon] = w.KillsByMeans[weapon] + 1
}