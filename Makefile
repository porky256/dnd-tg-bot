include build/.env
export

run-db:
	docker-compose up db

stop-db:
	docker-compose stop db


run-air:
	docker-compose up dnd-telegram-air

stop-air:
	docker-compose stop dnd-telegram-air

run:
	docker-compose up dnd-telegram

stop:
	docker-compose stop dnd-telegram

migration-up:
	migrate -path data/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}' up

migration-down:
	migrate -path data/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}' down

migration-fix:
	migrate -path data/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}' force 1
	yes | migrate -path data/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}' down
