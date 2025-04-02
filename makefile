include .env

migration:
	@if [ -z "$(name)" ]; then \
		echo "No name parameter passed. Try: make migration name='create_user_table'"; \
	else \
		migrate -source $(DB_URL) create -ext sql -dir "$(MIGRATION_PATH)" -seq $(name); \
	fi

migrate-up:
	migrate -verbose -database ${DB_URL} -path $(MIGRATION_PATH) up

migrate-down:
	migrate -verbose -database ${DB_URL} -path $(MIGRATION_PATH) down
