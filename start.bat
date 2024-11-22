@echo off
echo Starting the program...
cd .\spotfix\

:loop

start /B cmd /c "go run . exec --notRegistered"
if %errorlevel% equ 0 (
    echo notRegistered task executed successfully!
    timeout /t 1800 > nul
    goto loop
) else (
    echo Error in notRegistered task, waiting until the next execution begin...
    goto loop
)

timeout /t 1800 > nul