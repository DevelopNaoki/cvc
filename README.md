# cvc
## Overview
This is a web UI for controlling containers and VMs
but, now this can only use web ssh

## Premise
```
# go version
go version go1.16.4 linux/amd64
```

## How to use
first, clone this repository
```
git clone https://github.com/DevelopNaoki/cvc/
```

second, get some packages
```
go get github.com/yudai/gotty
```

if ```gotty``` can use, exec this command
```
go build
```
and 
```
./cvc
```

Installation finished and Open http://<your ip address>:3000 on browser
