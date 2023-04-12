---
title: OEP Title

oep-number: 10

creation-date: 2023-04-12

status: implementable

authors:

- "@Ashughorla"
- "@kasabe28"
- "@Rohit-0505"
- "@sujeet01"
- "@ushabelgur"

reviewers:

- "@ManuStoessel"
- "@gehoern"
- "@afritzler"
- "@adracus"

---

# OEP-10: Security Group

## Table of Contents

- [Summary](#summary)
- [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
- [Proposal](#proposal)
- [Alternatives](#alternatives)

## Summary

A security group controls the traffic that is allowed to reach and leave the 
resources that it is associated with. For example, if a security group is 
associated with an onmetal machine, it controls the inbound and outbound traffic for 
the machine. This proposal describes how to add a security group for `onmetal`'s 
`Machine` resources.

## Motivation

A security group acts as a virtual firewall for resources to control incoming and 
outgoing traffic. Inbound rules control the incoming traffic to resources, and 
outbound rules control the outgoing traffic from resources. Security group can 
only be associated with resources in the same Network. `onmetal`'s `securitygroup` 
API should offer a way to define such rules.

### Goals

* Add inbound and outbound rules in security group to monitor and filter incoming 
  and outgoing traffic from an onmetal machine.

### Non-Goals

* Cross-Namespace consumption of the `SecurityGroup` - A `SecurityGroup` is only
  available within one `Network`

## Proposal

Introduce a `networking.api.onmetal.de.SecurityGroup` resource that dynamically selects multiple target `Machine`s via
a `machineRef` `metav1.LabelSelector` (as e.g. in `AliasPrefix`es).

The `SecurityGroup` should provide rules to filter inbound and outbound traffic to onmetal machines.
`spec.rules` defines rules to filter incoming and outgoing network traffic.
`spec.rules.ingress` defines set of inbound rules. `spec.rules.egress` defines set of outbound rules.

Each `rule` consist of:
- `ip`: A single IP
- `prefix`: allowed IP range
- `ports`: `ports` defines an allow list of which traffic should be handled by a `SecurityGroup`. A `port` consists of
   a `protocol`, `port` and an optional `portEnd` to support port range filtering.

Ingress and Egress rules provided in `SecurityGroup` will be applied for all the `Machine`'s matching given label/labels.

Example Manifest:

```yaml
apiVersion: networking.api.onmetal.de/v1alpha1
kind: SecurityGroup
metadata:
  name: MySecurityGroup
spec:
  machineRef:
    matchLabels:
      app: web
  rules:
    ingress:
      - ip: 10.0.0.1 # It is possible to specify a literal IP
        prefix: 0.0.0.0/0 # or range of IP's
        ports: 
        - # protocols supported UDP, TCP, SCTP
          protocol: tcp
          # single port
          port: 80
        - protocol: udp
          # port range
          port: 1024
          portEnd: 2048
      - prefix: 10.0.0.0/16
        ports: 
        - protocol: udp
          # single port
          port: 3447
    egress:
      - prefix: 10.0.0.0/16
```

## Alternatives