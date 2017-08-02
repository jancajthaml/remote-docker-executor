package main

import (
	"fmt"

	"os"

    "strconv"

	"github.com/labstack/echo"
)

func container_state(name string) int {
	resolution, err := bash(fmt.Sprintf("is_running=$(docker inspect -f {{.State.Running}} %s 2>/dev/null); [ -z $is_running ] && echo void || echo $is_running", name))

    if err != nil {
    	fmt.Printf("Error when checking container state : %s\n", err)
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

func get_exposed_port() int {
    out, err1 := bash("get_exposed_port")

    if err1 != nil {
    	os.Exit(2)
    }

    port, err2 := strconv.Atoi(out)
    if err2 != nil {
        os.Exit(2)
    }

    return port
}

func get_ip() string {
    out, err := bash("get_ip")

    if err != nil {
    	os.Exit(2)
    }

    return out
}


func main() {
	ip := get_ip()
    port := get_exposed_port()

    fmt.Printf("Advertised IP : %s\n", ip)
    fmt.Printf("Exposed PORT  : %d\n", port)

	e := echo.New()

	e.GET("/containers/:name", containerInspect)
	e.DELETE("/containers/:name", containerRemove)
	e.POST("/containers/:name", containerCreate)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}