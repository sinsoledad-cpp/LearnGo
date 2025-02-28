# open_multiple_terminals.ps1

# 打开第一个终端并执行 go run
Start-Process "cmd.exe" -ArgumentList "/K cd C:\path\to\your\directory && go run your_program1.go"

# 打开第二个终端并执行 go run
Start-Process "cmd.exe" -ArgumentList "/K cd C:\path\to\your\directory && go run your_program2.go"

# 打开第三个终端并执行 go run
Start-Process "cmd.exe" -ArgumentList "/K cd C:\path\to\your\directory && go run your_program3.go"

# 可以继续添加更多终端...
