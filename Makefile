.PHONY: run
run:
	go1.17 run -gcflags="-G=3" . > README.md

.PHONY: setup
setup:
	go generate ./tools.go
	psql -q -h localhost -U $(DATABASE_USER) -c "drop database if exists $(DATABASE_NAME);"
	psql -q -h localhost -U $(DATABASE_USER) -c "create database $(DATABASE_NAME);"
	./bin/psqldef -h localhost -U $(DATABASE_USER) $(DATABASE_NAME) < schema.sql
