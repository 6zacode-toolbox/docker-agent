#!/bin/sh

docker ps  --format "{{json .}}" | jq -s