Remote docker executor
======================

Manage docker containers with remote executor via HTTP API.

## Start

`docker run -p 8181:8181 -v /var/run/docker.sock:/var/run/docker.sock -v /etc:/etc jancajthaml/docker_executor`

## Create service

`curl -X POST -H "Content-Type: application/json" -d '{"image": "nginx", "port": 8080}' "http://localhost:8181/services/${SERVICE_NAME}"`

## Inspect Service

`curl "http://localhost:8181/services/${SERVICE_NAME}"`

## Remove Service

`curl -X DELETE "http://localhost:8181/services/${SERVICE_NAME}"`
