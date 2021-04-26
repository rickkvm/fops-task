# fops-task

# Table of Contents

- [Overview](#Overview)
- [How to](#Howto)
- [Denendencies](#Dependencies)
- [Features](#Features)
- [Known issues](#Issues)

## Overview
This is a offline learning and coding project for TrendMicro github interview process. This project will implement [feature](##Features) list below.
The main idea of code architecture is based on [cobra](cobra) framework. The file structure will look like

        .
        ├── linecount
        │   └── count.go
        ├── checksum
        │   └── checksum.go
        ├── cmd
        │   ├── checksum.go
        │   ├── linecount.go
        │   ├── root.go
        │   └── version.go
        └── main.go

Whenever we need to add a new subcommand, we can add cmd interface inside `cmd` folder, and implement the feature as the sub-package of `fops` package.


## HowTo
Get `fops` source code and install related dependencies:

        go get -u github.com/rickkvm/fops-task

Change directory into the repo. You can build code with following command, in which `main.VERSION=prebuild` is the application build version

        go build -ldflags "-X main.VERSION=prebuild" . -o fops
        
You'll get the generated binary `fops` in the same folder. You can run the application with commands `./fops`
Output: 
```text
File Ops

Usage:
  fops [flags]
  fops [command]

Available Commands:
  linecount   Print line count of file
  checksum    Print checksum of file
  version     Show the version info
  help        Help about commands

Flags:
  -h, --help   help for fops
```

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

## Dependencies
* [cobra](cobra) - This project uses cobra as command line framework.

## Issues
* cannot determine text file newline separator (windows file will be incorrect)
* counts big file will exhausted memory

[cobra]:(https://github.com/spf13/cobra)
