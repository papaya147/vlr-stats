package stats

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/papaya147/vlr-stats/util"
	"github.com/papaya147/vlr-stats/writer"
)

type Match struct {
	Timestamp    int64
	TeamOneName  string `json:"team_one_name"`
	TeamTwoName  string `json:"team_two_name"`
	TeamOneScore string `json:"team_one_score"`
	TeamTwoScore string `json:"team_two_score"`
	MatchUrl     string `json:"match_url"`
	EventName    string `json:"event_name"`
	EventIconUrl string `json:"event_icon_url"`
	MatchTime    string `json:"match_time"`
	Eta          string `json:"eta"`
}

type Matches struct {
	Matches []Match `json:"matches"`
}

func (m Matches) GetHeaders() []string {
	return []string{"Time Added", "Team One Name", "Team Two Name", "Team One Score", "Team Two Score",
		"Match URL", "Event Name", "Event Icon URL", "Match Time", "ETA"}
}

func (m Matches) GetRecords() [][]string {
	records := [][]string{}
	for _, match := range m.Matches {
		records = append(records, []string{fmt.Sprintf("%d", time.Now().Unix()), match.TeamOneName,
			match.TeamTwoName, match.TeamOneScore, match.TeamTwoScore, match.MatchUrl, match.EventName,
			match.EventIconUrl, match.MatchTime, match.Eta})
	}
	return records
}

func (m Matches) save(file *os.File) error {
	return writer.AppendToCSV(m, file)
}

var matchResultsBaseUrl = `https://vlrgg.cyclic.app/api/matches/results`

func SaveMatchResults(ctx context.Context, file *os.File) error {
	var matches Matches
	if err := util.Request(ctx, matchResultsBaseUrl, &matches); err != nil {
		return err
	}

	return matches.save(file)
}
