# Overview

HipChat has an integrations feature that will send slash commands to the URL of
your choosing.

This is a Magic the Gathering service that responds to the command `/mtg card-search`.

# Installation

Tested on ubuntu 16.04. Assuming you are the Ubuntu user on an EC2 machine and have
your GOPATH set to `/home/ubuntu/go`:

```
sudo apt-get install supervisor
go get github.com/webercoder/go-mtg-hipchat-bot
go install github.com/webercoder/go-mtg-hipchat-bot
sudo cp $GOPATH/src/github.com/webercoder/go-mtg-hipchat-bot/other/supervisor.conf /etc/supervisor/conf.d/go-mtg-hipchat-bot.conf
sudo service supervisor start
sudo supervisorctl reread
sudo supervisorctl update
```

To update and restart directly on the server:

```
go get -u github.com/webercoder/go-mtg-hipchat-bot
go install github.com/webercoder/go-mtg-hipchat-bot
sudo supervisorctl restart go-mtg-hipchat-bot
```

That's it for now.
