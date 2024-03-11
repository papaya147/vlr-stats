package main

import (
	"context"
	"log"
	"time"

	"github.com/papaya147/parallelize"
	"github.com/papaya147/vlr-stats/stats"
	"github.com/papaya147/vlr-stats/util"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	group := parallelize.NewSyncGroup()
	parallelize.AddMethodWithoutArgs(group, loadRankings, ctx)
	parallelize.AddMethodWithoutArgs(group, loadPlayers, ctx)
	parallelize.AddMethodWithoutArgs(group, loadMatchResults, ctx)
	if err := group.Run(); err != nil {
		log.Println(err)
	}
}

func loadRankings(ctx context.Context) error {
	rankingsFile, err := util.CreateAppendFile("rankings.csv")
	if err != nil {
		return err
	}
	defer rankingsFile.Close()

	return stats.SaveRankings(ctx, "latin-america", rankingsFile)
}

func loadPlayers(ctx context.Context) error {
	playersFile, err := util.CreateAppendFile("players.csv")
	if err != nil {
		return err
	}
	defer playersFile.Close()

	return stats.SavePlayers(ctx, playersFile)
}

func loadMatchResults(ctx context.Context) error {
	matchesFile, err := util.CreateAppendFile("match-results.csv")
	if err != nil {
		return err
	}
	defer matchesFile.Close()

	return stats.SaveMatchResults(ctx, matchesFile)
}
