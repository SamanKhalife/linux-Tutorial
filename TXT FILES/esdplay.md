# esdplay

The esdplay command is used to play a sound file using the Enlightened Sound Daemon (ESD). ESD is a sound server that can be used to play sound files on Linux and Unix-like systems.

For example, to play the sound file sample.wav at 44100 Hz and 16 bits per sample, you would use the following command:

`esdplay -f wav -r 44100 -b 16 sample.wav`

To play the standard input at 22050 Hz and 8 bits per sample, you would use the following command:

`esdplay -f raw -r 22050 -b 8`


# help 

```
esdplay [options] [file]

Play a sound file using the Enlightened Sound Daemon (ESD).

Options:

-f, --format=FORMAT   Set the format of the sound file.
-r, --rate=RATE       Set the sample rate.
-b, --bits=BITS       Set the bits per sample.
-l, --loop=COUNT      Set the loop count.
-h, --help            Show this help message.

Plays the specified sound file using ESD. If no file is specified, the standard input is used.

```

