# go-solitaire : Solitaire Written in Go
This is a console-based version of Solitaire written in Go.

# install 
```
go get github.com/nirasan/go-solitaire
cd $GOPATH/src/github.com/nirasan/go-solitaire

go get github.com/Masterminds/glide
glide install
```

# Build
```
go build github.com/nirasan/go-solitaire/cmd/solitaire
```

# Execute
```
solitaire
```

# Playable sample for mac
https://github.com/nirasan/go-solitaire/blob/master/sample/solitaire

# How to play
* Move cursor: Arrow key or "hjkl"
* Select card: "space"
* Quit: "q" or "escape"

# `ls` mode
The output of the ls mode is to simulate the output of the ls command .
Please play while at work by all means .
```
solitaire --mode=ls
```