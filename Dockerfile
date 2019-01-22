FROM iron/go:dev
WORKDIR /app

ENV SRC_DIR=/go/src/github.com/srgupta5328/verbose-parakeet/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/
ENTRYPOINT ["./myapp"]