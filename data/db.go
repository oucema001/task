package data

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func Init() error {
	var err error
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	//	defer db.Close()
	err = createBucket(db)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func createBucket(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MyTasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func List(bucketName string) error {
	fmt.Println("wsel")
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("task %s : %s\n", k, v)
		}
		return nil
	})

}

func Add(args []string, bucketName string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(strconv.Itoa(count())), []byte(strings.Join(args, " ")))
		return err
	})
}

func Delete(args []string, bucketName string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Delete([]byte(strings.Join(args, " ")))
		return err
	})
}

func count() int {
	total := 0
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyTasks"))

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			total++
		}
		return nil
	})
	return total
}
