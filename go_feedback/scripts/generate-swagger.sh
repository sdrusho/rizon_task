#!/bin/bash
swag init -g cmd/feedback/main.go --exclude "./internal/db/generated"
chmod +x scripts/generate-swagger.sh
