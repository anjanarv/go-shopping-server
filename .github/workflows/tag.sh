#! /bin/bash

# Assumes that you tag versions with the version number (e.g., "1.1")
# and then the build number is that plus the number of commits since
# the tag (e.g., "1.1.17")

DESCRIBE=$(git describe --tags `git rev-list --tags --max-count=1`)

echo "first="$DESCRIBE

# increment the build number (ie 115 to 116)
MAJOR=`echo $DESCRIBE | awk '{split($0,a,"."); print a[1]}'`
MINOR=`echo $DESCRIBE | awk '{split($0,a,"."); print a[2]}'`
PATCH=`echo $DESCRIBE | awk '{split($0,a,"."); print a[3]}'`

logmsg=$(git log -1 --pretty=oneline ${hash} | awk '$0~var {print $2}' var="${hash}")

echo $logmsg

MAJOR=${MAJOR:1}
echo "MAJOR VERSION="$MAJOR

NEWVERSION=""
case "$logmsg" in
    "fix:") NEWVERSION=$MAJOR:$MINOR:$(($PATCH+1));;
    "feat:") NEWVERSION=$MAJOR:$(($MINOR+1)):$PATCH;;
    "BREAKING CHANGE:") NEWVERSION=v$(($MAJOR+1)):$MINOR:$PATCH;;
    "*") NEWVERSION=$MAJOR:$((MINOR+1)):$PATCH;;
esac

echo $NEWVERSION
export NEWTAG=$NEWVERSION
