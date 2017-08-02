package main

import (
	"fmt"
	"os/exec"
)

func bash(first string, args ...interface{}) (string, error) {
	var (
		out []byte
		err   error
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