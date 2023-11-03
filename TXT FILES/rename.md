# rename





# help 

```
Usage:
    file-rename [ -h|-m|-V ] [ -v ] [ -0 ] [ -n ] [ -f ] [ -d ] [ -u [enc]]
    [ -e|-E perlexpr]*|perlexpr [ files ]

Options:
    -v, --verbose
            Verbose: print names of files successfully renamed.

    -0, --null
            Use \0 as record separator when reading from STDIN.

    -n, --nono
            No action: print names of files to be renamed, but don't rename.

    -f, --force
            Over write: allow existing files to be over-written.

    --path, --fullpath
            Rename full path: including any directory component. DEFAULT

    -d, --filename, --nopath, --nofullpath
            Do not rename directory: only rename filename component of path.

    -h, --help
            Help: print SYNOPSIS and OPTIONS.

    -m, --man
            Manual: print manual page.

    -V, --version
            Version: show version number.

    -u, --unicode [encoding]
            Treat filenames as perl (unicode) strings when running the
            user-supplied code.

            Decode/encode filenames using encoding, if present.

            encoding is optional: if omitted, the next argument should be an
            option starting with '-', for instance -e.

    -e      Expression: code to act on files name.

            May be repeated to build up code (like "perl -e"). If no -e, the
            first argument is used as code.

    -E      Statement: code to act on files name, as -e but terminated by
            ';'.

```



## breakdown

```

```
