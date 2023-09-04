CURRENT_DIR=$(shell pwd)

include .build_info

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: swag_init
swag_init:
	swag init -g internal/transport/handlers/server.go -o api/openapi

.PHONY: mockgen
mockgen:
	mockgen -package mockdb -destination internal/core/repository/psql/mock/store.go github.com/abdukhashimov/go_api/internal/core/repository/psql/sqlc Querier

.PHONY: image_build
image_build:
	docker build -t ${PROJECT_NAME} .

.PHONY: run_dokcer_image
run_docker_image:
	docker run --network=host -e APPLICATION_MODE=dev -e PSQL_URI='postgres://postgres:postgres@localhost:5432/${PROJECT_NAME}' open http --port 9090

.PHONY: dev_environment_remove
dev_environment_remove:
	docker compose -f docker-compose.dev.yml down --volumes

.PHONY: createdb
createdb:
	docker exec -it ${PSQL_CONTAINER_NAME} createdb -U postgres ${PROJECT_NAME}

.PHONY: dropdb
dropdb:
	docker exec -it ${PSQL_CONTAINER_NAME} dropdb -U postgres ${PROJECT_NAME}

.PHONY: migrate_up
migrate_up:
	migrate -path migrations -database ${PSQL_URI} -verbose up

.PHONY: migrate_down
migrate_down:
	migrate -path migrations -database ${PSQL_URI} -verbose down 1

.PHONY: migrate_force
migrate_force:
	migrate -path migrations -database ${PSQL_URI} force ${MIGRATE_FORCE_VERSION}

.PHONY: migrate_file
migrate_file:
	migrate create -ext sql -dir ./migrations -seq ${FILE_NAME}

.PHONY: run
run:
	docker compose -f docker-compose.dev.yml up -d --build

.PHONY: stop
stop:
	docker compose -f docker-compose.dev.yml down

.PHONY: restart	
restart:
	docker compose -f docker-compose.dev.yml down
	docker compose -f docker-compose.dev.yml up -d --build

.PHONY: remove
remove:
	docker compose -f docker-compose.dev.yml down --volumes

.PHONY: logs
logs:
	docker compose -f docker-compose.dev.yml logs -f $(name)

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

copy-proto-module:
	rm -rf ${CURRENT_DIR}/api/protos
	rsync -rv --exclude=.git ${CURRENT_DIR}/ucode_protos/* ${CURRENT_DIR}/api/protos
	
gen-proto-module:
	./scripts/gen-proto.sh ${CURRENT_DIR}