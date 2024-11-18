@echo off
echo Starting the program...
cd .\spotfix\

:loop

echo Executing notRegistered task for 1 hour...

for /l %%i in (1,1,60) do (
    cmd /c go run . exec --notRegistered
    if %errorlevel% equ 0 (
        echo notRegistered task executed successfully!
    ) else (
        echo Error in notRegistered task, skipping to next task...
        goto workSchedule
    )
        timeout /t 60 > nul
)

:workSchedule
echo Executing workSchedule task for 1 hour...

for /l %%i in (1,1,60) do (
    cmd /c go run . exec --workSchedule
    if %errorlevel% equ 0 (
        echo workSchedule task executed successfully!
    ) else (
        echo Error in workSchedule task, skipping to next task...
        goto loop
    )  
        timeout /t 60 > nul
    )
    goto loop