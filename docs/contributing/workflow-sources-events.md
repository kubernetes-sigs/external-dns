<!--
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Sources, Listeners, and Events — Workflow

## Overview

ExternalDNS has three interlocking subsystems: **informers** (watch k8s resources),
**sources** (translate resources into desired DNS records), and **events** (emit k8s
events back against the originating resources when DNS changes are applied).

---

## Startup: Informer & Source Construction

```
main.go
  └─► source.BuildWithConfig(cfg)
        └─► for each --source flag:
              ├─ NewGatewaySource / NewIngressSource / NewServiceSource / ...
              │     ├─ Create SharedInformerFactory (scoped to namespace + label filter)
              │     ├─ Get typed informer  (e.g. GatewayInformer)
              │     ├─ MustAddIndexers     (annotation filter + controller match)
              │     ├─ MustAddEventHandler (DefaultEventHandler — primes the cache)
              │     ├─ factory.Start(ctx)
              │     └─ WaitForCacheSync
              └─► Returns Source implementation
```

---

## Runtime: Watch Loop (Informers → Controller)

```
k8s API Server
  │  (watch stream: add/update/delete)
  ▼
SharedInformer (per resource type)
  │  cache updated
  ▼
Event handler registered via source.AddEventHandler()
  │  calls handler()   ← a simple func()
  ▼
Controller.ScheduleRunOnce()
  │  debounced by --min-event-sync-interval (default 5s)
  ▼
Controller.RunOnce()
```

---

## RunOnce: Reconciliation Loop

```
Controller.RunOnce(ctx)
  │
  ├─1─ Registry.Records()
  │      └─► Provider.Records() → current DNS state (what's in the DNS provider)
  │
  ├─2─ Source.Endpoints(ctx)
  │      └─► for each k8s resource (read from informer cache):
  │            ├─ filter by annotation / label / controller
  │            ├─ extract hostnames (spec, annotations, fqdn-template)
  │            ├─ extract targets   (status.addresses / loadBalancer IPs)
  │            ├─ build []*endpoint.Endpoint
  │            └─ AttachRefObject(endpoints, ObjectReference{kind, ns, name, uid})
  │                    ▲
  │                    └── this tags each Endpoint with the k8s object that owns it
  │
  ├─3─ Registry.AdjustEndpoints(sourceEndpoints)
  │      └─► normalize / add TXT ownership records
  │
  ├─4─ plan.Plan{}.Calculate()
  │      └─► diff: Create / UpdateOld+UpdateNew / Delete
  │
  └─5─ Registry.ApplyChanges(plan.Changes)
         ├─► Provider.ApplyChanges() → mutate DNS provider
         └─► emitChangeEvent(EventEmitter, changes, reason)
```

---

## Events: Emission After DNS Change

```
emitChangeEvent(EventEmitter, plan.Changes, reason)
  │
  ├─ for ep in Changes.Create   → NewEventFromEndpoint(ep, ActionCreate,  reason)
  ├─ for ep in Changes.UpdateNew→ NewEventFromEndpoint(ep, ActionUpdate,  reason)
  └─ for ep in Changes.Delete   → NewEventFromEndpoint(ep, ActionDelete,  RecordDeleted)
                │
                ▼
  NewEventFromEndpoint(ep, action, reason)
    └─► ep.RefObjects()          ← the ObjectReferences attached by Source
          │
          └─► for each ObjectReference:
                eventForRef(ref) → eventsv1.Event{
                  regarding:  {kind, ns, name, uid}   ← the k8s resource
                  reason:     RecordReady | RecordError | RecordDeleted
                  action:     Created | Updated | Deleted
                  note:       "Synced DNS record <host> → <target>"
                }
                │
                ▼
              EventEmitter.Add(event)
                └─► workqueue → EventsV1Interface.Create()
                      └─► k8s API Server (Event object visible in kubectl describe)
```

---

## Source Responsibility: AttachRefObject

Every source **must** call `endpoint.AttachRefObject` so the events subsystem
knows which k8s object to emit the event against.

```go
// In source/gateway.go (gatewaySource)
endpoint.AttachRefObject(gwEndpoints, events.NewObjectReference(gw, types.Gateway))

// In source/gateway.go (gatewayRouteSource)
endpoint.AttachRefObject(routeEndpoints, events.NewObjectReference(rt.Object(), "gateway-"+kind))

// In source/ingress.go
endpoint.AttachRefObject(ingEndpoints, events.NewObjectReference(ing, types.Ingress))
```

If `AttachRefObject` is **not** called, `ep.RefObjects()` returns nil and
`NewEventFromEndpoint` produces no events — DNS changes happen silently with
no feedback in `kubectl describe <resource>`.

---

## Full Data Flow (single resource change)

```
Gateway resource updated in k8s
        │
        ▼
GatewayInformer detects delta
        │
        ▼
handler() called → Controller.ScheduleRunOnce()
        │
        ▼
Controller.RunOnce()
  ├─ Source.Endpoints() reads Gateway from cache
  │     builds Endpoint{DNSName: "foo.example.com", Targets: ["1.2.3.4"]}
  │     AttachRefObject → tags endpoint with Gateway{ns/name/uid}
  │
  ├─ plan.Calculate() → Changes{Create: [foo.example.com → 1.2.3.4]}
  │
  ├─ Provider.ApplyChanges() → Route53 / CloudDNS / ...
  │
  └─ emitChangeEvent()
        └─ ep.RefObjects() → Gateway{default/my-gw}
              └─ k8s Event: "Created DNS record foo.example.com → 1.2.3.4"
                    emitted ON the Gateway object
                    visible via: kubectl describe gateway my-gw
```

---

## Key Packages

| Package | Role |
|---------|------|
| `source/` | Read k8s resources → `[]*endpoint.Endpoint` |
| `source/informers/` | Informer helpers: transforms, indexers, event handlers |
| `controller/` | Orchestrates RunOnce loop |
| `pkg/events/` | Build + emit k8s Event objects via EventsV1 API |
| `plan/` | Diff current vs desired DNS state |
| `registry/` | Ownership tracking via TXT records |
| `provider/` | DNS provider API calls |
