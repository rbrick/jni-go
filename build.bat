@echo off
cd go/
REM You'll need to have a JDK installation that doesn't have spaces in the paths.
REM With the default installer it installs in "C:\Program Files" and Windows and/or GCC doesn't like that...
REM For this, i just moved the needed include files to go/ directory.
REM Improper...but it works.
REM set CGO_CFLAGS=-I/c/JDK_Includes/include/ -I/c/JDK_Includes/include/win32/
go build -buildmode=c-shared -o native-bukkit.dll
move /Y native-bukkit.dll ..
cd ..