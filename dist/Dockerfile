FROM alpine:3.7

# Working Directory
WORKDIR /home/influx2cortex

# Copy in the source
COPY influx2cortex ./influx2cortex

# Make shell scripts executable
RUN chmod +x ./influx2cortex

# Running
CMD ./influx2cortex -auth.enable=$AUTH_ENABLE -server.http-listen-address $LISTEN_ADDRESS -server.http-listen-port $LISTEN_PORT -write-endpoint $WRITE_ENDPOINT