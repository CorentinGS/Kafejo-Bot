export GOPATH="$(go env GOPATH)"
export PATH="${PATH}:${GOPATH}/bin"

echo "Starting go fumpt" && gofumpt -l -w . && \
echo "Starting gci" && gci write -s standard -s default  . && \
echo "Starting golangci-lint" &&   golangci-lint run && \
echo "Starting go mod tidy" && go mod tidy && \
echo "Starting go test" && \
go test -cover  ./...


