FROM golang:1.16.10-alpine3.13

WORKDIR /app

COPY . .

ENTRYPOINT ["tail", "-f", "/dev/null"]
#ENTRYPOINT ["/bin/sh", "-c" , "docker cp containerId:/sourceFilePath/someFile.txt C:/localMachineDestinationFolder"]