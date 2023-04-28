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