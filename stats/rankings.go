package stats

import (
	"fmt"
	"os"
	"time"

	"github.com/papaya147/vlr-stats/util"
	"github.com/papaya147/vlr-stats/writer"
)

type Team struct {
	TeamName      string `json:"team_name"`
	TeamLogo      string `json:"team_logo"`
	TeamUrl       string `json:"team_url"`
	TeamRank      string `json:"team_rank"`
	Rating        string `json:"rating_score"`
	WinStreak     string `json:"win_streak"`
	Record        string `json:"record"`
	TotalWinnings string `json:"total_winnings"`
}

type Teams struct {
	Teams []Team `json:"teams"`
}

func (t Teams) GetHeaders() []string {
	return []string{"Time Added", "Team Name", "Team URL", "Team Rank", "Rating", "Win Streak", "Record", "Total Winnings"}
}

func (t Teams) GetRecords() [][]string {
	records := [][]string{}
	for _, team := range t.Teams {
		records = append(records, []string{fmt.Sprintf("%d", time.Now().Unix()), team.TeamName,
			team.TeamUrl, team.TeamRank, team.Rating, team.WinStreak, team.Record, team.TotalWinnings})
	}
	return records
}

func (t Teams) save(file *os.File) error {
	return writer.AppendToCSV(t, file)
}

var rankingsBaseUrl = `https://vlrgg.cyclic.app/api/rankings/%s`

func SaveRankings(region string, file *os.File) error {
	var teams Teams
	if err := util.Request(fmt.Sprintf(rankingsBaseUrl, region), &teams); err != nil {
		return err
	}

	return teams.save(file)
}
