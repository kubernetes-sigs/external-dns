# FQDN Package Enhancement Proposal

**Date**: 2025-12-30
**Status**: PROPOSAL
**Package**: `sigs.k8s.io/external-dns/source/fqdn`

---

## Table of Contents

1. [Executive Summary](#executive-summary)
2. [Current Limitations](#current-limitations)
3. [Proposed Enhancements](#proposed-enhancements)
4. [Design Proposal](#design-proposal)
5. [Implementation Plan](#implementation-plan)
6. [Migration Strategy](#migration-strategy)
7. [Examples](#examples)

---

## Executive Summary

This proposal enhances the `fqdn` package to address current limitations and add powerful new features:

### **Key Enhancements**

1. ✅ **Template Interface** - Parse once, reuse everywhere (eliminates per-source parsing)
2. ✅ **Multi-Template Support** - Conditional templates for different sources/scenarios
3. ✅ **DNS Validation** - Prevent invalid hostnames in production
4. ✅ **Auto-Generated Documentation** - Extract examples from tests, document source support
5. ✅ **Advanced Template Functions** - More powerful templating capabilities
6. ✅ **Template Testing Framework** - Validate templates before deployment

---

## Current Limitations

### 1. **Template Parsed Per Source Constructor** 🔴

**Problem**: Every source parses the same template string repeatedly.

```go
// service.go:104-107
tmpl, err := fqdn.ParseTemplate(fqdnTemplate)

// ingress.go:76-79
tmpl, err := fqdn.ParseTemplate(fqdnTemplate)

// gateway.go:126-129
tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
```

**Impact**:

- Wasteful parsing (same template parsed 15+ times)
- Error handling in every constructor
- Can't share parsed templates
- No centralized validation

---

### 2. **Single Template String** 🔴

**Problem**: One template string for all sources; hard to write conditionals.

```go
// Current: One size fits all
--fqdn-template="{{.Name}}.{{.Namespace}}.example.com"

// What if you need:
// - Services: {{.Name}}.{{.Namespace}}.svc.example.com
// - Ingresses: {{.Name}}.{{.Namespace}}.ingress.example.com
// - Gateways: {{.Name}}.gateway.example.com
```

**Current Workaround** (ugly):

```go
{{if eq .Kind "Service"}}{{.Name}}.svc.example.com{{else if eq .Kind "Ingress"}}{{.Name}}.ingress.example.com{{else}}{{.Name}}.example.com{{end}}
```

**Issues**:

- Template becomes unreadable
- Error-prone
- Can't validate per-source
- No type safety

---

### 3. **No Template Documentation** 🟡

**Problem**: No way to know which sources support FQDN templates.

**Current State**:

- Manual documentation only
- No automated validation
- Examples scattered in tests
- Source support not discoverable

---

### 4. **No DNS Validation** 🟡

**Problem**: Templates can generate invalid DNS names.

```go
// These all "work" but produce invalid DNS:
{{.Name}}                           // "My_Service" (underscores invalid)
{{.Namespace | toUpper}}            // "PROD" (should be lowercase)
{{.Name}}.{{.Name}}.{{.Name}}...    // 300+ characters (exceeds DNS limit)
```

---

## Proposed Enhancements

### Enhancement 1: Template Interface & Registry 🔴 **HIGH PRIORITY**

#### **Design**

```go
// Package: source/fqdn

// Template represents a parsed, reusable FQDN template.
type Template interface {
    // Execute generates hostnames from a Kubernetes object
    Execute(ctx context.Context, obj kubeObject) ([]string, error)

    // Validate checks if template would produce valid DNS names
    Validate(obj kubeObject) error

    // String returns the original template string
    String() string

    // Functions returns available template functions
    Functions() []string
}

// TemplateRegistry manages parsed templates (singleton pattern).
type TemplateRegistry struct {
    templates map[string]Template
    mu        sync.RWMutex
}

// Global registry
var globalRegistry = NewTemplateRegistry()

// Register a template globally (called once at startup)
func Register(name string, templateStr string, opts ...Option) error {
    return globalRegistry.Register(name, templateStr, opts...)
}

// Get a registered template
func Get(name string) (Template, error) {
    return globalRegistry.Get(name)
}

// ParseTemplate creates a template (backward compatible)
func ParseTemplate(input string) (*template.Template, error) {
    // Keep existing function for compatibility
}
```

#### **Usage in Config**

```go
// store.go - Parse once at startup
func NewSourceConfig(cfg *externaldns.Config) (*Config, error) {
    // Register the default template
    if cfg.FQDNTemplate != "" {
        if err := fqdn.Register("default", cfg.FQDNTemplate); err != nil {
            return nil, fmt.Errorf("invalid FQDN template: %w", err)
        }
    }

    return &Config{
        FQDNTemplateName: "default", // Store name, not parsed template
        // ...
    }, nil
}
```

#### **Usage in Sources**

```go
// service.go - No parsing needed
func NewServiceSource(ctx context.Context, config *Config, ...) (Source, error) {
    // Get pre-parsed template from registry
    tmpl, err := fqdn.Get(config.FQDNTemplateName)
    if err != nil {
        return nil, err
    }

    return &serviceSource{
        fqdnTemplate: tmpl, // Use interface
        // ...
    }, nil
}

// Execute template
func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
    hostnames, err := sc.fqdnTemplate.Execute(context.Background(), svc)
    if err != nil {
        return nil, err
    }
    // ...
}
```

#### **Benefits**

- ✅ Parse once at startup
- ✅ Centralized error handling
- ✅ Type-safe interface
- ✅ Easy to mock/test
- ✅ Can add features (validation, caching) without changing sources

---

### Enhancement 2: Multi-Template Support 🔴 **HIGH PRIORITY**

#### **Design**

```go
// TemplateSet manages multiple templates for different sources/conditions.
type TemplateSet struct {
    templates map[string]Template
    selector  TemplateSelector
}

// TemplateSelector chooses which template to use based on object properties.
type TemplateSelector interface {
    Select(obj kubeObject) (string, error)
}

// Register a template set
func RegisterSet(name string, config TemplateSetConfig) error

type TemplateSetConfig struct {
    // Map of condition -> template
    Templates map[string]string

    // Default template if no condition matches
    Default string

    // Selector strategy
    Strategy SelectorStrategy
}

type SelectorStrategy string

const (
    StrategyByKind       SelectorStrategy = "kind"        // Select by .Kind
    StrategyByLabel      SelectorStrategy = "label"       // Select by label value
    StrategyByAnnotation SelectorStrategy = "annotation"  // Select by annotation
    StrategyByNamespace  SelectorStrategy = "namespace"   // Select by namespace
    StrategyCustom       SelectorStrategy = "custom"      // Custom logic
)
```

#### **Configuration Format**

**YAML Configuration**:

```yaml
# external-dns-config.yaml
fqdnTemplates:
  # Simple single template (backward compatible)
  default: "{{.Name}}.{{.Namespace}}.example.com"

  # Multi-template set
  multi:
    strategy: kind
    templates:
      Service: "{{.Name}}.{{.Namespace}}.svc.example.com"
      Ingress: "{{.Name}}.{{.Namespace}}.ingress.example.com"
      Gateway: "{{.Name}}.gateway.example.com"
      HTTPRoute: "{{.Name}}.route.example.com"
    default: "{{.Name}}.{{.Namespace}}.example.com"

  # Label-based selection
  byEnvironment:
    strategy: label
    selector: "environment"
    templates:
      production: "{{.Name}}.prod.example.com"
      staging: "{{.Name}}.staging.example.com"
      development: "{{.Name}}.dev.example.com"
    default: "{{.Name}}.example.com"

  # Namespace-based selection
  byNamespace:
    strategy: namespace
    templates:
      kube-system: "{{.Name}}.system.example.internal"
      default: "{{.Name}}.{{.Namespace}}.apps.example.com"
      production: "{{.Name}}.prod.example.com"
    default: "{{.Name}}.example.com"
```

**CLI Flags** (backward compatible):

```bash
# Simple (existing)
--fqdn-template="{{.Name}}.{{.Namespace}}.example.com"

# Multi-template (new)
--fqdn-template-set=kind \
--fqdn-template-service="{{.Name}}.svc.example.com" \
--fqdn-template-ingress="{{.Name}}.ingress.example.com" \
--fqdn-template-default="{{.Name}}.example.com"

# Or from file
--fqdn-template-config=/etc/external-dns/templates.yaml
```

#### **Usage Example**

```go
// Register multi-template set
err := fqdn.RegisterSet("multi", fqdn.TemplateSetConfig{
    Strategy: fqdn.StrategyByKind,
    Templates: map[string]string{
        "Service":   "{{.Name}}.{{.Namespace}}.svc.example.com",
        "Ingress":   "{{.Name}}.ingress.example.com",
        "Gateway":   "{{.Name}}.gateway.example.com",
    },
    Default: "{{.Name}}.example.com",
})

// Get template set
tmplSet, _ := fqdn.GetSet("multi")

// Execute - automatically selects correct template
service := &v1.Service{...}
hostnames, _ := tmplSet.Execute(context.Background(), service)
// Uses Service template: "my-service.default.svc.example.com"

ingress := &networkingv1.Ingress{...}
hostnames, _ = tmplSet.Execute(context.Background(), ingress)
// Uses Ingress template: "my-ingress.ingress.example.com"
```

#### **Implementation**

```go
// fqdn/template_set.go
type templateSet struct {
    templates map[string]Template
    selector  TemplateSelector
    default_  Template
}

func (ts *templateSet) Execute(ctx context.Context, obj kubeObject) ([]string, error) {
    // Select template based on object
    templateName, err := ts.selector.Select(obj)
    if err != nil {
        return nil, err
    }

    // Get template
    tmpl, ok := ts.templates[templateName]
    if !ok {
        tmpl = ts.default_
    }

    // Execute selected template
    return tmpl.Execute(ctx, obj)
}

// Built-in selectors
type kindSelector struct{}

func (s *kindSelector) Select(obj kubeObject) (string, error) {
    return obj.GetObjectKind().GroupVersionKind().Kind, nil
}

type labelSelector struct {
    labelKey string
}

func (s *labelSelector) Select(obj kubeObject) (string, error) {
    labels := obj.GetLabels()
    if value, ok := labels[s.labelKey]; ok {
        return value, nil
    }
    return "", fmt.Errorf("label %q not found", s.labelKey)
}

type namespaceSelector struct{}

func (s *namespaceSelector) Select(obj kubeObject) (string, error) {
    return obj.GetNamespace(), nil
}
```

#### **Benefits**

- ✅ Different templates per source type
- ✅ Conditional logic extracted from templates
- ✅ More readable configuration
- ✅ Type-safe selection
- ✅ Easier to test
- ✅ Backward compatible

---

### Enhancement 3: DNS Validation 🟡 **MEDIUM PRIORITY**

#### **Design**

```go
// Validator checks if hostnames are DNS-compliant.
type Validator interface {
    Validate(hostname string) error
}

// ValidationOptions configures hostname validation and sanitization.
type ValidationOptions struct {
    Enabled       bool
    AutoFix       bool   // Automatically fix invalid hostnames
    MaxLength     int    // Default: 253 (DNS standard)
    AllowUnicode  bool   // Support IDN (Internationalized Domain Names)
    StrictRFC1123 bool   // Enforce strict RFC 1123 compliance
}

// Built-in validators
var (
    RFC1123Validator = &rfc1123Validator{}
    LengthValidator  = &lengthValidator{maxLen: 253}
    IDNValidator     = &idnValidator{}
)

// Sanitizer fixes invalid hostnames
type Sanitizer interface {
    Sanitize(hostname string) (string, error)
}

var DefaultSanitizer = &defaultSanitizer{
    toLowercase:     true,
    replaceInvalid:  true,
    replacementChar: '-',
    truncate:        true,
    maxLength:       253,
}
```

#### **Usage**

```go
// Create template with validation
tmpl, err := fqdn.New("{{.Name}}.{{.Namespace}}.example.com",
    fqdn.WithValidation(true),
    fqdn.WithAutoFix(true),
)

// Execute with validation
hostnames, err := tmpl.Execute(ctx, service)
// Returns error if hostname invalid (when AutoFix=false)
// Or returns fixed hostname (when AutoFix=true)
```

#### **Example: Auto-Fix**

```go
// Input service name: "My_Service"
// Template: "{{.Name}}.example.com"

// Without validation:
// Output: "My_Service.example.com" ❌ (invalid: underscores)

// With validation + auto-fix:
// Output: "my-service.example.com" ✅ (fixed: lowercase, replaced underscore)
```

#### **Implementation**

```go
// fqdn/validator.go
type rfc1123Validator struct{}

func (v *rfc1123Validator) Validate(hostname string) error {
    if len(hostname) > 253 {
        return fmt.Errorf("hostname too long: %d chars (max 253)", len(hostname))
    }

    labels := strings.Split(hostname, ".")
    for _, label := range labels {
        if len(label) == 0 || len(label) > 63 {
            return fmt.Errorf("invalid label length: %q", label)
        }
        if !isValidDNSLabel(label) {
            return fmt.Errorf("invalid DNS label: %q", label)
        }
    }
    return nil
}

var dnsLabelRegex = regexp.MustCompile(`^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`)

func isValidDNSLabel(label string) bool {
    return dnsLabelRegex.MatchString(label)
}

// fqdn/sanitizer.go
type defaultSanitizer struct {
    toLowercase     bool
    replaceInvalid  bool
    replacementChar rune
    truncate        bool
    maxLength       int
}

func (s *defaultSanitizer) Sanitize(hostname string) (string, error) {
    if s.toLowercase {
        hostname = strings.ToLower(hostname)
    }

    if s.replaceInvalid {
        hostname = s.replaceInvalidChars(hostname)
    }

    if s.truncate && len(hostname) > s.maxLength {
        hostname = hostname[:s.maxLength]
        // Ensure doesn't end with dot or dash
        hostname = strings.TrimRight(hostname, ".-")
    }

    return hostname, nil
}

func (s *defaultSanitizer) replaceInvalidChars(hostname string) string {
    var result strings.Builder
    for _, r := range hostname {
        if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '.' || r == '-' {
            result.WriteRune(r)
        } else {
            result.WriteRune(s.replacementChar)
        }
    }
    return result.String()
}
```

---

### Enhancement 4: Auto-Generated Documentation 🟡 **MEDIUM PRIORITY**

#### **4.1 Generate Docs from Tests**

```go
// fqdn/doc_generator.go

// GenerateDocsFromTests extracts examples from test files.
func GenerateDocsFromTests(testFile string) (string, error)

// Example output:
/*
# FQDN Template Examples

## Simple Template
```

Template: {{ .Name }}.example.com
Input:    Service "test" in namespace "default"
Output:   ["test.example.com"]

```

## Multiple Hostnames
```

Template: {{.Name}}.example.com, {{.Name}}.example.org
Input:    Service "test"
Output:   ["test.example.com", "test.example.org"]

```

## Using Labels
```

Template: {{.Labels.environment}}.example.com
Input:    Service with label "environment=production"
Output:   ["production.example.com"]

```
*/
```

**Implementation**:

```go
// Parse test file and extract TestExecTemplate cases
func GenerateDocsFromTests(testFile string) (string, error) {
    fset := token.NewFileSet()
    node, err := parser.ParseFile(fset, testFile, nil, parser.ParseComments)
    if err != nil {
        return "", err
    }

    var examples []Example

    // Walk AST and find test cases
    ast.Inspect(node, func(n ast.Node) bool {
        // Find: tests := []struct { name, tmpl, obj, want }
        // Extract and format as examples
        return true
    })

    return formatExamples(examples), nil
}
```

#### **4.2 Document Source Support**

```go
// Generate table showing which sources support FQDN templates

// source/doc_generator.go

type SourceInfo struct {
    Name              string
    File              string
    SupportsFQDN      bool
    SupportsAnnotation bool
    TemplateFieldName string
}

// ScanSources analyzes source files and generates documentation.
func ScanSources(sourceDir string) ([]SourceInfo, error)

// Example output:
/*
# FQDN Template Support by Source

| Source | FQDN Template | Hostname Annotation | Template Field |
|--------|---------------|---------------------|----------------|
| Service | ✅ | ✅ | fqdnTemplate |
| Ingress | ✅ | ✅ | fqdnTemplate |
| Gateway | ✅ | ✅ | fqdnTemplate |
| HTTPRoute | ✅ | ✅ | fqdnTemplate |
| Pod | ✅ | ❌ | fqdnTemplate |
| Node | ❌ | ❌ | - |
| CRD | ✅ | ✅ | fqdnTemplate |
*/
```

**Implementation**:

```go
func ScanSources(sourceDir string) ([]SourceInfo, error) {
    var sources []SourceInfo

    // Read all *_source.go files
    files, _ := filepath.Glob(filepath.Join(sourceDir, "*source*.go"))

    for _, file := range files {
        fset := token.NewFileSet()
        node, _ := parser.ParseFile(fset, file, nil, 0)

        info := SourceInfo{
            File: filepath.Base(file),
            Name: extractSourceName(file),
        }

        // Check for fqdnTemplate field
        ast.Inspect(node, func(n ast.Node) bool {
            if structType, ok := n.(*ast.StructType); ok {
                for _, field := range structType.Fields.List {
                    for _, name := range field.Names {
                        if name.Name == "fqdnTemplate" {
                            info.SupportsFQDN = true
                            info.TemplateFieldName = "fqdnTemplate"
                        }
                        if name.Name == "ignoreHostnameAnnotation" {
                            info.SupportsAnnotation = true
                        }
                    }
                }
            }
            return true
        })

        sources = append(sources, info)
    }

    return sources, nil
}
```

#### **4.3 CLI Tool for Documentation**

```bash
# Generate documentation
$ external-dns doc generate --output docs/

# Generated files:
# - docs/fqdn-templates.md          (examples from tests)
# - docs/source-support.md          (which sources support templates)
# - docs/template-functions.md     (available template functions)
```

**Implementation**:

```go
// cmd/external-dns/doc.go

var docCmd = &cobra.Command{
    Use:   "doc",
    Short: "Generate documentation",
}

var docGenerateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate documentation from source code",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Generate template examples
        examples, _ := fqdn.GenerateDocsFromTests("source/fqdn/fqdn_test.go")
        ioutil.WriteFile("docs/fqdn-templates.md", []byte(examples), 0644)

        // Generate source support table
        sources, _ := source.ScanSources("source/")
        table := formatSourceTable(sources)
        ioutil.WriteFile("docs/source-support.md", []byte(table), 0644)

        // Generate function reference
        functions := fqdn.AvailableFunctions()
        reference := formatFunctionReference(functions)
        ioutil.WriteFile("docs/template-functions.md", []byte(reference), 0644)

        return nil
    },
}
```

---

### Enhancement 5: Advanced Template Functions 🟢 **LOW PRIORITY**

```go
// Add more powerful template functions

var advancedFuncs = template.FuncMap{
    // String manipulation
    "truncate":     truncate,
    "sha256":       sha256Hash,
    "base32":       base32Encode,
    "regexReplace": regexReplace,
    "split":        strings.Split,
    "join":         strings.Join,

    // Conditional
    "default":      defaultValue,
    "ternary":      ternary,
    "coalesce":     coalesce,

    // DNS specific
    "reverseDNS":   reverseDNS,
    "extractZone":  extractZone,
    "ensureSuffix": ensureSuffix,

    // Encoding
    "toBase64":     toBase64,
    "fromBase64":   fromBase64,
}
```

**Examples**:

```go
// Truncate long names
{{truncate 20 .Name}}.example.com

// Use default if label missing
{{default .Labels.env "prod"}}.example.com

// Ternary conditional
{{ternary (eq .Labels.env "prod") "production" "non-prod"}}.example.com

// Generate short consistent name
{{substr 0 8 (sha256 .Name)}}.example.com

// Reverse DNS
{{reverseDNS .Status.PodIP}}.in-addr.arpa

// Ensure suffix
{{ensureSuffix .Name ".example.com"}}
```

---

### Enhancement 6: Template Testing Framework 🟢 **LOW PRIORITY**

```go
// Template testing utilities

// TestTemplate validates a template against test cases
func TestTemplate(t *testing.T, tmpl Template, cases []TestCase) {
    for _, tc := range cases {
        t.Run(tc.Name, func(t *testing.T) {
            result, err := tmpl.Execute(context.Background(), tc.Object)

            if tc.ShouldError {
                assert.Error(t, err)
                return
            }

            assert.NoError(t, err)
            assert.Equal(t, tc.Expected, result)
        })
    }
}

// Template validation before deployment
func ValidateTemplate(tmpl Template, sampleObjects []kubeObject) error {
    for _, obj := range sampleObjects {
        if err := tmpl.Validate(obj); err != nil {
            return fmt.Errorf("template validation failed for %s: %w",
                obj.GetName(), err)
        }
    }
    return nil
}
```

**Usage**:

```go
// In tests
func TestMyTemplate(t *testing.T) {
    tmpl, _ := fqdn.New("{{.Name}}.{{.Namespace}}.example.com")

    fqdn.TestTemplate(t, tmpl, []fqdn.TestCase{
        {
            Name: "simple service",
            Object: &v1.Service{
                ObjectMeta: metav1.ObjectMeta{
                    Name: "api",
                    Namespace: "prod",
                },
            },
            Expected: []string{"api.prod.example.com"},
        },
    })
}
```

---

## Design Proposal

### New Package Structure

```
source/fqdn/
├── fqdn.go                      # Backward compatible API
├── template.go                  # Template interface & implementation
├── template_set.go              # Multi-template support
├── registry.go                  # Template registry (singleton)
├── selector.go                  # Template selectors
├── validator.go                 # DNS validation
├── sanitizer.go                 # Hostname sanitization
├── functions.go                 # Custom template functions
├── functions_advanced.go        # Advanced functions
├── doc_generator.go             # Documentation generation
├── testing.go                   # Testing utilities
│
├── fqdn_test.go                # Existing tests
├── template_test.go            # Template tests
├── template_set_test.go        # Multi-template tests
├── validator_test.go           # Validation tests
└── examples_test.go            # Example tests (for godoc)
```

### Core Interfaces

```go
// template.go
package fqdn

import (
    "context"
    "text/template"
)

// Template represents a parsed, reusable FQDN template.
type Template interface {
    // Execute generates hostnames from a Kubernetes object
    Execute(ctx context.Context, obj kubeObject) ([]string, error)

    // Validate checks if template would produce valid DNS names
    Validate(obj kubeObject) error

    // String returns the original template string
    String() string
}

// Option configures template behavior
type Option func(*templateOptions)

type templateOptions struct {
    validate    bool
    autoFix     bool
    maxLength   int
    allowUnicode bool
}

// Template creation
func New(templateStr string, opts ...Option) (Template, error)
func NewSet(config TemplateSetConfig) (Template, error)

// Options
func WithValidation(enabled bool) Option
func WithAutoFix(enabled bool) Option
func WithMaxLength(n int) Option
func WithUnicode(enabled bool) Option

// Registry functions
func Register(name string, templateStr string, opts ...Option) error
func RegisterSet(name string, config TemplateSetConfig) error
func Get(name string) (Template, error)
func GetSet(name string) (Template, error)

// Backward compatibility
func ParseTemplate(input string) (*template.Template, error)
func ExecTemplate(tmpl *template.Template, obj kubeObject) ([]string, error)
```

---

## Implementation Plan

### Phase 1: Foundation (Week 1-2)

**Goals**: Core infrastructure without breaking changes

1. **Create Template Interface** ✅
   - Define `Template` interface
   - Implement `template` struct (wraps `text/template.Template`)
   - Add `New()` constructor
   - Keep backward compatibility

2. **Add Template Registry** ✅
   - Implement `TemplateRegistry`
   - Add `Register()` and `Get()` functions
   - Thread-safe implementation

3. **Update Config** ✅
   - Modify `store.go` to register templates at startup
   - Keep `Config` structure backward compatible
   - Add `FQDNTemplateName string` field

4. **Tests** ✅
   - Unit tests for new components
   - Integration tests
   - Backward compatibility tests

**Deliverables**:

- Working Template interface
- Template registry
- All existing tests pass

---

### Phase 2: Multi-Template Support (Week 3-4)

**Goals**: Enable different templates per source

1. **Implement TemplateSet** ✅
   - Create `TemplateSet` implementation
   - Implement selectors (kind, label, namespace)
   - Add `RegisterSet()` and `GetSet()`

2. **Configuration Format** ✅
   - Design YAML configuration format
   - Add CLI flags for multi-template
   - Add config file loading

3. **Update Sources** ✅
   - Modify sources to use `Template` interface
   - Remove per-source parsing
   - Use registry to get templates

4. **Tests** ✅
   - Multi-template test cases
   - Selector tests
   - Source integration tests

**Deliverables**:

- Working multi-template support
- YAML configuration
- Updated sources

---

### Phase 3: Validation & Sanitization (Week 5-6)

**Goals**: Prevent invalid DNS names

1. **Implement Validators** ✅
   - RFC 1123 validator
   - Length validator
   - IDN validator

2. **Implement Sanitizers** ✅
   - Default sanitizer (lowercase, replace invalid)
   - Truncate sanitizer
   - IDN converter

3. **Integration** ✅
   - Add validation to Template.Execute()
   - Add options for auto-fix
   - Add validation to template creation

4. **Tests** ✅
   - Validator test cases
   - Sanitizer test cases
   - Edge cases (unicode, long names, etc.)

**Deliverables**:

- DNS validation
- Auto-fix functionality
- Comprehensive tests

---

### Phase 4: Documentation Generation (Week 7-8)

**Goals**: Auto-generate docs from code

1. **Test Parser** ✅
   - Parse test files using `go/ast`
   - Extract test cases
   - Format as examples

2. **Source Scanner** ✅
   - Scan source files
   - Identify FQDN support
   - Generate support matrix

3. **CLI Tool** ✅
   - Add `external-dns doc generate` command
   - Generate multiple doc files
   - Integrate into CI/CD

4. **Documentation** ✅
   - Write user guide
   - Add migration guide
   - Create examples

**Deliverables**:

- Documentation generator
- CLI tool
- Auto-generated docs

---

### Phase 5: Advanced Features (Week 9-10)

**Goals**: Power user features

1. **Advanced Template Functions** ✅
   - Implement additional functions
   - Add function registry
   - Document all functions

2. **Testing Framework** ✅
   - Template testing utilities
   - Validation helpers
   - Example generators

3. **Performance Optimization** ✅
   - Benchmark existing code
   - Optimize hot paths
   - Add caching if needed

**Deliverables**:

- Advanced template functions
- Testing framework
- Performance benchmarks

---

## Migration Strategy

### Backward Compatibility

**Guarantee**: All existing code continues to work without changes.

```go
// OLD CODE (still works):
tmpl, err := fqdn.ParseTemplate("{{.Name}}.example.com")
if err != nil {
    return err
}
hostnames, err := fqdn.ExecTemplate(tmpl, service)

// NEW CODE (recommended):
tmpl, err := fqdn.New("{{.Name}}.example.com")
if err != nil {
    return err
}
hostnames, err := tmpl.Execute(context.Background(), service)
```

### Migration Path

#### Step 1: Update Config (No Code Changes)

```go
// Before: Template parsed in every source
func NewServiceSource(..., fqdnTemplate string, ...) {
    tmpl, _ := fqdn.ParseTemplate(fqdnTemplate)
    // ...
}

// After: Template registered once at startup
func NewSourceConfig(cfg *externaldns.Config) (*Config, error) {
    if cfg.FQDNTemplate != "" {
        fqdn.Register("default", cfg.FQDNTemplate)
    }
    return &Config{
        FQDNTemplateName: "default",
    }, nil
}

func NewServiceSource(..., config *Config, ...) {
    tmpl, _ := fqdn.Get(config.FQDNTemplateName)
    // ...
}
```

#### Step 2: Add Multi-Template (Opt-In)

```yaml
# Use single template (existing behavior)
fqdn-template: "{{.Name}}.example.com"

# Or use multi-template (new feature)
fqdn-templates:
  strategy: kind
  templates:
    Service: "{{.Name}}.svc.example.com"
    Ingress: "{{.Name}}.ingress.example.com"
  default: "{{.Name}}.example.com"
```

#### Step 3: Enable Validation (Opt-In)

```yaml
fqdn-template: "{{.Name}}.example.com"
fqdn-validation:
  enabled: true
  auto-fix: true
```

### Deprecation Timeline

**v1.15.0** (Current):

- ✅ Add new Template interface
- ✅ Keep old API working
- 📢 Announce new features

**v1.16.0** (3 months):

- ✅ All features available
- 📢 Deprecate `ParseTemplate()` / `ExecTemplate()` (add deprecation comments)
- 📚 Update documentation

**v1.17.0** (6 months):

- ⚠️ Log warnings when using old API
- 📢 Final notice before removal

**v2.0.0** (12 months):

- ❌ Remove old `ParseTemplate()` / `ExecTemplate()` functions
- ✅ Template interface is only API

---

## Examples

### Example 1: Simple Migration

**Before**:

```go
// service.go (before)
type serviceSource struct {
    fqdnTemplate *template.Template
}

func NewServiceSource(ctx, client, ..., fqdnTemplate string, ...) (Source, error) {
    tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
    if err != nil {
        return nil, err
    }
    return &serviceSource{
        fqdnTemplate: tmpl,
    }, nil
}

func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
    hostnames, err := fqdn.ExecTemplate(sc.fqdnTemplate, svc)
    // ...
}
```

**After**:

```go
// service.go (after)
type serviceSource struct {
    fqdnTemplate fqdn.Template  // Interface instead of *template.Template
}

func NewServiceSource(ctx, client, config *Config, ...) (Source, error) {
    tmpl, err := fqdn.Get(config.FQDNTemplateName)
    if err != nil {
        return nil, err
    }
    return &serviceSource{
        fqdnTemplate: tmpl,
    }, nil
}

func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
    hostnames, err := sc.fqdnTemplate.Execute(context.Background(), svc)
    // ...
}

// store.go - Register template once
func NewSourceConfig(cfg *externaldns.Config) (*Config, error) {
    if cfg.FQDNTemplate != "" {
        if err := fqdn.Register("default", cfg.FQDNTemplate); err != nil {
            return nil, err
        }
    }
    return &Config{
        FQDNTemplateName: "default",
    }, nil
}
```

---

### Example 2: Multi-Template Configuration

**Configuration**:

```yaml
# config/templates.yaml
fqdnTemplates:
  # Different templates per source type
  byKind:
    strategy: kind
    templates:
      Service: "{{.Name}}.{{.Namespace}}.svc.cluster.local"
      Ingress: "{{.Name}}.{{.Namespace}}.ingress.example.com"
      Gateway: "{{.Name}}.gw.example.com"
      HTTPRoute: "{{.Name}}.route.example.com"
    default: "{{.Name}}.{{.Namespace}}.example.com"

  # Different templates per environment
  byEnvironment:
    strategy: label
    selector: "environment"
    templates:
      production: "{{.Name}}.prod.example.com"
      staging: "{{.Name}}.staging.example.com"
      development: "{{.Name}}.{{.Namespace}}.dev.example.com"
    default: "{{.Name}}.example.com"
```

**Usage**:

```go
// Load configuration
config, err := loadConfig("config/templates.yaml")

// Register template sets
for name, setConfig := range config.FQDNTemplates {
    fqdn.RegisterSet(name, setConfig)
}

// Use in sources
tmpl, _ := fqdn.GetSet("byKind")

// Automatically selects correct template based on Kind
service := &v1.Service{...}
hostnames, _ := tmpl.Execute(ctx, service)
// Output: "my-service.default.svc.cluster.local"

ingress := &networkingv1.Ingress{...}
hostnames, _ = tmpl.Execute(ctx, ingress)
// Output: "my-ingress.default.ingress.example.com"
```

---

### Example 3: Validation & Auto-Fix

**Configuration**:

```yaml
fqdn-template: "{{.Name}}.{{.Namespace}}.example.com"
fqdn-validation:
  enabled: true
  auto-fix: true
  max-length: 253
  strict-rfc1123: true
```

**Code**:

```go
// Register template with validation
fqdn.Register("default", "{{.Name}}.{{.Namespace}}.example.com",
    fqdn.WithValidation(true),
    fqdn.WithAutoFix(true),
    fqdn.WithMaxLength(253),
)

tmpl, _ := fqdn.Get("default")

// Invalid service name
service := &v1.Service{
    ObjectMeta: metav1.ObjectMeta{
        Name:      "My_Service",  // Invalid: uppercase + underscore
        Namespace: "PROD",        // Invalid: uppercase
    },
}

// Execute with auto-fix
hostnames, _ := tmpl.Execute(ctx, service)
// Output: ["my-service.prod.example.com"]  ✅ Fixed!

// Without auto-fix:
// Error: invalid hostname "My_Service.PROD.example.com": invalid characters
```

---

### Example 4: Auto-Generated Documentation

**Run Documentation Generator**:

```bash
$ external-dns doc generate --output docs/

Generating documentation...
✓ Extracted 25 examples from fqdn_test.go
✓ Scanned 21 source files
✓ Generated docs/fqdn-templates.md
✓ Generated docs/source-support.md
✓ Generated docs/template-functions.md
Done!
```

**Generated Output** (`docs/source-support.md`):

```markdown
# FQDN Template Support by Source

Auto-generated on 2025-12-30

| Source | FQDN Template | Hostname Annotation | Notes |
|--------|---------------|---------------------|-------|
| Service | ✅ | ✅ | Full support |
| Ingress | ✅ | ✅ | Full support |
| Gateway | ✅ | ✅ | Full support |
| HTTPRoute | ✅ | ✅ | Full support |
| TCPRoute | ✅ | ✅ | Full support |
| TLSRoute | ✅ | ✅ | Full support |
| UDPRoute | ✅ | ✅ | Full support |
| GRPCRoute | ✅ | ✅ | Full support |
| Pod | ✅ | ❌ | Template only |
| Node | ❌ | ❌ | Not supported |
| CRD | ✅ | ✅ | Full support |
| IstioGateway | ✅ | ✅ | Full support |
| IstioVirtualService | ✅ | ✅ | Full support |
| ContourHTTPProxy | ✅ | ✅ | Full support |

## Usage

Sources marked with ✅ in "FQDN Template" support the `--fqdn-template` flag.
Sources marked with ✅ in "Hostname Annotation" support hostname annotations.

For template syntax and examples, see [fqdn-templates.md](fqdn-templates.md).
```

---

## Summary

This proposal provides a comprehensive enhancement to the `fqdn` package addressing all current limitations:

### ✅ **Solved Problems**

1. **Template Interface** - Parse once, reuse everywhere
2. **Multi-Template Support** - Different templates per source/condition
3. **DNS Validation** - Prevent invalid hostnames
4. **Auto-Generated Docs** - Always up-to-date documentation
5. **Backward Compatibility** - Existing code continues to work

### 📊 **Impact**

- **Performance**: Parse templates once (15x reduction in parsing)
- **Maintainability**: Centralized template management
- **Usability**: Conditional templates without complex logic
- **Safety**: DNS validation prevents production issues
- **Documentation**: Always accurate, auto-generated

### 🎯 **Next Steps**

1. **Review** this proposal
2. **Approve** design decisions
3. **Begin** Phase 1 implementation
4. **Iterate** based on feedback

---

**Author**: External DNS Team
**Date**: 2025-12-30
**Status**: Awaiting Review
