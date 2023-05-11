# Start Golang code in the background
Start-Process -FilePath ".\main.exe" -ArgumentList "arguments" -RedirectStandardOutput "output.log" -RedirectStandardError "error.log" -NoNewWindow
