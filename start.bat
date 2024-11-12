@echo off

echo Starting the program...
cd .\spotfix\ 

:loop

echo Executing notRegistered task...
cmd /c go run . exec --notRegistered

if %errorlevel% equ 0 (
    echo notRegistered task executed successfully!
) else (
    echo Error in notRegistered task, skipping to next task...
)

timeout /t 1800

echo Executing workSchedule task...
cmd /c go run . exec --workSchedule

if %errorlevel% equ 0 (
    echo workSchedule task executed successfully!
) else (
    echo Error in workSchedule task, skipping to next task...
)

timeout /t 1800

goto loop