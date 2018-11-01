@echo off
go build && %~dp0\Waziup-API.exe -crt localhost.crt -key localhost.key
