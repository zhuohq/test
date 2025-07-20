@echo off
echo Starting Go Echo App...
echo.

REM 检查Go是否安装
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo Error: Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

REM 下载依赖
echo Downloading dependencies...
go mod download
if %errorlevel% neq 0 (
    echo Error: Failed to download dependencies
    pause
    exit /b 1
)

REM 整理依赖
echo Tidying dependencies...
go mod tidy
if %errorlevel% neq 0 (
    echo Error: Failed to tidy dependencies
    pause
    exit /b 1
)

REM 运行应用
echo Starting server on http://localhost:8080
echo Press Ctrl+C to stop the server
echo.
go run main.go

pause