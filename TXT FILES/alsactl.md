# alsactl 

The alsactl command in Linux is used to control the Advanced Linux Sound Architecture (ALSA) sound system. It can be used to set the volume of sound cards, mute sound cards, and control other aspects of the ALSA sound system.

Here are some examples of how to use the alsactl command:

```
# To set the volume of the master channel to 50%:
alsactl set Master 50%

# To mute the sound card:
alsactl mute

# To unmute the sound card:
alsactl unmute

# To show information about the master channel:
alsactl info Master
```


# help 

```
Usage: amidi options

-h, --help                      this help
-V, --version                   print current version
-l, --list-devices              list all hardware ports
-L, --list-rawmidis             list all RawMIDI definitions
-p, --port=name                 select port by name
-s, --send=file                 send the contents of a (.syx) file
-r, --receive=file              write received data into a file
-S, --send-hex="..."            send hexadecimal bytes
-d, --dump                      print received data as hexadecimal bytes
-T, --timestamp=...             adds a timestamp in front of each dumped message
                realtime
                monotonic
                raw
-t, --timeout=seconds           exits when no data has been received
                                for the specified duration
-a, --active-sensing            include active sensing bytes
-c, --clock                     include clock bytes
-i, --sysex-interval=mseconds   delay in between each SysEx message
root@Saman:~# ^C
root@Saman:~# alsactl --help
Usage: alsactl <options> command

global options:
  -h,--help        this help
  -d,--debug       debug mode
  -v,--version     print version of this program

Available state options:
  -f,--file #      configuration file (default /var/lib/alsa/asound.state)
  -a,--config-dir #  boot / hotplug configuration directory (default /var/lib/alsa)
  -l,--lock        use file locking to serialize concurrent access
  -L,--no-lock     do not use file locking to serialize concurrent access
  -O,--lock-state-file #  state lock file path (default /var/lock/asound.state.lock)
  -F,--force       try to restore the matching controls as much as possible
                     (default mode)
  -g,--ignore      ignore 'No soundcards found' error
  -P,--pedantic    do not restore mismatching controls (old default)
  -I,--no-init-fallback  
                   don't initialize even if restore fails
  -r,--runstate #  save restore and init state to this file (only errors)
                     default settings is 'no file set'
  -R,--remove      remove runstate file at first, otherwise append errors
  -p,--period #    store period in seconds for the daemon command
  -e,--pid-file #  pathname for the process id (daemon mode)

Available init options:
  -E,--env #=#     set environment variable for init phase (NAME=VALUE)
  -i,--initfile #  main configuration file for init phase
                     (default /usr/share/alsa/init/00main)
  -b,--background  run daemon in background
  -s,--syslog      use syslog for messages
  -n,--nice #      set the process priority (see 'man nice')
  -c,--sched-idle  set the process scheduling policy to idle (SCHED_IDLE)
  -D,--ucm-defaults  execute also the UCM 'defaults' section
  -U,--no-ucm      don't init with UCM
  -X,--ucm-nodev   show UCM no device errors
```

## breakdown

```
-c, --card=CARD: This option specifies the sound card to control.
-q, --quiet: This option suppresses all output.
-v, --verbose: This option increases verbosity.
-h, --help: This option shows this help message.
-V, --version: This option prints version information.
```


