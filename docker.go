package main

func docker(args ...string) (string, error) {
	return bash("docker", args)
}
