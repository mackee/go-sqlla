package postgresql

import "time"

type IdentityID int64

//genddl:table identities
//sqlla:table identities
//sqlla:plugin timeHooks create=CreatedAt,UpdatedAt update=UpdatedAt
type Identity struct {
	ID        IdentityID `db:"id,autoincrement,primarykey"`
	AccountID AccountID  `db:"account_id"`
	Email     string     `db:"email"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
