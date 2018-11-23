# Base container image
FROM golang:1.8-alpine AS builder

# Use Alpine's apk tool to install git
RUN apk --no-cache -U add git

# Install package manager
RUN go get -u github.com/kardianos/govendor

# Copy app files into container
WORKDIR /go/src/app
COPY . .

# Install dependencies
RUN govendor sync

# Build the app
RUN govendor build -o /go/src/app/weatherpin

#=================================================#

# Smallest container image
FROM alpine

# Copy built executable from base image to here
COPY --from=builder /go/src/app/weatherpin /
COPY --from=builder /go/src/app/conf.yml /

# Run the app
CMD ["/weatherpin"]