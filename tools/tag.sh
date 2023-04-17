#!/bin/bash
#CHANGED_PKG=$1
#if [ -z $1 ]; then
#    indent 4 "- [${RED}ERROR${RESET}] - Environment variable CHANGED_PKG is not defined or empty."
#    echo ""
#    exit 1
#fi
echo "TAGGING CHANGED PKG - "$CHANGED_PKG
CHANGED_PKG=sampleserver
DESCRIBE=$(git describe --match "sampleserver/v*" --tags `git rev-list --tags --max-count=1`) >/dev/null
NEWTAG=""
if [ -z "$DESCRIBE" ]; then
    echo "No previous tags exist in format ${CHANGED_PKG}/vx.x.x-rc*.Creating new version tag v1.0.0-rc1 ..."
    NEWTAG="v1.0.0-rc1"
else
    echo "Previous tags exist in format ${CHANGED_PKG}/vx.x.x"
    MAJOR=`echo $DESCRIBE | awk '{split($0,a,"."); print a[1]}'`
    MINOR=`echo $DESCRIBE | awk '{split($0,a,"."); print a[2]}'`
    PATCH=`echo $DESCRIBE | awk '{split($0,a,"."); print a[3]}'`
    PATCH=`echo $PATCH | awk '{split($0,a,"-"); print a[1]}'`
if [[ $DESCRIBE == *rc[0-9]* ]]; then
    echo "RC tags exist in format ${CHANGED_PKG}/vx.x.x-rc*"
    RC=`echo $DESCRIBE | awk '{split($0,a,"-"); print a[2]}' | tr -dc '0-9'`
  NEWTAG=$MAJOR.$MINOR.$PATCH-$(($RC+1));
else
  NEWTAG=$MAJOR.$(($MINOR+1)).$PATCH-rc1;
fi
fi
NEWTAG=$CHANGED_PKG/$NEWTAG
echo "New tag created: "$NEWTAG
tag=$(git tag -a $NEWTAG -m $NEWTAG)
push_tag=$(git push --tags)
echo "New tag successfully pushed to remote "$push_tag