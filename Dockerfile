FROM alpine:3.7

RUN apk update ca-certificates && apk add ca-certificates
RUN addgroup -S sat && adduser -S -g sat sat
USER sat:sat

COPY --chown=sat:sat s3analyser /home/sat/bin/s3analyser

ENTRYPOINT ["/home/sat/bin/s3analyser"]
