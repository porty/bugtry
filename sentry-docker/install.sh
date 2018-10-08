#!/bin/bash

set -euo pipefail

docker-compose run --rm sentry config generate-secret-key
SENTRY_SECRET_KEY='09n=0&n=k&14_x@mq8)%=lg765&cukiqj1(+mw1s^gqz!pjk=('
docker-compose run --rm -e SENTRY_SECRET_KEY="$SENTRY_SECRET_KEY" -e SENTRY_REDIS_HOST="redis" -e SENTRY_POSTGRES_HOST="postgres" sentry upgrade
