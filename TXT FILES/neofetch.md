# neofetch

Neofetch is a command-line system information tool that displays information about your operating system, hardware, and software. It is a powerful tool that can be used to quickly and easily get a snapshot of your system's status.

Neofetch is written in Python and is available for Linux, macOS, and Windows. It is a very popular tool, and there are many different themes and plugins available for it.

Here are some of the things that neofetch can display:

* Your operating system name and version
* Your kernel version
* Your CPU model and speed
* Your RAM usage
* Your GPU model and memory
* Your current shell
* Your current directory
* Your hostname
* Your IP address
* Your timezone

Neofetch is a powerful tool that can be used for a variety of tasks. It is a valuable tool for system administrators, developers, and anyone who needs to get quick and easy system information.

Here are some examples of how neofetch can be used:

* To get a quick snapshot of your system's status, you can run the following command:

```
neofetch
```

This command will display a basic overview of your system's status.

* To get a more detailed overview of your system's status, you can use the `-c` option to specify a theme. For example, to use the "ascii" theme, you would run the following command:

```
neofetch -c ascii
```

This command will display a more detailed overview of your system's status, using the "ascii" theme.

* To get a list of all the available themes, you can run the following command:

```
neofetch -l
```

This command will display a list of all the available themes.

* To get help on using neofetch, you can run the following command:

```
neofetch -h
```

This command will display help on using neofetch.

Neofetch is a powerful tool that can be used for a variety of tasks. It is a valuable tool for system administrators, developers, and anyone who needs to get quick and easy system information.



# help

```
Usage: neofetch func_name --option "value" --option "value"

Neofetch is a CLI system information tool written in BASH. Neofetch
displays information about your system next to an image, your OS logo,
or any ASCII file of your choice.

NOTE: Every launch flag has a config option.

Options:

INFO:
    func_name                   Specify a function name (second part of info() from config) to
                                quickly display only that function's information.

                                Example: neofetch uptime --uptime_shorthand tiny

                                Example: neofetch uptime disk wm memory

                                This can be used in bars and scripts like so:

                                memory="$(neofetch memory)"; memory="${memory##*: }"

                                For multiple outputs at once (each line of info in an array):

                                IFS=$'\n' read -d "" -ra info < <(neofetch memory uptime wm)

                                info=("${info[@]##*: }")

    --disable infoname          Allows you to disable an info line from appearing
                                in the output. 'infoname' is the function name from the
                                'print_info()' function inside the config file.
                                For example: 'info "Memory" memory' would be '--disable memory'

                                NOTE: You can supply multiple args. eg. 'neofetch --disable cpu gpu'

    --title_fqdn on/off         Hide/Show Fully Qualified Domain Name in title.
    --package_managers on/off   Hide/Show Package Manager names . (on, tiny, off)
    --os_arch on/off            Hide/Show OS architecture.
    --speed_type type           Change the type of cpu speed to display.
                                Possible values: current, min, max, bios,
                                scaling_current, scaling_min, scaling_max

                                NOTE: This only supports Linux with cpufreq.
    --speed_shorthand on/off    Whether or not to show decimals in CPU speed.

                                NOTE: This flag is not supported in systems with CPU speed less than
                                1 GHz.

    --cpu_brand on/off          Enable/Disable CPU brand in output.
    --cpu_cores type            Whether or not to display the number of CPU cores
                                Possible values: logical, physical, off

                                NOTE: 'physical' doesn't work on BSD.

    --cpu_speed on/off          Hide/Show cpu speed.
    --cpu_temp C/F/off          Hide/Show cpu temperature.

                                NOTE: This only works on Linux and BSD.

                                NOTE: For FreeBSD and NetBSD-based systems, you need to enable
                                coretemp kernel module. This only supports newer Intel processors.

    --distro_shorthand on/off   Shorten the output of distro (on, tiny, off)

                                NOTE: This option won't work in Windows (Cygwin)

    --kernel_shorthand on/off   Shorten the output of kernel

                                NOTE: This option won't work in BSDs (except PacBSD and PC-BSD)

    --uptime_shorthand on/off   Shorten the output of uptime (on, tiny, off)
    --refresh_rate on/off       Whether to display the refresh rate of each monitor
                                Unsupported on Windows
    --gpu_brand on/off          Enable/Disable GPU brand in output. (AMD/NVIDIA/Intel)
    --gpu_type type             Which GPU to display. (all, dedicated, integrated)

                                NOTE: This only supports Linux.

    --de_version on/off         Show/Hide Desktop Environment version
    --gtk_shorthand on/off      Shorten output of gtk theme/icons
    --gtk2 on/off               Enable/Disable gtk2 theme/font/icons output
    --gtk3 on/off               Enable/Disable gtk3 theme/font/icons output
    --shell_path on/off         Enable/Disable showing $SHELL path
    --shell_version on/off      Enable/Disable showing $SHELL version
    --disk_show value           Which disks to display.
                                Possible values: '/', '/dev/sdXX', '/path/to/mount point'

                                NOTE: Multiple values can be given. (--disk_show '/' '/dev/sdc1')

    --disk_subtitle type        What information to append to the Disk subtitle.
                                Takes: name, mount, dir, none

                                'name' shows the disk's name (sda1, sda2, etc)

                                'mount' shows the disk's mount point (/, /mnt/Local Disk, etc)

                                'dir' shows the basename of the disks's path. (/, Local Disk, etc)

                                'none' shows only 'Disk' or the configured title.

    --disk_percent on/off       Hide/Show disk percent.

    --ip_host url               URL to query for public IP
    --ip_timeout int            Public IP timeout (in seconds).
    --song_format format        Print the song data in a specific format (see config file).
    --song_shorthand on/off     Print the Artist/Album/Title on separate lines.
    --memory_percent on/off     Display memory percentage.
    --memory_unit kib/mib/gib   Memory output unit.
    --music_player player-name  Manually specify a player to use.
                                Available values are listed in the config file

TEXT FORMATTING:
    --colors x x x x x x        Changes the text colors in this order:
                                title, @, underline, subtitle, colon, info
    --underline on/off          Enable/Disable the underline.
    --underline_char char       Character to use when underlining title
    --bold on/off               Enable/Disable bold text
    --separator string          Changes the default ':' separator to the specified string.

COLOR BLOCKS:
    --color_blocks on/off       Enable/Disable the color blocks
    --col_offset auto/num      Left-padding of color blocks
    --block_width num           Width of color blocks in spaces
    --block_height num          Height of color blocks in lines
    --block_range num num       Range of colors to print as blocks

BARS:
    --bar_char 'elapsed char' 'total char'
                                Characters to use when drawing bars.
    --bar_border on/off         Whether or not to surround the bar with '[]'
    --bar_length num            Length in spaces to make the bars.
    --bar_colors num num        Colors to make the bar.
                                Set in this order: elapsed, total
    --cpu_display mode          Bar mode.
                                Possible values: bar, infobar, barinfo, off
    --memory_display mode       Bar mode.
                                Possible values: bar, infobar, barinfo, off
    --battery_display mode      Bar mode.
                                Possible values: bar, infobar, barinfo, off
    --disk_display mode         Bar mode.
                                Possible values: bar, infobar, barinfo, off

IMAGE BACKEND:
    --backend backend           Which image backend to use.
                                Possible values: 'ascii', 'caca', 'chafa', 'jp2a', 'iterm2',
                                'off', 'sixel', 'tycat', 'w3m', 'kitty'
    --source source             Which image or ascii file to use.
                                Possible values: 'auto', 'ascii', 'wallpaper', '/path/to/img',
                                '/path/to/ascii', '/path/to/dir/', 'command output' [ascii]

    --ascii source              Shortcut to use 'ascii' backend.

                                NEW: neofetch --ascii "$(fortune | cowsay -W 30)"

    --caca source               Shortcut to use 'caca' backend.
    --chafa source              Shortcut to use 'chafa' backend.
    --iterm2 source             Shortcut to use 'iterm2' backend.
    --jp2a source               Shortcut to use 'jp2a' backend.
    --kitty source              Shortcut to use 'kitty' backend.
    --pot source                Shortcut to use 'pot' backend.
    --pixterm source            Shortcut to use 'pixterm' backend.
    --sixel source              Shortcut to use 'sixel' backend.
    --termpix source            Shortcut to use 'termpix' backend.
    --tycat source              Shortcut to use 'tycat' backend.
    --w3m source                Shortcut to use 'w3m' backend.
    --off                       Shortcut to use 'off' backend (Disable ascii art).

    NOTE: 'source; can be any of the following: 'auto', 'ascii', 'wallpaper', '/path/to/img',
    '/path/to/ascii', '/path/to/dir/'

ASCII:
    --ascii_colors x x x x x x  Colors to print the ascii art
    --ascii_distro distro       Which Distro's ascii art to print

                                NOTE: AIX, Alpine, AlterLinux, Anarchy, Android, Antergos, antiX,
                                "AOSC OS", "AOSC OS/Retro", Apricity, ArcoLinux, ArchBox,
                                ARCHlabs, ArchStrike, XFerience, ArchMerge, Arch, Artix, Arya,
                                Bedrock, Bitrig, BlackArch, BLAG, BlankOn, BlueLight, bonsai, BSD,
                                BunsenLabs, Calculate, Carbs, CentOS, Chakra, ChaletOS, Chapeau,
                                Chrom, Cleanjaro, ClearOS, Clear_Linux, Clover, Condres,
                                Container_Linux, CRUX, Cucumber, Debian, Deepin, DesaOS, Devuan,
                                DracOS, DarkOs, DragonFly, Drauger, Elementary, EndeavourOS, Endless,
                                EuroLinux, Exherbo, Fedora, Feren, FreeBSD, FreeMiNT, Frugalware,
                                Funtoo, GalliumOS, Garuda, Gentoo, Pentoo, gNewSense, GNOME, GNU,
                                GoboLinux, Grombyang, Guix, Haiku, Huayra, Hyperbola, janus, Kali,
                                KaOS, KDE_neon, Kibojoe, Kogaion, Korora, KSLinux, Kubuntu, LEDE,
                                LFS, Linux_Lite, LMDE, Lubuntu, Lunar, macos, Mageia, MagpieOS,
                                Mandriva, Manjaro, Maui, Mer, Minix, LinuxMint, MX_Linux, Namib,
                                Neptune, NetBSD, Netrunner, Nitrux, NixOS, Nurunner, NuTyX,
                                OBRevenge, OpenBSD, openEuler, OpenIndiana, openmamba, OpenMandriva,
                                OpenStage, OpenWrt, osmc, Oracle, OS Elbrus, PacBSD, Parabola,
                                Pardus, Parrot, Parsix, TrueOS, PCLinuxOS, Peppermint, popos,
                                Porteus, PostMarketOS, Proxmox, Puppy, PureOS, Qubes, Radix,
                                Raspbian, Reborn_OS, Redstar, Redcore, Redhat, Refracted_Devuan,
                                Regata, Rosa, sabotage, Sabayon, Sailfish, SalentOS, Scientific,
                                Septor, SereneLinux, SharkLinux, Siduction, Slackware, SliTaz,
                                SmartOS, Solus, Source_Mage, Sparky, Star, SteamOS, SunOS,
                                openSUSE_Leap, openSUSE_Tumbleweed, openSUSE, SwagArch, Tails,
                                Trisquel, Ubuntu-Budgie, Ubuntu-GNOME, Ubuntu-MATE, Ubuntu-Studio,
                                Ubuntu, Venom, Void, Obarun, windows10, Windows7, Xubuntu, Zorin,
                                and IRIX have ascii logos
                                NOTE: Arch, Ubuntu, Redhat, and Dragonfly have 'old' logo variants.

                                NOTE: Use '{distro name}_old' to use the old logos.

                                NOTE: Ubuntu has flavor variants.

                                NOTE: Change this to Lubuntu, Kubuntu, Xubuntu, Ubuntu-GNOME,
                                Ubuntu-Studio, Ubuntu-Mate  or Ubuntu-Budgie to use the flavors.

                                NOTE: Arcolinux, Dragonfly, Fedora, Alpine, Arch, Ubuntu,
                                CRUX, Debian, Gentoo, FreeBSD, Mac, NixOS, OpenBSD, android,
                                Antrix, CentOS, Cleanjaro, ElementaryOS, GUIX, Hyperbola,
                                Manjaro, MXLinux, NetBSD, Parabola, POP_OS, PureOS,
                                Slackware, SunOS, LinuxLite, OpenSUSE, Raspbian,
                                postmarketOS, and Void have a smaller logo variant.

                                NOTE: Use '{distro name}_small' to use the small variants.

    --ascii_bold on/off         Whether or not to bold the ascii logo.
    -L, --logo                  Hide the info text and only show the ascii logo.

IMAGE:
    --loop                      Redraw the image constantly until Ctrl+C is used. This fixes issues
                                in some terminals emulators when using image mode.
    --size 00px | --size 00%    How to size the image.
                                Possible values: auto, 00px, 00%, none
    --crop_mode mode            Which crop mode to use
                                Takes the values: normal, fit, fill
    --crop_offset value         Change the crop offset for normal mode.
                                Possible values: northwest, north, northeast,
                                west, center, east, southwest, south, southeast

    --xoffset px                How close the image will be to the left edge of the
                                window. This only works with w3m.
    --yoffset px                How close the image will be to the top edge of the
                                window. This only works with w3m.
    --bg_color color            Background color to display behind transparent image.
                                This only works with w3m.
    --gap num                   Gap between image and text.

                                NOTE: --gap can take a negative value which will move the text
                                closer to the left side.

    --clean                     Delete cached files and thumbnails.

OTHER:
OTHER:
    --config /path/to/config    Specify a path to a custom config file
    --config none               Launch the script without a config file
    --no_config                 Don't create the user config file.
    --print_config              Print the default config file to stdout.
    --stdout                    Turn off all colors and disables any ASCII/image backend.
    --help                      Print this text and exit
    --version                   Show neofetch version
    -v                          Display error messages.
    -vv                         Display a verbose log for error reporting.

DEVELOPER:
    --gen-man                   Generate a manpage for Neofetch in your PWD. (Requires GNU help2man)

```
