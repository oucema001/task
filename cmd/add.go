package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "use this command to add a task in the database",
	Long:  "",
	Run:   addTasks,
}

func init() {
}

func addTasks(cmd *cobra.Command, args []string) {
	//fmt.Println("haha that adding a list")
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyTasks"))
		err := b.Put([]byte(strconv.Itoa(count(db))), []byte(strings.Join(args, " ")))
		return err
	})
}

func count(db *bolt.DB) int {
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
