#!/bin/sh

echo --- Test new archive
curl -sSL -X POST http://localhost:8080/v1/archives | json_pp

echo --- Test get all archives
curl -sSL -X GET  http://localhost:8080/v1/archives | json_pp

echo --- Test get archive by ID
curl -sSL -X GET  http://localhost:8080/v1/archives/0 | json_pp

echo --- Test new backup
curl -sSL -X POST http://localhost:8080/v1/backups | json_pp

echo --- Test get all backups
curl -sSL -X GET  http://localhost:8080/v1/backups | json_pp

echo --- Test get backup by ID
curl -sSL -X GET  http://localhost:8080/v1/backups/0 | json_pp

echo --- Test get backup by invalid ID
curl -sSL -X GET  http://localhost:8080/v1/backups/999 | json_pp

echo --- Test new restoration
curl -sSL -X POST http://localhost:8080/v1/restorations | json_pp

echo --- Test get all restorations
curl -sSL -X GET  http://localhost:8080/v1/restorations | json_pp

echo --- Test get restoration by ID
curl -sSL -X GET  http://localhost:8080/v1/restorations/0 | json_pp

echo --- Test get restoration by invalid ID
curl -sSL -X GET  http://localhost:8080/v1/restorations/9999 | json_pp

echo --- Test new log export request
curl -sSL -X POST http://localhost:8080/v1/logs | json_pp

echo --- Test get all file/disk quotas
curl -sSL -X GET  http://localhost:8080/v1/quotas | json_pp

echo --- Test new synchronization request
curl -sSL -X POST http://localhost:8080/v1/synchronizations | json_pp
