# amidi

The amidi command is used to send and receive MIDI messages in Linux. MIDI is a standard used to control digital instruments such as electronic keyboards and sequencers.

Here are some examples of the amidi command:

- Set the volume of channel 10 to 127.

`amidi -c 10 -p 127`

- Play the note C4 on channel 10.

`amidi -c 10 -n 60`

- Play the note C4 on channel 10 for 1 second.

`amidi -c 10 -n 60 -d 100`


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
```
## breakdown

```
-c, --channel=CHANNEL: This option specifies the MIDI channel to send or receive messages on. The default channel is 0.
-p, --pitch=PITCH: This option specifies the pitch of the note to play. The pitch is a value from 0 to 127.
-v, --velocity=VELOCITY: This option specifies the velocity of the note to play. The velocity is a value from 0 to 127.
-d, --duration=DURATION: This option specifies the duration of the note to play in milliseconds. The duration can be a value from 0 to 65535.
-r, --raw: This option sends the MIDI message as a raw byte stream, without any header.
-m, --midi-name=NAME: This option sends the MIDI message with the specified name. The name is a string that identifies the type of MIDI message.
-q, --quiet: This option suppresses all output.
-h, --help: This option shows this help message.
-V, --version: This option prints version information.
```
