package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func bash(first string, args ...interface{}) (string, error) {
	var (
		out []byte
		err error
	)

	if len(args) == 0 {
		fmt.Printf("$ : `%s`\n", first)
		out, err = exec.Command("bash", "-c", first).Output()
	} else {
		switch args[0].(type) {
		case string:
			cmd := args[0].(string)
			fmt.Printf("$ : `%s %s`\n", first, cmd)
			out, err = exec.Command(first, cmd).Output()
		case []string:
			a := args[0].([]string)
			cmd := make([]string, len(a))
			for i, v := range a {
				cmd[i] = fmt.Sprint(v)
			}

			fmt.Printf("$ : `%s %v`\n", first, cmd)
			out, err = exec.Command(first, cmd...).Output()
		}
	}

	if len(out) == 0 {
		return "", err
	}
	return string(out[:len(out)-1]), err
}

func get_exposed_port() (int, error) {
	var (
		out  string
		err  error
		port int
	)

	out, err = bash("docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{(index $conf 0).HostPort}}{{end}}' $(cat /etc/hosts | tail -1 | cut -f2) 2> /dev/null || echo 0")

	if err != nil {
		return 0, err
	}

	port, err = strconv.Atoi(out)
	if err != nil {
		return 0, err
	}

	return port, err
}

func get_ip() (string, error) {
	return bash("docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(cat /etc/hosts | tail -1 | cut -f2) 2> /dev/null || echo 0.0.0.0")
}
