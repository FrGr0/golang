@echo off

set PATH=%PATH%;C:\apps\go\bin;C:\apps\MinGW\mingw32\bin
set GOPATH=%GOPATH%;C:\FGros\DEV-GO\lib

echo compilation: %1

go build "%1"

echo execution: %~n1.exe

call "%~n1.exe"

pause