# go-TValerts ![tvAlerts logo](https://github.com/TheLinuxGuy/go-tvalerts/blob/master/tvAlerts-72x72-icon.png "tvAlerts")
This is an application written in Go to monitor the Teamviewer log file and send Pushover notifications immediately once a new connection or disconnection is detected.

This application may run on any platform but since I mostly use teamviewer for Windows that is the only platform I bothered to test. I may test on OSX if I have time later.

## How to use and setup go-TValerts

1. Build the binary from source or simply [download the .exe](https://github.com/TheLinuxGuy/go-tvalerts/blob/master/go-TValerts.exe)
2. edit example.json with the important data needed to run:
* **TVlogfile**: Path to Teamviewer log file, it depends on your version. For TV 12 it is C:\Program Files (x86)\TeamViewer\TeamViewer12_Logfile.log
* **LastConnLogLine**: Leave this at 0. Program will change this value every time it finds a matching new connection on TeamViewer. It uses the variable to track new alarms. 
* **LastDiscLogLine**: Leave this at 0. Similar to above but for disconnects.
* **PushoverUserKey**: Find this userkey immediately after login to pushover.net. It's called "Your user key"
* **PushoverToken**: In your Pushover.net control panel create a new application. After you hit save you will have a unique API Token/Key.

### Creating the go-TValerts Pushover application

1. Login to your pushover.net account (or create a new account if you don't have one)
2. In the control panel upon login at the bottom you should see "Your applications" with a link to create a new one
3. After following the link you will need to create an application, setup as follows:
* **Name**: tvAlerts
* **Type**: application
* **Description**: go-TValerts teamviewer log monitor
* **URL**: https://github.com/TheLinuxGuy/go-tvalerts
* **Icon**: download and use [this logo](https://github.com/TheLinuxGuy/go-tvalerts/blob/master/tvAlerts-72x72-icon.png)
* Agree to the Pushover terms of service and click create application.
* Now grab the API key for the newly created app and add it to PushoverToken in config.json (if you downloaded example.json rename it config.json)

### License

This is open source (read LICENSE.md), if you feel like contributing or adding enhancements feel free to do so. 

#### About the author
https://desantolo.com

### Screenshots
![Screenshot1](https://github.com/TheLinuxGuy/go-tvalerts/blob/master/screenshots/pushoverView.PNG)
![Screenshot2](https://github.com/TheLinuxGuy/go-tvalerts/blob/master/screenshots/notificationcenterView.PNG)
