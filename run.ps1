Set-ExecutionPolicy RemoteSigned


# Check if Nmap is installed 
if (!(Get-Command "nmap" -ErrorAction SilentlyContinue)) { 

    # Download the Nmap installer 
    $installerPath = "$env:TEMP\nmap-7.92-setup.exe" 
    Invoke-WebRequest -Uri "https://nmap.org/dist/nmap-7.92-setup.exe" -OutFile $installerPath 
    
    # Run the installer 
    Start-Process -FilePath $installerPath -Wait 

    # Add Nmap to the PATH environment variable 
    $nmapPath = "C:\Program Files (x86)\Nmap" 
    $env:Path += ";$nmapPath" 
    [Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::Machine) 
    
} 
    
# Test Nmap installation 
nmap -V 

# Set the current working directory to the location of the script
Set-Location $PSScriptRoot

# Build the binary from the Go source code
go build -o .\main.exe .\cmd\main.go

# Define variables
$serviceName = "GoAgent"
$displayName = "GoAgent"
$description = "My custom service"
$binaryPath = Join-Path $PSScriptRoot "main.exe"

# Check if the service exists
if (Get-Service -Name $serviceName -ErrorAction SilentlyContinue) {
    # Stop the service if it's running
    if ((Get-Service -Name $serviceName).Status -eq "Running") {
        Stop-Service -Name $serviceName
    }

    # Delete the service
    & sc.exe delete $serviceName
    Write-Host "The '$serviceName' service has been deleted."

      # The service does not exist, so create it
      Write-Host "Creating the '$displayName' service, pointing to executable '$binaryPath'"
      New-Service -Name $serviceName -BinaryPathName $binaryPath -DisplayName $displayName -Description $description -StartupType Automatic
  
      # Start the service
      Write-Host "Starting the '$displayName' service..."
      # Start-Service -Name $serviceName

}else {
        # The service does not exist, so create it
        Write-Host "Creating the '$displayName' service, pointing to executable '$binaryPath'"
        New-Service -Name $serviceName -BinaryPathName $binaryPath -DisplayName $displayName -Description $description -StartupType Automatic
    
        # Start the service
        Write-Host "Starting the '$displayName' service..."
        # Start-Service -Name $serviceName
}

