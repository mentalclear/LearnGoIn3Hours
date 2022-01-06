### GOPATH environment variable setup

1. In my environment I have a `GOPATH` directory that contains my Go code here: ~/go/src 
2. So for this one to work with custom packages I need to add the `GOPATH` pointing to that specific directory like so:
`export GOPATH=$HOME/go` there's no need to add  /src it will find it automatically.
3. Then just add that one to ~/.bashrc
4. A module GO111MODULE needs to be off so Go would look into GOPATH for packages. Do that like so: 
`go env -w GO111MODULE=off`
