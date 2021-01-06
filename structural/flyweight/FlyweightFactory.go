package flyweight

import (
	"time"
)

type teamFlyweightFactory struct {
	teams map[string]*Team
}

func NewTeamFactory() *teamFlyweightFactory {
	return &teamFlyweightFactory{
		teams: make(map[string]*Team),
	}
}

func (t *teamFlyweightFactory) GetTeam(name string) *Team {
	team := t.teams[name]
	if team == nil {
		team = &Team{
			ID:   t.generateTeamId(),
			Name: name,
		}

		t.teams[name] = team

		return team
	}

	return team
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.teams)
}

func (t *teamFlyweightFactory) generateTeamId() int {
	return int(time.Now().UnixNano()) % 1000000
}
