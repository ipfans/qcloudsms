# qcloudsms
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ipfans/qcloudsms)
[![Go Report Card](https://goreportcard.com/badge/github.com/ipfans/qcloudsms?style=flat-square)](https://goreportcard.com/report/github.com/ipfans/qcloudsms)
[![Build Status](https://travis-ci.org/ipfans/qcloudsms.svg?style=flat-square&branch=master)](https://travis-ci.org/ipfans/qcloudsms)
[![codecov](https://codecov.io/gh/ipfans/qcloudsms/branch/master/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ipfans/qcloudsms)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ipfans/qcloudsms/master/LICENSE)

QCloud SMS Service(https://www.qcloud.com/product/sms) Golang SDK.

## Installtion

```
go get -u -v github.com/ipfans/qcloudsms
```

## Status

Pre-alpha. Lots of functionality is knowingly missing or broken. API may change in v1.0.

## Features

* [ ] SMS
  * [x] sendsms    发送短信
  * [x] sendmultisms2    发送多条短信
  * [ ] smscallback    短信状态回调
  * [ ] smsreply    短信上行
  * [ ] pullstatus    拉取短信状态
  * [x] pullstatus4mobile    拉取单个手机短信状态

* [ ] Improve test coverage
