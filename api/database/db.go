package database

import api "grpc-chat/api/proto"

type DB struct {
	db []api.Message
}

func (db *DB) New() (*DB, error) {
	db.db = make([]api.Message, 0, 5)
	return db, nil
}

func (db *DB) PutIn(message *api.Message) {
	db.db = append(db.db, *message)
}

func (db *DB) GetDB() []api.Message {
	return db.db
}
