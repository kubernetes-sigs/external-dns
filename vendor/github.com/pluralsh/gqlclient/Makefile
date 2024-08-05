<<<<<<< HEAD
generate: ## resync client with current graph endpoint
	hack/gen-api-client.sh
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
generate-in-container: ## resync client with current graph endpoint
	hack/gen-api-client.sh

update-schema: ## download schema from plural
	curl -L https://github.com/pluralsh/plural/raw/master/schema/schema.graphql --output schema/schema.graphql

generate: update-schema
	go run github.com/Yamashou/gqlgenc
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
