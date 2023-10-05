# aplay

The aplay command in Linux is used to play audio files on a sound card.

Here are some examples of how to use the aplay command:

```
# To play the audio file `mysong.wav` on the first channel at a playback rate of 44100 Hz:
aplay -c 1 -r 44100 mysong.wav

# To play the audio file `mysong.wav` in quiet mode:
aplay -q mysong.wav

# To show help for the aplay command:
aplay -h
```

# help 

```
Usage: aplay [OPTION]... [FILE]...

-h, --help              help
    --version           print current version
-l, --list-devices      list all soundcards and digital audio devices
-L, --list-pcms         list device names
-D, --device=NAME       select PCM by name
-q, --quiet             quiet mode
-t, --file-type TYPE    file type (voc, wav, raw or au)
-c, --channels=#        channels
-f, --format=FORMAT     sample format (case insensitive)
-r, --rate=#            sample rate
-d, --duration=#        interrupt after # seconds
-s, --samples=#         interrupt after # samples per channel
-M, --mmap              mmap stream
-N, --nonblock          nonblocking mode
-F, --period-time=#     distance between interrupts is # microseconds
-B, --buffer-time=#     buffer duration is # microseconds
    --period-size=#     distance between interrupts is # frames
    --buffer-size=#     buffer duration is # frames
-A, --avail-min=#       min available space for wakeup is # microseconds
-R, --start-delay=#     delay for automatic PCM start is # microseconds 
                        (relative to buffer size if <= 0)
-T, --stop-delay=#      delay for automatic PCM stop is # microseconds from xrun
-v, --verbose           show PCM structure and setup (accumulative)
-V, --vumeter=TYPE      enable VU meter (TYPE: mono or stereo)
-I, --separate-channels one file for each channel
-i, --interactive       allow interactive operation from stdin
-m, --chmap=ch1,ch2,..  Give the channel map to override or follow
    --disable-resample  disable automatic rate resample
    --disable-channels  disable automatic channel conversions
    --disable-format    disable automatic format conversions
    --disable-softvol   disable software volume control (softvol)
    --test-position     test ring buffer position
    --test-coef=#       test coefficient for ring buffer position (default 8)
                        expression for validation is: coef * (buffer_size / 2)
    --test-nowait       do not wait for ring buffer - eats whole CPU
    --max-file-time=#   start another output file when the old file has recorded
                        for this many seconds
    --process-id-file   write the process ID here
    --use-strftime      apply the strftime facility to the output file name
    --dump-hw-params    dump hw_params of the device
    --fatal-errors      treat all errors as fatal
```

## breakdown

```
-c, --channels=CHANNELS: This option specifies the number of channels to use. The default value is 1.
-r, --rate=RATE: This option specifies the playback rate. The default value is 44100.
-q, --quiet: This option suppresses all output.
-f, --format=FORMAT: This option specifies the output format. The default format is S16_LE.
-d, --device=DEVICE: This option specifies the sound card to use. The default device is /dev/dsp.
-M, --mmap: This option uses memory mapped I/O.
--list-devices: This option lists the available sound cards.
--verbose: This option makes the output more verbose.
-h, --help: This option shows this help message.
-V, --version: This option prints version information.
```
