package windowsagent

import (
	"fmt"

	ps "github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/winservices"
	"golang.org/x/sys/windows/svc"
)

func WindowsServices() {
	listServices, _ := winservices.NewService("")
	fmt.Printf("All listServices: ", listServices)

	processList, _ := ps.Processes()
	fmt.Printf("All ps: ", processList)

	for x := range processList {
		var process ps.Process
		process = processList[x]
		fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())

		// do os.* stuff on the pid
	}

	statusHandler := svc.StatusHandle()
	fmt.Printf("All statusHandler: ", statusHandler)

}
