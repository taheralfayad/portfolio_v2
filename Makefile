sync-db:
	@DATE=$(shell date +%Y-%m-%d) && \
	echo "=== Starting fresh local containers ===" && \
	docker compose down -v && \
	docker compose up -d && \
	sleep 5 && \
	echo "=== Dumping remote DB ($$DATE) ===" && \
	ssh portfolio "export PATH=\$$PATH:/snap/bin && cd /home/taher/portfolio_v2 && docker compose -f docker-compose.prod.yml exec -T db pg_dump -U taher -d portfolio > /tmp/backup_$$DATE.sql" && \
	echo "=== Copying to local ===" && \
	scp portfolio:/tmp/backup_$$DATE.sql ./backup_$$DATE.sql && \
	echo "=== Restoring local DB ===" && \
	docker compose exec -T db psql -U taher -d portfolio < backup_$$DATE.sql && \
	echo "=== Cleaning up remote ===" && \
	ssh portfolio "export PATH=\$$PATH:/snap/bin && rm /tmp/backup_$$DATE.sql" && \
	echo "Done! Saved as backup_$$DATE.sql"

create-user:
	docker exec -it portfolio-backend-1 go run scripts/create_user.go create

create-migration:
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		echo "Error: MIGRATION_NAME is required"; \
		exit 1; \
	fi
	docker compose run migrate create -ext sql -dir ./migrations -seq $(MIGRATION_NAME)
