# esdsample

The esdsample command is used to play a sample file using the Enlightened Sound Daemon (ESD). ESD is a sound server that can be used to play sound files on Linux and Unix-like systems.

For example, to play the sample file sample.wav at 44100 Hz and 16 bits per sample, you would use the following command:

`esdsample -f wav -r 44100 -b 16 sample.wav`

To play the standard input at 22050 Hz and 8 bits per sample, you would use the following command:

`esdsample -f raw -r 22050 -b 8`


# help 

```
esdsample [options] [file]

Play a sample file using the Enlightened Sound Daemon (ESD).

Options:

-f, --format=FORMAT   Set the format of the sample file.
-r, --rate=RATE       Set the sample rate.
-b, --bits=BITS       Set the bits per sample.
-l, --loop=COUNT      Set the loop count.
-h, --help            Show this help message.

Plays the specified sample file using ESD. If no file is specified, the standard input is used.

```



## breakdown

```
-f, --format=FORMAT: This option specifies the format of the sample file. The supported formats are:
raw: Raw PCM audio
wav: WAV audio
-r, --rate=RATE: This option specifies the sample rate. The default sample rate is 44100 Hz.
-b, --bits=BITS: This option specifies the bits per sample. The supported bits per sample are:
8: 8 bits per sample
16: 16 bits per sample
-l, --loop=COUNT: This option specifies the loop count. The default loop count is 0, which means that the sample will not loop.
-h, --help: This option shows this help message.
```
