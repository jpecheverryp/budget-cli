db/migrate-create:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

db/migrate-up:
	migrate -path ./migrations -database ${BUDGET_CLI_DB_URL} up

