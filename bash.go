package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func bash(first string, args ...interface{}) (string, error) {
	var (
		out []byte
		err error
	)

	if len(args) == 0 {
		fmt.Printf("executing : `%s`\n", first)
		out, err = exec.Command("bash", "-c", first).Output()
	} else {
		switch args[0].(type) {
		case string:
			cmd := args[0].(string)
			fmt.Printf("executing : `%s %s`\n", first, cmd)
			out, err = exec.Command(first, cmd).Output()
		case []string:
			a := args[0].([]string)
			cmd := make([]string, len(a))
			for i, v := range a {
				cmd[i] = fmt.Sprint(v)
			}

			fmt.Printf("executing : `%s %v`\n", first, cmd)
			out, err = exec.Command(first, cmd...).Output()
		}
	}

	if len(out) == 0 {
		return "", err
	}
	return string(out[:len(out)-1]), err
}

func get_exposed_port() int {
	var (
		out  string
		err  error
		port int
	)

	out, err = bash("docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{(index $conf 0).HostPort}}{{end}}' $(cat /etc/hosts | tail -1 | cut -f2) 2> /dev/null || echo 0")

	if err != nil {
		os.Exit(2)
	}

	port, err = strconv.Atoi(out)
	if err != nil {
		os.Exit(2)
	}

	return port
}

func get_ip() string {
	var (
		ip  string
		err error
	)

	ip, err = bash("docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(cat /etc/hosts | tail -1 | cut -f2) 2> /dev/null || echo 0.0.0.0")

	if err != nil {
		os.Exit(2)
	}

	return ip
}
