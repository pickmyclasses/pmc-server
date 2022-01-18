go mod tidy
cd ./model/migrate || exit
go run migrate.go || echo >> "migrate failed"
cd ../../ || exit
go run main.go || echo >> "start server failed"
