# uniform
[![Build Status](https://travis-ci.org/takecy/uniform.svg?branch=master)](https://travis-ci.org/takecy/uniform)
[![Go](https://img.shields.io/badge/language-go-blue.svg?style=flat)](https://golang.org/)
[![GoDoc](https://godoc.org/github.com/takecy/uniform?status.svg)](https://godoc.org/github.com/takecy/uniform)

:tshirt: Deployed resource (in EC2 instances) version is uniform ?

<br/>
## Overview
Check version of deployed resources on multiple servers.  
If you deployed resources(ex)api binary) to multiple server,  
all resources, Is the same version on all servers?

<br/>
## Usage
Set your AWS Key, Secret to Environment Variables.
```bash
$ export UNIFORM_AWS_KEY=your_key
$ export UNIFORM_AWS_SECRET=your_secret
$ export UNIFORM_AWS_REGION=ap-northeast-1
```

Show usage
```bash
$ uniform
```
Check deployed binary version
```bash
$ uniform -t Environment=Production,Type=API,Name=1a01
```


<br/>
## Features
- [x] Check resource version
 - [x] Specific version API endpoint
 - [x] Specific reagion
 - [x] Specific Tags
- [x] Support Autoscaling

<br/>
## Development
```bash
$ git clone git@github.com:takecy/uniform.git
$ cd uniform
$ make prepare
$ make test
```

<br/>
## LICENSE
MIT
