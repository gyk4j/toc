#!/bin/sh

IP=localhost
PORT=8080
API=v1
BASE=http://$IP:$PORT/$API

echo --- Summary
echo $BASE

echo --- Test new backup
curl -s -X POST $BASE/backups | json_pp

echo --- Test get all backups
curl -s -X GET  $BASE/backups | json_pp

echo --- Test get backup by ID
curl -s -X GET  $BASE/backups/0 | json_pp

echo --- Test get backup by invalid ID
curl -sI -X GET  $BASE/backups/9999

echo --- Test new restoration
curl -s -X POST $BASE/restorations \
  -d "$(curl -s -X GET  $BASE/backups/0)" \
  -H "Content-Type: application/json" \
  | json_pp

echo --- Test get all restorations
curl -s -X GET  $BASE/restorations | json_pp

echo --- Test get restoration by ID
curl -s -X GET  $BASE/restorations/0 | json_pp

echo --- Test get restoration by invalid ID
curl -sI -X GET  $BASE/restorations/9999

echo --- Test new transfer
curl -s -X POST $BASE/transfers \
  -d "$(curl -s -X GET  $BASE/backups/0)" \
  -H "Content-Type: application/json" \
  | json_pp

echo --- Test get all transfers
curl -s -X GET  $BASE/transfers | json_pp

echo --- Test get transfer by ID
curl -s -X GET  $BASE/transfers/0 | json_pp

echo --- Test get transfer by invalid ID
curl -sI -X GET  $BASE/transfers/9999

