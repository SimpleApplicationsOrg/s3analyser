FROM alpine:3.7
RUN apk update ca-certificates && apk add ca-certificates
ADD https://github.com/SimpleApplicationsOrg/s3analyser/releases/download/$CACHE_TAG/s3analyser /home/sat/bin/s3analyser
RUN chmod +x /home/sat/bin/s3analyser
ENTRYPOINT ["/home/sat/bin/s3analyser"]
