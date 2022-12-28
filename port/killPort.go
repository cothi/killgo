package port

import (
	"fmt"
	"os"
	"os/exec"
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

	fmt.Println("Done! ")
	wg.Wait()
}


func KillPorts() {
	//args := os.Args
	c := make(chan []string, 1)
	var portList []string
	var wg sync.WaitGroup
	wg.Add(1)


	s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	s.Suffix = "  loading find Port"
	go AllPort(c, &wg)
	
	s.Start()

	portList = <-c
	s.Stop()

	ans := surv.MultiSelectPort(portList)

	for _, v := range ans.Port {
		if v != "" {
			splitSelect := strings.Split(strings.TrimSpace(v), " ")
			PortKill(utils.SplitToInt(utils.GetLastIndex(splitSelect)))
			fmt.Println("kill port ", utils.SplitToInt(utils.GetLastIndex(splitSelect)))
		}
	}
	//wg.Done()

	fmt.Println("Done! ")
	wg.Wait()
}


// port kill
// @arg port int
func PortKill(port int) {
	command := fmt.Sprintf("lsof -i tcp:%d | grep LISTEN | awk '{print $2}' | xargs kill", port)
	fmt.Println("search loading ")
	_, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		fmt.Println("not complete")
		return
	}
	fmt.Printf("complete kill port %d \n", port)
}