package sqlla

import ()

type Selector interface {
	ToSelectSQL() (string, []interface{}, error)
}

type Updater interface {
	ToUpdateSQL() (string, []interface{}, error)
}

type Inserter interface {
	ToInsertSQL() (string, []interface{}, error)
}

type Deleter interface {
	ToDeleteSQL() (string, []interface{}, error)
}
