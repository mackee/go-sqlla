package example

import (
	"github.com/mackee/go-sqlla/v2"
)

// UserAndTeam is User and Team
//+sqlla: join
type UserAndTeam struct {
	User User `sqlla:"table=user"`
	Team Team `sqlla:"table=team"`
}

func (ut UserAndTeam) _innerJoin(methods sqlla.JoinMethods) []sqlla.JoinCondition {
	return []sqlla.JoinCondition{
		methods.On(ut.User.TeamID == ut.Team.ID),
	}
}
