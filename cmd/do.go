package cmd

import (
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Use this command to finish a task in the database",
	Long:  "",
	Run:   doTasks,
}

func init() {

}

func doTasks(cmd *cobra.Command, args []string) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyTasks"))
		err := b.Delete([]byte(strings.Join(args, " ")))
		return err
	})
}
