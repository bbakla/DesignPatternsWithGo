package main

import (
	"fmt"
	"github.com/bbakla/DesignPatternsWithGo/structural/flyweight"
)

type ChampionsCup struct {
	flyweight.Organization
}

type League struct {
	flyweight.Organization
}

func main() {
	factory := flyweight.NewTeamFactory()
	championsCupTeamNames := []string{"fenerbahce", "barcelona", "liverpool"}
	championsCupTournament := flyweight.NewTournament(factory, "preparationTournament", championsCupTeamNames)
	championsCup := ChampionsCup{championsCupTournament}

	fmt.Printf("Name of the teams in %s are %v\n", championsCup.GetTournamentName(), championsCup.GetTeams())

	leagueTeams := []string{"fenerbahce", "besiktas", "samsunspor"}
	leagueTeamsTournament := flyweight.NewTournament(factory, "super lig", leagueTeams)
	league := League{leagueTeamsTournament}

	fmt.Printf("Name of the teams in %s are %v\n", league.Organization.GetTournamentName(), league.GetTeams())

	fmt.Printf("Number of teams that are created should be 5. Actual is %d\n", factory.GetNumberOfObjects())

}
