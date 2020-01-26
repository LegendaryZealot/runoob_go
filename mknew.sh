#!/bin/bash

parameterNum=1;

if [ $# == $parameterNum ]
then
    mkdir $1;
    cp ./Makefile ./$1/;
    touch ./$1/main.go;
else
    echo "bad parameter(s)!";
fi