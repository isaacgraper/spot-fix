@echo off

setlocal enabledelayedexpansion

set BOT_EXEC="C:\bot.exec"

echo Starting the bot tasks!
echo.
echo Starting the bot executable...
    
"%BOT_EXEC%" exec --notRegistered
if %errorlevel% equ 0 (
    echo Bot process completed successfully with the first instruction.
    echo.

    "%BOT_EXEC%" exec --workSchedule
    if %errorlevel% equ 0 (
        rem Add more instructions here if needed
        echo Bot process has been successfully terminated with both instructions!
        echo.
        pause
    ) else (
        echo Something went wrong with the second instruction (workSchedule)...
        echo.
        pause
        exit /b %errorlevel%
    )
) else (
    echo Something went wrong with the bot executable (notRegistered)...
    echo.
    pause
    exit /b %errorlevel%
)
pause
cls