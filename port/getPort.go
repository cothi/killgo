package port

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/cothi/port/utils"
)


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
