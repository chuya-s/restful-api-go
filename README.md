# vscode

## Language Server
Code > Preferences > Settings
check `go.useLanguageServer`

## GO111MODULE
setting.json in vscode
(cmd+shift+P with command pallet)
```
 "go.toolsEnvVars": {"GO111MODULE": "on"},
```

## gopls
setting.json in vscode
(cmd+shift+P with command pallet)
check `gopls: Language Server from Google`

## Reference
https://qiita.com/chanhama/items/a21ca7d5cd43d6f3f90d

# Rest Architecture
## JSON
encodeing/json
## Marshalling
data structures in GO into JSON
## Getting Started With A Basic API

# go mod
## go mod init
`$ go mod init example.com/XXX`
## go get
`$ go get github.com/gorilla/mux`


# Docker (Local)
## Dockerfile
edit Dockerfile
## docker build
`$ docker build -t backen .`
## docker run
`$ docker run -it -p 8080:8080 backend`
