#!/bin/bash
day=$1
if [ -z ${day} ]
then
    echo "missing advent day"
    exit 1
fi
if [ -z ${TOKEN} ]
then
    echo "missing advent login token"
    exit 1
fi
curl "https://adventofcode.com/2020/day/${day}/input" -H "cookie: session=${TOKEN}"
