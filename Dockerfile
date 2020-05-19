FROM znly/protoc

ENV GOMPLATE_VERSION=2.6.0

RUN apk add --no-cache wget ca-certificates bash

ADD generate.sh .

RUN chmod +x generate.sh

VOLUME /tmp/proto

# Remove entrypoint set in parent container
ENTRYPOINT []

CMD ["./generate.sh"]
