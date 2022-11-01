package database

import (
	"encoding/binary"
	"log"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

// Creating db *bolt.DB as global variable for the pacakge 'database'
// This will make the db functions accessible outside the package
// So functions in the 'cmd' pacakge will have access to the db connection (pointer).
var db *bolt.DB
var taskBucket = []byte("tasks")
var dbPath = setDBPath("tasks.db")

type Task struct {
	Key int
	Value string
}

func Connect() {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Connected to database:%v", dbPath)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		if err != nil {
			log.Fatalf("Failed to create a bucket %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Could not create bucket: %v", err)
	}
}

func Close() {
	db.Close()
}

func CreateTask(task string) (int, error) {
	var id int
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)

		value := []byte(task)

		err := b.Put(key, value)
		return err
	})
	return id, nil
}

func DeleteTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		
		err := b.Delete(itob(id))
		if err != nil {
			log.Fatalf("Failed to delete task with id:%d", id)
		}
		return nil
	})
}

func ListTasks() []Task{
	tasks := []Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key: btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Could not list tasks: %v", err)
	}
	return tasks
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func setDBPath(name string) string{
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd + "/" + name
}