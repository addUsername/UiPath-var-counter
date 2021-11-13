FROM golang:1.9.7-alpine

WORKDIR /app

COPY . .

#RUN go mod tidy
RUN go get github.com/gookit/color

ENTRYPOINT ["tail", "-f", "/dev/null"]
#ENTRYPOINT ["/bin/sh", "-c" , "docker cp containerId:/sourceFilePath/someFile.txt C:/localMachineDestinationFolder"]