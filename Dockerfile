FROM alpine

# Set destination for COPY
WORKDIR /app

COPY ./build/note-go .
COPY ./data ./data
COPY ./templates ./templates

RUN chmod +x ./note-go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["./note-go"]
