package repository

import (
	bolt "go.etcd.io/bbolt"
)

// la db we para hacer consultas ricolinas e.e
var Db *bolt.DB

func OpenDb() error {
	var err error
	Db, err = bolt.Open("cypherdb.db", 0600, nil)
	if err != nil {
		return err
	}
	if err = Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("plans"))
		return err
	}); err != nil {
		return err
	}
	return nil
}
