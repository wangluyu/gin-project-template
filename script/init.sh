#!/usr/bin/env bash
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get gopkg.in/yaml.v2
go get -u github.com/go-redis/redis
go get github.com/sirupsen/logrus
go get github.com/dgrijalva/jwt-go

dir=/home
git clone https://github.com/go-swagger/go-swagger "$dir"
cd "$dir"
go install ./cmd/swagger