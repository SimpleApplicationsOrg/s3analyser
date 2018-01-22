#!/usr/bin/env bash

set -e

docker build . -t s3analyser:${TRAVIS_COMMIT}
docker tag s3analyser s3analyser/s3analyser:${TRAVIS_TAG};
docker login -u "$DOCKER_SAT_USERNAME" -p "$DOCKER_SAT_PASSWORD";
docker push s3analyser/s3analyser:${TRAVIS_TAG};
