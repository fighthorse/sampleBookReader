@echo off
echo Starting Go Service...
start /B go run interface/cmd/main.go
echo interface-Go Service started.

@echo off
echo Starting Go Service...
start /B go run admin/cmd/main.go
echo  admin-Go Service started.