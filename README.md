# fops-task

# Table of Contents

- [Overview](#Overview)
- [How to](##Howto)
- [Features](##Features)
- [Known issues](##Issues)

## Overview
This is a offline learning and coding project for TrendMicro github interview process.
add something

## HowTo
Get `fops` source code and install related dependencies:

        go get -u github.com/rickkvm/fops-task

You can build code with following command, in which `main.VERSION=prebuild` is the application build version

        go build -ldflags "-X main.VERSION=prebuild" .


## Features
This project will implemnt following features  
- [x] linecount  
- [x] checksum
  * [x] md5
  * [x] sha1
  * [x] sha256
- [x] version
- [x] help information
  * [x] linecount
  * [x] checksum

## Issues
* cannot determine text file newline separator (windows file will be incorrect)
* counts big file will exhausted memory
