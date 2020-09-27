package db

import (
	"encoding/binary"
	"fmt"
	"time"

	bolt "github.com/coreos/bbolt"
)

var taskBucket = []byte("task")
var db *bolt.DB

//Task struct to add task in the taskBucket
type Task struct {
	Key   int
	Value string
}

//Init the func
func Init(dbPath string) error {
	var err error
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	fn := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}
	// fn is a clousure
	return db.Update(fn)
}

//CreateTask comment
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	fmt.Println("itob", b)
	return b
}

func btoi(b []byte) int {
	fmt.Println("btoi", binary.BigEndian.Uint64(b))
	return int(binary.BigEndian.Uint64(b))
}

//AllTasks function to view
func AllTasks() ([]Task, error) {
	fmt.Println("222222222")
	var tasks []Task

	fn := func(tx *bolt.Tx) error {
		bukt := tx.Bucket(taskBucket)
		cur := bukt.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	}

	err := db.View(fn)

	// err := db.View(func(tx *bolt.Tx) error {
	// 	bukt := tx.Bucket(taskBucket)
	// 	cur := bukt.Cursor()
	// 	for k, v := cur.First(); k != nil; k, v = cur.Next() {
	// 		tasks = append(tasks, Task{
	// 			Key:   btoi(k),
	// 			Value: string(v),
	// 		})
	// 	}
	// 	return nil
	// })
	fmt.Println("Therr err....", err)
	if err != nil {
		return nil, err
	}
	fmt.Println("Therr task....", tasks)
	return tasks, nil
}

//DeleteTask function to delete
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(taskBucket)
		return bkt.Delete(itob(key))

	})
}
