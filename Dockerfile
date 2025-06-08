FROM golang:1.24

WORKDIR /CMS

COPY . .

# Make sure .bin exists and build binary
RUN mkdir -p .bin && go build -o .bin/couppn-api ./cmd/main.go

# Run the binary
CMD [ "./.bin/couppn-api" ]
