# Check specific port and answer health check.

it supports to run as daemon (windows service.

Use the following library.  
https://github.com/kardianos/service


## Build.

```
GOOS=windows GOARCH=386 go build health.go
```

## Usage

how to install,start,stop,uninstall

```
PS C:\> .\health.exe install
PS C:\> .\health.exe start
PS C:\> .\health.exe stop
PS C:\> .\health.exe uninstall
```
