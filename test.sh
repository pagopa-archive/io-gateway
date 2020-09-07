#!/bin/bash
TAG=$(git tag --points-at HEAD)
if test -z "$TAG"
then VER=$(git branch --show-current)
else VER=$TAG
fi
set -e
source source-me-first
if ! iogw/iogw --version 2>&1 | grep $DOCKER_USER/$VER
then 
   echo "FAIL: Version mismatch: exepected $DOCKER_USER/$VER got $(iogw/iogw --version 2>&1)"
   exit 1 
fi
iogw/iogw stop
rm -Rvf $HOME/tmp-iogw-test
docker pull library/redis:5
echo "****** INIT"
iogw/iogw init $HOME/tmp-iogw-test pagopa/io-sdk-javascript --io-apikey=123456890 --wskprops
echo "****** START"
iogw/iogw -v start --skip-pull-images --skip-docker-version --skip-open-browser
echo "****** BUILD"
docker exec --user=$UID iogw-theia bash -c 'bash /home/project/build.sh'
echo "****** STATUS"
iogw/iogw status
CHECK=ISPXNB32R82Y766F
DATA="${1:-$HOME/tmp-iogw-test/data/data.xlsx}"
URL="http://localhost:3280/api/v1/web/guest/iosdk/import"
JSON='{"file": "'$(base64 $DATA | tr -d '\n')'"}'
HEAD="Content-Type: application/json"
if curl -s $URL -H "$HEAD" -d "$JSON"  | grep $CHECK >/dev/null
then echo SUCCESS ; exit 0
else echo FAIL ; exit 1
fi
iogw/iogw stop
