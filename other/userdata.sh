#!/bin/bash
sudo -u ubuntu bash -c 'GOPATH=/home/ubuntu/go go get -u github.com/webercoder/go-mtg-hipchat-bot'
sudo -u ubuntu bash -c 'GOPATH=/home/ubuntu/go go install github.com/webercoder/go-mtg-hipchat-bot'
sudo supervisorctl restart go-mtg-hipchat-bot
