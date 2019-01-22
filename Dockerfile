#iron/go image is an alpine image that has go tools in it
FROM iron/go:dev
WORKDIR /app

ENV SRC_DIR=/go/src/github.com/srgupta5328/verbose-parakeet/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/
ENTRYPOINT ["./myapp"]

# How to run the Dockerfile for local development
# docker build -t srgupta5328/verbose-parakeet .
# docker run --rm -p 8080:8080 srgupta5328/verbose-parakeet