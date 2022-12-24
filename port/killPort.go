package port

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/cothi/port/surv"
	"github.com/cothi/port/utils"
)

func KillPort() {
	//args := os.Args
	c := make(chan []string, 1)
	var portList []string
	var wg sync.WaitGroup
	wg.Add(1)


	if len(os.Args) < 2 {
		s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
		s.Suffix = "  loading find Port"
		go AllPort(c, &wg)
		
		s.Start()

		portList = <-c
		s.Stop()

		ans := surv.SelectPort(portList)
		splitSelect := strings.Split(strings.TrimSpace(ans.Port), " ")
		PortKill(utils.SplitToInt(utils.GetLastIndex(splitSelect)))
	}
	//wg.Done()
	//port, _ := strconv.Atoi(args[1])
	//portKill(port)
	//allPort()
	//port_kill(10)
	wg.Wait()
}
