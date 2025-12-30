# External DNS Documentation Audit Findings

**Date:** 2025-12-22
**Purpose:** Comprehensive analysis of the docs folder to identify areas needing improvement

---

## Executive Summary

The External DNS project has excellent coverage of basic usage and provider-specific tutorials (40+ providers), but lacks comprehensive troubleshooting, migration, security, and performance tuning guides. Several inconsistencies and outdated references need addressing.

---

## 1. Current Documentation Coverage

### ✅ Excellent Coverage Areas

- **Tutorials**: 40+ provider-specific tutorials (AWS, Azure, Google Cloud, Cloudflare, etc.)
- **Sources**: 20+ Kubernetes resource sources (Ingress, Service, Gateway API, Istio, Traefik, etc.)
- **Advanced Features**: Split horizon DNS, FQDN templating, TTL management, rate limiting, NAT64, import records
- **Registry Systems**: TXT, DynamoDB, and AWS Service Discovery registries
- **Annotations**: Detailed annotation reference with examples
- **Flags**: Auto-generated comprehensive flag documentation
- **Metrics**: Auto-generated metrics reference
- **Contributing**: Developer guides, design documentation, deprecation policy

---

## 2. Issues Found: Outdated, Unclear, or Confusing Content

### 🔴 High Priority

#### 2.1 FQDN Templating Contradiction

**File:** `docs/advanced/fqdn-templating.md`
**Issue:**

- Line 240: States `--fqdn-template` ignores `external-dns.alpha.kubernetes.io/hostname` annotations
- Earlier sections (Line 299): Discusses combining them with `--combine-fqdn-annotation`
- **Impact:** Users confused about actual behavior

#### 2.2 Gateway Documentation Confusion

**Files:** `docs/sources/gateway.md` vs `docs/sources/gateway-api.md`
**Issue:**

- Two separate files with overlapping/similar content
- Unclear which is canonical or if one is deprecated
- API version support inconsistencies (v1alpha2, v1beta1, v1)
- **Impact:** Users unsure which guide to follow

#### 2.3 Incomplete TODO Markers

**File:** `docs/advanced/fqdn-templating.md`
**Lines:** 34, 66
**Issue:** Published documentation contains unfinished TODO sections
**Impact:** Looks unprofessional, missing information

### 🟡 Medium Priority

#### 2.4 Outdated Traefik References

**File:** `docs/faq.md`
**Lines:** 63-65
**Issue:**

- References Traefik v1.7 (old version)
- Points to deprecated Helm chart at `helm/charts`
- Should reference Traefik v2.x+ with current configuration
- **Impact:** Users get wrong/outdated information

#### 2.5 Providers Table Lacks Context

**File:** `docs/providers.md`
**Issue:**

- Table shows provider capabilities (Zone Cache, Dry Run, Default TTL) without explanation
- No documentation of what "n/a" means in each column
- No links to detailed provider documentation
- No guidance on which providers are recommended/maintained
- **Impact:** Users can't make informed provider selection

#### 2.6 Vague Release Schedule

**File:** `docs/release.md`
**Lines:** 5-6
**Issue:**

- States "Currently we don't release regularly" - too vague
- References asking in Slack for release timing
- No clear versioning roadmap
- **Impact:** Users uncertain about stability and update frequency

#### 2.7 Auto-Generated Documentation Warning

**File:** `docs/flags.md`
**Lines:** 3-5
**Issue:**

- States "THIS FILE MUST NOT BE EDITED BY HAND"
- Confusing for end users (it's for maintainers)
- Should clarify this is auto-generated with timestamp
- **Impact:** Minor UX issue

---

## 3. Missing Documentation

### 🔴 Critical Gaps

#### 3.1 No Troubleshooting Guide

**Missing:** Comprehensive troubleshooting documentation
**Needs:**

- Common errors (DNS records not created, permission issues, etc.)
- Debug logging guidance
- Step-by-step diagnostic procedures
- Provider-specific troubleshooting
- **Impact:** Users struggle with common issues, increase support burden

#### 3.2 No Migration Guide

**Missing:** Migration from other DNS controllers to External DNS
**Needs:**

- Migration from Mate/Molecule/Kops
- Handling existing DNS records during migration
- Rollback procedures
- Testing migration before cutover
- **Impact:** Barrier to adoption for users with existing DNS automation

#### 3.3 No Security/RBAC Deep Dive

**Current State:** Only basic RBAC examples in tutorials
**Needs:**

- Detailed security best practices
- Least-privilege policies
- Multi-tenant cluster guidance
- Audit recommendations
- Secret management best practices
- **Impact:** Users may deploy with excessive permissions, security risks

### 🟡 Important Gaps

#### 3.4 No Performance Tuning Guide

**Missing:** Optimization guidance for large deployments
**Needs:**

- Sync interval tuning
- Batch size configuration
- Caching strategies
- Resource limits recommendations
- Monitoring and alerting strategies
- **Impact:** Poor performance in large clusters

#### 3.5 No High Availability Guide

**Missing:** Multi-instance deployment guidance
**Needs:**

- Leader election configuration
- Running multiple instances safely
- Failover scenarios
- **Impact:** Users risk DNS conflicts or gaps in coverage

#### 3.6 No Webhook Provider Guide

**Current State:** Webhook provider listed but minimal documentation
**Needs:**

- Schema documentation for webhook requests/responses
- Example webhook implementations
- Authentication/authorization patterns
- **Impact:** Limited webhook adoption

#### 3.7 Missing Complex Configuration Examples

**Needs:**

- Combining multiple sources with filters
- Complex annotation-prefix scenarios beyond split-horizon
- Service mesh integrations
- **Impact:** Users struggle with advanced use cases

---

## 4. Inconsistencies and Errors

### 4.1 Terminology Inconsistencies

- `--txt-owner-id` vs `txt-owner-id` (inconsistent hyphenation)
- `zone` vs `hosted zone` used interchangeably without definition

### 4.2 API Version References

- FAQ.md references unversioned APIs
- Gateway-api.md mentions v1alpha2 (possibly deprecated) with unclear support timeline

### 4.3 Flag Documentation Mismatches

- flags.md says providers/sources are "required" but some deployments use defaults
- Unclear mutual exclusivity (e.g., `--domain-filter` vs `--regex-domain-filter`)

### 4.4 Example Configuration Inconsistencies

- Mix of `--provider=aws` (CLI) vs `provider: aws` (YAML) without context
- Some Helm chart examples missing namespace specifications

### 4.5 Registry Documentation Gaps

- Mentions four registries but no pros/cons comparison
- TXT registry is default but reasons not explained
- DynamoDB registry costs/benefits not documented

---

## 5. Areas Needing Better Explanation

### 5.1 Configuration Precedence

**File:** `docs/advanced/configuration-precedence.md`
**Issue:**

- Good flowchart exists but lacks concrete examples
- Mermaid diagram should be validated
- Missing explanation of when annotations override vs ignored

### 5.2 FQDN Templating

**File:** `docs/advanced/fqdn-templating.md`
**Issues:**

- Complex examples need step-by-step walkthroughs
- Warning about subdomain-only hostnames (line 278) needs better prominence
- Source support matrix clarity

### 5.3 Registry Selection

**Missing:** Decision tree for choosing between TXT, DynamoDB, noop, aws-sd registries
**Missing:** Migration guidance between registry types

### 5.4 Source Selection

**Missing:** Guidance on choosing between multiple sources
**Missing:** Performance implications of different source combinations

### 5.5 DNS Record Types

**Issues:**

- `--managed-record-types` documentation unclear on valid combinations
- No explanation of ALIAS vs CNAME for AWS
- Limited docs on MX, NS, SRV, TXT record management

### 5.6 Provider Setup Complexity

**Issues:**

- Provider tutorials vary significantly in depth (10 lines vs 100+)
- No consistent structure across provider documentation
- Missing common troubleshooting steps per provider

---

## 6. Documentation Quality Matrix

| Category | Status | Details |
|----------|--------|---------|
| Basic Usage | ✅ Excellent | Clear tutorials, well-documented |
| Provider Support | ✅ Good | 40+ providers with tutorials |
| Advanced Features | ✅ Good | Split-horizon, templating, TTL documented |
| Troubleshooting | ❌ Poor | Lacks comprehensive guide |
| Migration Path | ❌ Missing | No guidance for users switching from other solutions |
| Security/RBAC | ⚠️ Minimal | Only basic examples provided |
| Performance Tuning | ❌ Missing | No optimization guides |
| High Availability | ❌ Missing | No multi-instance deployment guide |
| Configuration | ⚠️ Moderate | Some complexity poorly explained |
| API/Schema Docs | ⚠️ Partial | Annotations documented but webhook schema missing |

---

## 7. Recommended Actions

### 🔴 High Priority (Address First)

1. **Fix FQDN Templating Contradiction**
   - File: `docs/advanced/fqdn-templating.md:240`
   - Clarify interaction between `--fqdn-template` and hostname annotations
   - Provide clear examples of when each is used vs combined

2. **Consolidate Gateway Documentation**
   - Files: `docs/sources/gateway.md` and `docs/sources/gateway-api.md`
   - Clarify which is canonical or merge them
   - Update API version support timeline

3. **Complete TODO Markers**
   - File: `docs/advanced/fqdn-templating.md:34,66`
   - Finish incomplete sections

4. **Create Troubleshooting Guide**
   - New file: `docs/troubleshooting.md`
   - Include common errors, debug procedures, provider-specific issues

5. **Add Registry Comparison Guide**
   - New section or file explaining TXT vs DynamoDB vs noop vs aws-sd
   - Decision tree for selection
   - Migration procedures

### 🟡 Medium Priority

6. **Create Migration Guide**
   - New file: `docs/migration.md`
   - Cover migration from Mate, Molecule, Kops, custom solutions

7. **Add Security/RBAC Deep Dive**
   - New file: `docs/security-best-practices.md`
   - Least-privilege examples
   - Multi-tenant guidance
   - Secret management

8. **Create Performance Tuning Guide**
   - New file: `docs/performance-tuning.md`
   - Large cluster optimization
   - Resource limits
   - Monitoring setup

9. **Update Outdated References**
   - File: `docs/faq.md:63-65`
   - Update Traefik to v2.x+
   - Fix Helm chart references

10. **Document Webhook Provider**
    - Enhance `docs/tutorials/webhook.md` (if exists) or create it
    - Add schema documentation
    - Provide example implementations

### ⚪ Low Priority (Nice to Have)

11. **Clarify Provider Table**
    - File: `docs/providers.md`
    - Explain columns and "n/a" values
    - Add links to detailed provider docs

12. **Standardize Example Formats**
    - Consistent use of CLI vs YAML vs Helm
    - Add context for when to use each format

13. **Add Source Selection Decision Tree**
    - Help users choose appropriate sources
    - Document performance implications

14. **Update Release Documentation**
    - File: `docs/release.md:5-6`
    - Provide clearer timeline or process

15. **Add Complex Configuration Examples**
    - Multiple sources with filters
    - Advanced annotation-prefix scenarios
    - Service mesh integrations

---

## 8. File References for Quick Access

### Files with Issues

- `docs/advanced/fqdn-templating.md` (lines 34, 66, 240, 278, 299)
- `docs/sources/gateway.md` (vs gateway-api.md conflict)
- `docs/sources/gateway-api.md` (vs gateway.md conflict)
- `docs/faq.md` (lines 63-65)
- `docs/providers.md` (table needs explanation)
- `docs/release.md` (lines 5-6)
- `docs/flags.md` (lines 3-5)
- `docs/advanced/configuration-precedence.md` (needs examples)

### Files to Create

- `docs/troubleshooting.md` ⭐ High Priority
- `docs/migration.md` ⭐ High Priority
- `docs/security-best-practices.md` ⭐ Medium Priority
- `docs/performance-tuning.md` ⭐ Medium Priority
- `docs/high-availability.md` ⭐ Medium Priority
- `docs/registry-comparison.md` or section in existing registry docs

---

## Next Steps

1. Review this audit with the team
2. Prioritize which issues to tackle first
3. Assign owners to each documentation improvement
4. Create GitHub issues for tracking
5. Set timeline for addressing high-priority items

---

**Note:** This audit was generated through comprehensive automated analysis of the docs folder. All line numbers and file references were accurate as of 2025-12-22.
