############################
# STEP 1 prepare essential resources
############################
FROM gcr.io/iaas-gcr-reg-prd-ad3d/build/csa/alpine-tools:3.18.0 as builder

USER root

# needed for running locally, should probably go in docker-compose
COPY --from=us.gcr.io/iaas-gcr-reg-prd-ad3d/cicd/base/cicd-cert-base:latest /*.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates

USER efx_container_user

############################
# STEP 2 build the binary
############################
FROM gcr.io/iaas-gcr-reg-prd-ad3d/build/golang:1.21-alpine as gobin

USER root

RUN apk add git make

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER efx_container_user

WORKDIR /home/efx_container_user/

# Copy dependencies and test files
COPY --chown=efx_container_user:efx_container_user go.mod go.sum ./
COPY --chown=efx_container_user:efx_container_user . ./

RUN chown -R efx_container_user:efx_container_user . && mkdir bin/

RUN CGO_ENABLED=0 make build-docker

############################
# STEP 2 build a small image
############################
FROM gcr.io/iaas-gcr-reg-prd-ad3d/golden/scratch:1.0

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /home/efx_container_user /home/efx_container_user

# Copy our static executable
COPY --from=gobin --chown=efx_container_user /home/efx_container_user/bin/main /home/efx_container_user/main

# Use an unprivileged user.
USER efx_container_user

# Run the hello binary.
ENTRYPOINT ["/home/efx_container_user/main"]
