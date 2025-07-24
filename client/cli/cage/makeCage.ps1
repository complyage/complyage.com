$Env:GOOS   = "windows"
$Env:GOARCH = "amd64"

go build -o cage.exe .
Copy-Item .\cage.exe ..\..\ -Force
