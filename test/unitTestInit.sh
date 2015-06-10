#!/bin/sh

docker kill `docker ps -a -q`
docker rm `docker ps -a -q`
docker rmi `docker images -a -q`

docker pull busybox:latest
BUSYBOX=`docker create busybox`

docker commit $BUSYBOX test_1
docker commit $BUSYBOX test_2
docker commit $BUSYBOX test_3
docker commit $BUSYBOX test_4
docker commit $BUSYBOX test_5
docker commit $BUSYBOX test_6
docker commit $BUSYBOX test_7
TEST_1=`docker create test_1`
TEST_2=`docker create test_2`
TEST_3=`docker create test_3`
TEST_5=`docker create test_5`

docker commit $TEST_1 test_1_1
docker commit $TEST_1 test_1_2
docker commit $TEST_1 test_1_3
docker commit $TEST_1 test_1_4
TEST_1_1=`docker create test_1_1`
TEST_1_2=`docker create test_1_2`
TEST_1_3=`docker create test_1_3`

docker commit $TEST_2 test_2_1
docker commit $TEST_2 test_2_2
docker commit $TEST_2 test_2_3
docker commit $TEST_2 test_2_4
docker commit $TEST_2 test_2_5
TEST_2_1=`docker create test_2_1`
TEST_2_2=`docker create test_2_2`
TEST_2_5=`docker create test_2_5`

docker commit $TEST_3 test_3_1
docker commit $TEST_3 test_3_2
TEST_3_1=`docker create test_3_1`
TEST_3_2=`docker create test_3_2`

docker commit $TEST_5 test_5_1
docker commit $TEST_5 test_5_2
docker commit $TEST_5 test_5_3

docker commit $TEST_1_1 test_1_1_1
docker commit $TEST_1_1 test_1_1_2
docker commit $TEST_1_1 test_1_1_3

docker commit $TEST_1_2 test_1_2_1
docker commit $TEST_1_2 test_1_2_2
docker commit $TEST_1_2 test_1_2_3
docker commit $TEST_1_2 test_1_2_4

docker commit $TEST_1_3 test_1_3_1
docker commit $TEST_1_3 test_1_3_2
docker commit $TEST_1_3 test_1_3_3
docker commit $TEST_1_3 test_1_3_4
docker commit $TEST_1_3 test_1_3_5
docker commit $TEST_1_3 test_1_3_6

docker commit $TEST_2_2 test_2_2_1
docker commit $TEST_2_2 test_2_2_2
docker commit $TEST_2_2 test_2_2_3
docker commit $TEST_2_2 test_2_2_4

docker commit $TEST_3_1 test_3_1_1
docker commit $TEST_3_1 test_3_1_2
docker commit $TEST_3_1 test_3_1_3

docker commit $TEST_3_2 test_3_3_1
docker commit $TEST_3_2 test_3_3_2
docker commit $TEST_3_2 test_3_3_3
docker commit $TEST_3_2 test_3_3_4
docker commit $TEST_3_2 test_3_3_5
docker commit $TEST_3_2 test_3_3_6