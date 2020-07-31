package cmd

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the tasks in the database",
	Long:  "",
	Run:   listTasks,
}

func init() {

}

func listTasks(cmd *cobra.Command, args []string) {
	//fmt.Println("haha that a list")
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyTasks"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("task %s : %s\n", k, v)
		}
		return nil
	})
}
