build:
	docker compose build

run:
	docker compose up

down:
	docker compose down

# sample usage: make create-migration migration-name=add_user_table
create-migration:
	goose -dir ./backend/storage/migrations create $(migration-name) sql

goose-up:
	goose -dir ./backend/storage/migrations postgres "user=postgres password=postgres host=localhost port=5434 dbname=hackutd2024 sslmode=disable" up

goose-down:
	goose -dir ./backend/storage/migrations postgres "user=postgres password=postgres host=localhost port=5434 dbname=hackutd2024 sslmode=disable" down

# only works if you have psql installed. If so, password is postgres
psql:
	psql -p 5434 -U postgres -d hackutd2024