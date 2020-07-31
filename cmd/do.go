package cmd

import (
	"log"

	"github.com/oucema001/task/data"
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
	err := data.Delete(args, "MyTasks")
	if err != nil {
		log.Fatal(err)
	}
}
