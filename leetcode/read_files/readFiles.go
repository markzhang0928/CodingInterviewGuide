package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func parseCgroupFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", nil
	}
	defer f.Close()

	return parseCgroupFromReader(f)
}

func parseCgroupFromReader(r io.Reader) (string, error) {
	var (
		s = bufio.NewScanner(r)
	)
	for s.Scan() {
		if err := s.Err(); err != nil {
			return "", err
		}
		var (
			text  = s.Text()
			parts = strings.SplitN(text, ":", 3)
		)
		if len(parts) < 3 {
			return "", fmt.Errorf("invalid cgroup entry: %q", text)
		}
		// text is like "0::/user.slice/user-1001.slice/session-1.scope"
		if parts[0] == "0" && parts[1] == "" {
			return parts[2], nil
		}
	}
	return "", errors.New("cgroup path not found")
}

func main() {

	parseCgroupFile("/proc/self/cgroup")
}
