# Go Echo App PowerShell启动脚本

Write-Host "Starting Go Echo App..." -ForegroundColor Green
Write-Host ""

# 检查Go是否安装
try {
    $goVersion = go version
    Write-Host "Go version: $goVersion" -ForegroundColor Cyan
} catch {
    Write-Host "Error: Go is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please install Go from https://golang.org/dl/" -ForegroundColor Yellow
    Read-Host "Press Enter to exit"
    exit 1
}

# 下载依赖
Write-Host "Downloading dependencies..." -ForegroundColor Yellow
go mod download
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to download dependencies" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# 整理依赖
Write-Host "Tidying dependencies..." -ForegroundColor Yellow
go mod tidy
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to tidy dependencies" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# 运行应用
Write-Host "Starting server on http://localhost:8080" -ForegroundColor Green
Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Cyan
Write-Host ""
go run main.go