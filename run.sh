go mod tidy
cd ./model/migrate_scripts || exit
go run migrate_scripts.go || echo >> "migrate failed"
cd ../../ || exit

# initialize swagger2
go get -u github.com/swaggo/swag/cmd/swag
swag init --parseDependency --parseInternal

go run main.go || echo >> "start server failed"
