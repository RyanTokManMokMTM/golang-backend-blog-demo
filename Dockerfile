FROM golang

# setting working dir
RUN mkdir -p /demo
WORKDIR /demo

#Copy entire project to container
Copy . .

# download all module
RUN go mod download

# build
RUN go build -o server

EXPOSE 8000
#run the server
ENTRYPOINT ["./server"]
