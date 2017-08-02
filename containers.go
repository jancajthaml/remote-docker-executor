package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

type Container struct {
	Image  string `json:"image"`
	Command []string `json:"cmd"`
	Volumes []string `json:"volumes"`
}

func containerInspect(c echo.Context) error {
	name := c.Param("name")
	out, err := docker("inspect", name)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	return c.String(http.StatusOK, out)
}

func containerRemove(c echo.Context) error {
	name := c.Param("name")
	out, err := docker("rm", "--force", name)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	return c.String(http.StatusOK, out)
}

func containerCreate(c echo.Context) error {
	name := c.Param("name")
	container := new(Container)

	if err := c.Bind(container); err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	d:
	switch state := container_state(name); state {
	case 0: // restart
		_, err2 := docker("restart", name)
		if err2 != nil {
			fmt.Printf("Error while restarting container : %s\n", err2)
			return c.String(http.StatusBadRequest, "failed to restart container")	
		}
		break d
	case 1: // already running
		break d
	default: // create
		var cmd = []string { "run", "--name", name }
		for _, volume := range container.Volumes {
			cmd = append(cmd, "-v", volume)
		}
		cmd = append(cmd, "-d", container.Image)
		cmd = append(cmd, container.Command...)

		_, err2 := docker(cmd...)
		if err2 != nil {
			fmt.Printf("Error while creating container : %s\n", err2)
			return c.String(http.StatusBadRequest, "failed to start container")	
		}
		break d
	}

	return c.String(http.StatusOK, "")
}