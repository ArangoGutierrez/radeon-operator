# Build the manager binary
FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.16-openshift-4.8 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Build
RUN make build

# Create production image for running the operator
FROM registry.ci.openshift.org/ocp/4.8:base
WORKDIR /
COPY --from=builder /workspace/radeon-operator .
COPY build/assets /opt/device-plugin

RUN useradd  -r -u 499 nonroot
RUN getent group nonroot || groupadd -o -g 499 nonroot 

ENTRYPOINT ["/radeon-operator"]
