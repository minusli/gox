cd `dirname $0`
rm cover.html
go test -coverprofile=cover.data ./...
go tool cover -html=cover.data -o cover.html
rm cover.data