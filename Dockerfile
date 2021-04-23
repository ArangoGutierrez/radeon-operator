# Build the manager binary
FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.16-openshift-4.8 as builder

WORKDIR /workspace

# Build
COPY . .
RUN make build

# Create production image for running the operator
FROM registry.ci.openshift.org/ocp/4.8:base
WORKDIR /
COPY --from=builder /workspace/radeon-operator .
COPY build/assets /opt/device-plugin

RUN useradd  -r -u 499 nonroot
RUN getent group nonroot || groupadd -o -g 499 nonroot 

ENTRYPOINT ["/radeon-operator"]
