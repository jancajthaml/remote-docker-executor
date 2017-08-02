Remote docker executor
======================

Manage docker containers with remote executor via HTTP API.

## Start

`docker run -p 8181:8181 -v /var/run/docker.sock:/var/run/docker.sock jancajthaml/remote-docker-executor`

## Containers management
### Create Container

`curl -X POST -H "Content-Type: application/json" -d '{
  "image": "golang:1.7",
  "cmd": ["bash"],
  "volumes": ["/var/run/docker.sock:/var/run/docker.sock", "/dev/shm:/dev/shm"]
}' "http://localhost:8181/containers/${SERVICE_NAME}"`

### Inspect Container

`curl "http://localhost:8181/containers/${SERVICE_NAME}"`

### Remove Container

`curl -X DELETE "http://localhost:8181/containers/${SERVICE_NAME}"`
