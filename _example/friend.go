package example

import "github.com/mackee/go-sqlla/v2"

// FriendID is Friend's surrogate key
type FriendID int64

// Friend is ...
//+table: friend
type Friend struct {
	ID      FriendID `db:"id"`
	User1ID UserID   `db:"user1_id"`
	User2ID UserID   `db:"user1_id"`
}

// FriendWithUsers is management friend with users
//+sqlla: join
type FriendWithUsers struct {
	Friend Friend `sqlla:"table=friend"`
	User1  User   `sqlla:"table=user"`
	User2  User   `sqlla:"table=user"`
}

func (uf FriendWithUsers) _innerJoin(methods sqlla.JoinMethods) []sqlla.JoinCondition {
	return []sqlla.JoinCondition{
		methods.On(uf.Friend.User1ID == uf.User1.ID),
		methods.On(uf.Friend.User2ID == uf.User2.ID),
	}
}
