-include .env
export $(shell sed 's/=.*//' .env)

LATEST_DUMP=/tmp/dump_latest.gz
BACKUP?=dump_latest.gz

.ONESHELL:

api:
	docker-compose up -d elastic db
	cd cmd/api && go run .

api-tester:
	docker-compose up -d elastic db
	cd scripts/api_tester && go run .

indexer:
	docker-compose up -d elastic
	cd cmd/indexer && go run .

metrics:
	docker-compose up -d elastic db
	cd cmd/metrics && go run .

seo:
ifeq ($(BCD_ENV), development)
	cd scripts/nginx && go run .
else
	docker-compose exec api seo
	docker-compose restart gui
endif

migration:
ifeq ($(BCD_ENV), development)
	cd scripts/migration && go run .
else
	docker-compose exec api migration
endif

rollback:
ifeq ($(BCD_ENV), development)
	cd scripts/bcdctl && go run . rollback -n $(NETWORK) -l $(LEVEL)
else
	docker-compose exec api bcdctl rollback -n $(NETWORK) -l $(LEVEL)
endif

list-metrics:
ifeq ($(BCD_ENV), development)
	cd scripts/bcdctl && go run . list_services
else
	docker-compose exec api bcdctl list_services
endif

s3-creds:
	docker-compose exec elastic bash -c 'bin/elasticsearch-keystore add --force --stdin s3.client.default.access_key <<< "$$AWS_ACCESS_KEY_ID"'
	docker-compose exec elastic bash -c 'bin/elasticsearch-keystore add --force --stdin s3.client.default.secret_key <<< "$$AWS_SECRET_ACCESS_KEY"'
ifeq ($(BCD_ENV), development)
	cd scripts/bcdctl && go run . reload_secure_settings
else
	docker-compose exec api bcdctl reload_secure_settings
endif

s3-repo:
ifeq ($(BCD_ENV), development)
	cd scripts/bcdctl && go run . create_repository
else
	docker-compose exec api bcdctl create_repository
endif

s3-db-restore:
	echo "Database restore..."
ifeq (,$(wildcard $(LATEST_DUMP)))
	aws s3 cp --profile bcd s3://bcd-db-snaps/$(BACKUP) $(LATEST_DUMP)
endif

	docker-compose exec -T db dropdb -U $(POSTGRES_USER) --if-exists $(POSTGRES_DB)
	gunzip -dc $(LATEST_DUMP) | docker-compose exec -T db psql -U $(POSTGRES_USER) -v ON_ERROR_STOP=on $(POSTGRES_DB)
	rm $(LATEST_DUMP)

s3-es-restore:
	echo "Elasticsearch restore..."
ifeq ($(BCD_ENV), development)
	cd scripts/bcdctl && go run . restore
else
	docker-compose exec api bcdctl restore
endif

s3-contracts-restore:
	echo "Contracts restore..."
ifeq (,$(wildcard /tmp/contracts.tar.gz))
	aws s3 cp --profile bcd s3://bcd-db-snaps/contracts.tar.gz /tmp/contracts.tar.gz
endif	
	rm -rf $(SHARE_PATH)/contracts/
	mkdir $(SHARE_PATH)/contracts/
	tar -C $(SHARE_PATH)/contracts/ -xzf /tmp/contracts.tar.gz
	rm /tmp/contracts.tar.gz

s3-db-snapshot:
	echo "Database snapshot..."
	docker-compose exec db pg_dump $(POSTGRES_DB) --create -U $(POSTGRES_USER) | gzip -c > $(LATEST_DUMP)	
	aws s3 mv --profile bcd $(LATEST_DUMP) s3://bcd-db-snaps/dump_latest.gz

s3-es-snapshot:
	echo "Elasticsearch snapshot..."
ifeq ($(BCD_ENV), development)
	cd scripts/bcdctl && go run . snapshot
else
	docker-compose exec api bcdctl snapshot
endif

s3-contracts-snapshot:
	echo "Packing contracts..."
	cd $(SHARE_PATH)/contracts
	tar -zcvf /tmp/contracts.tar.gz .
	aws s3 mv --profile bcd /tmp/contracts.tar.gz s3://bcd-db-snaps/contracts.tar.gz

s3-list:
	echo "Database snapshots"
	aws s3 ls --profile bcd s3://bcd-db-snaps

	echo "Elasticsearch snapshots"
	aws s3 ls --profile bcd s3://bcd-elastic-snapshots

es-reset:
	docker-compose rm -s -v -f elastic || true
	docker volume rm $$(docker volume ls -q | grep esdata | grep $$COMPOSE_PROJECT_NAME) || true
	docker-compose up -d elastic

test:
	go test ./...

lint:
	golangci-lint run

test-api:
	# to install newman:
	# npm install -g newman
	newman run ./scripts/newman/tests.json -e ./scripts/newman/env.json

docs:
	# wget https://github.com/swaggo/swag/releases/download/v1.7.0/swag_1.7.0_Linux_x86_64.tar.gz
	# tar -zxvf swag_1.7.0_Linux_x86_64.tar.gz
	# sudo cp swag /usr/bin/swag
	cd cmd/api && swag init --parseDependency --parseInternal --parseDepth 2

stable:
	TAG=master docker-compose up -d api metrics indexer

db-dump:
	docker-compose exec db pg_dump -c bcd > dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql

db-restore:
	docker-compose exec db psql --username $$POSTGRES_USER -v ON_ERROR_STOP=on bcd < $(BACKUP)

ps:
	docker ps --format "table {{.Names}}\t{{.RunningFor}}\t{{.Status}}\t{{.Ports}}"

sandbox:
	COMPOSE_PROJECT_NAME=bcdbox TAG=master docker-compose -f docker-compose.sandbox.yml up -d

flextesa-sandbox:
	COMPOSE_PROJECT_NAME=bcdbox TAG=master docker-compose -f docker-compose.flextesa.yml up -d

sandbox-down:
	COMPOSE_PROJECT_NAME=bcdbox docker-compose -f docker-compose.sandbox.yml down

sandbox-clear:
	COMPOSE_PROJECT_NAME=bcdbox docker-compose -f docker-compose.sandbox.yml down -v

gateway:
	COMPOSE_PROJECT_NAME=bcdhub TAG=master docker-compose -f docker-compose.gateway.yml up -d

gateway-down:
	COMPOSE_PROJECT_NAME=bcdhub docker-compose -f docker-compose.gateway.yml down

gateway-clear:
	COMPOSE_PROJECT_NAME=bcdhub docker-compose -f docker-compose.gateway.yml down -v
