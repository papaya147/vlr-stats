package main

import (
	"log"

	"github.com/papaya147/vlr-stats/stats"
	"github.com/papaya147/vlr-stats/util"
)

func main() {
	if err := loadRankings(); err != nil {
		log.Println("error from loading rankings:", err)
	}
	if err := loadPlayers(); err != nil {
		log.Println("error from loading players:", err)
	}
	if err := loadMatchResults(); err != nil {
		log.Println("error from loading match results:", err)
	}
}

func loadRankings() error {
	rankingsFile, err := util.CreateAppendFile("rankings.csv")
	if err != nil {
		return err
	}
	defer rankingsFile.Close()

	return stats.SaveRankings("latin-america", rankingsFile)
}

func loadPlayers() error {
	playersFile, err := util.CreateAppendFile("players.csv")
	if err != nil {
		return err
	}
	defer playersFile.Close()

	return stats.SavePlayers(playersFile)
}

func loadMatchResults() error {
	matchesFile, err := util.CreateAppendFile("match-results.csv")
	if err != nil {
		return err
	}
	defer matchesFile.Close()

	return stats.SaveMatchResults(matchesFile)
}
