$StartTime = (Get-Date).AddHours(-1)
$EndTime = Get-Date

Get-WinEvent -FilterHashtable @{
    LogName   = 'Security'
    StartTime = $StartTime
    EndTime   = $EndTime
} |
Select-Object -Property @{Name = 'LogName'; Expression = { $_.LogName } }, @{Name = 'EventCode'; Expression = { $_.Id } }, @{Name = 'SourceName'; Expression = { $_.ProviderName } }, Message, @{Name = 'TimeGenerated'; Expression = { $_.TimeCreated } }, @{Name = 'EventType'; Expression = { $_.LevelDisplayName } } |
ConvertTo-Json -Compress |
Out-File -FilePath ".\Files.json"