PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

.PHONY: init-docsify
init-docsify:
	@docsify init ./docs/guide

.PHONY: swagger
swagger:
	@swag init -g cmd/api/main.go -o ./docs/guide/
	@rm -f docs/guide/swagger.yaml
	@curl -X POST "https://converter.swagger.io/api/convert" -H "Accept:application/yaml" -H "Content-Type:application/yaml" --data-binary @docs/swagger.yaml >> docs/guide/swagger.yaml