# SQL-Judge-Go-Service
For build LINUX: $Env:GOOS = 'linux' ; $Env:GOARCH = 'amd64' ; go build cmd/main.go
For build Windows: $Env:GOOS = 'windows' ; $Env:GOARCH = 'amd64' ; go build cmd/main.go

For run LINUX: nohup ./main &
For run Windows: main.exe

Need folders: configs/congif.yml
              db/test_db/
