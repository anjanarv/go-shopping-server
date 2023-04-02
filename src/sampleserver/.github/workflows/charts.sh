#! /bin/bash

# Get latest docker tag
tagToUse=`curl -L --fail "https://hub.docker.com/v2/repositories/anjanarv83/sampleserver/tags/?page_size=1000" | \
jq '.results | .[] | .name' -r | \
sed 's/latest//' | \
sort --version-sort | \
tail -n 1`

echo $tagToUse

# tagToUse="v1.1.5-RC1" --testing

MAJOR=`echo $tagToUse | awk '{split($0,a,"."); print a[1]}'`
MINOR=`echo $tagToUse | awk '{split($0,a,"."); print a[2]}'`
PATCH=`echo $tagToUse | awk '{split($0,a,"."); print a[3]}'`

if [[ $tagToUse == *-RC[0-9]* ]] # if there is already an RC docker image tag
then
  echo "increasing RC"
  rcversion=`echo $tagToUse |grep -Eo '[0-9]+$'`
  echo "rcversion is="$rcversion
  PATCH=`echo $PATCH | awk '{split($0,a,"-"); print a[1]}'`
  tagToUse=$MAJOR.$MINOR.$PATCH-RC$(($rcversion+1))
else    # else, if the latest tag does not have an RC in it
  tagToUse=$MAJOR.$MINOR.$(($PATCH+1))-RC1;
fi

echo $tagToUse

# update the latest image tag in values.yaml
sed -i '' 's|^  tag:.*|  tag: '$tagToUse'|' ../../deployment/values.yaml

