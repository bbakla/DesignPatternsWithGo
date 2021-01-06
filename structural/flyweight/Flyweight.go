package flyweight

import "time"

type Team struct {
	ID             int
	Name           string
	Players        []Player
	HistoricalData []HistoricalData
}

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
}

type Match struct {
	Date         time.Time
	VisitorID    uint64
	HostID       uint64
	LocalScore   byte
	VisitorScore byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}
