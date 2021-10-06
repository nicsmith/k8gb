# k8gb - Kubernetes Global Balancer

A Global Service Load Balancing solution with a focus on having cloud native qualities and work natively in a Kubernetes context.


Global load balancing, commonly referred to as GSLB (Global Server Load Balancing) solutions, has been typically the domain of proprietary network software and hardware vendors and installed and managed by siloed network teams.

k8gb is a completely open source, cloud native, global load balancing solution for Kubernetes.

k8gb focuses on load balancing traffic across geographically dispersed Kubernetes clusters using multiple load balancing strategies to meet requirements such as region failover for high availability.

Global load balancing for any Kubernetes Service can now be enabled and managed by any operations or development teams in the same Kubernetes native way as any other custom resource.

## Key Differentiators

* Load balancing is based on timeproof DNS protocol which is perfect for global scope and extremely reliable
* No dedicated management cluster and no single point of failure
* Kubernetes native application health checks utilizing status of Liveness and Readiness probes for load balancing decisions
* Configuration with a single Kubernetes CRD of Gslb kind

## Quick Start

Simply run

```sh
make deploy-full-local-setup
```

It will deploy two local [k3s](https://k3s.io/) clusters via [k3d](https://k3d.io/), [expose associated CoreDNS service for UDP DNS traffic](./docs/exposing_dns.md)), and install k8gb with test applications and two sample Gslb resources on top.

This setup is adapted for local scenarios and works without external DNS provider dependency.

Consult with [local playground](/docs/local.md) documentation to learn all the details of experimenting with local setup.

Optionally, you can run `make deploy-prometheus` and check the metrics on the test clusters (http://localhost:8080, http://localhost:8081).
