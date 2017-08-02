package main

import (
	"fmt"
	"github.com/labstack/echo"
)

func container_state(name string) int {
	resolution, err := bash(fmt.Sprintf("is_running=$(docker inspect -f {{.State.Running}} %s 2>/dev/null); [ -z $is_running ] && echo void || echo $is_running", name))

	if err != nil {
		return -1
	}

	switch resolution {
	case "true":
		return 1
	case "false":
		return 0
	default:
		return -1
	}
}

func main() {
	ip, err := get_ip()
	if err != nil {
		panic("failed to obtain ip")
	}

	port, err := get_exposed_port()
	if err != nil {
		panic("failed to obtain port")
	}

	fmt.Printf("Advertised IP : %s\n", ip)
	fmt.Printf("Exposed PORT  : %d\n", port)

	e := echo.New()

	e.GET("/containers/:name", containerInspect)
	e.DELETE("/containers/:name", containerRemove)
	e.POST("/containers/:name", containerCreate)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
