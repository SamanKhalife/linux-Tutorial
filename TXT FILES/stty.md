# 

The "stty" command in Linux is used to change and display terminal settings. It allows you to control various aspects of your terminal, such as line settings, input and output processing, control characters, and more.

# help 

```
Usage: stty [-F DEVICE | --file=DEVICE] [SETTING]...
  or:  stty [-F DEVICE | --file=DEVICE] [-a|--all]
  or:  stty [-F DEVICE | --file=DEVICE] [-g|--save]
Print or change terminal characteristics.

Mandatory arguments to long options are mandatory for short options too.
  -a, --all          print all current settings in human-readable form
  -g, --save         print all current settings in a stty-readable form
  -F, --file=DEVICE  open and use the specified DEVICE instead of stdin
      --help     display this help and exit
      --version  output version information and exit

Optional - before SETTING indicates negation.  An * marks non-POSIX
settings.  The underlying system defines which settings are available.

Special characters:
 * discard CHAR  CHAR will toggle discarding of output
   eof CHAR      CHAR will send an end of file (terminate the input)
   eol CHAR      CHAR will end the line
 * eol2 CHAR     alternate CHAR for ending the line
   erase CHAR    CHAR will erase the last character typed
   intr CHAR     CHAR will send an interrupt signal
   kill CHAR     CHAR will erase the current line
 * lnext CHAR    CHAR will enter the next character quoted
   quit CHAR     CHAR will send a quit signal
 * rprnt CHAR    CHAR will redraw the current line
   start CHAR    CHAR will restart the output after stopping it
   stop CHAR     CHAR will stop the output
   susp CHAR     CHAR will send a terminal stop signal
 * swtch CHAR    CHAR will switch to a different shell layer
 * werase CHAR   CHAR will erase the last word typed

Special settings:
   N             set the input and output speeds to N bauds
 * cols N        tell the kernel that the terminal has N columns
 * columns N     same as cols N
 * [-]drain      wait for transmission before applying settings (on by default)
   ispeed N      set the input speed to N
 * line N        use line discipline N
   min N         with -icanon, set N characters minimum for a completed read
   ospeed N      set the output speed to N
 * rows N        tell the kernel that the terminal has N rows
 * size          print the number of rows and columns according to the kernel
   speed         print the terminal speed
   time N        with -icanon, set read timeout of N tenths of a second

Control settings:
   [-]clocal     disable modem control signals
   [-]cread      allow input to be received
 * [-]crtscts    enable RTS/CTS handshaking
   csN           set character size to N bits, N in [5..8]
   [-]cstopb     use two stop bits per character (one with '-')
   [-]hup        send a hangup signal when the last process closes the tty
   [-]hupcl      same as [-]hup
   [-]parenb     generate parity bit in output and expect parity bit in input
   [-]parodd     set odd parity (or even parity with '-')
 * [-]cmspar     use "stick" (mark/space) parity

Input settings:
   [-]brkint     breaks cause an interrupt signal
   [-]icrnl      translate carriage return to newline
   [-]ignbrk     ignore break characters
   [-]igncr      ignore carriage return
   [-]ignpar     ignore characters with parity errors
 * [-]imaxbel    beep and do not flush a full input buffer on a character
   [-]inlcr      translate newline to carriage return
   [-]inpck      enable input parity checking
   [-]istrip     clear high (8th) bit of input characters
 * [-]iutf8      assume input characters are UTF-8 encoded
 * [-]iuclc      translate uppercase characters to lowercase
 * [-]ixany      let any character restart output, not only start character
   [-]ixoff      enable sending of start/stop characters
   [-]ixon       enable XON/XOFF flow control
   [-]parmrk     mark parity errors (with a 255-0-character sequence)
   [-]tandem     same as [-]ixoff

Output settings:
 * bsN           backspace delay style, N in [0..1]
 * crN           carriage return delay style, N in [0..3]
 * ffN           form feed delay style, N in [0..1]
 * nlN           newline delay style, N in [0..1]
 * [-]ocrnl      translate carriage return to newline
 * [-]ofdel      use delete characters for fill instead of NUL characters
 * [-]ofill      use fill (padding) characters instead of timing for delays
 * [-]olcuc      translate lowercase characters to uppercase
 * [-]onlcr      translate newline to carriage return-newline
 * [-]onlret     newline performs a carriage return
 * [-]onocr      do not print carriage returns in the first column
   [-]opost      postprocess output
 * tabN          horizontal tab delay style, N in [0..3]
 * tabs          same as tab0
 * -tabs         same as tab3
 * vtN           vertical tab delay style, N in [0..1]

Local settings:
   [-]crterase   echo erase characters as backspace-space-backspace
 * crtkill       kill all line by obeying the echoprt and echoe settings
 * -crtkill      kill all line by obeying the echoctl and echok settings
 * [-]ctlecho    echo control characters in hat notation ('^c')
   [-]echo       echo input characters
 * [-]echoctl    same as [-]ctlecho
   [-]echoe      same as [-]crterase
   [-]echok      echo a newline after a kill character
 * [-]echoke     same as [-]crtkill
   [-]echonl     echo newline even if not echoing other characters
 * [-]echoprt    echo erased characters backward, between '\' and '/'
 * [-]extproc    enable "LINEMODE"; useful with high latency links
 * [-]flusho     discard output
   [-]icanon     enable special characters: erase, kill, werase, rprnt
   [-]iexten     enable non-POSIX special characters
   [-]isig       enable interrupt, quit, and suspend special characters
   [-]noflsh     disable flushing after interrupt and quit special characters
 * [-]prterase   same as [-]echoprt
 * [-]tostop     stop background jobs that try to write to the terminal
 * [-]xcase      with icanon, escape with '\' for uppercase characters

Combination settings:
 * [-]LCASE      same as [-]lcase
   cbreak        same as -icanon
   -cbreak       same as icanon
   cooked        same as brkint ignpar istrip icrnl ixon opost isig
                 icanon, eof and eol characters to their default values
   -cooked       same as raw
   crt           same as echoe echoctl echoke
   dec           same as echoe echoctl echoke -ixany intr ^c erase 0177
                 kill ^u
 * [-]decctlq    same as [-]ixany
   ek            erase and kill characters to their default values
   evenp         same as parenb -parodd cs7
   -evenp        same as -parenb cs8
 * [-]lcase      same as xcase iuclc olcuc
   litout        same as -parenb -istrip -opost cs8
   -litout       same as parenb istrip opost cs7
   nl            same as -icrnl -onlcr
   -nl           same as icrnl -inlcr -igncr onlcr -ocrnl -onlret
   oddp          same as parenb parodd cs7
   -oddp         same as -parenb cs8
   [-]parity     same as [-]evenp
   pass8         same as -parenb -istrip cs8
   -pass8        same as parenb istrip cs7
   raw           same as -ignbrk -brkint -ignpar -parmrk -inpck -istrip
                 -inlcr -igncr -icrnl -ixon -ixoff -icanon -opost
                 -isig -iuclc -ixany -imaxbel -xcase min 1 time 0
   -raw          same as cooked
   sane          same as cread -ignbrk brkint -inlcr -igncr icrnl
                 icanon iexten echo echoe echok -echonl -noflsh
                 -ixoff -iutf8 -iuclc -ixany imaxbel -xcase -olcuc -ocrnl
                 opost -ofill onlcr -onocr -onlret nl0 cr0 tab0 bs0 vt0 ff0
                 isig -tostop -ofdel -echoprt echoctl echoke -extproc -flusho,
                 all special characters to their default values

```

