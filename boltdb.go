package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func openDB(fileName string) *bolt.DB {
	// Open the data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(fileName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func initBucket(db *bolt.DB, name string) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}
		return nil
	})
}

func writeToDB(db *bolt.DB, data Job, bucket string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		encoded, err := json.Marshal(data)
		if err != nil {
			return err
		}
		fmt.Println(data.Hash)
		err = b.Put([]byte(data.Hash), encoded)
		return err
	})
}

func readFromDB(db *bolt.DB, key, bucket string) []byte {
	v := []byte("")
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v = b.Get([]byte(key))
		return nil
	})

	return v

}
