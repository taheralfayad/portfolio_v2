create-user:
	docker exec -it portfolio-backend-1 go run scripts/create_user.go create

create-migration:
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		echo "Error: MIGRATION_NAME is required"; \
		exit 1; \
	fi
	docker compose run migrate create -ext sql -dir ./migrations -seq $(MIGRATION_NAME)
