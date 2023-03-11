# Dev version (for simplicity of developing)

include .env

test:
	(cd AuthService && go test -v ./...) &&
	(cd CatalogService && go test -v ./...)

up:
		docker-compose -f docker-compose.dev.yml up --build -d

down:
		docker-compose -f docker-compose.dev.yml down -v

# Prod version

up-prod:
		docker-compose -f docker-compose.prod.yml up --build -d

down-prod:
		docker-compose -f docker-compose.prod.yml down -v