FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
ADD ./app /app/
ADD ./s3_access_keys.csv /app/
ADD ./key.json /app/
# ADD ./zoneinfo.zip /usr/lsocal/go/lib/time/
ADD ./demo.html /app/
ENTRYPOINT ["./app"]