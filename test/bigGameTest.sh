#!/bin/bash

MAXCONTAINER_LVL_2="10"
MAXCONTAINER_LVL_3="10"
MAXCONTAINER_LINKED="10"
MAXCONTAINER_VOLUME="10"

docker kill `docker ps -a -q`
docker rm `docker ps -a -q`
docker rmi `docker images -a -q`

echo "=========================="
echo "    Download BusyBox      "
echo "=========================="

docker pull busybox:latest
echo "=========================="
echo " Create BusyBox container "
echo "=========================="

BUSYBOX=`docker create busybox`

echo "========================="
echo "  Create lvl2 container  "
echo "========================="

count=1

while [ "$count" -lt "$MAXCONTAINER_LVL_2" ]
do
    echo "Create test_$count container"
    docker commit $BUSYBOX test_$count
    container[$count]="`docker run -d --name test_$count test_$count tail -f /bin/echo`"
    ((count++))
done

echo "========================="
echo "  Create lvl3 container  "
echo "========================="

countA=1
countB=1

while [ "$countA" -lt "$MAXCONTAINER_LVL_2" ]
do
    while [ "$countB" -lt "$MAXCONTAINER_LVL_3" ]
    do
        echo "Create test_${countA}_${countB} container"
        docker commit ${container[${countA}]} test_${countA}_${countB}
        container[$count]="`docker run -d --name test_${countA}_${countB} test_${countA}_${countB} tail -f /bin/echo`"
        (( count++ ))
        (( countB++ ))
    done
    countB=1
    (( countA++ ))
done

echo "========================="
echo " Create linked container "
echo "========================="

i=0

while [ "$i" -lt "$MAXCONTAINER_LINKED" ]
do
    i=$((RANDOM%($MAXCONTAINER_LVL_2-1)+1))
    c=$((RANDOM%($MAXCONTAINER_LVL_3-1)+1))
    docker run -d -P --link test_${i}_${c}:test_${i}_${c} busybox /bin/sh
    ((i++))
done

echo "========================="
echo " Create volume container "
echo "========================="

i=0

while [ "$i" -lt "$MAXCONTAINER_VOLUME" ]
do
    i=$((RANDOM%($MAXCONTAINER_LVL_2-1)+1))
    c=$((RANDOM%($MAXCONTAINER_LVL_3-1)+1))
    docker run -d --volumes-from test_${i}_${c} busybox:latest /bin/sh
    ((i++))
done