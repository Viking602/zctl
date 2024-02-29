package add

import (
	"fmt"
	"github.com/spf13/cobra"
	"zctl/utils"
)

var CmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a new resource",
	Long:  `Add a new resource to the Project Zero application.`,
	Run:   run,
}

var logicName string
var handlerName string
var projectName string

func init() {
	CmdAdd.Flags().StringVarP(&logicName, "logic", "l", "", "The name of the logic to add")
	CmdAdd.Flags().StringVarP(&handlerName, "handler", "s", "", "The name of the handler to add")
	CmdAdd.Flags().StringVarP(&projectName, "project", "p", "", "The name of the project to add the resource to")
}

func run(_ *cobra.Command, args []string) {

	var l service
	l.ProjectName = projectName
	l.LogicName = utils.FirstUpper(logicName)
	l.ServiceName = utils.FirstUpper(handlerName)

	if err := l.createFile(fmt.Sprintf("inernal/handlers/%v_%v_hander.go", handlerName, logicName), "handler.tmpl"); err != nil {
		fmt.Println(err)
		return
	}

	if err := l.createFile(fmt.Sprintf("inernal/logic/%v_%v_logic.go", handlerName, logicName), "logic.tmpl"); err != nil {
		fmt.Println(err)
		return
	}
}
