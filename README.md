# go_windows_agent
Windows OS hardware information with Go

## Need Help for improving project level
> I am new with Go language and creating windows agent, If anyone can support that will be really great help

1) Make each information with certain interval like 20 seconds
2) Call all packages in single file under main function
3) Want to send information to any API (kind of microservices or will build one app with multiple API for consume information.)

# Steps to install

## In Windows

- Pull the Zip file from the repo
- Unzip the file
- Run the following command 
`start /B path/to/the/agent`

<br>
<br>
<br>

# Run agent in the background from source code

## Move to the directory where the `main.go` exists
`cd cmd\main.go`

## Build the source file using the following command
```
go build
```

## To create the bin as a background service:
```
sc create agent-world binPath= "C:\Users\davet\OneDrive\Desktop\go_windows_agent\cmd\agent\agent.exe" start= delayed-auto DisplayName= "Agent Background Service"
sc create go-agent binPath= "C:\Users\davet\OneDrive\Desktop\go_windows_agent\go_windows_agent.exe" start= delayed-auto DisplayName= "Go Agent Background Service" 
```

## To start the service in the background
```
sc start agent-world 
```

## To delete the service
```
sc delete agent-world
```


NOTE: In case of any error please follow to the blog/article:
https://www.partitionwizard.com/clone-disk/windows-could-not-start-the-service-on-local-computer-error-1053.html
