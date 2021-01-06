/*
This is actually context
*/
package flyweight

type Organization interface {
	GetTeams() []Team
	GetTournamentName() string
}

type Tournament struct {
	name      string
	teamNames []string
	factory   *teamFlyweightFactory
}

func NewTournament(factory *teamFlyweightFactory, name string, teamNames []string) Organization {

	return &Tournament{
		name:      name,
		teamNames: teamNames,
		factory:   factory,
	}
}

func (s *Tournament) GetTeams() []Team {
	var teams []Team
	//factory := NewTeamFactory()

	for _, name := range s.teamNames {
		teams = append(teams, *s.factory.GetTeam(name))
	}

	return teams
}

func (s *Tournament) GetTournamentName() string {
	return s.name
}
