#!/usr/bin/env sh

set -eo pipefail

if [ -n "${COOKIES_URL}" ]; then
    wget -O /tmp/cookies.txt $COOKIES_URL
    cat /tmp/cookies.txt
fi;

exec presseclub
