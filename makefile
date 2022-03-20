createdb:
	createdb -p 5432 -U postgres -h localhost book_store

migrateup:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/book_store?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/book_store?sslmode=disable" -verbose down

migratecreate:
	migrate create -ext psql -dir migrations -seq $(MIGRATION_NAME)