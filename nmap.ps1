# Check if Nmap is installed
if (!(Get-Command "nmap" -ErrorAction SilentlyContinue)) {
    # Check if Chocolatey is installed
    if (!(Get-Command "choco.exe" -ErrorAction SilentlyContinue)) {
        # Install Chocolatey
        Set-ExecutionPolicy Bypass -Scope Process -Force
        iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
    }
    
    # Install Nmap using Chocolatey
    choco install nmap -y
}