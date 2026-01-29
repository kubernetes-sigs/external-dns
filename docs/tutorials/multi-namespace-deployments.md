# Multi-Namespace Deployments

This guide explains how to configure external-dns to watch multiple namespaces in a multi-tenant cluster.

## Overview

By default, external-dns operates in one of two modes:
- **Single namespace**: `--namespace=production` (watches only one namespace)
- **Cluster-wide**: `--namespace=""` (watches all namespaces)

For multi-tenant clusters, you may want to watch a **selected set of namespaces** without granting cluster-wide access.

## Use Cases

### Multi-Tenant Clusters
Each tenant has a set of namespaces identified by labels. You want:
- One external-dns instance per tenant
- Each instance only managing DNS for that tenant's namespaces
- Isolation between tenants

### Gateway + Application Separation
Your Gateway (Istio, Kong, etc.) runs in one namespace, while application services run in separate tenant namespaces. You need external-dns to watch both.

## Solution: Namespace Label Selector

Use `--namespace-label-selector` to select namespaces by label.

### Example

```bash
external-dns --source=service --namespace-label-selector=team=platform
```


### Basic Example

**Label your namespaces:**
```bash
kubectl label namespace prod tenant=acme
kubectl label namespace staging tenant=acme
kubectl label namespace dev tenant=acme
