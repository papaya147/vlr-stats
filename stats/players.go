package stats

import (
	"fmt"
	"os"
	"time"

	"github.com/papaya147/vlr-stats/util"
	"github.com/papaya147/vlr-stats/writer"
)

type Player struct {
	PlayerName                       string `json:"player_name"`
	PlayerLink                       string `json:"player_link"`
	PlayerTeamInitials               string `json:"player_team_initials"`
	PlayerCountryInitials            string `json:"player_country_initials"`
	RoundsPlayed                     string `json:"rounds_played"`
	Rating                           string `json:"rating"`
	AverageCombatScore               string `json:"average_combat_score"`
	KillsDeaths                      string `json:"kills_deaths"`
	KillAssistTradeSurvivePercentage string `json:"kill_assist_trade_survive_percentage"`
	AverageDamagePerRound            string `json:"average_damage_per_round"`
	KillsPerRound                    string `json:"kills_per_round"`
	AssistsPerRound                  string `json:"assists_per_round"`
	FirstKillsPerRound               string `json:"first_kills_per_round"`
	FirstDeathsPerRound              string `json:"first_deaths_per_round"`
	HeadshotPercentage               string `json:"headshot_percentage"`
	ClutchSuccessPercentage          string `json:"clutch_success_percentage"`
	MaxKillsInSingleMap              string `json:"max_kills_in_single_map"`
	Kills                            string `json:"kills"`
	Deaths                           string `json:"deaths"`
}

type Players struct {
	Players []Player `json:"players"`
}

func (p Players) GetHeaders() []string {
	return []string{"Time Added", "Player Name", "Player Link", "Player Team Initials",
		"Player Country Initials", "Rounds Played", "Rating", "Average Combat Score",
		"Kills Deaths", "Kill Assist Trade Survive Percentage", "Average Damage Per Round",
		"Kills Per Round", "Assists Per Round", "First Kills Per Round", "First Deaths Per Round",
		"Headshot Percentage", "Clutch Success Percentage", "Max Kills In Single Map", "Kills", "Deaths",
	}
}

func (p Players) GetRecords() [][]string {
	records := [][]string{}
	for _, player := range p.Players {
		records = append(records, []string{fmt.Sprintf("%d", time.Now().Unix()), player.PlayerName,
			player.PlayerLink, player.PlayerTeamInitials, player.PlayerCountryInitials, player.RoundsPlayed,
			player.Rating, player.AverageCombatScore, player.KillsDeaths, player.KillAssistTradeSurvivePercentage,
			player.AverageDamagePerRound, player.KillsPerRound, player.AssistsPerRound, player.FirstKillsPerRound,
			player.FirstDeathsPerRound, player.HeadshotPercentage, player.ClutchSuccessPercentage,
			player.MaxKillsInSingleMap, player.Kills, player.Deaths,
		})
	}
	return records
}

func (p Players) save(file *os.File) error {
	return writer.AppendToCSV(p, file)
}

var playersBaseUrl = `https://vlrgg.cyclic.app/api/players`

func SavePlayers(file *os.File) error {
	var players Players
	if err := util.Request(playersBaseUrl, &players); err != nil {
		return err
	}

	return players.save(file)
}
