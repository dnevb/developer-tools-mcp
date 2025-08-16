FROM alpine:latest

WORKDIR /root/

COPY corekit-mcp .

ENTRYPOINT ["./corekit-mcp"]
