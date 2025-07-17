createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=internal/sql/migrations -database "postgres://root:root@localhost:5432/agent_calendar?sslmode=disable" -verbose up

migratedown:
	migrate -path=internal/sql/migrations -database "postgres://root:root@localhost:5432/agent_calendar?sslmode=disable" -verbose down

.PHONY: migrate migratedown createmigration