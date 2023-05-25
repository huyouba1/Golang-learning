package views

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"todolist/commands/command"
)

type commandView struct {
}

func (v *commandView) Menu(cmds map[int]*command.Command) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"选项", "说明"})
	for i := 1; i <= len(cmds); i++ {
		table.Append([]string{
			strconv.Itoa(i),
			cmds[i].Name,
		})
	}

	table.Render()
	//ioutils.Output(strings.Repeat("*", 15))
	//for i := 1; i <= len(cmds); i++ {
	//	ioutils.Output(fmt.Sprintf("%d. %s", i, cmds[i].Name))
	//}
	//ioutils.Output(strings.Repeat("*", 15))
}

var CommandView = new(commandView)
