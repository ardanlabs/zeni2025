# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

run-db:
	CGO_ENABLED=0 go run api/services/sales/main.go --db-host localhost | CGO_ENABLED=0 go run api/tooling/logfmt/main.go

run:
	CGO_ENABLED=0 go run api/services/sales/main.go | CGO_ENABLED=0 go run api/tooling/logfmt/main.go

run-help:
	CGO_ENABLED=0 go run api/services/sales/main.go --help

curl:
	curl -i http://localhost:3000/test

# ==============================================================================
# Define dependencies

GOLANG          := golang:1.24
ALPINE          := alpine:3.21
KIND            := kindest/node:v1.32.0
POSTGRES        := postgres:17.4
GRAFANA         := grafana/grafana:11.6.0
PROMETHEUS      := prom/prometheus:v3.2.0
TEMPO           := grafana/tempo:2.7.0
LOKI            := grafana/loki:3.4.0
PROMTAIL        := grafana/promtail:3.4.0

KIND_CLUSTER    := ardan-starter-cluster
NAMESPACE       := sales-system
SALES_APP       := sales
AUTH_APP        := auth
BASE_IMAGE_NAME := localhost/ardanlabs
VERSION         := 0.0.1
SALES_IMAGE     := $(BASE_IMAGE_NAME)/$(SALES_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)

# VERSION       := "0.0.1-$(shell git rev-parse --short HEAD)"

# ==============================================================================
# Building containers

build: sales

sales:
	docker build \
		-f zarf/docker/dockerfile.sales \
		-t $(SALES_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

# ==============================================================================
# Running from within k8s/kind

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER) & \
	wait;

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

# ------------------------------------------------------------------------------

dev-load:
	kind load docker-image $(SALES_IMAGE) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/database | kubectl apply -f -
	kubectl rollout status --namespace=$(NAMESPACE) --watch --timeout=120s sts/database

	kustomize build zarf/k8s/dev/sales | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(SALES_APP) --timeout=120s --for=condition=Ready

dev-restart:
	kubectl rollout restart deployment $(SALES_APP) --namespace=$(NAMESPACE)

dev-run: build dev-up dev-load dev-apply

dev-update: build dev-load dev-restart

dev-update-apply: build dev-load dev-apply

dev-logs:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) --all-containers=true -f --tail=100 --max-log-requests=6 | go run api/tooling/logfmt/main.go -service=$(SALES_APP)

dev-logs-db:
	kubectl logs --namespace=$(NAMESPACE) -l app=database --all-containers=true -f --tail=100

dev-logs-init:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) -f --tail=100 -c init-migrate-seed

# ------------------------------------------------------------------------------

dev-describe-deployment:
	kubectl describe deployment --namespace=$(NAMESPACE) $(SALES_APP)

dev-describe-sales:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(SALES_APP)

dev-describe-database:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=database

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

# ==============================================================================
# Metrics and Tracing

metrics:
	expvarmon -ports="localhost:3010" -vars="build,requests,goroutines,errors,panics,mem:memstats.HeapAlloc,mem:memstats.HeapSys,mem:memstats.Sys"

statsviz:
	open http://localhost:3010/debug/statsviz

# ==============================================================================
# Administration

pgcli:
	pgcli postgresql://postgres:postgres@localhost

# ==============================================================================
# Hitting endpoints

users:
	curl -i http://localhost:3000/v1/users?page=1&rows=2
