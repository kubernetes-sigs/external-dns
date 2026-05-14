generate-in-container: ## resync client with current graph endpoint
	hack/gen-api-client.sh

update-schema: ## download schema from plural
	curl -L https://github.com/pluralsh/plural/raw/master/schema/schema.graphql --output schema/schema.graphql

generate: update-schema
	go run github.com/Yamashou/gqlgenc
