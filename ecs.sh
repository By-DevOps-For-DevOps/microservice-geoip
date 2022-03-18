#!/usr/bin/env bash
if [ "$DEPLOY_ENVIRONMENT" != "release" ]; then
  GITHUB_TOKEN=${GITHUB_TOKEN}
  git config --global user.email ${GITHUB_EMAIL}
  git config --global user.name ${GITHUB_USERNAME}
  git clone https://github.com/microservices-today/ecs-api ecs  --branch ${CICD_VERSION:-1.0.1}
fi
. ./ecs/build.sh
