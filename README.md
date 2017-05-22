# go-tvalerts
Monitor Teamviewer log file on Windows and send Pushover alerts when anyone connects to your computer.

##### Where Teamviewer logs are on Windows:
C:\Program Files (x86)\TeamViewer\TeamViewer12_Logfile.log

Goals for this project:
- Writen in Golang
- Use config.json file for variables (incl log file location, pushover API info)


TODO:
- Read and update config.json after every run (to recall and avoid sending duplicate alerts)
- Function to parse thru log file, find matches for new connect or disconnects
- 