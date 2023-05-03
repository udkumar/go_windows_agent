 # Check if Chocolatey is installed
 if (!(Get-Command "choco.exe" -ErrorAction SilentlyContinue)) {
    # Install Chocolatey
    Set-ExecutionPolicy Bypass -Scope Process -Force
    iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
}

# Install Nmap using Chocolatey
choco install nmap -y --force

# Test Nmap installation 
nmap -V 

# Set the current working directory to the location of the script
Set-Location $PSScriptRoot

# Check if main.exe exists in the current directory
if (Test-Path -Path ".\main.exe") {
    # If it does, ask the user if they want to rebuild
    $buildResponse = Read-Host "main.exe already exists. Do you want to rebuild? (y/n)"
    if ($buildResponse -eq "y" -or $buildResponse -eq "Y") {
        # If the user wants to rebuild, run the go build command
        go build -o .\main.exe .\cmd\main.go
    } else {
        # If the user doesn't want to rebuild, skip the go build command
        Write-Host "Skipping build."
    }
} else {
    # If main.exe doesn't exist, run the go build command
    go build -o .\main.exe .\cmd\main.go
}

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

