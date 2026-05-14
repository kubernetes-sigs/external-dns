# Egoscale v3

Exoscale API Golang wrapper

**Egoscale v3** is based on a generator written from scratch with [libopenapi](https://github.com/pb33f/libopenapi).

The core base of the generator is using libopenapi to parse and read the [Exoscale OpenAPI spec](https://api-ch-gva-2.exoscale.com/v2/openapi.json) and then generate the code from it.

## Installation

Install the following dependencies:

```shell
go get "github.com/exoscale/egoscale/v3"
```

Add the following import:

```golang
import "github.com/exoscale/egoscale/v3"
```
## Examples

```Golang
package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	v3 "github.com/exoscale/egoscale/v3"
	"github.com/exoscale/egoscale/v3/credentials"
)

func main() {
	creds := credentials.NewEnvCredentials()
	// OR
	creds = credentials.NewStaticCredentials("EXOxxx..", "...")

	client, err := v3.NewClient(creds)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	op, err := client.CreateInstance(ctx, v3.CreateInstanceRequest{
		Name:     "egoscale-v3",
		DiskSize: 50,
		// Ubuntu 24.04 LTS
		Template: &v3.Template{ID: v3.UUID("cbd89eb1-c66c-4637-9483-904d7e36c318")},
		// Medium type
		InstanceType: &v3.InstanceType{ID: v3.UUID("b6e9d1e8-89fc-4db3-aaa4-9b4c5b1d0844")},
	})
	if err != nil {
		log.Fatal(err)
	}

	op, err = client.Wait(ctx, op, v3.OperationStateSuccess)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := client.GetInstance(ctx, op.Reference.ID)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(instance)
}
```

### Findable

Most of the list request `ListX()` return a type containing the list of the resource requested and a method `FindX()` to be able to retrieve a resource by its `name` or `id` most of the time.

```Golang
pools, err := client.ListInstancePools(ctx)
if err != nil {
	log.Fatal(err)
}
pool, err := pools.FindInstancePool("my-pool-example")
if err != nil {
	log.Fatal(err)
}

fmt.Println(pool.Name)
```

## Development

### Generate Egoscale v3

From the root repo
```Bash
make pull-oapi-spec # Optional(to pull latest Exoscale Open-API spec)
make generate
```

### Debug generator output

```Bash
mkdir test
GENERATOR_DEBUG=client make generate > test/client.go
GENERATOR_DEBUG=schemas make generate > test/schemas.go
GENERATOR_DEBUG=operations make generate > test/operations.go
```

### OpenAPI Extensions

The generator support two types of extension:
- `x-go-type` to specify a type definition in Golang.

	OpenAPI Spec
	```yaml
	api-endpoint:
	  type: string
	  x-go-type: Endpoint
	  description: Zone API endpoint
	```
	Generated code
	```Golang
	type Endpoint string

	type Zone struct {
		APIEndpoint Endpoint // Here is the generated type definition.
		...
	}
	```
- `x-go-findable` to specify which fields in the findable resource to fetch
	OpenAPI Spec
	```yaml
	elastic-ip:
      type: object
      properties:
        id:
          type: string
		  x-go-findable: "1"
          description: Elastic IP ID
        ip:
          type: string
		  x-go-findable: "2"
          description: Elastic IP address
	```
	Generated code
	```Golang
	// FindElasticIP attempts to find an ElasticIP by idOrIP.
	func (l ListElasticIPSResponse) FindElasticIP(idOrIP string) (ElasticIP, error) {
		for i, elem := range l.ElasticIPS {
			if string(elem.ID) == idOrIP || string(elem.IP) == idOrIP {
				return l.ElasticIPS[i], nil
			}
		}

		return ElasticIP{}, fmt.Errorf("%q not found in ListElasticIPSResponse: %w", idOrIP, ErrNotFound)
	}
	```

## Generator Overrides System

The Egoscale v3 generator incorporates an overrides system to preserve backwards compatibility in the Go API when the OpenAPI specification changes, such as renaming schemas or references. This ensures that existing code using the SDK does not break due to type, field name, or JSON tag changes.

The system consists of a unified overrides system working during code generation, scoped to schemas with overrides to avoid affecting new APIs.

### Components

#### Schema Property and Reference Overrides
- **Location**: `generator/helpers/helpers.go` (`SchemaPropertyOverrides` map)
- **Purpose**: Unified map for property names and reference paths in legacy schemas to preserve backwards compatibility. Schemas with `Refs` are treated as legacy (applying ref overrides); others use spec defaults. Property overrides are applied per-schema as needed.
- **Example**:
  ```go
  var SchemaPropertyOverrides = map[string]*Overrides{
      "AttachInstanceToElasticIPRequest": {
          Props: map[string]string{
              "instance-target": "instance",
          },
          Refs: map[string]string{
              "#/components/schemas/ssh-key-ref": "SSHKey",
              "#/components/schemas/instance-ref": "InstanceTarget",
          },
      },
      "CreateInstance": {
          Props: nil,
          Refs: map[string]string{
              // ... ref overrides map
          },
      },
  }
  ```
  - In `AttachInstanceToElasticIPRequest`, a property `"instance-target"` generates `Instance` with JSON `"instance"`, and refs like `"#/components/schemas/ssh-key-ref"` generate `SSHKey`.
  - Schemas without entries or with `Refs: nil` use spec defaults.

#### Backwards Compatibility Full Types
- **Purpose**: Generates full struct definitions for old type names (e.g., `Template`) directly from the OpenAPI spec to maintain backwards compatibility, ensuring old types have complete field sets for operations like listing.
- **Generation**: All schemas in the spec are generated, including full types (e.g., `Template` with 18 fields from `template`) and ref types (e.g., `TemplateRef` with 1 field from `template-ref`).

#### Special Aliases
- **Location**: `generator/schemas/schemas.go` (hardcoded in `Generate` function)
- **Purpose**: Adds specific type aliases for complex backwards compatibility cases where aliasing is appropriate.
- **Example**:
  ```go
  type InstanceTarget = InstanceRef
  type BlockStorageSnapshotTarget = BlockStorageSnapshotRef
  type BlockStorageVolumeTarget = BlockStorageVolumeRef
  ```
  - Provides aliases like `InstanceTarget` for operations requiring distinct field names.

### How It Works
1. **Schema Processing**: Generates all schemas from the OpenAPI spec in alphabetical order, including full types (e.g., `Template`) and ref types (e.g., `TemplateRef`).
2. **Conditional Application**: For schemas with overrides in `SchemaPropertyOverrides`, applies `Refs` (if present) and `Props` (if present); otherwise, uses spec defaults.
3. **Type and Field Generation**: References and properties are resolved with overrides, producing consistent Go structs.
4. **Special Aliases Addition**: Special aliases are appended globally.

### Usage in Development
- Modify `SchemaPropertyOverrides` in `generator/helpers/helpers.go` for spec updates (add/update `Props` or `Refs` per schema).
- Regenerate with `make generate`, build with `go build ./...`, and inspect with `GENERATOR_DEBUG=schemas make generate`.

This system enables the SDK to evolve with the OpenAPI spec while preserving API stability for existing users.
