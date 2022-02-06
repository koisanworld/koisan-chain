SYSTEM = $(shell uname -s)
HOST = host.docker.internal
DOCKER_IMAGE = compverse/blockscout_prod:latest
BS_CONTAINER_NAME = blockscout
PG_CONTAINER_NAME = postgres
PG_CONTAINER_IMAGE = postgres:12.5
THIS_FILE = $(lastword $(MAKEFILE_LIST))

ifeq ($(SYSTEM), Linux)
	HOST=localhost
endif

DB_URL = postgresql://postgres:ethereum@$(HOST):5432/explorer?ssl=false
BLOCKSCOUT_CONTAINER_PARAMS = -e 'MIX_ENV=prod' \
							  -e 'DATABASE_URL=$(DB_URL)'
ifeq ($(SYSTEM), Linux)
	BLOCKSCOUT_CONTAINER_PARAMS += --network=host
endif
ifdef NETWORK
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'NETWORK=$(NETWORK)'
endif
ifdef SUBNETWORK
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'SUBNETWORK=$(SUBNETWORK)'
endif
ifdef LOGO_TEST
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'LOGO_TEST=$(LOGO_TEST)'
endif
ifdef ETHEREUM_JSONRPC_VARIANT
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'ETHEREUM_JSONRPC_VARIANT=$(ETHEREUM_JSONRPC_VARIANT)'
endif
ifdef ETHEREUM_JSONRPC_HTTP_URL
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'ETHEREUM_JSONRPC_HTTP_URL=$(ETHEREUM_JSONRPC_HTTP_URL)'
endif
ifdef ETHEREUM_JSONRPC_TRACE_URL
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'ETHEREUM_JSONRPC_TRACE_URL=$(ETHEREUM_JSONRPC_TRACE_URL)'
endif
ifdef ETHEREUM_JSONRPC_WS_URL
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'ETHEREUM_JSONRPC_WS_URL=$(ETHEREUM_JSONRPC_WS_URL)'
endif
ifdef COIN
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'COIN=$(COIN)'
endif
ifdef BLOCK_TRANSFORMER
	BLOCKSCOUT_CONTAINER_PARAMS += -e 'BLOCK_TRANSFORMER=$(BLOCK_TRANSFORMER)'
endif
ifdef PORT
    BLOCKSCOUT_CONTAINER_PARAMS += -e 'PORT=$(PORT)'
endif
ifdef PG_DATADIR
  PG_CONTAINER_PARAMS  += -v $(PG_DATADIR):/var/lib/postgresql/data
endif

HAS_BLOCKSCOUT_IMAGE := $(shell docker images | grep -sw ${DOCKER_IMAGE})

migrate_only:
	@echo "==> Running migrations"
	@docker run --rm \
					$(BLOCKSCOUT_CONTAINER_PARAMS) \
					$(DOCKER_IMAGE) /bin/sh -c "echo $$MIX_ENV && mix do ecto.create, ecto.migrate"

migrate: postgres
	@$(MAKE) -f $(THIS_FILE) migrate_only


PG_EXIST := $(shell docker ps -a --filter name=${PG_CONTAINER_NAME} | grep ${PG_CONTAINER_NAME})
PG_STARTED := $(shell docker ps --filter name=${PG_CONTAINER_NAME} | grep ${PG_CONTAINER_NAME})

postgres:
ifdef PG_EXIST
	@echo "==> Checking PostrgeSQL container"
ifdef PG_STARTED
	@echo "==> PostgreSQL Already started"
	@$(MAKE) -f $(THIS_FILE) migrate_only
else
	@echo "==> Starting PostgreSQL container"
	@docker start $(PG_CONTAINER_NAME)
	@$(MAKE) -f $(THIS_FILE) migrate_only
endif
else
	@echo "==> Creating new PostgreSQL container"
	@docker run -d --name $(PG_CONTAINER_NAME) \
					-e POSTGRES_PASSWORD="ethereum" \
					-e POSTGRES_USER="postgres" \
					-e POSTGRES_HOST_AUTH_METHOD="trust" \
					-p 5432:5432 \
					$(PG_CONTAINER_PARAMS) \
					$(PG_CONTAINER_IMAGE)
	@sleep 1
	@$(MAKE) -f $(THIS_FILE) migrate_only
endif

start: postgres
	@echo "==> Starting blockscout"
	@docker run -d --name $(BS_CONTAINER_NAME) \
					$(BLOCKSCOUT_CONTAINER_PARAMS) \
					$(DOCKER_IMAGE) /bin/sh -c "mix phx.server"

run: start

.PHONY: migrate \
		start \
		postgres \
		run


