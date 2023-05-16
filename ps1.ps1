Install-Module -Name 'Microsoft.PowerShell.Management'


# Check if the service already exists
$service = Get-Service "GoAgent"

# If the service exists, stop it and delete it
if ($service -ne $null) {
    Stop-Service $service.Name
    Remove-Service $service.Name
}

# Create a new service
New-Service -Name "GoAgent" -DisplayName "GoAgent" -StartupType Automatic -BinaryPath "C:\Users\Administrator\Desktop\go_windows_agent\main.exe"

# Start the service
Start-Service "GoAgent"