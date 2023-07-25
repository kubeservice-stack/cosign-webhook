FROM gcr.io/projectsigstore/cosign:v2.1.1 as cosign-bin

FROM alpine
COPY --from=cosign-bin /ko-app/cosign /usr/local/bin/cosign
ENTRYPOINT [ "cosign" ]