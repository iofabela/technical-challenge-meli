PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

.PHONY: init-docsify
init-docsify:
	@docsify init ./docs/guide

.PHONY: swagger
swagger:
	@swag init -g cmd/api/main.go -o ./docs/guide/
	@rm -f docs/guide/swagger.yaml
	@curl -X POST "https://converter.swagger.io/api/convert" -H "Accept:application/yaml" -H "Content-Type:application/yaml" --data-binary @docs/swagger.yaml >> docs/guide/swagger.yaml

.PHONY: start # para construir la imagen y abrir el docsify
start:
	docker build -t gin-sqlite-app .
	docsify serve ./docs/guide/ --open

.PHONY: run # para ejecutar el contenedor
run:
	docker run -p 8080:8080 -v $(pwd)/items.db:/app/items.db gin-sqlite-app
	open http://localhost:8080/docs/index.html

.PHONY: get-data-sql
get-data-sql:
	sqlite3 items.db -cmd "SELECT * FROM items;" ".exit"

.PHONY: stop # para detener el contenedor
stop:
	docker stop gin-sqlite-app
	npx kill-port 3000

