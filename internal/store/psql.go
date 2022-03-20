package store

type PSQLStore struct {
	conn string
}

func NewPSQLStore(conn string) *PSQLStore {
	return &PSQLStore{
		conn,
	}
}

