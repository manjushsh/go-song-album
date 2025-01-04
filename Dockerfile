# https://www.docker.com/blog/developing-go-apps-docker/

# Specifies a parent image
FROM golang:1.23.4-alpine3.21

# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
# RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN cd server && go mod download
 
# Builds your app with optional configuration
RUN cd server && go build -o /godocker
 
# Tells Docker which network port your container listens on
EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
CMD [ “/godocker” ]