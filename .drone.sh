#!/bin/sh

# build a static binary with the build number and extra features.
go build -ldflags '-extldflags "-static" -X github.com/mblink/drone/version.VersionDev=build.'${DRONE_BUILD_NUMBER} -o release/drone-server github.com/mblink/drone/cmd/drone-server
GOOS=linux GOARCH=amd64 CGO_ENABLED=0         go build -ldflags '-X github.com/mblink/drone/version.VersionDev=build.'${DRONE_BUILD_NUMBER} -o release/drone-agent             github.com/mblink/drone/cmd/drone-agent
GOOS=linux GOARCH=arm64 CGO_ENABLED=0         go build -ldflags '-X github.com/mblink/drone/version.VersionDev=build.'${DRONE_BUILD_NUMBER} -o release/linux/arm64/drone-agent github.com/mblink/drone/cmd/drone-agent
GOOS=linux GOARCH=arm   CGO_ENABLED=0 GOARM=7 go build -ldflags '-X github.com/mblink/drone/version.VersionDev=build.'${DRONE_BUILD_NUMBER} -o release/linux/arm/drone-agent   github.com/mblink/drone/cmd/drone-agent
