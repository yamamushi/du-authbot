# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/yamamushi/du-authbot

# Create our shared volume
RUN mkdir /du-authbot

# Run our dependency installation for Opus Encoding/Decoding
#RUN apt-get update && \
#        DEBIAN_FRONTEND=noninteractive apt-get install -y libav-tools opus-tools -f && \
#        apt-get clean && \
#        rm -rf /var/lib/apt/lists/


# Get the du-discordbot dependencies inside the container.
RUN go get github.com/bwmarrin/discordgo
RUN go get github.com/BurntSushi/toml
RUN go get github.com/coreos/bbolt
RUN go get github.com/asdine/storm
RUN go get github.com/gofrs/uuid
RUN go get github.com/wcharczuk/go-chart
RUN go get gopkg.in/mgo.v2


# Install and run du-authbot
RUN go install github.com/yamamushi/du-authbot

# Run the bot by default when the container starts.
WORKDIR /du-authbot
ENTRYPOINT /go/bin/du-authbot

VOLUME /du-authbot