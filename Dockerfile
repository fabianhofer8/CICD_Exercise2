FROM golang:1.20-alpine

# Set maintainer label: maintainer=[YOUR-EMAIL]
LABEL maintainer="s2310455008@fhooe.at"

# Set working directory: `/src`
WORKDIR /src

COPY *.mod ./
COPY *.go ./
COPY *.sum ./

# List items in the working directory (ls)
RUN ls

# Build the GO app as myapp binary and move it to /usr/
RUN go build -o myapp

#Expose port 8888
EXPOSE 8888

# Run the service myapp when a container of this image is launched
CMD ["./myapp"]
