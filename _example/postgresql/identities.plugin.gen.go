// Code generated by github.com/mackee/go-sqlla/v2/cmd/sqlla. DO NOT EDIT.
package postgresql

import "time"

func (i Identity) DefaultInsertHook(q identityInsertSQL) (identityInsertSQL, error) {
	now := time.Now()
	return q.
		ValueCreatedAt(now).
		ValueUpdatedAt(now), nil
}

func (i Identity) DefaultUpdateHook(q identityUpdateSQL) (identityUpdateSQL, error) {
	now := time.Now()
	return q.
		SetUpdatedAt(now), nil
}
