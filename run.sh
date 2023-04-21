go build -o bert cmd/web/*.go
./bert -dbname=postgres -dbuser=postgres -dbpass=bert -production=false -cache=false 