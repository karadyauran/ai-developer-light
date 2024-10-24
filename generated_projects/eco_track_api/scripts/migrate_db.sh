#!/bin/bash

set -e

DATABASE_URL="postgres://eco_user:eco_pass@localhost:5432/eco_db?sslmode=disable"

migrate -path ./sql/migrations -database $DATABASE_URL up