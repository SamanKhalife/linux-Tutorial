# esdctl

The esdctl command is a command-line tool for controlling the Enlightened Sound Daemon (ESD). ESD is a sound server that can be used to play sound files on Linux and Unix-like systems.

he esdctl command takes a command as its first argument. The supported commands are:

status: Displays the status of ESD.
play: Plays a sound file.
stop: Stops a sound file.
pause: Pauses a sound file.
resume: Resumes a sound file.
volume: Sets the volume of ESD.
list-devices: Lists all of the available devices.
info: Displays information about a device.



The esdctl command is a command-line tool for controlling the Enlightened Sound Daemon (ESD). ESD is a sound server that can be used to play sound files on Linux and Unix-like systems.

The syntax for the esdctl command is:

esdctl [options] [command]
The esdctl command takes a command as its first argument. The supported commands are:

status: Displays the status of ESD.
play: Plays a sound file.
stop: Stops a sound file.
pause: Pauses a sound file.
resume: Resumes a sound file.
volume: Sets the volume of ESD.
list-devices: Lists all of the available devices.
info: Displays information about a device.
For example, to display the status of ESD, you would use the following command:

esdctl status
To play the sound file sample.wav, you would use the following command:

`esdctl play sample.wav`

To stop the sound file sample.wav, you would use the following command:

`esdctl stop sample.wav`



# help 

```
esdctl [options] [command]

Control the Enlightened Sound Daemon (ESD).

Commands:

status          Displays the status of ESD.
play            Plays a sound file.
stop            Stops a sound file.
pause           Pauses a sound file.
resume          Resumes a sound file.
volume          Sets the volume of ESD.
list-devices    Lists all of the available devices.
info            Displays information about a device.

Options:

-h, --help     Show this help message.
```
