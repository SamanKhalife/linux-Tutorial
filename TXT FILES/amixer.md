# amixer


The amixer command in Linux is used to control the volume of sound cards. It is a powerful tool that can be used to adjust the volume of individual sound channels, as well as the overall volume of the sound card.

Here are some examples of how to use the amixer command:

```
# To show the current volume levels:
amixer

# To increase the volume of the master channel by 10%:
amixer -c 0 sset Master 10%+

# To mute the sound card:
amixer -c 0 sset Master toggle

# To show help for the amixer command:
amixer -h
```

# help 

```
Usage: amixer <options> [command]

Available options:
  -h,--help       this help
  -c,--card N     select the card
  -D,--device N   select the device, default 'default'
  -d,--debug      debug mode
  -n,--nocheck    do not perform range checking
  -v,--version    print version of this program
  -q,--quiet      be quiet
  -i,--inactive   show also inactive controls
  -a,--abstract L select abstraction level (none or basic)
  -s,--stdin      Read and execute commands from stdin sequentially
  -R,--raw-volume Use the raw value (default)
  -M,--mapped-volume Use the mapped volume
```

## breakdown

```
-c, --card=CARD: This option specifies the sound card to control.
-q, --quiet: This option suppresses all output.
-v, --verbose: This option increases verbosity.
-h, --help: This option shows this help message.
-V, --version: This option prints version information.
```

