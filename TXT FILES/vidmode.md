# vidmode

The vidmode command in Linux is used to display the current video mode and to change the video mode.

For example, the following command will display the current video mode:

`vidmode`

The following command will change the video mode to 1024x768x16:

`vidmode -h 1024 -w 768 -b 16`

# help 

```
vidmode [options]

Display or change the video mode.

Options:

-h, --height=HEIGHT   Set the height of the screen.
-w, --width=WIDTH   Set the width of the screen.
-b, --bits=DEPTH   Set the color depth.
-r, --refresh=RATE   Set the refresh rate.
-f, --full-screen   Set the screen to full-screen mode.
-s, --save   Save the current video mode.
-l, --list   List the available video modes.

Examples:

    vidmode
    vidmode -h 1024 -w 768 -b 16
    vidmode -s
    vidmode -l
```



## breakdown

```
-h, --height=HEIGHT: This option sets the height of the screen.
-w, --width=WIDTH: This option sets the width of the screen.
-b, --bits=DEPTH: This option sets the color depth.
-r, --refresh=RATE: This option sets the refresh rate.
-f, --full-screen: This option sets the screen to full-screen mode.
-s, --save: This option saves the current video mode.
-l, --list: This option lists the available video modes.
```
