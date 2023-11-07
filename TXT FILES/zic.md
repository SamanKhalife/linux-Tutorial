# zic


In the context of Linux systems, "zic" stands for "zone information compiler." It is a command-line utility used to compile time zone data in the form of textual descriptions into a binary format that can be used by various software applications and libraries.

The time zone data is stored in a specific format defined by the IANA (Internet Assigned Numbers Authority) Time Zone Database. This database contains information about different time zones, their offsets from Coordinated Universal Time (UTC), daylight saving time rules, historical changes, and more.

The zic command reads the time zone data files, typically found in the "/usr/share/zoneinfo" directory, and compiles them into a binary file format understood by the system. This binary data is then used by applications and services to correctly handle time zone conversions and adjustments.

In summary, zic is a Linux command-line utility used to compile time zone data into a binary format for efficient handling of time zone conversions and adjustments by software applications.



# help

```
zic: usage is zic [ --version ] [ --help ] [ -v ] \
        [ -b {slim|fat} ] [ -d directory ] [ -l localtime ] [ -L leapseconds ] \
        [ -p posixrules ] [ -r '[@lo][/@hi]' ] [ -t localtime-link ] \
        [ filename ... ]
```

# man

```
NAME
       zic - timezone compiler

SYNOPSIS
       zic [ option ... ] [ filename ... ]

DESCRIPTION
       The  zic program reads text from the file(s) named on the command line and creates the time conversion information files
       specified in this input.  If a filename is “-”, standard input is read.

OPTIONS
       --version
              Output version information and exit.

       --help Output short usage message and exit.

       -b bloat
              Output backward-compatibility data as specified by bloat.  If bloat is fat, generate additional data entries that
              work  around  potential  bugs or incompatibilities in older software, such as software that mishandles the 64-bit
              generated data.  If bloat is slim, keep the output files small; this can help check for the bugs  and  incompati‐
              bilities.   Although the default is currently fat, this is intended to change in future zic versions, as software
              that mishandles the 64-bit data typically mishandles timestamps after the year 2038 anyway.  Also see the -r  op‐
              tion for another way to shrink output size.

       -d directory
              Create  time  conversion information files in the named directory rather than in the standard directory named be‐
              low.

       -l timezone
              Use timezone as local time.  zic will act as if the input contained a link line of the form

                   Link  timezone  localtime

       -L leapsecondfilename
              Read leap second information from the file with the given name.  If this option is not used, no leap  second  in‐
              formation appears in output files.

       -p timezone
              Use  timezone's rules when handling nonstandard TZ strings like "EET-2EEST" that lack transition rules.  zic will
              act as if the input contained a link line of the form

                   Link  timezone  posixrules

              This feature is obsolete and poorly supported.  Among other things it should not be used for timestamps after the
              year 2037, and it should not be combined with -b slim if timezone's transitions are at standard time or Universal
              Time (UT) instead of local time.

       -r [@lo][/@hi]
              Reduce the size of output files by limiting their applicability to timestamps in the range from lo (inclusive) to
              hi  (exclusive),  where  lo  and  hi  are  possibly-signed  decimal counts of seconds since the Epoch (1970-01-01
              00:00:00 UTC).  Omitted counts default to extreme values.  For example, “zic -r @0” omits data intended for nega‐
              tive  timestamps (i.e., before the Epoch), and “zic -r @0/@2147483648” outputs data intended only for nonnegative
              timestamps that fit into 31-bit signed integers.  On platforms with GNU date, “zic -r @$(date  +%s)”  omits  data
              intended for past timestamps.  Also see the -b slim option for another way to shrink output size.
      -t file
              When  creating  local  time information, put the configuration link in the named file rather than in the standard
              location.

       -v     Be more verbose, and complain about the following situations:

```
