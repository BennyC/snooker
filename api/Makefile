migrate-up:
	migrate -database "postgres://postgres:example@0.0.0.0:54321/snooker?sslmode=disable" -path ./db/migrations up
migrate-down:
	migrate -database "postgres://postgres:example@0.0.0.0:54321/snooker?sslmode=disable" -path ./db/migrations down
integration:
	go test ./... -run "Integration"
unit:
	go test ./... -run "Unit"