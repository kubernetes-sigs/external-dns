```yaml
---
title: "FQDN Package Enhancement"
version: v1alpha1
authors: @ivankatliarchuk
creation-date: 2025-12-31
status: draft
tags: ["proposal", "fqdn", "templating"]
cmd: ["markdown-toc-creator docs/proposal/005-fqdn-template-enhancement.md --add-horizontal-rules=false"]
---
```

# FQDN Package Enhancement

## Table of Contents

<!--TOC-->

**Table of Contents**

- [FQDN Package Enhancement](#fqdn-package-enhancement)
  - [Table of Contents](#table-of-contents)
  - [Summary](#summary)
    - [Enhancement Overview](#enhancement-overview)
  - [Motivation](#motivation)
    - [Current Problem](#current-problem)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
  - [Proposal](#proposal)
    - [Current Limitations](#current-limitations)
      - [Limitation 1: Template Parsed Per Source Constructor](#limitation-1-template-parsed-per-source-constructor)
      - [Limitation 2: Two Different Constructor Patterns](#limitation-2-two-different-constructor-patterns)
      - [Limitation 3: No Type Safety](#limitation-3-no-type-safety)
      - [Limitation 4: Single Template for All Sources](#limitation-4-single-template-for-all-sources)
      - [Limitation 5: No DNS Validation](#limitation-5-no-dns-validation)
      - [Limitation 6: No Template Documentation](#limitation-6-no-template-documentation)
      - [Limitation 7: Limited Template Functions](#limitation-7-limited-template-functions)
      - [Limitation 8: No Template Testing Support](#limitation-8-no-template-testing-support)
    - [User Stories](#user-stories)
      - [Story 1: Platform Engineer Debugging Template Errors](#story-1-platform-engineer-debugging-template-errors)
      - [Story 2: Developer Adding FQDN Support to New Source](#story-2-developer-adding-fqdn-support-to-new-source)
      - [Story 3: SRE Implementing Template Validation](#story-3-sre-implementing-template-validation)
  - [Enhancement 1: Template Interface & Registry (High Priority)](#enhancement-1-template-interface--registry-high-priority)
    - [Motivation](#motivation-1)
    - [Design](#design)
      - [Architecture Overview](#architecture-overview)
      - [Package Structure](#package-structure)
  - [Package Structure](#package-structure-1)
    - [Migration Strategy](#migration-strategy)
      - [Phase 1: Infrastructure](#phase-1-infrastructure)
      - [Phase 2: Source Migration](#phase-2-source-migration)
      - [Phase 3: Verification](#phase-3-verification)
    - [Behavior](#behavior)
      - [Normal Operation](#normal-operation)
      - [Thread Safety](#thread-safety)
      - [Memory Impact](#memory-impact)
    - [Drawbacks](#drawbacks)
  - [Alternatives](#alternatives)
    - [Alternative 1: Keep Current Approach (Do Nothing)](#alternative-1-keep-current-approach-do-nothing)
    - [Alternative 2: Lazy Initialization](#alternative-2-lazy-initialization)
    - [Alternative 3: Config-Level Caching](#alternative-3-config-level-caching)
    - [Alternative 5: Combine Registry + Dependency Injection](#alternative-5-combine-registry--dependency-injection)
    - [API](#api)
      - [Core Interfaces](#core-interfaces)
      - [Template Creation](#template-creation)
      - [Registry API](#registry-api)
  - [Enhancement 2: Multi-Template Support (High Priority)](#enhancement-2-multi-template-support-high-priority)
    - [Motivation](#motivation-2)
    - [Design](#design-1)
    - [Configuration](#configuration)
    - [Implementation Example](#implementation-example)
      - [**Usage Example**](#usage-example)
    - [Benefits](#benefits)
    - [User Stories](#user-stories-1)
      - [Story 4: Multi-Zone DNS Management](#story-4-multi-zone-dns-management)
      - [Story 5: Environment-Based Templates with Alternates](#story-5-environment-based-templates-with-alternates)
      - [Story 6: DNS Migration](#story-6-dns-migration)
  - [Enhancement 3: DNS Validation (Medium Priority)](#enhancement-3-dns-validation-medium-priority)
    - [Motivation](#motivation-3)
    - [Design](#design-2)
    - [Implementation](#implementation)
    - [Configuration](#configuration-1)
    - [Integration Points](#integration-points)
    - [Usage Examples](#usage-examples)
    - [Benefits](#benefits-1)
    - [User Stories](#user-stories-2)
      - [Story 7: Production DNS Validation](#story-7-production-dns-validation)
      - [Story 8: International Service Names](#story-8-international-service-names)
  - [Enhancement 4: Auto-Generated Documentation (Medium Priority)](#enhancement-4-auto-generated-documentation-medium-priority)
    - [Motivation](#motivation-4)
    - [Design](#design-3)
      - [Package Structure](#package-structure-2)
      - [Source Interfaces](#source-interfaces)
      - [Godoc-Style Comments](#godoc-style-comments)
      - [Structured Test Tags](#structured-test-tags)
    - [Implementation](#implementation-1)
      - [Generator Tool](#generator-tool)
      - [Extract Examples from Tests](#extract-examples-from-tests)
      - [Scan Source Interfaces](#scan-source-interfaces)
      - [Function Reference from Godoc](#function-reference-from-godoc)
    - [Test Enforcement](#test-enforcement)
    - [Usage](#usage)
    - [Example Output](#example-output)
    - [Benefits](#benefits-2)
  - [Enhancement 5: Advanced Template Functions (Low Priority)](#enhancement-5-advanced-template-functions-low-priority)
    - [Motivation](#motivation-5)
      - [Current Limitation](#current-limitation)
    - [Design](#design-4)
      - [Function Organization](#function-organization)
      - [Core Functions Module](#core-functions-module)
      - [String Functions](#string-functions)
      - [Conditional Functions](#conditional-functions)
      - [DNS Functions](#dns-functions)
      - [Encoding Functions](#encoding-functions)
    - [Configuration](#configuration-2)
    - [Deprecation Plan](#deprecation-plan)
    - [Benchmarking](#benchmarking)
    - [Usage Examples](#usage-examples-1)
    - [Future Considerations](#future-considerations)
  - [Enhancement 6: FQDN Template Execution in Informer SetTransform (Medium Priority)](#enhancement-6-fqdn-template-execution-in-informer-settransform-medium-priority)
    - [Motivation](#motivation-6)
      - [Current Limitation](#current-limitation-1)
    - [Design](#design-5)
      - [Core Concept](#core-concept)
      - [Annotation Storage](#annotation-storage)
      - [Transform Implementation](#transform-implementation)
      - [Reading FQDNs in Endpoints()](#reading-fqdns-in-endpoints)
    - [Re-execution Behavior](#re-execution-behavior)
    - [Configuration](#configuration-3)
    - [Implementation](#implementation-2)
      - [Apply to All Sources](#apply-to-all-sources)
      - [Helper Functions](#helper-functions)
    - [Performance Testing](#performance-testing)
    - [Benefits](#benefits-3)
    - [Usage Examples](#usage-examples-2)
    - [Performance Goals](#performance-goals)
  - [Enhancement 7: Target FQDN Template Support (Medium Priority)](#enhancement-7-target-fqdn-template-support-medium-priority)
    - [Motivation](#motivation-7)
      - [Current Limitation](#current-limitation-2)
    - [Design](#design-6)
      - [Annotation-Based Target Templates](#annotation-based-target-templates)
      - [Configuration-Based Target Templates](#configuration-based-target-templates)
      - [Implementation](#implementation-3)
    - [Configuration](#configuration-4)
      - [CLI Flags](#cli-flags)
      - [YAML Configuration](#yaml-configuration)
      - [Annotations](#annotations)
    - [Validation](#validation)
    - [Benefits](#benefits-4)
    - [Usage Examples](#usage-examples-3)
    - [Error Handling](#error-handling)
    - [Integration with Enhancement 6 (SetTransform)](#integration-with-enhancement-6-settransform)
    - [Testing](#testing)
- [Implementation Examples](#implementation-examples)
  - [Enhancement 1: Template Interface and Registry](#enhancement-1-template-interface-and-registry)
  - [Enhancement 2: Multi-Template Support with Selectors](#enhancement-2-multi-template-support-with-selectors)
  - [Enhancement 3: DNS Validation and Sanitization](#enhancement-3-dns-validation-and-sanitization)
  - [Enhancement 4: Auto-Generated Documentation](#enhancement-4-auto-generated-documentation)
  - [Enhancement 5: Advanced Template Functions](#enhancement-5-advanced-template-functions)
  - [Enhancement 6: FQDN Template Execution in Informer SetTransform](#enhancement-6-fqdn-template-execution-in-informer-settransform)
  - [Enhancement 7: Target FQDN Template Support](#enhancement-7-target-fqdn-template-support)

<!--TOC-->

## Summary

This proposal introduces comprehensive enhancements to the `fqdn` package in external-dns to address current limitations and provide powerful new capabilities for FQDN template management.
The enhancements are organized by priority and build upon each other to create a robust, flexible, and maintainable template system.

### Enhancement Overview

<!-- TODO: update this as well -->

| # | Enhancement | Priority | Description |
|---|-------------|----------|-------------|
| 1 | Template Interface & Registry | 🔴 High | Parse templates once, reuse everywhere |
| 2 | Multi-Template Support | 🔴 High | Different templates per source/condition |
| 3 | DNS Validation | 🟡 Medium | Prevent invalid hostnames in production |
| 4 | Auto-Generated Documentation | 🟡 Medium | Generate docs from code and tests |
| 5 | Advanced Template Functions | 🟢 Low | Extended template function library |
| 6 | Template Testing Framework | 🟢 Low | Testing utilities and helpers |

**Key Benefits**:

- Parse templates once instead of 10+ times (Enhancement 1)
- Different templates for different sources/scenarios (Enhancement 2)
- Automatic DNS validation and sanitization (Enhancement 3)
- Always up-to-date documentation (Enhancement 4)
- Powerful template capabilities (Enhancement 5)
- Easy template testing (Enhancement 6)
- Easier integration of FQDN templating with every source
- Foundation for future extensibility

## Motivation

### Current Problem

FQDN templates are parsed repeatedly across the codebase:

```go
// service.go:104
tmpl, err := fqdn.ParseTemplate(fqdnTemplate)

// ingress.go:76
tmpl, err := fqdn.ParseTemplate(fqdnTemplate)

// gateway.go:126
tmpl, err := fqdn.ParseTemplate(config.FQDNTemplate)

// ... 10+ more sources
```

**Verified occurrences** (grep results):

- source/node.go:60
- source/contour_httpproxy.go:65
- source/istio_gateway.go:76
- source/service.go:104
- source/gateway.go:126
- source/ingress.go:76
- source/istio_virtualservice.go:78
- source/skipper_routegroup.go:197
- source/openshift_route.go:70
- source/pod.go:123

**Impact**:

1. **Performance**: Same template parsed 10+ times at startup
2. **Maintainability**: Error handling duplicated across sources
3. **Extensibility**: Adding features (validation, caching) requires changing all sources
4. **Testing**: Difficult to mock/test template behavior
5. **Type Safety**: Direct use of `*text/template.Template` instead of domain-specific interface

### Goals

1. **Parse Once** (Enhancement 1): Templates parsed exactly once at startup and registered globally
2. **Centralized Management** (Enhancement 1): Single location for template registration, validation, and retrieval
3. **Type Safety** (Enhancement 1): Introduce domain-specific `Template` interface hiding implementation details
4. **Multi-Template Support** (Enhancement 2): Enable different templates for different sources, namespaces, or labels
5. **DNS Validation** (Enhancement 3): Automatically validate and sanitize generated hostnames to prevent invalid DNS records
6. **Auto-Documentation** (Enhancement 4): Generate documentation from code, tests, and source analysis
7. **Extended Functions** (Enhancement 5): Provide comprehensive template function library for complex use cases
8. **Testing Support** (Enhancement 6): Make it easy to test and validate templates before deployment
9. **Backward Compatibility**: Existing code continues to work without changes
10. **Source Integration**: Make it trivial to add FQDN template support to any source

### Non-Goals

1. **Redesigning text/template**: We wrap Go's standard template package, not replace it
2. **Breaking Changes**: All existing APIs remain functional during migration period
3. **Performance Micro-Optimization**: Focus is on architecture and features, not micro-optimizations
4. **Template Language Extensions**: We don't add new template syntax, only functions
5. **Dynamic Template Reloading**: Templates are registered at startup, not runtime
6. **Per-Provider Template Logic**: Provider-specific logic belongs in providers, not templates

## Proposal

### Current Limitations

#### Limitation 1: Template Parsed Per Source Constructor

**Code Evidence**:

```bash
$ grep -n "fqdn.ParseTemplate" source/*.go
# Returns 10+ matches across different source files
```

Every source that supports FQDN templates includes this pattern:

```go
tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
if err != nil {
    return nil, err
}
```

**Problems**:

- Wasteful: Same template string parsed multiple times
- Scattered: Error handling in every constructor
- Inflexible: Can't share parsed templates
- Untestable: Difficult to inject mock templates

#### Limitation 2: Two Different Constructor Patterns

**Pattern A**: Direct string parameter (older sources)

```go
// service.go, ingress.go, pod.go, etc.
func NewServiceSource(..., fqdnTemplate string, ...) (Source, error) {
    tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
    // ...
}
```

**Pattern B**: Config struct parameter (newer sources)

```go
// gateway.go
func NewGatewayRouteSource(..., config *Config) (Source, error) {
    tmpl, err := fqdn.ParseTemplate(config.FQDNTemplate)
    // ...
}
```

**Problem**: Solution must address both patterns.

#### Limitation 3: No Type Safety

Sources store raw `*template.Template`:

```go
type serviceSource struct {
    fqdnTemplate *template.Template  // Generic text/template type
}
```

**Problems**:

- No domain-specific methods (e.g., `Validate()`)
- Exposes text/template implementation details
- Difficult to extend with new functionality

#### Limitation 4: Single Template for All Sources

Currently, one template string applies to all sources:

```bash
--fqdn-template="{{.Name}}.{{.Namespace}}.example.com"
```

**Problems**:

- Services, Ingresses, and Gateways all use the same pattern
- Complex conditionals needed: `{{if eq .Kind "Service"}}...{{else}}...{{end}}`
- Templates become unreadable and error-prone
- Can't have different templates for different environments or namespaces

**Current Workaround** (ugly):

```go
{{if eq .Kind "Service"}}{{.Name}}.svc.example.com{{else if eq .Kind "Ingress"}}{{.Name}}.ingress.example.com{{else}}{{.Name}}.example.com{{end}}
```

#### Limitation 5: No DNS Validation

Templates can generate invalid DNS names that fail at runtime:

```go
// These all "work" but produce invalid DNS:
{{.Name}}                           // "My_Service" (underscores invalid)
{{.Namespace | toUpper}}            // "PROD" (should be lowercase)
{{.Name}}.{{.Name}}.{{.Name}}...    // 300+ characters (exceeds DNS limit)
```

**Impact**:

- Invalid DNS records created in production
- Silent failures or cryptic DNS errors
- Debugging requires manual DNS testing

#### Limitation 6: No Template Documentation

**Problems**:

- No way to know which sources support FQDN templates
- Examples scattered across tests and docs
- Manual documentation becomes outdated
- No automated validation of documentation accuracy

#### Limitation 7: Limited Template Functions

Current template functions from source/fqdn/fqdn.go:30-44:

- `contains`, `trimPrefix`, `trimSuffix`, `trim`, `toLower`
- `replace`, `isIPv6`, `isIPv4`

**Missing capabilities**:

- String truncation for long names
- Hashing for short identifiers
- Conditional/default value helpers
- DNS-specific utilities (reverse DNS, zone extraction)
- Encoding/decoding functions

#### Limitation 8: No Template Testing Support

**Problems**:

- Difficult to test templates before deployment
- No standard way to validate templates
- Manual testing required
- No test helpers or utilities

### User Stories

#### Story 1: Platform Engineer Debugging Template Errors

*As a platform engineer*, I deploy external-dns with a complex FQDN template. When the template has a syntax error, I receive 10+ error messages (one per source) all reporting the same parsing failure. This clutters logs and makes debugging harder.

**Desired Experience**: Template validation happens once at startup with a single clear error message.

#### Story 2: Developer Adding FQDN Support to New Source

*As a developer*, I'm adding a new source type and want to support FQDN templates. Currently, I must:

1. Add `fqdnTemplate string` parameter to constructor
2. Call `fqdn.ParseTemplate()`
3. Handle errors
4. Store `*template.Template` in source struct
5. Call `fqdn.ExecTemplate()` when generating endpoints

**Desired Experience**:

```go
// Simple - just get pre-parsed template from config
tmpl, err := fqdn.Get(config.FQDNTemplateName)
source.fqdnTemplate = tmpl
```

#### Story 3: SRE Implementing Template Validation

*As an SRE*, I want to add DNS validation to templates to prevent invalid hostnames in production. Currently, I would need to modify 10+ source files to add validation logic.

**Desired Experience**: Add validation once in the Template implementation, automatically applies to all sources.

## Enhancement 1: Template Interface & Registry (High Priority)

### Motivation

This enhancement addresses Limitations 1-3 by introducing a centralized template management system. Instead of each source parsing templates independently, templates are parsed once at startup and stored in a global registry.

### Design

#### Architecture Overview

```yml
┌─────────────────────────────────────────────────────────┐
│                   Startup / Config                       │
│                                                         │
│  1. Parse fqdn-template flag/config                       │
│  2. Register template: fqdn.Register("default", tmpl)   │
└─────────────────────────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────┐
│              Template Registry (Singleton)              │
│                                                         │
│  map["default"] → Template Interface                    │
│    └─ Wraps *text/template.Template                     │
│    └─ Provides Execute(), Validate(), etc.              │
└─────────────────────────────────────────────────────────┘
          │              │              │
          ▼              ▼              ▼
    ┌─────────┐    ┌─────────┐    ┌─────────┐
    │ Service │    │ Ingress │    │ Gateway │
    │ Source  │    │ Source  │    │ Source  │
    │         │    │         │    │         │
    │ tmpl =  │    │ tmpl =  │    │ tmpl =  │
    │ Get()   │    │ Get()   │    │ Get()   │
    └─────────┘    └─────────┘    └─────────┘
```

#### Package Structure

```bash
source/fqdn/
├── fqdn.go              # Backward compatible API (ParseTemplate, ExecTemplate)
├── template.go          # NEW: Template interface & implementation
├── registry.go          # NEW: Global template registry
├── options.go           # NEW: Template options (for future extensions)
│
├── fqdn_test.go        # Existing tests
├── template_test.go    # NEW: Template interface tests
└── registry_test.go    # NEW: Registry tests
```

## Package Structure

All enhancements will be implemented within the `source/fqdn/` package:

```bash
source/fqdn/
├── fqdn.go                      # Backward compatible API
├── template.go                  # Template interface & implementation (Enhancement 1)
├── registry.go                  # Template registry (Enhancement 1)
├── options.go                   # Template options (Enhancement 1)
├── template_set.go              # Multi-template support (Enhancement 2)
├── selector.go                  # Template selectors (Enhancement 2)
├── validator.go                 # DNS validation (Enhancement 3)
├── sanitizer.go                 # Hostname sanitization (Enhancement 3)
├── functions.go                 # Current template functions
├── functions_advanced.go        # Advanced functions (Enhancement 5)
├── doc_generator.go             # Documentation generation (Enhancement 4)
├── testing.go                   # Testing utilities (Enhancement 6)
│
├── fqdn_test.go                # Existing tests
├── template_test.go            # Template tests
├── template_set_test.go        # Multi-template tests
├── validator_test.go           # Validation tests
└── examples_test.go            # Example tests (for godoc)
```

### Migration Strategy

#### Phase 1: Infrastructure

**Step 1**: Update `source/store.go` to register templates at startup. The Config struct will maintain backward compatibility with the existing FQDNTemplate field while adding a new FQDNTemplateName field for registry-based templates. During initialization, templates are registered with the global registry.

**Deliverables**:

- Template interface implemented
- Registry working with tests
- Template registration in config
- All existing tests pass

#### Phase 2: Source Migration

**Step 2.1**: Migrate Pattern B sources (config-based)

These sources already accept `*Config`, making migration straightforward:

- source/gateway.go

Sources will retrieve templates from the registry using `fqdn.Get()` instead of parsing them directly with `fqdn.ParseTemplate()`.

**Step 2.2**: Update source struct field types

Source structs will change from using the generic `*template.Template` type to the domain-specific `fqdn.Template` interface.

**Step 2.3**: Migrate sources (string-based)

Sources with complex parameter lists will be simplified to accept a `*Config` parameter instead of individual string parameters, enabling cleaner retrieval of templates from the registry.

**Sources to migrate** (Pattern A):

- source/service.go
- source/ingress.go
- source/pod.go
- source/node.go
- source/istio_gateway.go
- source/istio_virtualservice.go
- source/contour_httpproxy.go
- source/openshift_route.go
- source/skipper_routegroup.go

**Step 2.4**: Update template execution calls

```go
// BEFORE
hostnames, err := fqdn.ExecTemplate(sc.fqdnTemplate, service)

// AFTER
hostnames, err := sc.fqdnTemplate.Execute(ctx, service)
```

**Deliverables**:

- All sources use Template interface
- No source calls `fqdn.ParseTemplate()` directly
- All tests updated and passing

#### Phase 3: Verification

An integration test will verify that templates are parsed only once at startup, even when multiple sources are created. This ensures the registry is working correctly and templates are being shared across sources.

### Behavior

#### Normal Operation

1. **Startup**: external-dns parses config, registers FQDN template as "default"
2. **Source Creation**: Each source retrieves "default" template from registry
3. **Endpoint Generation**: Sources call `template.Execute()` to generate hostnames
4. **Error Handling**: Template errors reported once at startup, not per-source

#### Thread Safety

- Registry uses `sync.RWMutex` for concurrent access
- `Register()` and `Update()` use write lock
- `Get()` uses read lock (allows concurrent reads)
- Template execution is thread-safe (text/template is immutable after parsing)

#### Memory Impact

**Before**: 10 source constructors × template size ≈ 10× memory
**After**: 1 template in registry ≈ 1× memory

For typical template (~1KB), savings are negligible. Real benefit is architecture and extensibility.

### Drawbacks

1. **Global State**: Registry is a singleton (global variable)
   - **Mitigation**: Provide `NewTemplateRegistry()` for testing with custom registries

2. **Breaking Changes**: Source constructor signatures change
   - **Mitigation**: Keep deprecated wrappers for one major version
3. **Unused Context**: `Execute(ctx)` may be unnecessary now
   - **Mitigation**: Document that it's for future extensibility, can ignore for now

## Alternatives

### Alternative 1: Keep Current Approach (Do Nothing)

**Pros**:

- No changes required
- No migration effort
- No risk of regressions

**Cons**:

- Continues wasteful parsing
- Difficult to add new features
- Scattered error handling
- Poor developer experience

**Verdict**: Not recommended - problem is significant enough to warrant fixing

### Alternative 2: Lazy Initialization

Parse template on first use instead of at startup:

```go
type serviceSource struct {
    fqdnTemplateStr string
    fqdnTemplate    *template.Template
    parseOnce       sync.Once
}

func (s *serviceSource) getTemplate() (*template.Template, error) {
    var err error
    s.parseOnce.Do(func() {
        s.fqdnTemplate, err = fqdn.ParseTemplate(s.fqdnTemplateStr)
    })
    return s.fqdnTemplate, err
}
```

**Pros**:

- Templates only parsed if actually used
- No global registry

**Cons**:

- Still parses once per source (10× duplication)
- Errors deferred to runtime instead of startup
- Doesn't solve the architecture problem

**Verdict**: Doesn't address core issue

### Alternative 3: Config-Level Caching

Cache parsed template in `source.Config`:

```go
type Config struct {
    FQDNTemplate       string
    parsedFQDNTemplate *template.Template  // Cached
}

func NewSourceConfig(cfg *externaldns.Config) (*Config, error) {
    tmpl, err := fqdn.ParseTemplate(cfg.FQDNTemplate)
    return &Config{
        FQDNTemplate:       cfg.FQDNTemplate,
        parsedFQDNTemplate: tmpl,
    }, nil
}
```

**Pros**:

- Simple implementation
- Parse once
- No global state

**Cons**:

- Exposes `*template.Template` directly (no abstraction)
- Can't add features like validation
- Config becomes stateful (harder to serialize/debug)

**Verdict**: Decent middle ground, but less extensible than registry

### Alternative 5: Combine Registry + Dependency Injection

Registry for registration, explicit passing to sources:

```go
// main.go
fqdn.Register("default", cfg.FQDNTemplate)

// sources.go
func NewServiceSource(..., templateName string) (Source, error) {
    tmpl, _ := fqdn.Get(templateName)
}
```

**Pros**:

- Explicit which template each source uses
- Supports multi-template easily

**Cons**:

- More verbose
- Easy to pass wrong template name

**Verdict**: This is essentially the proposed approach (templateName stored in Config)

### API

#### Core Interfaces

The Template interface provides a domain-specific abstraction for generating DNS hostnames from Kubernetes objects.
It includes an Execute method that takes a context and Kubernetes object, returning a slice of hostname strings, and a String method for debugging. Options are provided through a functional options pattern.

**Rationale for Context**:

- Future-proofing: Enables timeout/cancellation for complex templates
- Consistency: Matches Go idioms for functions that may block
- Extensibility: Required for future features like remote template fetching

**Note**: If context proves unnecessary, it can be ignored in implementations.

#### Template Creation

```go
// New creates a new Template from a template string.
// Returns error if template parsing fails.
func New(templateStr string, opts ...Option) (Template, error)
```

#### Registry API

```go
// registry.go

// Register adds a template to the global registry.
// Returns error if:
// - name is empty
// - templateStr is invalid
// - name already exists (call Unregister first, or use Update)
func Register(name string, templateStr string, opts ...Option) error

// MustRegister is like Register but panics on error.
// Useful for startup registration where failure should be fatal.
func MustRegister(name string, templateStr string, opts ...Option)

// Get retrieves a registered template by name.
// Returns error if template not found.
func Get(name string) (Template, error)

// List returns all registered template names.
func List() []string

// Clear removes all templates (useful for testing).
func Clear()
```

See Implementation Examples section at the end of this document for detailed code.

## Enhancement 2: Multi-Template Support (High Priority)

### Motivation

This enhancement addresses Limitation 4 by enabling different templates for different sources, namespaces, labels, or other conditions. Additionally, it supports **multiple templates per resource** to generate multiple DNS records (e.g., primary + alternate names, multi-zone support, migration scenarios).

**Current limitation**:

1. Solves Real Pain Point: Addresses the ugly workaround of complex conditionals in templates

```bash
// Current ugly approach
{{if eq .Kind "Service"}}...{{else if eq .Kind "Ingress"}}...{{end}}
```

2. Well-Designed Interface: The TemplateSelector interface is clean and extensible
3. Multiple Selection Strategies: Covers common use cases (kind, label, annotation, namespace)
4. Backward Compatible: Existing single template approach continues to work
5. Builds on `Enhancement 1`: Correctly depends on Template interface and registry

**Key capabilities**:

- Different templates for different sources/conditions
- Multiple DNS names per resource
- Cleaner configuration without complex template conditionals
- Backward compatible with single template approach

### Design

The TemplateSet manages multiple templates for different sources or conditions. It implements the Template interface and uses a TemplateSelector to choose which template group to apply based on object properties.
The Execute method handles template selection, execution of multiple templates, and automatic deduplication of generated hostnames.

See Implementation Examples section for detailed code.

### Configuration

**YAML Configuration** (with array support):

```yaml
# external-dns-config.yaml
fqdnTemplates:
  # Strategy for template selection (applied globally)
  strategy:
  - default  # Options: kind, label, namespace, annotation, all, source

  # Default templates - supports multiple templates per resource (loaded when specified)
  default:
    - "{{.Name}}.{{.Namespace}}.example.com"
    - "{{.Name}}.example.org"  # Generate additional DNS record

  # Multi-template set by kind
  multi: # Fallback for unknown kinds always default
    strategy: kind
    templates:
      service:
        - "{{.Name}}.{{.Namespace}}.svc.example.com"
        - "{{.Name}}.svc.example.com"  # Short form
      ingress:
        - "{{.Name}}.{{.Namespace}}.ingress.example.com"
      gateway:
        - "{{.Name}}.gateway.example.com"
      httpoute:
        - "{{.Name}}.route.example.com"

    bySource:
      multi: # Fallback for source always default
      strategy: source
      templates:
        service:
          - "{{.Name}}.{{.Status.LoadBalancer.Ingress.IP}}.svc.example.com"
        node:
          - "{{.Name}}.{{.Namespace}}.node.tld"
        traefik-proxy:
          - "{{.Name}}.treafik.io"

  # Label-based selection
  byEnvironment: # Fallback for unknown kinds always default
    strategy: label # labels, annotations, environment variable
    selector: "environment"  # Label key to check environment=staging
    templates:
      production:
        - "{{.Name}}.prod.example.com"
        - "{{.Name}}.production.example.com"  # Alternate name
      staging:
        - "{{.Name}}.staging.example.com"
      development:
        - "{{.Name}}.dev.example.com"

  # Namespace-based selection
  byNamespace: # Fallback for unknown kinds always default
    strategy: namespace
    templates:
      kube-system:
        - "{{.Name}}.system.example.internal"
      default:  # The namespace named "default"
        - "{{.Name}}.{{.Namespace}}.apps.example.com"
      production:
        - "{{.Name}}.prod.example.com"
```

**CLI Flags**:

```bash
# Simple (existing)
--fqdn-template="{{.Name}}.{{.Namespace}}.example.com"

# Multiple simple templates (backward compatible + new)
--fqdn-template="{{.Name}}.{{.Namespace}}.example.com" \
--fqdn-template="{{.Name}}.example.org" \
--fqdn-template="{{.Name}}.internal"
# Each resource generates 3 DNS records

# Load from config file (recommended for complex setups)
--fqdn-template-config=/etc/external-dns/templates.yaml

# Mixed approach (both are loaded)
--fqdn-template="{{.Name}}.fallback.com" \
--fqdn-template-config=/etc/external-dns/templates.yaml
```

**Flag priority and merging**:

1. Templates from `--fqdn-template-config` are loaded first
2. Templates from `--fqdn-template` flags are added
3. All registered templates are executed (additive behavior)
4. Host Duplicates trigger warnings and are deduplicated

### Implementation Example

#### **Usage Example**

```go
// Register multi-template set
err := fqdn.RegisterSet("multi", fqdn.TemplateSetConfig{
    Strategy: fqdn.StrategyByKind,
    Templates: map[string]string{
        "service":   "{{.Name}}.{{.Namespace}}.svc.example.com",
        "ingress":   "{{.Name}}.ingress.example.com",
        "gateway":   "{{.Name}}.gateway.example.com",
    },
    Default: "{{.Name}}.example.com",
})

// Get template set
tmplSet, _ := fqdn.GetSet("multi")
// Execute - automatically selects correct template
hostnames, _ := tmplSet.Execute(context.Background(), &v1.Service{...})
// Uses Service template: "my-service.default.svc.example.com"

hostnames, _ = tmplSet.Execute(context.Background(), networkingv1.Ingress{...})
// Uses Ingress template: "my-ingress.ingress.example.com"
```

See Implementation Examples section for config loading and duplicate detection code.

### Benefits

- Different templates per source type, namespace, or label
- Multiple DNS records per resource (primary + alternates, multi-zone)
- Conditional logic extracted from templates
- More readable configuration
- Type-safe selection
- Backward compatible (single string → array of one)
- Duplicate detection with warnings
- Migration support (old + new DNS names during transitions)

### User Stories

#### Story 4: Multi-Zone DNS Management

*As a platform engineer*, I need to create DNS records in both example.com and example.org zones for the same services.

**Solution**:

```yaml
fqdnTemplates:
  default:
    - "{{.Name}}.{{.Namespace}}.example.com"
    - "{{.Name}}.{{.Namespace}}.example.org"
```

Result: Each service gets 2 DNS records automatically.

#### Story 5: Environment-Based Templates with Alternates

*As a DevOps engineer*, production services need both "prod" and "production" DNS names for compatibility.

**Solution**:

```yaml
fqdnTemplates:
  byEnvironment:
    strategy: label
    selector: "environment"
    templates:
      production:
        - "{{.Name}}.prod.example.com"
        - "{{.Name}}.production.example.com"
```

#### Story 6: DNS Migration

*As an SRE*, I'm migrating from old naming scheme to new, need both during transition.

**Solution**:

```yaml
fqdnTemplates:
  default:
    - "{{.Name}}.{{.Namespace}}.example.com"  # New format
    - "{{.Name}}-{{.Namespace}}.example.com"  # Old format (migration)
```

Both DNS records exist during migration period.

## Enhancement 3: DNS Validation (Medium Priority)

### Motivation

This enhancement addresses Limitation 5 by providing automatic DNS validation and sanitization. Templates can currently generate invalid DNS names that fail at runtime. This enhancement validates hostnames **after generation** and optionally sanitizes them to ensure RFC 1123 compliance.

**Current limitation**:

1. Addresses Real Problem: Prevents invalid DNS names from being created in production

- Underscores: My_Service → my-service
- Uppercase: PROD → prod
- Over-length: 300 chars → truncated to 253

2. Clean Separation: Validator (checks) vs Sanitizer (fixes)
3. RFC 1123 Compliance: Uses correct DNS label regex
4. Options Pattern: Flexible configuration with WithValidation(), WithAutoFix() options
5. Builds on Enhancement 1: Integrates with Template interface via options

**Key capabilities**:

- Validate generated hostnames against DNS standards
- Automatic sanitization (lowercase, replace invalid chars)
- Support for Internationalized Domain Names (IDN/Punycode RFC 3492)
- Configurable error handling (skip, fail, or ignore)
- Integration with prefix/suffix flags

### Design

Validation happens after template execution, not during template creation. The system provides Validator and Sanitizer interfaces with configurable options for handling invalid hostnames. Validation can either skip, fail, or auto-fix invalid DNS names.

See Implementation Examples section for detailed code.

### Implementation

**RFC 1123 Validator** (with detailed error messages):

```go
// fqdn/validator.go
type rfc1123Validator struct{}

func (v *rfc1123Validator) Validate(hostname string) error {
    if len(hostname) > 253 {
        return fmt.Errorf("hostname too long: %d chars (max 253)", len(hostname))
    }

    if hostname == "" {
        return fmt.Errorf("hostname cannot be empty")
    }

    labels := strings.Split(hostname, ".")
    for i, label := range labels {
        if err := v.validateLabel(label, i); err != nil {
            return err
        }
    }
    return nil
}

func (v *rfc1123Validator) validateLabel(label string, index int) error {
    if len(label) == 0 {
        return fmt.Errorf("label %d is empty", index)
    }
    if len(label) > 63 {
        return fmt.Errorf("label %d too long: %d chars (max 63): %q", index, len(label), label)
    }

    // Check first character
    if !isAlphanumeric(rune(label[0])) {
        return fmt.Errorf("label %d must start with alphanumeric: %q", index, label)
    }

    // Check last character
    if !isAlphanumeric(rune(label[len(label)-1])) {
        return fmt.Errorf("label %d must end with alphanumeric: %q", index, label)
    }

    // Check middle characters
    for j, r := range label {
        if !isAlphanumeric(r) && r != '-' {
            return fmt.Errorf("label %d contains invalid character at position %d: %q", index, j, string(r))
        }
    }

    return nil
}

func isAlphanumeric(r rune) bool {
    return (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}
```

**Sanitizer with validation** (sanitize then validate):

```go
// fqdn/sanitizer.go
import "golang.org/x/net/idna"

type defaultSanitizer struct {
    toLowercase     bool
    replaceInvalid  bool
    replacementChar rune
    maxLength       int
    allowUnicode    bool
}

func (s *defaultSanitizer) Sanitize(hostname string, opts ValidationOptions) (string, error) {
    original := hostname

    // Step 1: Handle Unicode (Punycode encoding per RFC 3492)
    if opts.AllowUnicode && containsUnicode(hostname) {
        var err error
        hostname, err = idna.ToASCII(hostname)
        if err != nil {
            return "", fmt.Errorf("failed to convert Unicode hostname to Punycode: %w", err)
        }
    }

    // Step 2: Lowercase
    if s.toLowercase {
        hostname = strings.ToLower(hostname)
    }

    // Step 3: Replace invalid characters
    if s.replaceInvalid {
        hostname = s.replaceInvalidChars(hostname)
    }

    // Step 4: Validate result
    if err := RFC1123Validator.Validate(hostname); err != nil {
        return "", fmt.Errorf("sanitization of %q failed to produce valid hostname %q: %w",
            original, hostname, err)
    }

    return hostname, nil
}

func (s *defaultSanitizer) replaceInvalidChars(hostname string) string {
    var result strings.Builder
    prevWasDash := false

    for i, r := range hostname {
        if isValidDNSChar(r) {
            result.WriteRune(r)
            prevWasDash = (r == '-')
        } else {
            // Don't create consecutive dashes or leading/trailing dashes
            if !prevWasDash && i > 0 && i < len(hostname)-1 {
                result.WriteRune(s.replacementChar)
                prevWasDash = true
            }
        }
    }

    return result.String()
}

func isValidDNSChar(r rune) bool {
    return (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '.' || r == '-'
}

func containsUnicode(s string) bool {
    for _, r := range s {
        if r > 127 {
            return true
        }
    }
    return false
}
```

### Configuration

**YAML Configuration** (loaded via `--fqdn-template-config`):

```yaml
# external-dns-config.yaml
fqdnTemplates:
  # ... template configuration ...

# Validation configuration
fqdnValidation:
  enabled: true
  autoFix: true            # Automatically sanitize invalid hostnames
  maxLength: 253           # DNS standard maximum
  onTooLong: error         # Options: error | ignore
  allowUnicode: true       # Support IDN via Punycode (RFC 3492)
  strictRFC1123: true      # Enforce strict RFC 1123 compliance
  onError: skip            # Options: skip | fail
```

**No additional CLI flags** - all configuration via YAML file.

### Integration Points

**Implementation decision** (to be finalized during implementation):

**Option A**: Validation in source layer after template execution

```go
func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
    // Execute template
    hostnames, err := sc.fqdnTemplate.Execute(ctx, svc)
    if err != nil {
        return nil, err
    }

    // Validate and sanitize
    if sc.config.ValidationEnabled {
        hostnames, err = fqdn.ValidateAndSanitize(hostnames, sc.config.ValidationOpts)
        if err != nil {
            return nil, err
        }
    }

    return createEndpoints(hostnames), nil
}
```

**Option B**: Validation at registry level

```go
type TemplateRegistry struct {
    templates         map[string]Template
    validationEnabled bool
    validationOpts    ValidationOptions
    mu                sync.RWMutex
}

func (r *TemplateRegistry) ExecuteAndValidate(name string, ctx context.Context, obj kubeObject) ([]string, error) {
    tmpl, _ := r.Get(name)
    hostnames, err := tmpl.Execute(ctx, obj)
    if err != nil {
        return nil, err
    }

    if r.validationEnabled {
        return ValidateAndSanitize(hostnames, r.validationOpts)
    }
    return hostnames, nil
}
```

**Note**: Final decision between Options A and B will be made during implementation based on code organization and testing requirements.

### Usage Examples

**Example 1: Auto-fix invalid characters**

```go
// Input service name: "My_Service"
// Template: "{{.Name}}.example.com"

// Without validation:
// Output: "My_Service.example.com" (invalid: uppercase, underscore)

// With validation + auto-fix:
// Output: "my-service.example.com" (fixed: lowercase, replaced underscore)
```

**Example 2: Unicode/IDN support**

```go
// Input service name: "服务" (Chinese for "service")
// Template: "{{.Name}}.example.com"
// Validation: allowUnicode=true

// Without validation:
// Output: "服务.example.com" (invalid: non-ASCII)

// With Punycode encoding:
// Output: "xn--vuq861b.example.com" (Punycode per RFC 3492)
```

**Example 3: Too-long hostname handling**

```go
// Input: 300-character service name
// Template: "{{.Name}}.{{.Namespace}}.example.com"
// Total length: 330 characters

// With onTooLong=error:
// Error: "hostname too long: 330 chars (max 253)"

// With onTooLong=ignore:
// Warning logged, hostname used as-is (may fail at DNS provider)
```

### Benefits

- ✅ Prevents invalid DNS records in production
- ✅ Automatic sanitization of common issues
- ✅ Support for Unicode/IDN via Punycode (RFC 3492)
- ✅ Configurable error handling
- ✅ Integration with prefix/suffix flags
- ✅ Detailed error messages for debugging
- ✅ Validates after generation (cleaner separation)
- ✅ Configuration via YAML only (no flag proliferation)

### User Stories

#### Story 7: Production DNS Validation

*As an SRE*, I've had incidents where Kubernetes services with names like "My_Service" or "TEST-SERVICE" created invalid DNS records that failed silently. I want external-dns to automatically fix these issues or alert me.

**Solution**:

```yaml
fqdnValidation:
  enabled: true
  autoFix: true
  onError: skip
```

Invalid names are automatically fixed:

- `My_Service` → `my-service`
- `TEST-SERVICE` → `test-service`
- `under_score` → `under-score`

#### Story 8: International Service Names

*As a platform engineer in Japan*, our services have Japanese names. I need DNS records that work with international characters.

**Solution**:

```yaml
fqdnValidation:
  enabled: true
  allowUnicode: true
```

Unicode is converted to Punycode:

- `サービス.example.com` → `xn--pck0a3d7a.example.com`

## Enhancement 4: Auto-Generated Documentation (Medium Priority)

### Motivation

This enhancement addresses Limitation 6 by automatically generating documentation from code comments, tests, and source analysis. This ensures documentation stays up-to-date and accurate, enforced by tests.

**Current limitation**:

1. **No automated doc generation**: Examples scattered across tests and docs
2. **Manual documentation**: Becomes outdated quickly
3. **No enforcement**: Nothing prevents docs from becoming stale
4. **No source matrix**: Hard to know which sources support FQDN templates
5. **Function reference missing**: Users don't know what template functions exist
6. **Version tracking**: No way to know when features were added

**Key capabilities**:

- Generate docs from code comments and structured tests
- Enforce freshness via go tests
- Source support matrix via interface compliance
- Function reference from godoc comments
- Version tracking with `@since` tags
- Single source of truth (code is the documentation)

### Design

**Not FQDN package responsibility** - Documentation generation is a separate build tool.

#### Package Structure

```bash
internal/gen/docs/fqdn/
├── main.go              # Doc generator CLI
├── templates.go         # Extract from tests and @example tags
├── sources.go           # Scan source interfaces
├── functions.go         # Parse godoc comments
└── testdata/            # Test fixtures

docs/advanced/fqdn/      # Generated output (DO NOT EDIT)
├── templates.md         # Template examples
├── sources.md           # Source support matrix
└── functions.md         # Function reference
```

Similar to existing `internal/gen/docs/metrics` package.

#### Source Interfaces

Sources declare support via interface compliance:

```go
// source/source.go

// FQDNTemplateSupport indicates a source supports FQDN templates.
type FQDNTemplateSupport interface {
    SupportsFQDNTemplate() bool
    SupportsHostnameAnnotation() bool
}

// Example implementation
// source/service.go
func (s *serviceSource) SupportsFQDNTemplate() bool {
    return true
}

func (s *serviceSource) SupportsHostnameAnnotation() bool {
    return !s.ignoreHostnameAnnotation
}
```

#### Godoc-Style Comments

```go
// source/fqdn/functions.go

// truncate limits a string to maxLen characters.
//
// Example:
//
//	{{truncate 10 .Name}}
//	Input: "very-long-service-name"
//	Output: "very-long-"
//
// @since v0.20.0
func truncate(maxLen int, s string) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen]
}
```

#### Structured Test Tags

```go
// source/fqdn/fqdn_test.go

// @example Simple Template
// @description Basic hostname generation
// @since v0.20.0
func TestSimpleTemplate(t *testing.T) {
    tmpl := New("{{.Name}}.example.com")

    // Template: {{.Name}}.example.com
    // Input: Service "test" in namespace "default"
    // Output: ["test.example.com"]

    result, _ := tmpl.Execute(ctx, testService)
    assert.Equal(t, []string{"test.example.com"}, result)
}

// @example Multi-Zone Templates
// @description Generate DNS records in multiple zones
// @since v0.21.0
func TestMultiZoneTemplate(t *testing.T) {
    config := TemplateSetConfig{
        Default: []string{
            "{{.Name}}.example.com",
            "{{.Name}}.example.org",
        },
    }

    // Template: Multiple templates in array
    // Input: Service "api" in namespace "prod"
    // Output: ["api.example.com", "api.example.org"]

    // ...
}
```

### Implementation

#### Generator Tool

```go
// internal/gen/docs/fqdn/main.go

package main

import (
    "flag"
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "io/ioutil"
    "path/filepath"
)

func main() {
    outputDir := flag.String("output", "docs/advanced/fqdn", "Output directory")
    flag.Parse()

    // Generate all documentation
    if err := generateDocs(*outputDir); err != nil {
        panic(err)
    }
}

func generateDocs(outputDir string) error {
    // 1. Generate template examples from tests
    examples, err := extractExamplesFromTests("source/fqdn/fqdn_test.go")
    if err != nil {
        return err
    }
    if err := writeMarkdown(filepath.Join(outputDir, "templates.md"), examples); err != nil {
        return err
    }

    // 2. Generate source support matrix
    sources, err := scanSourceInterfaces("source/")
    if err != nil {
        return err
    }
    if err := writeMarkdown(filepath.Join(outputDir, "sources.md"), sources); err != nil {
        return err
    }

    // 3. Generate function reference from godoc
    functions, err := extractFunctionDocs("source/fqdn/functions.go")
    if err != nil {
        return err
    }
    if err := writeMarkdown(filepath.Join(outputDir, "functions.md"), functions); err != nil {
        return err
    }

    return nil
}
```

#### Extract Examples from Tests

```go
// internal/gen/docs/fqdn/templates.go

type Example struct {
    Name        string
    Description string
    Since       string
    Template    string
    Input       string
    Output      string
}

func extractExamplesFromTests(testFile string) (string, error) {
    fset := token.NewFileSet()
    node, err := parser.ParseFile(fset, testFile, nil, parser.ParseComments)
    if err != nil {
        return "", err
    }

    var examples []Example

    // Look for @example tags in comments
    for _, commentGroup := range node.Comments {
        text := commentGroup.Text()
        if strings.Contains(text, "@example") {
            example := parseExampleComment(text)
            examples = append(examples, example)
        }
    }

    // Also parse inline comments in test functions
    ast.Inspect(node, func(n ast.Node) bool {
        if fn, ok := n.(*ast.FuncDecl); ok {
            if strings.HasPrefix(fn.Name.Name, "Test") {
                example := parseTestFunction(fn)
                if example != nil {
                    examples = append(examples, *example)
                }
            }
        }
        return true
    })

    return formatExamplesMarkdown(examples), nil
}

func parseExampleComment(text string) Example {
    // Parse:
    // @example Simple Template
    // @description Basic hostname generation
    // @since v0.14.0

    lines := strings.Split(text, "\n")
    example := Example{}

    for _, line := range lines {
        if strings.Contains(line, "@example") {
            example.Name = strings.TrimSpace(strings.TrimPrefix(line, "@example"))
        }
        if strings.Contains(line, "@description") {
            example.Description = strings.TrimSpace(strings.TrimPrefix(line, "@description"))
        }
        if strings.Contains(line, "@since") {
            example.Since = strings.TrimSpace(strings.TrimPrefix(line, "@since"))
        }
    }

    return example
}
```

#### Scan Source Interfaces

```go
// internal/gen/docs/fqdn/sources.go

type SourceInfo struct {
    Name                       string
    File                       string
    SupportsFQDN               bool
    SupportsHostnameAnnotation bool
    Since                      string
}

func scanSourceInterfaces(sourceDir string) (string, error) {
    var sources []SourceInfo

    files, _ := filepath.Glob(filepath.Join(sourceDir, "*.go"))

    for _, file := range files {
        fset := token.NewFileSet()
        node, _ := parser.ParseFile(fset, file, nil, parser.ParseComments)

        info := SourceInfo{
            File: filepath.Base(file),
            Name: extractSourceName(file),
        }

        // Check for interface implementations
        ast.Inspect(node, func(n ast.Node) bool {
            if fn, ok := n.(*ast.FuncDecl); ok {
                if fn.Name.Name == "SupportsFQDNTemplate" {
                    info.SupportsFQDN = true
                    info.Since = extractSinceTag(fn.Doc)
                }
                if fn.Name.Name == "SupportsHostnameAnnotation" {
                    info.SupportsHostnameAnnotation = true
                }
            }
            return true
        })

        sources = append(sources, info)
    }

    return formatSourceTable(sources), nil
}

func extractSinceTag(doc *ast.CommentGroup) string {
    if doc == nil {
        return ""
    }

    for _, comment := range doc.List {
        if strings.Contains(comment.Text, "@since") {
            parts := strings.Fields(comment.Text)
            for i, part := range parts {
                if part == "@since" && i+1 < len(parts) {
                    return parts[i+1]
                }
            }
        }
    }
    return ""
}
```

#### Function Reference from Godoc

```go
// internal/gen/docs/fqdn/functions.go

type FunctionDoc struct {
    Name        string
    Signature   string
    Description string
    Example     string
    Since       string
}

func extractFunctionDocs(file string) (string, error) {
    fset := token.NewFileSet()
    node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
    if err != nil {
        return "", err
    }

    var functions []FunctionDoc

    ast.Inspect(node, func(n ast.Node) bool {
        if fn, ok := n.(*ast.FuncDecl); ok {
            // Extract from godoc comments
            if fn.Doc != nil {
                doc := FunctionDoc{
                    Name:        fn.Name.Name,
                    Signature:   formatSignature(fn.Type),
                    Description: extractDescription(fn.Doc),
                    Example:     extractExample(fn.Doc),
                    Since:       extractSinceTag(fn.Doc),
                }
                functions = append(functions, doc)
            }
        }
        return true
    })

    return formatFunctionReference(functions), nil
}
```

### Test Enforcement

```go
// source/fqdn/doc_test.go

func TestDocsUpToDate(t *testing.T) {
    // Run doc generator
    cmd := exec.Command("go", "run", "internal/gen/docs/fqdn/main.go",
        "-output", "docs/advanced/fqdn")
    cmd.Dir = repoRoot()
    output, err := cmd.CombinedOutput()
    if err != nil {
        t.Fatalf("Failed to generate docs: %v\n%s", err, output)
    }

    // Check if any files changed
    cmd = exec.Command("git", "diff", "--exit-code", "docs/advanced/fqdn/")
    cmd.Dir = repoRoot()
    if err := cmd.Run(); err != nil {
        t.Fatalf(`
Documentation is out of date!

Please run from repo root:
    go run internal/gen/docs/fqdn/main.go

Or run:
    make generate-docs

Then commit the updated files in docs/advanced/fqdn/
`)
    }
}
```

### Usage

**Generate docs**:

```bash
# From repo root
$ go run internal/gen/docs/fqdn/main.go

# Or via Makefile
$ make generate-docs
```

**Verify docs are current** (in tests):

```bash
go test ./source/fqdn/... -run TestDocsUpToDate
```

**Generated files** (DO NOT EDIT):

```bash
docs/advanced/fqdn/
├── templates.md         # Examples from tests
├── sources.md           # Source support matrix
└── functions.md         # Function reference
```

### Example Output

**docs/advanced/fqdn/templates.md**:

```markdown
# FQDN Template Examples

_This file is auto-generated. DO NOT EDIT._
_Last updated: 2025-12-31_
_Generated by: internal/gen/docs/fqdn_

## Simple Template

**Since**: v0.20.0

Basic hostname generation.

**Template**:
```

{{.Name}}.example.com

```bash

**Input**: Service "test" in namespace "default"

**Output**: `["test.example.com"]`

---

## Multi-Zone Templates

**Since**: v0.21.0

Generate DNS records in multiple zones.

**Template**:
```yaml
default:
  - "{{.Name}}.example.com"
  - "{{.Name}}.example.org"
```

**Input**: Service "api" in namespace "prod"

**Output**: `["api.example.com", "api.example.org"]`

```bash

**docs/advanced/fqdn/sources.md**:
```markdown
# FQDN Template Support by Source

_This file is auto-generated. DO NOT EDIT._

| Source | FQDN Template | Hostname Annotation | Since |
|--------|---------------|---------------------|-------|
| Service | ✅ | ✅ | v0.16.0 |
| Ingress | ✅ | ✅ | v0.17.0 |
| Gateway | ✅ | ✅ | v0.21.0 |
| Pod | ✅ | ❌ | v0.15.0 |
| Node | ❌ | ❌ | - |
```

**docs/advanced/fqdn/functions.md**:

```markdown
# Template Function Reference

_This file is auto-generated. DO NOT EDIT._

## truncate

**Since**: v0.16.0

**Signature**: `truncate(maxLen int, s string) string`

Limits a string to maxLen characters.

**Example**:
```

{{truncate 10 .Name}}
Input: "very-long-service-name"
Output: "very-long-"

```bash
```

### Benefits

- ✅ Documentation always up to date
- ✅ Tests enforce freshness
- ✅ Single source of truth (code)
- ✅ Godoc-style comments (Go standard)
- ✅ Interface compliance for sources
- ✅ Version tracking with @since tags
- ✅ Markdown output in docs/advanced/fqdn/
- ✅ Similar to existing metrics doc generation

## Enhancement 5: Advanced Template Functions (Low Priority)

### Motivation

#### Current Limitation

The existing template function library is minimal and insufficient for complex hostname generation patterns:

- Only 8 basic functions (contains, trimPrefix, trimSuffix, trim, toLower, replace, isIPv4, isIPv6)
- No conditional logic (default, ternary, coalesce)
- No string manipulation beyond trim operations (truncate, split, join)
- No DNS-specific utilities (reverseDNS, )
- Functions not documented with godoc-style comments
- No benchmarking or performance visibility
- Function naming doesn't follow Helm/Sprig conventions

This enhancement provides an extended library of vetted template functions for real-world use cases while maintaining security and following established conventions.

### Design

#### Function Organization

Functions are organized by category in separate files:

```bash
source/fqdn/
├── functions.go           # Core/existing functions, FuncMap registration
└── functions_bench_test.go # Benchmarks for all functions
```

#### Core Functions Module

Function registration is moved to package level with `customFuncs()` that returns a FuncMap with all categories registered (string, conditional, DNS, encoding).

#### String Functions

Provides `truncate`, `split`, and `join` functions for string manipulation with godoc-style comments and examples.

#### Conditional Functions

Provides `default`, `ternary`, and `coalesce` functions for conditional logic in templates.

#### DNS Functions

Provides DNS-specific utilities like `reverseDNS` for DNS-related template operations.

See Implementation Examples section for detailed function code.

### Configuration

No additional configuration needed. Functions are automatically available in all templates.

### Deprecation Plan

**Sprig Convention Adoption**:

- v0.21.0: Add `lower` function, keep `toLower` working
- v0.22.0: Add deprecation warning when `toLower` is used
- v0.24.0: Remove `toLower` function

### Benchmarking

All functions include benchmarks to track performance:

```go
// source/fqdn/functions_bench_test.go

func BenchmarkTruncate(b *testing.B) {
    input := "very-long-service-name-that-needs-truncation"
    for i := 0; i < b.NB; i++ {
        _ = truncate(20, input)
    }
}

func BenchmarkSHA256(b *testing.B) {
    input := "my-service-name"
    for i := 0; i < b.NB; i++ {
        _ = sha256Hash(input)
    }
}

// ... benchmarks for all functions
```

Benchmark results documented in godoc for each function.

### Usage Examples

**Truncate long names**:

```yaml
fqdnTemplates:
  default:
    - "{{truncate 20 .Name}}.example.com"

# Input: Service "very-long-service-name"
# Output: "very-long-service-n.example.com"
```

**Use default for missing labels**:

```yaml
fqdnTemplates:
  default:
    - "{{default .Labels.env "prod"}}.example.com"

# Input: Service with no "env" label
# Output: "prod.example.com"
```

**Conditional logic**:

```yaml
fqdnTemplates:
  default:
    - "{{ternary (eq .Labels.env "prod") "production" "staging"}}.example.com"

# Input: Service with env="prod"
# Output: "production.example.com"
```

**Reverse DNS**:

```yaml
fqdnTemplates:
  default:
    - "{{reverseDNS .Status.PodIP}}.in-addr.arpa"

# Input: Pod with IP "192.0.2.1"
# Output: "1.2.0.192.in-addr.arpa"
```

### Future Considerations

**Sprig Integration** (optional, as function library grows):

- Consider adopting full Sprig library for comprehensive template functions
- Evaluate security implications of regex functions (currently excluded)
- Maintain backward compatibility during migration

## Enhancement 6: FQDN Template Execution in Informer SetTransform (Medium Priority)

### Motivation

#### Current Limitation

FQDN template execution happens on every `Endpoints()` call, which is invoked repeatedly during reconciliation loops:

- Template parsing and execution occurs multiple times for the same object
- CPU intensive for large clusters with many resources
- No caching mechanism for generated FQDNs
- Inefficient when templates produce static results for unchanging objects
- Performance degrades with complex templates (multiple functions, conditionals)
- Sources keep full object in cache when using templates (memory overhead)

Example from source/pod.go:81-85:

```go
if fqdnTemplate == "" {
    // SetTransform reduces memory when no template
} else {
    // No transform - keep full object for template execution
}
```

This enhancement moves template execution to **Informer SetTransform**, executing templates **once when objects are added/modified** in the cache, significantly reducing CPU usage.

### Design

#### Core Concept

Execute FQDN template in `SetTransform()` and cache results in object annotations. This shifts template execution from **reconciliation-time** to **cache-time**.

**Current flow**:

```bash
Object → Informer → Cache (full object) → Endpoints() → Execute Template → FQDNs
                                              ↑
                                     Called repeatedly
```

**Proposed flow**:

```bash
Object → Informer → SetTransform → Execute Template → Cache (with FQDNs in annotations)
                         ↑
                    Called once

Cache → Endpoints() → Read FQDNs from annotations
              ↑
         Fast lookup
```

#### Annotation Storage

```go
const (
    GeneratedFQDNPrefix = "external-dns.alpha.kubernetes.io/generated-fqdns"
    // Kubernetes has 63 characters limitation on the annotation name
    // and 256KB on the total name+value.
)

// Example annotations after transform:
// external-dns.alpha.kubernetes.io/generated-fqdns: "api.prod.example.com,api.prod.example.org,api-v2.prod.example.com"
```

**Annotation chunking**:

```go
func storeGeneratedFQDNs(annotations map[string]string, fqdns []string) {
    const maxChunkSize = 200 // Leave room for overhead

    // Clear existing generated FQDN annotations
    for key := range annotations {
        if strings.HasPrefix(key, GeneratedFQDNPrefix) {
            delete(annotations, key)
        }
    }

    // Chunk and store
    joined := strings.Join(fqdns, ",")
    for i := 0; i < len(joined); i += maxChunkSize {
        end := i + maxChunkSize
        if end > len(joined) {
            end = len(joined)
        }
        key := fmt.Sprintf("%s%d", GeneratedFQDNPrefix, i/maxChunkSize)
        annotations[key] = joined[i:end]
    }
}
```

#### Transform Implementation

```go
// source/service.go

func NewServiceSource(
    ctx context.Context,
    kubeClient kubernetes.Interface,
    config *Config,
) (Source, error) {
    // ... informer setup ...

    serviceInformer := informerFactory.Core().V1().Services()

    // Get template from registry (Enhancement 1)
    var tmpl fqdn.Template
    if config.FQDNTemplateName != "" {
        var err error
        tmpl, err = fqdn.Get(config.FQDNTemplateName)
        if err != nil {
            return nil, fmt.Errorf("failed to get template %q: %w", config.FQDNTemplateName, err)
        }
    }

    // Apply transform if template is configured
    if tmpl != nil {
        _ = serviceInformer.Informer().SetTransform(func(i any) (any, error) {
            svc, ok := i.(*v1.Service)
            if !ok {
                return nil, fmt.Errorf("object is not a service")
            }

            // Check if already transformed (idempotent check)
            if _, exists := svc.Annotations[GeneratedFQDNPrefix+"0"]; exists {
                return svc, nil
            }

            // Execute template - context captured from outer scope
            hostnames, err := tmpl.Execute(ctx, svc)
            if err != nil {
                // Log error but don't skip record addition
                log.Errorf("Failed to execute FQDN template for service %s/%s: %v",
                    svc.Namespace, svc.Name, err)
                return svc, nil
            }

            // Create minimal service with generated FQDNs in annotations
            transformed := &v1.Service{
                ObjectMeta: metav1.ObjectMeta{
                    Name:      svc.Name,
                    Namespace: svc.Namespace,
                    // Copy existing annotations
                    Annotations: make(map[string]string),
                },
                Spec: v1.ServiceSpec{
                    Type: svc.Spec.Type,
                    // Only keep fields needed for endpoint generation
                },
                Status: svc.Status, // For LoadBalancer IPs
            }

            // Copy original annotations
            for k, v := range svc.Annotations {
                transformed.Annotations[k] = v
            }

            // Store generated FQDNs in chunked annotations
            storeGeneratedFQDNs(transformed.Annotations, hostnames)

            return transformed, nil
        })
    } else {
        // No template - apply memory optimization transform
        _ = serviceInformer.Informer().SetTransform(func(i any) (any, error) {
            svc, ok := i.(*v1.Service)
            if !ok {
                return nil, fmt.Errorf("object is not a service")
            }

            // Similar to pod.go pattern - keep minimal fields
            return &v1.Service{
                ObjectMeta: metav1.ObjectMeta{
                    Name:        svc.Name,
                    Namespace:   svc.Namespace,
                    Annotations: svc.Annotations,
                },
                Spec: v1.ServiceSpec{
                    Type: svc.Spec.Type,
                },
                Status: svc.Status,
            }, nil
        })
    }

    // ... rest of source setup ...
}
```

#### Reading FQDNs in Endpoints()

```go
func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
    // Read generated FQDNs from annotations
    hostnames := readGeneratedFQDNs(svc.Annotations)

    if len(hostnames) == 0 {
        log.Debugf("No generated FQDNs found for service %s/%s", svc.Namespace, svc.Name)
        return nil, nil
    }

    resource := fmt.Sprintf("service/%s/%s", svc.Namespace, svc.Name)
    ttl := annotations.TTLFromAnnotations(svc.Annotations, resource)
    targets := annotations.TargetsFromTargetAnnotation(svc.Annotations)

    if len(targets) == 0 {
        targets = extractLoadBalancerTargets(svc)
    }

    providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(svc.Annotations)

    var endpoints []*endpoint.Endpoint
    for _, hostname := range hostnames {
        endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
    }

    return endpoints, nil
}

func readGeneratedFQDNs(annotations map[string]string) []string {
    var chunks []string

    // Read all FQDN annotation chunks in order
    for i := 0; ; i++ {
        key := fmt.Sprintf("%s%d", GeneratedFQDNPrefix, i)
        chunk, exists := annotations[key]
        if !exists {
            break
        }
        chunks = append(chunks, chunk)
    }

    if len(chunks) == 0 {
        return nil
    }

    // Join chunks and split by comma
    joined := strings.Join(chunks, "")
    return strings.Split(joined, ",")
}
```

### Re-execution Behavior

**Transform is called** when:

- Object is initially added to cache
- Object is modified (any field change triggers transform)
- Informer re-syncs (periodic or on reconnect)

**Idempotent design**: Check if FQDNs already exist to avoid re-execution on every informer event.

**Caveat**: Transform executes even when irrelevant fields change (e.g., status updates). This is acceptable because:

- Idempotent check prevents duplicate work
- Template execution is still more efficient than current approach
- Most updates in Kubernetes are status updates, which we skip via idempotent check

### Configuration

No new configuration needed. Behavior is automatically enabled when:

1. Enhancement 1 (Template Registry) is implemented
2. Source is configured with a template name

### Implementation

#### Apply to All Sources

Add SetTransform to all sources that support FQDN templates:

**Pattern A sources** (string parameter):

- source/service.go ✅
- source/ingress.go ✅
- source/pod.go ✅ (already has transform for memory, extend it)
- source/node.go ✅
- source/istio_gateway.go ✅
- source/istio_virtualservice.go ✅
- source/contour_httpproxy.go ✅
- source/openshift_route.go ✅
- source/skipper_routegroup.go ✅

**Pattern B sources** (config parameter):

- source/gateway.go ✅

#### Helper Functions

Create shared utilities in `source/fqdn/transform.go`:

```go
// source/fqdn/transform.go

const (
    GeneratedFQDNPrefix = "external-dns.alpha.kubernetes.io/generated-fqdns-"
    MaxAnnotationChunk  = 200
)

// StoreGeneratedFQDNs saves hostnames in chunked annotations.
func StoreGeneratedFQDNs(annotations map[string]string, fqdns []string) {
    // Clear existing
    for key := range annotations {
        if strings.HasPrefix(key, GeneratedFQDNPrefix) {
            delete(annotations, key)
        }
    }

    if len(fqdns) == 0 {
        return
    }

    // Chunk and store
    joined := strings.Join(fqdns, ",")
    for i := 0; i < len(joined); i += MaxAnnotationChunk {
        end := i + MaxAnnotationChunk
        if end > len(joined) {
            end = len(joined)
        }
        key := fmt.Sprintf("%s%d", GeneratedFQDNPrefix, i/MaxAnnotationChunk)
        annotations[key] = joined[i:end]
    }
}

// ReadGeneratedFQDNs retrieves hostnames from chunked annotations.
func ReadGeneratedFQDNs(annotations map[string]string) []string {
    var chunks []string
    for i := 0; ; i++ {
        key := fmt.Sprintf("%s%d", GeneratedFQDNPrefix, i)
        chunk, exists := annotations[key]
        if !exists {
            break
        }
        chunks = append(chunks, chunk)
    }

    if len(chunks) == 0 {
        return nil
    }

    joined := strings.Join(chunks, "")
    return strings.Split(joined, ",")
}

// IsAlreadyTransformed checks if object has generated FQDNs.
func IsAlreadyTransformed(annotations map[string]string) bool {
    _, exists := annotations[GeneratedFQDNPrefix+"0"]
    return exists
}
```

### Performance Testing

Add comprehensive benchmarks to measure improvement:

```go
// source/fqdn/transform_bench_test.go

func BenchmarkTemplateExecutionCurrent(b *testing.B) {
    // Simulate current approach: execute template on every Endpoints() call
    tmpl := setupTemplate("{{.Name}}.{{.Namespace}}.example.com")
    svc := createTestService("api", "prod")

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = tmpl.Execute(context.Background(), svc)
    }
}

func BenchmarkTemplateExecutionTransform(b *testing.B) {
    // Simulate new approach: read from annotations
    svc := createTestService("api", "prod")
    StoreGeneratedFQDNs(svc.Annotations, []string{"api.prod.example.com"})

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ReadGeneratedFQDNs(svc.Annotations)
    }
}

func BenchmarkLargeCluster(b *testing.B) {
    // Test with large number of services
    const numServices = 10000

    services := make([]*v1.Service, numServices)
    for i := 0; i < numServices; i++ {
        services[i] = createTestService(fmt.Sprintf("svc-%d", i), "default")
    }

    tmpl := setupTemplate("{{.Name}}.{{.Namespace}}.example.com")

    b.Run("Current", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            for _, svc := range services {
                _, _ = tmpl.Execute(context.Background(), svc)
            }
        }
    })

    b.Run("Transform", func(b *testing.B) {
        // Pre-compute FQDNs in annotations
        for _, svc := range services {
            hostnames, _ := tmpl.Execute(context.Background(), svc)
            StoreGeneratedFQDNs(svc.Annotations, hostnames)
        }

        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            for _, svc := range services {
                _ = ReadGeneratedFQDNs(svc.Annotations)
            }
        }
    })
}

func BenchmarkComplexTemplate(b *testing.B) {
    // Test with complex template using multiple functions
    tmpl := setupTemplate("{{truncate 8 (sha256 .Name)}}.{{default .Labels.env \"prod\"}}.example.com")
    svc := createTestService("my-very-long-service-name", "default")
    svc.Labels = map[string]string{"env": "production"}

    b.Run("Current", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _, _ = tmpl.Execute(context.Background(), svc)
        }
    })

    b.Run("Transform", func(b *testing.B) {
        hostnames, _ := tmpl.Execute(context.Background(), svc)
        StoreGeneratedFQDNs(svc.Annotations, hostnames)

        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            _ = ReadGeneratedFQDNs(svc.Annotations)
        }
    })
}
```

### Benefits

- ✅ **Significant CPU reduction**: Template executed once instead of on every reconciliation
- ✅ **Memory optimization**: Transform can strip unused fields (similar to pod.go pattern)
- ✅ **Faster reconciliation**: Endpoints() becomes simple annotation lookup
- ✅ **Scales better**: Performance improvement increases with cluster size
- ✅ **Applied uniformly**: All sources benefit from optimization
- ✅ **No configuration needed**: Automatically enabled with templates
- ✅ **Error resilient**: Template failures don't block object caching

### Usage Examples

**Example 1: Service with template**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: additional.example.com
spec:
  type: LoadBalancer
  # ...

# After SetTransform (in cache):
# annotations:
#   external-dns.alpha.kubernetes.io/hostname: additional.example.com
#   external-dns.alpha.kubernetes.io/generated-fqdns-0: api.prod.example.com,api.prod.example.org
```

**Example 2: Large hostname list (chunked)**

```yaml
# Template generates many FQDNs (e.g., multi-zone deployment)
# annotations after transform:
#   external-dns.alpha.kubernetes.io/generated-fqdns-0: "svc.zone1.example.com,svc.zone2.example.com,svc.zone3.example.com,svc.zone4.example.com"
#   external-dns.alpha.kubernetes.io/generated-fqdns-1: "svc.zone10.example.com"
```

**Example 3: Template execution error**

```yaml
# Service with missing label required by template
# Template: {{.Labels.required}}.example.com
# Service has no "required" label

# Logs: ERROR Failed to execute FQDN template for service default/api: template execution failed
# Object still added to cache (no FQDNs generated)
# No endpoints created for this service until template succeeds
```

### Performance Goals

Based on benchmarks, expected improvements:

- **Current approach**: ~50,000 ns/op for simple template execution
- **Transform approach**: ~500 ns/op for annotation lookup
- **Expected improvement**: ~100x faster for Endpoints() calls
- **Large clusters** (10,000 services): Minutes saved per reconciliation cycle

## Enhancement 7: Target FQDN Template Support (Medium Priority)

### Motivation

#### Current Limitation

DNS records have two components: **Name** and **Value** (target).

**Current capabilities**:

- ✅ **Name (hostname)** can be templated via `--fqdn-template` or `external-dns.alpha.kubernetes.io/hostname`
- ❌ **Value (target)** is limited to static annotation (`external-dns.alpha.kubernetes.io/target`) or auto-detected from object status

**Limitations**:

- No template support for target values
- Cannot dynamically generate CNAME targets based on object fields
- Cannot create DNS records pointing to cluster-internal load balancers with templated names
- No access to object metadata when specifying targets
- Must hardcode target values in annotations

**Example use case**:

```yaml
# Current limitation: Need to hardcode target for each service
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: api.example.com
    external-dns.alpha.kubernetes.io/target: lb-prod.cluster.internal  # ❌ Hardcoded
```

**Desired**:

```yaml
# With target-fqdn template: Dynamic target based on namespace
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: api.example.com
    external-dns.alpha.kubernetes.io/target-fqdn: "lb-{{.Namespace}}.cluster.internal"  # ✅ Templated
# Result: api.example.com CNAME lb-prod.cluster.internal
```

This enhancement adds **template support for target values**, enabling dynamic CNAME/A record creation based on object fields.

### Design

#### Annotation-Based Target Templates

Add new annotation: `external-dns.alpha.kubernetes.io/target-fqdn`

**Key properties**:

- Supports template syntax (same as FQDN templates)
- Access to object fields (`.Name`, `.Namespace`, `.Labels`, etc.)
- Mutually exclusive with `external-dns.alpha.kubernetes.io/target`
- Template functions available (truncate, sha256, default, etc.)

**Mutual exclusivity**:

```go
// source/annotations/processors.go

func TargetsFromAnnotations(annotations map[string]string, obj kubeObject, targetFQDNTemplate fqdn.Template) (endpoint.Targets, error) {
    // Check for static target annotation
    staticTarget := annotations[TargetKey]
    templatedTarget := annotations[TargetFQDNKey]

    // Validate mutual exclusivity
    if staticTarget != "" && templatedTarget != "" {
        return nil, fmt.Errorf("annotations %q and %q are mutually exclusive", TargetKey, TargetFQDNKey)
    }

    // Use static target
    if staticTarget != "" {
        return TargetsFromTargetAnnotation(annotations), nil
    }

    // Use templated target
    if templatedTarget != "" {
        tmpl, err := fqdn.ParseTemplate(templatedTarget)
        if err != nil {
            return nil, fmt.Errorf("invalid target-fqdn template: %w", err)
        }
        targets, err := tmpl.Execute(context.Background(), obj)
        if err != nil {
            return nil, fmt.Errorf("failed to execute target-fqdn template: %w", err)
        }
        return endpoint.Targets(targets), nil
    }

    // Fallback to auto-detected targets (LoadBalancer IP, Ingress status, etc.)
    return nil, nil
}
```

#### Configuration-Based Target Templates

Support default target templates in YAML config:

```yaml
# external-dns-config.yaml

# Global default target template (applied to all sources)
fqdnTargetTemplate: "lb-{{.Namespace}}.cluster.internal"

# Per-source target templates
sources:
  service:
    fqdnTargetTemplate: "svc-lb-{{.Namespace}}.cluster.internal"
  ingress:
    fqdnTargetTemplate: "ingress-lb-{{.Namespace}}.cluster.internal"
  gateway:
    fqdnTargetTemplate: "gateway-{{.Name}}.cluster.internal"
```

**Priority order** (highest to lowest):

1. Annotation: `external-dns.alpha.kubernetes.io/target-fqdn`
2. Annotation: `external-dns.alpha.kubernetes.io/target` (static)
3. Per-source config: `sources.<source>.fqdnTargetTemplate`
4. Global config: `fqdnTargetTemplate`
5. Auto-detected from object status (LoadBalancer IP, Ingress status, etc.)

#### Implementation

**Annotation constants**:

```go
// source/annotations/annotations.go

const (
    // Existing
    TargetKey = AnnotationKeyPrefix + "target"  // Static target

    // New
    TargetFQDNKey = AnnotationKeyPrefix + "target-fqdn"  // Templated target
)
```

**Config structure**:

```go
// source/store.go

type Config struct {
    // Existing fields
    FQDNTemplate       string
    FQDNTemplateName   string

    // New: Target template support
    FQDNTargetTemplate string  // Default target template (global)
}

// source/service.go (example)

type serviceSource struct {
    // Existing fields
    fqdnTemplate fqdn.Template

    // New: Target template
    fqdnTargetTemplate fqdn.Template  // Per-source target template
}

func NewServiceSource(ctx context.Context, kubeClient kubernetes.Interface, config *Config) (Source, error) {
    // ... existing setup ...

    var targetTmpl fqdn.Template
    if config.FQDNTargetTemplate != "" {
        var err error
        targetTmpl, err = fqdn.ParseTemplate(config.FQDNTargetTemplate)
        if err != nil {
            return nil, fmt.Errorf("invalid fqdn-target-template: %w", err)
        }
    }

    return &serviceSource{
        // ... existing fields ...
        fqdnTargetTemplate: targetTmpl,
    }, nil
}
```

**Target resolution with templates**:

```go
// source/service.go

func (sc *serviceSource) endpointsFromService(svc *v1.Service) ([]*endpoint.Endpoint, error) {
    // ... hostname generation ...

    // Resolve targets with template support
    targets, err := sc.resolveTargets(svc)
    if err != nil {
        return nil, err
    }

    // ... create endpoints ...
}

func (sc *serviceSource) resolveTargets(svc *v1.Service) (endpoint.Targets, error) {
    annotations := svc.Annotations

    // Priority 1: Check for target-fqdn annotation (templated)
    if targetFQDNTemplate, exists := annotations[annotations.TargetFQDNKey]; exists {
        // Check mutual exclusivity
        if _, hasStatic := annotations[annotations.TargetKey]; hasStatic {
            return nil, fmt.Errorf("service %s/%s has both %q and %q annotations (mutually exclusive)",
                svc.Namespace, svc.Name, annotations.TargetKey, annotations.TargetFQDNKey)
        }

        // Parse and execute annotation template
        tmpl, err := fqdn.ParseTemplate(targetFQDNTemplate)
        if err != nil {
            return nil, fmt.Errorf("invalid target-fqdn annotation: %w", err)
        }

        targets, err := tmpl.Execute(context.Background(), svc)
        if err != nil {
            return nil, fmt.Errorf("failed to execute target-fqdn template: %w", err)
        }

        return endpoint.Targets(targets), nil
    }

    // Priority 2: Check for static target annotation
    if staticTargets := annotations.TargetsFromTargetAnnotation(annotations); len(staticTargets) > 0 {
        return staticTargets, nil
    }

    // Priority 3: Use source-level target template (from config)
    if sc.fqdnTargetTemplate != nil {
        targets, err := sc.fqdnTargetTemplate.Execute(context.Background(), svc)
        if err != nil {
            return nil, fmt.Errorf("failed to execute source target template: %w", err)
        }
        return endpoint.Targets(targets), nil
    }

    // Priority 4: Auto-detect from LoadBalancer status
    return extractLoadBalancerTargets(svc), nil
}
```

### Configuration

#### CLI Flags

```bash
# Global default target template
--fqdn-target-template="lb-{{.Namespace}}.cluster.internal"

# Or via config file
--fqdn-template-config=config.yaml
```

#### YAML Configuration

```yaml
# external-dns-config.yaml

# Global default
fqdnTargetTemplate: "lb-{{.Namespace}}.cluster.internal"

# Per-source overrides
sources:
  service:
    fqdnTargetTemplate: "svc-{{.Name}}-{{.Namespace}}.lb.cluster.internal"
  ingress:
    fqdnTargetTemplate: "ingress-{{.Namespace}}.lb.cluster.internal"
```

#### Annotations

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  annotations:
    # Templated target (new)
    external-dns.alpha.kubernetes.io/target-fqdn: "lb-{{.Namespace}}.cluster.internal"

    # ❌ Cannot use both - will error
    # external-dns.alpha.kubernetes.io/target: "lb-prod.cluster.internal"
```

### Validation

**Mutual exclusivity check**:

```go
func validateTargetAnnotations(annotations map[string]string) error {
    hasStatic := annotations[annotations.TargetKey] != ""
    hasTemplated := annotations[annotations.TargetFQDNKey] != ""

    if hasStatic && hasTemplated {
        return fmt.Errorf("cannot specify both %q and %q annotations",
            annotations.TargetKey, annotations.TargetFQDNKey)
    }
    return nil
}
```

**Template syntax validation**:

```go
func validateTargetFQDNTemplate(templateStr string) error {
    _, err := fqdn.ParseTemplate(templateStr)
    if err != nil {
        return fmt.Errorf("invalid target-fqdn template syntax: %w", err)
    }
    return nil
}
```

### Benefits

- ✅ **Dynamic target generation**: Create targets based on object metadata
- ✅ **Reduced hardcoding**: No need to specify targets for each object
- ✅ **Consistent patterns**: Use templates for both names and targets
- ✅ **Namespace isolation**: Different targets per namespace automatically
- ✅ **Cluster-internal CNAMEs**: Point to internal load balancers with templated names
- ✅ **Backward compatible**: Existing `target` annotation still works
- ✅ **Template function support**: Use truncate, sha256, default, etc.

### Usage Examples

TODO:

- fetch IP from fqdn if cname
- if contains https://github.com/kubernetes-sigs/external-dns/issues/5661

**Example 1: Namespace-based load balancer CNAME**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: api.example.com
    external-dns.alpha.kubernetes.io/target-fqdn: "lb-{{.Namespace}}.cluster.internal"
spec:
  type: LoadBalancer

# Result:
# api.example.com CNAME lb-prod.cluster.internal
```

**Example 2: Service name-based internal routing**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: payment-service
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: payment.example.com
    external-dns.alpha.kubernetes.io/target-fqdn: "{{.Name}}.{{.Namespace}}.svc.cluster.local"
spec:
  type: ClusterIP

# Result:
# payment.example.com CNAME payment-service.prod.svc.cluster.local
```

**Example 3: Label-based routing**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  labels:
    region: us-west-2
  annotations:
    external-dns.alpha.kubernetes.io/hostname: api.example.com
    external-dns.alpha.kubernetes.io/target-fqdn: "lb-{{.Labels.region}}.cluster.internal"

# Result:
# api.example.com CNAME lb-us-west-2.cluster.internal
```

**Example 4: Conditional target with default**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  labels:
    lb-override: special-lb
  annotations:
    external-dns.alpha.kubernetes.io/hostname: api.example.com
    external-dns.alpha.kubernetes.io/target-fqdn: "{{default .Labels.lb-override \"default-lb\"}}.cluster.internal"

# Result:
# api.example.com CNAME special-lb.cluster.internal
```

**Example 5: Multi-zone targets**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: api.example.com
    # Comma-separated for multiple targets
    external-dns.alpha.kubernetes.io/target-fqdn: "lb-{{.Namespace}}.us-west-2.cluster.internal,lb-{{.Namespace}}.us-east-1.cluster.internal"

# Result:
# api.example.com CNAME lb-prod.us-west-2.cluster.internal
# api.example.com CNAME lb-prod.us-east-1.cluster.internal
```

**Example 6: Global default in config**

```yaml
# Config file
fqdnTargetTemplate: "lb-{{.Namespace}}.cluster.internal"

# Service (no target annotation needed)
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: api.example.com
spec:
  type: LoadBalancer

# Result (uses global default):
# api.example.com CNAME lb-prod.cluster.internal
```

**Example 7: Per-source config**

```yaml
# Config file
sources:
  service:
    fqdnTargetTemplate: "svc-lb-{{.Namespace}}.internal"
  ingress:
    fqdnTargetTemplate: "ingress-lb-{{.Namespace}}.internal"

# Service
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: api.example.com

# Result:
# api.example.com CNAME svc-lb-prod.internal

# Ingress
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: web
  namespace: prod
  annotations:
    external-dns.alpha.kubernetes.io/hostname: web.example.com

# Result:
# web.example.com CNAME ingress-lb-prod.internal
```

### Error Handling

**Invalid template syntax**:

```bash
ERROR: Service prod/api has invalid target-fqdn template "{{.InvalidField}}":
template execution failed: can't evaluate field InvalidField
```

**Template execution failure**:

```bash
ERROR: Failed to execute target-fqdn template for service prod/api:
template execution failed: missing required label "region"
```

### Integration with Enhancement 6 (SetTransform)

Target templates should also be executed in SetTransform for performance:

```go
_ = serviceInformer.Informer().SetTransform(func(i any) (any, error) {
    svc := i.(*v1.Service)

    // Execute FQDN template (Enhancement 6)
    if fqdnTemplate != nil {
        hostnames, _ := fqdnTemplate.Execute(ctx, svc)
        fqdn.StoreGeneratedFQDNs(svc.Annotations, hostnames)
    }

    // Execute target FQDN template (Enhancement 7)
    if targetFQDNTemplate, exists := svc.Annotations[annotations.TargetFQDNKey]; exists {
        tmpl, _ := fqdn.ParseTemplate(targetFQDNTemplate)
        targets, _ := tmpl.Execute(ctx, svc)

        // Store in separate annotation
        svc.Annotations["external-dns.alpha.kubernetes.io/generated-targets-0"] =
            strings.Join(targets, ",")
    }

    return svc, nil
})
```

### Testing

```go
// source/annotations/processors_test.go

func TestTargetFQDNAnnotation(t *testing.T) {
    tests := []struct {
        name        string
        svc         *v1.Service
        expected    endpoint.Targets
        expectError bool
    }{
        {
            name: "templated target with namespace",
            svc: &v1.Service{
                ObjectMeta: metav1.ObjectMeta{
                    Name:      "api",
                    Namespace: "prod",
                    Annotations: map[string]string{
                        annotations.TargetFQDNKey: "lb-{{.Namespace}}.cluster.internal",
                    },
                },
            },
            expected: endpoint.Targets{"lb-prod.cluster.internal"},
        },
        {
            name: "mutual exclusivity error",
            svc: &v1.Service{
                ObjectMeta: metav1.ObjectMeta{
                    Annotations: map[string]string{
                        annotations.TargetKey:     "static.example.com",
                        annotations.TargetFQDNKey: "{{.Name}}.cluster.internal",
                    },
                },
            },
            expectError: true,
        },
        {
            name: "multiple targets",
            svc: &v1.Service{
                ObjectMeta: metav1.ObjectMeta{
                    Name:      "api",
                    Namespace: "prod",
                    Annotations: map[string]string{
                        annotations.TargetFQDNKey: "lb-{{.Namespace}}.us-west-2.internal,lb-{{.Namespace}}.us-east-1.internal",
                    },
                },
            },
            expected: endpoint.Targets{
                "lb-prod.us-west-2.internal",
                "lb-prod.us-east-1.internal",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            targets, err := resolveTargets(tt.svc)
            if tt.expectError {
                assert.Error(t, err)
                return
            }
            assert.NoError(t, err)
            assert.Equal(t, tt.expected, targets)
        })
    }
}
```

---

# Implementation Examples

This section contains detailed Go code implementations for each enhancement. The enhancements above focus on motivation, design concepts, and usage, while this section provides the technical implementation details.

## Enhancement 1: Template Interface and Registry

**Core Interfaces**

```go
// template.go
package fqdn

import (
    "context"
    "text/template"
)

// Template represents a parsed, reusable FQDN template.
// It provides a domain-specific interface for generating DNS hostnames
// from Kubernetes objects.
type Template interface {
    // Execute generates hostnames from a Kubernetes object.
    // Returns a slice of hostname strings or an error if template execution fails.
    Execute(ctx context.Context, obj kubeObject) ([]string, error)

    // String returns the original template string for debugging/logging.
    String() string
}

// kubeObject represents any Kubernetes object that can be used in templates.
type kubeObject interface {
    runtime.Object
    metav1.Object
}

// Option configures template behavior.
type Option func(*templateOptions)

type templateOptions struct {
    // Future: validation, max length, etc.
}
```

**Template Creation**

```go
// New creates a new Template from a template string.
// Returns error if template parsing fails.
func New(templateStr string, opts ...Option) (Template, error)
```

**Registry API**

```go
// registry.go

// Register adds a template to the global registry.
// Returns error if:
// - name is empty
// - templateStr is invalid
// - name already exists (call Unregister first, or use Update)
func Register(name string, templateStr string, opts ...Option) error

// MustRegister is like Register but panics on error.
// Useful for startup registration where failure should be fatal.
func MustRegister(name string, templateStr string, opts ...Option)

// Get retrieves a registered template by name.
// Returns error if template not found.
func Get(name string) (Template, error)

// List returns all registered template names.
func List() []string

// Clear removes all templates (useful for testing).
func Clear()
```

**Registry Implementation**

```go
// registry.go

// TemplateRegistry manages parsed templates.
type TemplateRegistry struct {
    templates map[string]Template
    mu        sync.RWMutex
}

var (
    globalRegistry     = NewTemplateRegistry()
    ErrAlreadyExists   = errors.New("template already exists")
    ErrNotFound        = errors.New("template not found")
    ErrEmptyName       = errors.New("template name cannot be empty")
    ErrInvalidTemplate = errors.New("invalid template string")
)

func NewTemplateRegistry() *TemplateRegistry {
    return &TemplateRegistry{
        templates: make(map[string]Template),
    }
}

func (r *TemplateRegistry) Register(name string, templateStr string, opts ...Option) error {
    if name == "" {
        return ErrEmptyName
    }

    tmpl, err := New(templateStr, opts...)
    if err != nil {
        return fmt.Errorf("%w: %v", ErrInvalidTemplate, err)
    }

    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.templates[name]; exists {
        return fmt.Errorf("%w: %q", ErrAlreadyExists, name)
    }

    r.templates[name] = tmpl
    return nil
}

func (r *TemplateRegistry) Get(name string) (Template, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    tmpl, exists := r.templates[name]
    if !exists {
        return nil, fmt.Errorf("%w: %q", ErrNotFound, name)
    }

    return tmpl, nil
}
```

**Template Implementation**

```go
// template.go

// textTemplate wraps text/template.Template to implement Template interface.
type textTemplate struct {
    tmpl *template.Template
    raw  string
    opts templateOptions
}

func New(templateStr string, opts ...Option) (Template, error) {
    if templateStr == "" {
        return nil, nil
    }

    // Apply options
    options := templateOptions{}
    for _, opt := range opts {
        opt(&options)
    }

    // Parse template with standard functions
    tmpl, err := ParseTemplate(templateStr)
    if err != nil {
        return nil, err
    }

    return &textTemplate{
        tmpl: tmpl,
        raw:  templateStr,
        opts: options,
    }, nil
}

func (t *textTemplate) Execute(ctx context.Context, obj kubeObject) ([]string, error) {
    // Reuse existing ExecTemplate logic
    return ExecTemplate(t.tmpl, obj)
}

func (t *textTemplate) String() string {
    return t.raw
}
```

**Store Configuration**

```go
// store.go

type Config struct {
    // ... existing fields ...
    FQDNTemplate     string  // Keep for backward compatibility
    FQDNTemplateName string  // NEW: Template registry name
}

func NewSourceConfig(cfg *externaldns.Config) (*Config, error) {
    config := &Config{
        FQDNTemplate: cfg.FQDNTemplate,
        // ... other fields ...
    }

    // Register template at startup
    if cfg.FQDNTemplate != "" {
        if err := fqdn.Register("default", cfg.FQDNTemplate); err != nil {
            return nil, fmt.Errorf("invalid FQDN template: %w", err)
        }
        config.FQDNTemplateName = "default"
    }

    return config, nil
}
```

**Gateway Source Migration**

```go
// gateway.go - BEFORE
func NewGatewayRouteSource(..., config *Config) (Source, error) {
    tmpl, err := fqdn.ParseTemplate(config.FQDNTemplate)  // ❌ Remove
    if err != nil {
        return nil, err
    }
    // ...
}

// gateway.go - AFTER
func NewGatewayRouteSource(..., config *Config) (Source, error) {
    var tmpl fqdn.Template  // ✅ Use interface
    var err error

    if config.FQDNTemplateName != "" {
        tmpl, err = fqdn.Get(config.FQDNTemplateName)  // ✅ Get from registry
        if err != nil {
            return nil, err
        }
    }
    // ...
}
```

**Source Struct Field Type Update**

```go
// BEFORE
type gatewayRouteSource struct {
    fqdnTemplate *template.Template  // Generic type
}

// AFTER
type gatewayRouteSource struct {
    fqdnTemplate fqdn.Template  // Domain interface
}
```

**Service Source Migration**

```go
// service.go - BEFORE
func NewServiceSource(
    ctx context.Context,
    kubeClient kubernetes.Interface,
    namespace, annotationFilter, fqdnTemplate string,
    // ... 10+ more parameters
) (Source, error)

// service.go - AFTER
func NewServiceSource(
    ctx context.Context,
    kubeClient kubernetes.Interface,
    config *Config,  // ✅ Simplified signature
) (Source, error) {
    tmpl, err := fqdn.Get(config.FQDNTemplateName)
    // ...
}
```

**Template Execution Update**

```go
// BEFORE
hostnames, err := fqdn.ExecTemplate(sc.fqdnTemplate, service)

// AFTER
hostnames, err := sc.fqdnTemplate.Execute(ctx, service)
```

**Integration Test**

```go
func TestTemplateParsedOnce(t *testing.T) {
    .... setup ...

    config, err := NewSourceConfig(cfg)
    require.NoError(t, err)

    // Create multiple sources
    NewServiceSource(ctx, client, config)
    NewIngressSource(ctx, client, config)
    NewGatewayRouteSource(ctx, clients, config)

    // Verify template parsed only once
    assert.Equal(t, 1, parseCount, "Template should be parsed exactly once")
}
```

***Alternative Implementations***

**Lazy Initialization Alternative**

```go
type serviceSource struct {
    fqdnTemplateStr string
    fqdnTemplate    *template.Template
    parseOnce       sync.Once
}

func (s *serviceSource) getTemplate() (*template.Template, error) {
    var err error
    s.parseOnce.Do(func() {
        s.fqdnTemplate, err = fqdn.ParseTemplate(s.fqdnTemplateStr)
    })
    return s.fqdnTemplate, err
}
```

**Config-Level Caching Alternative**

```go
type Config struct {
    FQDNTemplate       string
    parsedFQDNTemplate *template.Template  // Cached
}

func NewSourceConfig(cfg *externaldns.Config) (*Config, error) {
    tmpl, err := fqdn.ParseTemplate(cfg.FQDNTemplate)
    return &Config{
        FQDNTemplate:       cfg.FQDNTemplate,
        parsedFQDNTemplate: tmpl,
    }, nil
}
```

**Registry + Dependency Injection Alternative**

```go
// main.go
fqdn.Register("default", cfg.FQDNTemplate)

// sources.go
func NewServiceSource(..., templateName string) (Source, error) {
    tmpl, _ := fqdn.Get(templateName)
}
```

**Backward Compatibility**

```go
// fqdn.go - Keep existing functions

// ParseTemplate parses a template string.
// Deprecated: Use New() for new code. This function is maintained for
// backward compatibility and will be removed in follow-up releases.
func ParseTemplate(input string) (*template.Template, error) {
    // Existing implementation unchanged
}

// ExecTemplate executes a template against a Kubernetes object.
// Deprecated: Use Template.Execute() for new code. This function is maintained
// for backward compatibility and will be removed in follow-up releases.
func ExecTemplate(tmpl *template.Template, obj kubeObject) ([]string, error) {
    // Existing implementation unchanged
}
```

## Enhancement 2: Multi-Template Support with Selectors

**TemplateSet Design**

```go
// TemplateSet manages multiple templates for different sources/conditions.
// Implements the Template interface, allowing it to be used anywhere a Template is expected.
type TemplateSet struct {
    templates        map[string][]Template  // condition -> array of templates
    defaultTemplates []Template             // fallback templates
    selector         TemplateSelector
}

// TemplateSelector chooses which template group to use based on object properties.
type TemplateSelector interface {
    Select(obj kubeObject) (string, error)
}

// Register a template set
func RegisterSet(name string, config TemplateSetConfig) error

type TemplateSetConfig struct {
    // Map of condition -> template(s)
    // Supports both single string and array of strings
    Templates map[string][]string

    // Default template(s) if no condition matches
    // Supports both single string and array of strings
    Default []string

    // Selector strategy
    Strategy SelectorStrategy

    // Selector key (for label/annotation strategies)
    Selector string
}

type SelectorStrategy string

const (
    StrategyByKind       SelectorStrategy = "kind"        // Select by .Kind
    StrategyByLabel      SelectorStrategy = "label"       // Select by label value
    StrategyByAnnotation SelectorStrategy = "annotation"  // Select by annotation
    StrategyByNamespace  SelectorStrategy = "namespace"   // Select by namespace
)
```

**TemplateSet Execute Implementation**

```go
func (ts *TemplateSet) Execute(ctx context.Context, obj kubeObject) ([]string, error) {
    // Select which template group to use
    templateName, err := ts.selector.Select(obj)
    if err != nil || templateName == "" {
        // Use default templates
        return ts.executeTemplates(ctx, obj, ts.defaultTemplates)
    }

    // Get template(s) for selected condition
    tmpls, ok := ts.templates[templateName]
    if !ok {
        tmpls = ts.defaultTemplates
    }

    return ts.executeTemplates(ctx, obj, tmpls)
}

func (ts *TemplateSet) executeTemplates(ctx context.Context, obj kubeObject, tmpls []Template) ([]string, error) {
    var allHostnames []string
    seen := make(map[string]bool)

    for _, tmpl := range tmpls {
        hostnames, err := tmpl.Execute(ctx, obj)
        if err != nil {
            return nil, err
        }

        // Deduplicate and warn on duplicates
        for _, hostname := range hostnames {
            if seen[hostname] {
                log.Warnf("Duplicate hostname %q generated by template %q for %s/%s, ignoring duplicate",
                    hostname, tmpl.String(), obj.GetNamespace(), obj.GetName())
                continue
            }
            seen[hostname] = true
            allHostnames = append(allHostnames, hostname)
        }
    }

    return allHostnames, nil
}

func (ts *TemplateSet) String() string {
    return fmt.Sprintf("TemplateSet[strategy=%s]", ts.selector)
}
```

**Config Loading with Normalization**

```go
func parseTemplateSetConfig(data map[interface{}]interface{}) TemplateSetConfig {
    config := TemplateSetConfig{
        Templates: make(map[string][]string),
    }

    if strategy, ok := data["strategy"].(string); ok {
        config.Strategy = SelectorStrategy(strategy)
    }

    if selector, ok := data["selector"].(string); ok {
        config.Selector = selector
    }

    if defaultVal, ok := data["default"]; ok {
        config.Default = parseTemplateValue(defaultVal)
    }

    if templates, ok := data["templates"].(map[interface{}]interface{}); ok {
        for k, v := range templates {
            if key, ok := k.(string); ok {
                config.Templates[key] = parseTemplateValue(v)
            }
        }
    }

    return config
}
```

**Duplicate Detection and Deduplication**

```go
func deduplicateHostnames(hostnames []string, context string) []string {
    seen := make(map[string]bool)
    unique := make([]string, 0, len(hostnames))

    for _, hostname := range hostnames {
        if seen[hostname] {
            log.Debuf("Duplicate hostname %q in %s, ignoring duplicate", hostname, context)
            continue
        }
        seen[hostname] = true
        unique = append(unique, hostname)
    }

    return unique
}
```

## Enhancement 3: DNS Validation and Sanitization

**Validator and Sanitizer Interfaces**

```go
// Validator checks if hostnames are DNS-compliant.
type Validator interface {
    Validate(hostname string) error
}

// Sanitizer fixes invalid hostnames and validates the result.
type Sanitizer interface {
    // Sanitize transforms hostname to be DNS-compliant.
    // Returns error if sanitization fails to produce valid hostname.
    Sanitize(hostname string) (string, error)
}

// ValidationOptions configures hostname validation and sanitization.
type ValidationOptions struct {
    Enabled       bool
    AutoFix       bool   // Automatically sanitize invalid hostnames
    MaxLength     int    // Default: 253 (DNS standard)
    OnTooLong     LengthErrorAction  // What to do if hostname > MaxLength
    AllowUnicode  bool   // Support IDN via Punycode (RFC 3492)
    StrictRFC1123 bool   // Enforce strict RFC 1123 compliance
    OnError       ValidationErrorAction
}

type LengthErrorAction string

const (
    LengthErrorActionIgnore LengthErrorAction = "ignore" // Use as-is (may fail at DNS provider)
    LengthErrorActionError  LengthErrorAction = "error"  // Return error, don't create record
)

type ValidationErrorAction string

const (
    ValidationErrorActionSkip ValidationErrorAction = "skip" // Log warning, skip invalid hostname
    ValidationErrorActionFail ValidationErrorAction = "fail" // Return error, stop processing
)

// Built-in validators
var (
    RFC1123Validator = &rfc1123Validator{}
    LengthValidator  = &lengthValidator{maxLen: 253}
    IDNValidator     = &idnValidator{}
)

// DefaultSanitizer provides sensible defaults for hostname sanitization
var DefaultSanitizer = &defaultSanitizer{
    toLowercase:     true,
    replaceInvalid:  true,
    replacementChar: '-',
    maxLength:       253,
    allowUnicode:    false,
}
```

**ValidateAndSanitize Function**

```go
// ValidateAndSanitize processes hostnames after template execution.
// This is called by sources or registry, not by templates themselves.
func ValidateAndSanitize(hostnames []string, opts ValidationOptions) ([]string, error) {
    if !opts.Enabled {
        return hostnames, nil
    }

    var result []string
    validator := RFC1123Validator

    for _, hostname := range hostnames {
        // Apply prefix/suffix if configured (external-dns flags)
        hostname = applyPrefixSuffix(hostname)

        // Validate
        if err := validator.Validate(hostname); err != nil {
            if !opts.AutoFix {
                return handleValidationError(hostname, err, opts.OnError)
            }

            // Sanitize invalid hostname
            sanitized, err := DefaultSanitizer.Sanitize(hostname, opts)
            if err != nil {
                return handleValidationError(hostname, err, opts.OnError)
            }
            hostname = sanitized
        }

        // Check length
        if len(hostname) > opts.MaxLength {
            if opts.OnTooLong == LengthErrorActionError {
                return nil, fmt.Errorf("hostname too long: %q (%d chars, max %d)",
                    hostname, len(hostname), opts.MaxLength)
            }
            // LengthErrorActionIgnore: use as-is, may fail at provider
            log.Warnf("Hostname exceeds max length: %q (%d chars, max %d)",
                hostname, len(hostname), opts.MaxLength)
        }

        result = append(result, hostname)
    }

    return result, nil
}

func handleValidationError(hostname string, err error, action ValidationErrorAction) ([]string, error) {
    switch action {
    case ValidationErrorActionSkip:
        log.Warnf("Skipping invalid hostname %q: %v", hostname, err)
        return []string{}, nil
    case ValidationErrorActionFail:
        return nil, fmt.Errorf("invalid hostname %q: %w", hostname, err)
    default:
        return nil, fmt.Errorf("unknown validation error action: %q", action)
    }
}

// applyPrefixSuffix applies external-dns --txt-prefix and --txt-suffix flags
func applyPrefixSuffix(hostname string) string {
    // Integration with existing external-dns configuration
    // Implementation depends on how prefix/suffix are configured
    return hostname
}
```

**RFC 1123 Validator**

```go
// fqdn/validator.go
type rfc1123Validator struct{}

func (v *rfc1123Validator) Validate(hostname string) error {
    if len(hostname) > 253 {
        return fmt.Errorf("hostname too long: %d chars (max 253)", len(hostname))
    }

    if hostname == "" {
        return fmt.Errorf("hostname cannot be empty")
    }

    labels := strings.Split(hostname, ".")
    for i, label := range labels {
        if err := v.validateLabel(label, i); err != nil {
            return err
        }
    }
    return nil
}

func (v *rfc1123Validator) validateLabel(label string, index int) error {
    if len(label) == 0 {
        return fmt.Errorf("label %d is empty", index)
    }
    if len(label) > 63 {
        return fmt.Errorf("label %d too long: %d chars (max 63): %q", index, len(label), label)
    }

    // Check first character
    if !isAlphanumeric(rune(label[0])) {
        return fmt.Errorf("label %d must start with alphanumeric: %q", index, label)
    }

    // Check last character
    if !isAlphanumeric(rune(label[len(label)-1])) {
        return fmt.Errorf("label %d must end with alphanumeric: %q", index, label)
    }

    // Check middle characters
    for j, r := range label {
        if !isAlphanumeric(r) && r != '-' {
            return fmt.Errorf("label %d contains invalid character at position %d: %q", index, j, string(r))
        }
    }

    return nil
}

func isAlphanumeric(r rune) bool {
    return (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}
```

**Sanitizer Implementation**

```go
// fqdn/sanitizer.go
import "golang.org/x/net/idna"

type defaultSanitizer struct {
    toLowercase     bool
    replaceInvalid  bool
    replacementChar rune
    maxLength       int
    allowUnicode    bool
}

func (s *defaultSanitizer) Sanitize(hostname string, opts ValidationOptions) (string, error) {
    original := hostname

    // Step 1: Handle Unicode (Punycode encoding per RFC 3492)
    if opts.AllowUnicode && containsUnicode(hostname) {
        var err error
        hostname, err = idna.ToASCII(hostname)
        if err != nil {
            return "", fmt.Errorf("failed to convert Unicode hostname to Punycode: %w", err)
        }
    }

    // Step 2: Lowercase
    if s.toLowercase {
        hostname = strings.ToLower(hostname)
    }

    // Step 3: Replace invalid characters
    if s.replaceInvalid {
        hostname = s.replaceInvalidChars(hostname)
    }

    // Step 4: Validate result
    if err := RFC1123Validator.Validate(hostname); err != nil {
        return "", fmt.Errorf("sanitization of %q failed to produce valid hostname %q: %w",
            original, hostname, err)
    }

    return hostname, nil
}

func (s *defaultSanitizer) replaceInvalidChars(hostname string) string {
    var result strings.Builder
    prevWasDash := false

    for i, r := range hostname {
        if isValidDNSChar(r) {
            result.WriteRune(r)
            prevWasDash = (r == '-')
        } else {
            // Don't create consecutive dashes or leading/trailing dashes
            if !prevWasDash && i > 0 && i < len(hostname)-1 {
                result.WriteRune(s.replacementChar)
                prevWasDash = true
            }
        }
    }

    return result.String()
}

func isValidDNSChar(r rune) bool {
    return (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '.' || r == '-'
}

func containsUnicode(s string) bool {
    for _, r := range s {
        if r > 127 {
            return true
        }
    }
    return false
}
```

## Enhancement 4: Auto-Generated Documentation

**Source Interfaces**

```go
// source/source.go

// FQDNTemplateSupport indicates a source supports FQDN templates.
type FQDNTemplateSupport interface {
    SupportsFQDNTemplate() bool
    SupportsHostnameAnnotation() bool
}

// Example implementation
// source/service.go
func (s *serviceSource) SupportsFQDNTemplate() bool {
    return true
}

func (s *serviceSource) SupportsHostnameAnnotation() bool {
    return !s.ignoreHostnameAnnotation
}
```

**Godoc-Style Comments**

```go
// source/fqdn/functions.go

// truncate limits a string to maxLen characters.
//
// Example:
//
//	{{truncate 10 .Name}}
//	Input: "very-long-service-name"
//	Output: "very-long-"
//
// @since v0.20.0
func truncate(maxLen int, s string) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen]
}
```

**Structured Test Tags**

```go
// source/fqdn/fqdn_test.go

// @example Simple Template
// @description Basic hostname generation
// @since v0.20.0
func TestSimpleTemplate(t *testing.T) {
    tmpl := New("{{.Name}}.example.com")

    // Template: {{.Name}}.example.com
    // Input: Service "test" in namespace "default"
    // Output: ["test.example.com"]

    result, _ := tmpl.Execute(ctx, testService)
    assert.Equal(t, []string{"test.example.com"}, result)
}

// @example Multi-Zone Templates
// @description Generate DNS records in multiple zones
// @since v0.21.0
func TestMultiZoneTemplate(t *testing.T) {
    config := TemplateSetConfig{
        Default: []string{
            "{{.Name}}.example.com",
            "{{.Name}}.example.org",
        },
    }

    // Template: Multiple templates in array
    // Input: Service "api" in namespace "prod"
    // Output: ["api.example.com", "api.example.org"]

    // ...
}
```

**Generator Tool**

```go
// internal/gen/docs/fqdn/main.go

package main

import (
    "flag"
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "io/ioutil"
    "path/filepath"
)

func main() {
    outputDir := flag.String("output", "docs/advanced/fqdn", "Output directory")
    flag.Parse()

    // Generate all documentation
    if err := generateDocs(*outputDir); err != nil {
        panic(err)
    }
}

func generateDocs(outputDir string) error {
    // 1. Generate template examples from tests
    examples, err := extractExamplesFromTests("source/fqdn/fqdn_test.go")
    if err != nil {
        return err
    }
    if err := writeMarkdown(filepath.Join(outputDir, "templates.md"), examples); err != nil {
        return err
    }

    // 2. Generate source support matrix
    sources, err := scanSourceInterfaces("source/")
    if err != nil {
        return err
    }
    if err := writeMarkdown(filepath.Join(outputDir, "sources.md"), sources); err != nil {
        return err
    }

    // 3. Generate function reference from godoc
    functions, err := extractFunctionDocs("source/fqdn/functions.go")
    if err != nil {
        return err
    }
    if err := writeMarkdown(filepath.Join(outputDir, "functions.md"), functions); err != nil {
        return err
    }

    return nil
}
```

**Extract Examples from Tests**

```go
// internal/gen/docs/fqdn/templates.go

type Example struct {
    Name        string
    Description string
    Since       string
    Template    string
    Input       string
    Output      string
}

func extractExamplesFromTests(testFile string) (string, error) {
    fset := token.NewFileSet()
    node, err := parser.ParseFile(fset, testFile, nil, parser.ParseComments)
    if err != nil {
        return "", err
    }

    var examples []Example

    // Look for @example tags in comments
    for _, commentGroup := range node.Comments {
        text := commentGroup.Text()
        if strings.Contains(text, "@example") {
            example := parseExampleComment(text)
            examples = append(examples, example)
        }
    }

    // Also parse inline comments in test functions
    ast.Inspect(node, func(n ast.Node) bool {
        if fn, ok := n.(*ast.FuncDecl); ok {
            if strings.HasPrefix(fn.Name.Name, "Test") {
                example := parseTestFunction(fn)
                if example != nil {
                    examples = append(examples, *example)
                }
            }
        }
        return true
    })

    return formatExamplesMarkdown(examples), nil
}

func parseExampleComment(text string) Example {
    // Parse:
    // @example Simple Template
    // @description Basic hostname generation
    // @since v0.14.0

    lines := strings.Split(text, "\n")
    example := Example{}

    for _, line := range lines {
        if strings.Contains(line, "@example") {
            example.Name = strings.TrimSpace(strings.TrimPrefix(line, "@example"))
        }
        if strings.Contains(line, "@description") {
            example.Description = strings.TrimSpace(strings.TrimPrefix(line, "@description"))
        }
        if strings.Contains(line, "@since") {
            example.Since = strings.TrimSpace(strings.TrimPrefix(line, "@since"))
        }
    }

    return example
}
```

**Scan Source Interfaces**

```go
// internal/gen/docs/fqdn/sources.go

type SourceInfo struct {
    Name                       string
    File                       string
    SupportsFQDN               bool
    SupportsHostnameAnnotation bool
    Since                      string
}

func scanSourceInterfaces(sourceDir string) (string, error) {
    var sources []SourceInfo

    files, _ := filepath.Glob(filepath.Join(sourceDir, "*.go"))

    for _, file := range files {
        fset := token.NewFileSet()
        node, _ := parser.ParseFile(fset, file, nil, parser.ParseComments)

        info := SourceInfo{
            File: filepath.Base(file),
            Name: extractSourceName(file),
        }

        // Check for interface implementations
        ast.Inspect(node, func(n ast.Node) bool {
            if fn, ok := n.(*ast.FuncDecl); ok {
                if fn.Name.Name == "SupportsFQDNTemplate" {
                    info.SupportsFQDN = true
                    info.Since = extractSinceTag(fn.Doc)
                }
                if fn.Name.Name == "SupportsHostnameAnnotation" {
                    info.SupportsHostnameAnnotation = true
                }
            }
            return true
        })

        sources = append(sources, info)
    }

    return formatSourceTable(sources), nil
}

func extractSinceTag(doc *ast.CommentGroup) string {
    if doc == nil {
        return ""
    }

    for _, comment := range doc.List {
        if strings.Contains(comment.Text, "@since") {
            parts := strings.Fields(comment.Text)
            for i, part := range parts {
                if part == "@since" && i+1 < len(parts) {
                    return parts[i+1]
                }
            }
        }
    }
    return ""
}
```

**Function Reference from Godoc**

```go
// internal/gen/docs/fqdn/functions.go

type FunctionDoc struct {
    Name        string
    Signature   string
    Description string
    Example     string
    Since       string
}

func extractFunctionDocs(file string) (string, error) {
    fset := token.NewFileSet()
    node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
    if err != nil {
        return "", err
    }

    var functions []FunctionDoc

    ast.Inspect(node, func(n ast.Node) bool {
        if fn, ok := n.(*ast.FuncDecl); ok {
            // Extract from godoc comments
            if fn.Doc != nil {
                doc := FunctionDoc{
                    Name:        fn.Name.Name,
                    Signature:   formatSignature(fn.Type),
                    Description: extractDescription(fn.Doc),
                    Example:     extractExample(fn.Doc),
                    Since:       extractSinceTag(fn.Doc),
                }
                functions = append(functions, doc)
            }
        }
        return true
    })

    return formatFunctionReference(functions), nil
}
```

**Test Enforcement**

```go
// source/fqdn/doc_test.go

func TestDocsUpToDate(t *testing.T) {
    // Run doc generator
    cmd := exec.Command("go", "run", "internal/gen/docs/fqdn/main.go",
        "-output", "docs/advanced/fqdn")
    cmd.Dir = repoRoot()
    output, err := cmd.CombinedOutput()
    if err != nil {
        t.Fatalf("Failed to generate docs: %v\n%s", err, output)
    }

    // Check if any files changed
    cmd = exec.Command("git", "diff", "--exit-code", "docs/advanced/fqdn/")
    cmd.Dir = repoRoot()
    if err := cmd.Run(); err != nil {
        t.Fatalf(`
Documentation is out of date!

Please run from repo root:
    go run internal/gen/docs/fqdn/main.go

Or run:
    make generate-docs

Then commit the updated files in docs/advanced/fqdn/
`)
    }
}
```

## Enhancement 5: Advanced Template Functions

**Core Functions Module**

```go
// source/fqdn/functions.go

// customFuncs returns all template functions available to FQDN templates.
// This function map is registered globally and used by all templates.
func customFuncs() template.FuncMap {
    funcs := template.FuncMap{
        // String functions (existing)
        "contains":   strings.Contains,
        "trimPrefix": strings.TrimPrefix,
        "trimSuffix": strings.TrimSuffix,
        "trim":       strings.TrimSpace,
        "replace":    replace,

        // Sprig convention: support both names, deprecate old names later
        "toLower":    strings.ToLower, // @deprecated Use "lower" instead
        "lower":      strings.ToLower, // @since v0.21.0

        // IP validation (existing)
        "isIPv6":     isIPv6String,
        "isIPv4":     isIPv4String,
    }

    // Register additional function categories
    registerStringFuncs(funcs)
    registerConditionalFuncs(funcs)
    registerDNSFuncs(funcs)
    registerEncodingFuncs(funcs)

    return funcs
}

// ParseTemplate creates a template with all custom functions registered.
func ParseTemplate(input string) (*template.Template, error) {
    if input == "" {
        return nil, nil
    }
    return template.New("endpoint").Funcs(customFuncs()).Parse(input)
}
```

**String Functions**

```go
// source/fqdn/functions_string.go

// registerStringFuncs adds string manipulation functions.
func registerStringFuncs(funcs template.FuncMap) {
    funcs["truncate"] = truncate
    funcs["split"] = split
    funcs["join"] = join
}

// truncate limits a string to maxLen characters.
//
// Example:
//
//   {{truncate 10 .Name}}
//   Input: "very-long-service-name"
//   Output: "very-long-"
//
// @since v0.21.0
func truncate(maxLen int, s string) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen]
}

// split divides a string by separator.
//
// Example:
//
//   {{index (split "-" .Name) 0}}
//   Input: "api-service-v2"
//   Output: "api"
//
// @since v0.21.0
func split(sep, s string) []string {
    return strings.Split(s, sep)
}

// join concatenates strings with separator.
//
// Example:
//
//   {{join "-" .Namespace .Name}}
//   Input: namespace="prod", name="api"
//   Output: "prod-api"
//
// @since v0.21.0
func join(sep string, parts ...string) string {
    return strings.Join(parts, sep)
}
```

**Conditional Functions**

```go
// source/fqdn/functions_conditional.go

// registerConditionalFuncs adds conditional logic functions.
func registerConditionalFuncs(funcs template.FuncMap) {
    funcs["default"] = defaultValue
    funcs["ternary"] = ternary
    funcs["coalesce"] = coalesce
}

// defaultValue returns fallback if value is empty.
//
// Example:
//
//   {{default .Labels.env "prod"}}.example.com
//   Input: Labels.env=""
//   Output: "prod.example.com"
//
// @since v0.21.0
func defaultValue(value, fallback string) string {
    if value == "" {
        return fallback
    }
    return value
}

// ternary returns trueVal if condition is true, else falseVal.
//
// Example:
//
//   {{ternary (eq .Labels.env "prod") "production" "staging"}}.example.com
//   Input: Labels.env="prod"
//   Output: "production.example.com"
//
// @since v0.21.0
func ternary(condition bool, trueVal, falseVal string) string {
    if condition {
        return trueVal
    }
    return falseVal
}

// coalesce returns the first non-empty string.
//
// Example:
//
//   {{coalesce .Labels.env .Namespace "default"}}.example.com
//   Input: Labels.env="", Namespace="prod"
//   Output: "prod.example.com"
//
// @since v0.21.0
func coalesce(values ...string) string {
    for _, v := range values {
        if v != "" {
            return v
        }
    }
    return ""
}
```

**DNS Functions**

```go
// source/fqdn/functions_dns.go

// registerDNSFuncs adds DNS-specific utility functions.
func registerDNSFuncs(funcs template.FuncMap) {
    funcs["reverseDNS"] = reverseDNS
    funcs["extractZone"] = extractZone
    funcs["ensureSuffix"] = ensureSuffix
}

// reverseDNS generates reverse DNS hostname for an IP address.
//
// Example:
//
//   {{reverseDNS .Status.PodIP}}.in-addr.arpa
//   Input: "192.0.2.1"
//   Output: "1.2.0.192.in-addr.arpa"
//
// @since v0.21.0
func reverseDNS(ip string) (string, error) {
    addr, err := netip.ParseAddr(ip)
    if err != nil {
        return "", fmt.Errorf("invalid IP address: %w", err)
    }

    if addr.Is4() {
        octets := strings.Split(addr.String(), ".")
        // Reverse octets
        for i, j := 0, len(octets)-1; i < j; i, j = i+1, j-1 {
            octets[i], octets[j] = octets[j], octets[i]
        }
        return strings.Join(octets, "."), nil
    }

    // IPv6 reverse DNS (ip6.arpa format)
    // Implementation details...
    return "", fmt.Errorf("IPv6 reverse DNS not yet implemented")
}

// extractZone extracts the DNS zone from a hostname.
//
// Example:
//
//   {{extractZone "api.prod.example.com" 2}}
//   Output: "example.com"
//
// @since v0.21.0
func extractZone(hostname string, levels int) string {
    parts := strings.Split(hostname, ".")
    if len(parts) <= levels {
        return hostname
    }
    return strings.Join(parts[len(parts)-levels:], ".")
}

// ensureSuffix ensures hostname ends with given suffix.
//
// Example:
//
//   {{ensureSuffix .Name ".example.com"}}
//   Input: "api" or "api.example.com"
//   Output: "api.example.com"
//
// @since v0.21.0
func ensureSuffix(hostname, suffix string) string {
    if strings.HasSuffix(hostname, suffix) {
        return hostname
    }
    return hostname + suffix
}
```

**Encoding Functions**

```go
// source/fqdn/functions_encoding.go

// registerEncodingFuncs adds encoding and hashing functions.
//
// Security Consideration: These functions should be used carefully.
// Encoding user-controlled input can introduce security issues if not
// properly validated. Always validate hostnames after template execution.
func registerEncodingFuncs(funcs template.FuncMap) {
    funcs["toBase64"] = toBase64
    funcs["fromBase64"] = fromBase64
    funcs["base32"] = base32Encode
    funcs["sha256"] = sha256Hash
}

// toBase64 encodes a string to base64.
//
// Example:
//
//   {{toBase64 .Name}}.example.com
//   Input: "api"
//   Output: "YXBp.example.com"
//
// Security: Validate resulting hostname for DNS compliance.
//
// @since v0.21.0
func toBase64(s string) string {
    return base64.StdEncoding.EncodeToString([]byte(s))
}

// fromBase64 decodes a base64 string.
//
// Example:
//
//   {{fromBase64 .Labels.encoded}}.example.com
//   Input: Labels.encoded="YXBp"
//   Output: "api.example.com"
//
// Security: Validate input is properly base64-encoded.
// Returns empty string on decode error.
//
// @since v0.21.0
func fromBase64(s string) string {
    decoded, err := base64.StdEncoding.DecodeString(s)
    if err != nil {
        return ""
    }
    return string(decoded)
}

// base32Encode encodes a string to base32 (DNS-safe encoding).
//
// Example:
//
//   {{base32 .Name}}.example.com
//   Input: "api"
//   Output: "mfqwc.example.com"
//
// @since v0.21.0
func base32Encode(s string) string {
    return strings.ToLower(base32.StdEncoding.EncodeToString([]byte(s)))
}

// sha256Hash generates SHA256 hash of input.
//
// Example:
//
//   {{truncate 8 (sha256 .Name)}}.example.com
//   Input: "my-service"
//   Output: "a3c5b2d1.example.com"
//
// Security: Use for generating consistent short identifiers.
// Not for cryptographic purposes.
//
// @since v0.21.0
func sha256Hash(s string) string {
    h := sha256.Sum256([]byte(s))
    return hex.EncodeToString(h[:])
}
```

**Benchmarking**

```go
// source/fqdn/functions_bench_test.go

func BenchmarkTruncate(b *testing.B) {
    input := "very-long-service-name-that-needs-truncation"
    for i := 0; i < b.NB; i++ {
        _ = truncate(20, input)
    }
}

func BenchmarkSHA256(b *testing.B) {
    input := "my-service-name"
    for i := 0; i < b.N; i++ {
        _ = sha256Hash(input)
    }
}

// ... benchmarks for all functions
```

## Enhancement 6: FQDN Template Execution in Informer SetTransform

**Annotation Storage**

```go
const (
    GeneratedFQDNPrefix = "external-dns.alpha.kubernetes.io/generated-fqdns-"
    // Max annotation size: 256 chars
    // Total annotations limit: 16KB per object
)

// Example annotations after transform:
// external-dns.alpha.kubernetes.io/generated-fqdns-0: "api.prod.example.com,api.prod.example.org"
// external-dns.alpha.kubernetes.io/generated-fqdns-1: "api-v2.prod.example.com"
```

**Annotation Chunking**

```go
func storeGeneratedFQDNs(annotations map[string]string, fqdns []string) {
    const maxChunkSize = 200 // Leave room for overhead

    // Clear existing generated FQDN annotations
    for key := range annotations {
        if strings.HasPrefix(key, GeneratedFQDNPrefix) {
            delete(annotations, key)
        }
    }

    // Chunk and store
    joined := strings.Join(fqdns, ",")
    for i := 0; i < len(joined); i += maxChunkSize {
        end := i + maxChunkSize
        if end > len(joined) {
            end = len(joined)
        }
        key := fmt.Sprintf("%s%d", GeneratedFQDNPrefix, i/maxChunkSize)
        annotations[key] = joined[i:end]
    }
}
```

**Transform Implementation**

```go
// source/service.go

func NewServiceSource(
    ctx context.Context,
    kubeClient kubernetes.Interface,
    config *Config,
) (Source, error) {
    // ... informer setup ...

    serviceInformer := informerFactory.Core().V1().Services()

    // Get template from registry (Enhancement 1)
    var tmpl fqdn.Template
    if config.FQDNTemplateName != "" {
        var err error
        tmpl, err = fqdn.Get(config.FQDNTemplateName)
        if err != nil {
            return nil, fmt.Errorf("failed to get template %q: %w", config.FQDNTemplateName, err)
        }
    }

    // Apply transform if template is configured
    if tmpl != nil {
        _ = serviceInformer.Informer().SetTransform(func(i any) (any, error) {
            svc, ok := i.(*v1.Service)
            if !ok {
                return nil, fmt.Errorf("object is not a service")
            }

            // Check if already transformed (idempotent check)
            if _, exists := svc.Annotations[GeneratedFQDNPrefix+"0"]; exists {
                return svc, nil
            }

            // Execute template - context captured from outer scope
            hostnames, err := tmpl.Execute(ctx, svc)
            if err != nil {
                // Log error but don't skip record addition
                log.Errorf("Failed to execute FQDN template for service %s/%s: %v",
                    svc.Namespace, svc.Name, err)
                return svc, nil
            }

            // Create minimal service with generated FQDNs in annotations
            transformed := &v1.Service{
                ObjectMeta: metav1.ObjectMeta{
                    Name:      svc.Name,
                    Namespace: svc.Namespace,
                    // Copy existing annotations
                    Annotations: make(map[string]string),
                },
                Spec: v1.ServiceSpec{
                    Type: svc.Spec.Type,
                    // Only keep fields needed for endpoint generation
                },
                Status: svc.Status, // For LoadBalancer IPs
            }

            // Copy original annotations
            for k, v := range svc.Annotations {
                transformed.Annotations[k] = v
            }

            // Store generated FQDNs in chunked annotations
            storeGeneratedFQDNs(transformed.Annotations, hostnames)

            return transformed, nil
        })
    } else {
        // No template - apply memory optimization transform
        _ = serviceInformer.Informer().SetTransform(func(i any) (any, error) {
            svc, ok := i.(*v1.Service)
            if !ok {
                return nil, fmt.Errorf("object is not a service")
            }

            // Similar to pod.go pattern - keep minimal fields
            return &v1.Service{
                ObjectMeta: metav1.ObjectMeta{
                    Name:        svc.Name,
                    Namespace:   svc.Namespace,
                    Annotations: svc.Annotations,
                },
                Spec: v1.ServiceSpec{
                    Type: svc.Spec.Type,
                },
                Status: svc.Status,
            }, nil
        })
    }

    // ... rest of source setup ...
}
```

**Reading FQDNs in Endpoints()**

```go
func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
    // Read generated FQDNs from annotations
    hostnames := readGeneratedFQDNs(svc.Annotations)

    if len(hostnames) == 0 {
        log.Debugf("No generated FQDNs found for service %s/%s", svc.Namespace, svc.Name)
        return nil, nil
    }

    resource := fmt.Sprintf("service/%s/%s", svc.Namespace, svc.Name)
    ttl := annotations.TTLFromAnnotations(svc.Annotations, resource)
    targets := annotations.TargetsFromTargetAnnotation(svc.Annotations)

    if len(targets) == 0 {
        targets = extractLoadBalancerTargets(svc)
    }

    providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(svc.Annotations)

    var endpoints []*endpoint.Endpoint
    for _, hostname := range hostnames {
        endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
    }

    return endpoints, nil
}

func readGeneratedFQDNs(annotations map[string]string) []string {
    var chunks []string

    // Read all FQDN annotation chunks in order
    for i := 0; ; i++ {
        key := fmt.Sprintf("%s%d", GeneratedFQDNPrefix, i)
        chunk, exists := annotations[key]
        if !exists {
            break
        }
        chunks = append(chunks, chunk)
    }

    if len(chunks) == 0 {
        return nil
    }

    // Join chunks and split by comma
    joined := strings.Join(chunks, "")
    return strings.Split(joined, ",")
}
```

**Helper Functions**

```go
// source/fqdn/transform.go

const (
    GeneratedFQDNPrefix = "external-dns.alpha.kubernetes.io/generated-fqdns-"
    MaxAnnotationChunk  = 200
)

// StoreGeneratedFQDNs saves hostnames in chunked annotations.
func StoreGeneratedFQDNs(annotations map[string]string, fqdns []string) {
    // Clear existing
    for key := range annotations {
        if strings.HasPrefix(key, GeneratedFQDNPrefix) {
            delete(annotations, key)
        }
    }

    if len(fqdns) == 0 {
        return
    }

    // Chunk and store
    joined := strings.Join(fqdns, ",")
    for i := 0; i < len(joined); i += MaxAnnotationChunk {
        end := i + MaxAnnotationChunk
        if end > len(joined) {
            end = len(joined)
        }
        key := fmt.Sprintf("%s%d", GeneratedFQDNPrefix, i/MaxAnnotationChunk)
        annotations[key] = joined[i:end]
    }
}

// ReadGeneratedFQDNs retrieves hostnames from chunked annotations.
func ReadGeneratedFQDNs(annotations map[string]string) []string {
    var chunks []string
    for i := 0; ; i++ {
        key := fmt.Sprintf("%s%d", GeneratedFQDNPrefix, i)
        chunk, exists := annotations[key]
        if !exists {
            break
        }
        chunks = append(chunks, chunk)
    }

    if len(chunks) == 0 {
        return nil
    }

    joined := strings.Join(chunks, "")
    return strings.Split(joined, ",")
}

// IsAlreadyTransformed checks if object has generated FQDNs.
func IsAlreadyTransformed(annotations map[string]string) bool {
    _, exists := annotations[GeneratedFQDNPrefix+"0"]
    return exists
}
```

**Performance Testing**

```go
// source/fqdn/transform_bench_test.go

func BenchmarkTemplateExecutionCurrent(b *testing.B) {
    // Simulate current approach: execute template on every Endpoints() call
    tmpl := setupTemplate("{{.Name}}.{{.Namespace}}.example.com")
    svc := createTestService("api", "prod")

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = tmpl.Execute(context.Background(), svc)
    }
}

func BenchmarkTemplateExecutionTransform(b *testing.B) {
    // Simulate new approach: read from annotations
    svc := createTestService("api", "prod")
    StoreGeneratedFQDNs(svc.Annotations, []string{"api.prod.example.com"})

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ReadGeneratedFQDNs(svc.Annotations)
    }
}

func BenchmarkLargeCluster(b *testing.B) {
    // Test with large number of services
    const numServices = 10000

    services := make([]*v1.Service, numServices)
    for i := 0; i < numServices; i++ {
        services[i] = createTestService(fmt.Sprintf("svc-%d", i), "default")
    }

    tmpl := setupTemplate("{{.Name}}.{{.Namespace}}.example.com")

    b.Run("Current", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            for _, svc := range services {
                _, _ = tmpl.Execute(context.Background(), svc)
            }
        }
    })

    b.Run("Transform", func(b *testing.B) {
        // Pre-compute FQDNs in annotations
        for _, svc := range services {
            hostnames, _ := tmpl.Execute(context.Background(), svc)
            StoreGeneratedFQDNs(svc.Annotations, hostnames)
        }

        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            for _, svc := range services {
                _ = ReadGeneratedFQDNs(svc.Annotations)
            }
        }
    })
}

func BenchmarkComplexTemplate(b *testing.B) {
    // Test with complex template using multiple functions
    tmpl := setupTemplate("{{truncate 8 (sha256 .Name)}}.{{default .Labels.env \"prod\"}}.example.com")
    svc := createTestService("my-very-long-service-name", "default")
    svc.Labels = map[string]string{"env": "production"}

    b.Run("Current", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _, _ = tmpl.Execute(context.Background(), svc)
        }
    })

    b.Run("Transform", func(b *testing.B) {
        hostnames, _ := tmpl.Execute(context.Background(), svc)
        StoreGeneratedFQDNs(svc.Annotations, hostnames)

        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            _ = ReadGeneratedFQDNs(svc.Annotations)
        }
    })
}
```

## Enhancement 7: Target FQDN Template Support

**Annotation Constants**

```go
// source/annotations/annotations.go

const (
    // Existing
    TargetKey = AnnotationKeyPrefix + "target"  // Static target

    // New
    TargetFQDNKey = AnnotationKeyPrefix + "target-fqdn"  // Templated target
)
```

**Config Structure**

```go
// source/store.go

type Config struct {
    // Existing fields
    FQDNTemplate       string
    FQDNTemplateName   string

    // New: Target template support
    FQDNTargetTemplate string  // Default target template (global)
}

// source/service.go (example)

type serviceSource struct {
    // Existing fields
    fqdnTemplate fqdn.Template

    // New: Target template
    fqdnTargetTemplate fqdn.Template  // Per-source target template
}

func NewServiceSource(ctx context.Context, kubeClient kubernetes.Interface, config *Config) (Source, error) {
    // ... existing setup ...

    var targetTmpl fqdn.Template
    if config.FQDNTargetTemplate != "" {
        var err error
        targetTmpl, err = fqdn.ParseTemplate(config.FQDNTargetTemplate)
        if err != nil {
            return nil, fmt.Errorf("invalid fqdn-target-template: %w", err)
        }
    }

    return &serviceSource{
        // ... existing fields ...
        fqdnTargetTemplate: targetTmpl,
    }, nil
}
```

**Target Resolution**

```go
// source/service.go

func (sc *serviceSource) endpointsFromService(svc *v1.Service) ([]*endpoint.Endpoint, error) {
    // ... hostname generation ...

    // Resolve targets with template support
    targets, err := sc.resolveTargets(svc)
    if err != nil {
        return nil, err
    }

    // ... create endpoints ...
}

func (sc *serviceSource) resolveTargets(svc *v1.Service) (endpoint.Targets, error) {
    annotations := svc.Annotations

    // Priority 1: Check for target-fqdn annotation (templated)
    if targetFQDNTemplate, exists := annotations[annotations.TargetFQDNKey]; exists {
        // Check mutual exclusivity
        if _, hasStatic := annotations[annotations.TargetKey]; hasStatic {
            return nil, fmt.Errorf("service %s/%s has both %q and %q annotations (mutually exclusive)",
                svc.Namespace, svc.Name, annotations.TargetKey, annotations.TargetFQDNKey)
        }

        // Parse and execute annotation template
        tmpl, err := fqdn.ParseTemplate(targetFQDNTemplate)
        if err != nil {
            return nil, fmt.Errorf("invalid target-fqdn annotation: %w", err)
        }

        targets, err := tmpl.Execute(context.Background(), svc)
        if err != nil {
            return nil, fmt.Errorf("failed to execute target-fqdn template: %w", err)
        }

        return endpoint.Targets(targets), nil
    }

    // Priority 2: Check for static target annotation
    if staticTargets := annotations.TargetsFromTargetAnnotation(annotations); len(staticTargets) > 0 {
        return staticTargets, nil
    }

    // Priority 3: Use source-level target template (from config)
    if sc.fqdnTargetTemplate != nil {
        targets, err := sc.fqdnTargetTemplate.Execute(context.Background(), svc)
        if err != nil {
            return nil, fmt.Errorf("failed to execute source target template: %w", err)
        }
        return endpoint.Targets(targets), nil
    }

    // Priority 4: Auto-detect from LoadBalancer status
    return extractLoadBalancerTargets(svc), nil
}
```

**Mutual Exclusivity with Static Target**

```go
// source/annotations/processors.go

func TargetsFromAnnotations(annotations map[string]string, obj kubeObject, targetFQDNTemplate fqdn.Template) (endpoint.Targets, error) {
    // Check for static target annotation
    staticTarget := annotations[TargetKey]
    templatedTarget := annotations[TargetFQDNKey]

    // Validate mutual exclusivity
    if staticTarget != "" && templatedTarget != "" {
        return nil, fmt.Errorf("annotations %q and %q are mutually exclusive", TargetKey, TargetFQDNKey)
    }

    // Use static target
    if staticTarget != "" {
        return TargetsFromTargetAnnotation(annotations), nil
    }

    // Use templated target
    if templatedTarget != "" {
        tmpl, err := fqdn.ParseTemplate(templatedTarget)
        if err != nil {
            return nil, fmt.Errorf("invalid target-fqdn template: %w", err)
        }
        targets, err := tmpl.Execute(context.Background(), obj)
        if err != nil {
            return nil, fmt.Errorf("failed to execute target-fqdn template: %w", err)
        }
        return endpoint.Targets(targets), nil
    }

    // Fallback to auto-detected targets (LoadBalancer IP, Ingress status, etc.)
    return nil, nil
}
```

**Validation Functions**

```go
func validateTargetAnnotations(annotations map[string]string) error {
    hasStatic := annotations[annotations.TargetKey] != ""
    hasTemplated := annotations[annotations.TargetFQDNKey] != ""

    if hasStatic && hasTemplated {
        return fmt.Errorf("cannot specify both %q and %q annotations",
            annotations.TargetKey, annotations.TargetFQDNKey)
    }
    return nil
}

func validateTargetFQDNTemplate(templateStr string) error {
    _, err := fqdn.ParseTemplate(templateStr)
    if err != nil {
        return fmt.Errorf("invalid target-fqdn template syntax: %w", err)
    }
    return nil
}
```
