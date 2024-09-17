#!/bin/sh

echo --- Test hello world
curl -sSL -X GET  http://localhost:8080/v1/

echo --- Test new archive
curl -sSL -X POST http://localhost:8080/v1/archives | json_pp

echo --- Test get all archives
curl -sSL -X GET  http://localhost:8080/v1/archives | json_pp

echo --- Test get archive by ID
curl -sSL -X GET  http://localhost:8080/v1/archives/0 | json_pp


