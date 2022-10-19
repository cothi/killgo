package port

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/cothi/port/utils"
)

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

func AllPort(c chan<- []string, wg *sync.WaitGroup) {

	// command set
	command := "lsof -i tcp | grep LISTEN | awk '{printf \"%-15s %s| \", $1, $2}'"
	output, err := exec.Command("/bin/sh", "-c", command).Output()

	// error check
	utils.ErrorCheck(err)
	portList := strings.Split(fmt.Sprintf("%s", output), "|")

	// return port list
	c <- portList
	close(c)

	// go sync done
	wg.Done()
}
