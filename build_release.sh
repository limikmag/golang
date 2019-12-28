#!/usr/bin/env bash
echo $limikmag_github_token
docker run --rm --privileged \
  -v /home/michal/Desktop/repos/golang:/go/src/github.com/limikmag/golang \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -w /go/src/github.com/limikmag/golang \
  -e GITHUB_TOKEN=${limikmag_github_token} \
  goreleaser/goreleaser release
