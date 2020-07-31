package cmd

import (
	"log"

	"github.com/oucema001/task/data"

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
	err := data.List("MyTasks")
	if err != nil {
		log.Fatal(err)
	}
}
