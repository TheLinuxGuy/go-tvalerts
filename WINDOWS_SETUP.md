# Setup go-TValerts on Windows

1. Compile from source or download simply [download the .exe](https://github.com/TheLinuxGuy/go-tvalerts/blob/master/go-TValerts.exe)
2. Create a directory in your preferred location, suggestion: C:\Program Files\go-tvalerts
3. Place .exe in the new folder as well as the example.json file which you need to rename to config.json
4. Read and follow the instructions on README.md on creating a pushover application and filling in the config.json settings
5. Once config.json and the .exe are in the same folder. Try to run the .exe - you may get a Windows smartscreen warning that you will have to agree.
![windows smartscreen](https://github.com/TheLinuxGuy/go-tvalerts/blob/master/screenshots/windows_smartscreen.PNG)
6. You may get some alerts to your phone now if you ran the .exe and it found matches, assuming your API keys are all correct.
7. Create a scheduled task to run every minute to have the program parse the log file. (Google if you need instructions or [simply import my task here](https://github.com/TheLinuxGuy/go-tvalerts/blob/master/windows_scheduledTask-go-tvalerts.xml)

Voila :)