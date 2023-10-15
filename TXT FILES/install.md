# 

The "install" command in Linux is used to copy files and set their permissions. It provides a convenient way to install files from a source location to a destination location while preserving their attributes.

# help 
```
Usage: install [OPTION]... [-T] SOURCE DEST
  or:  install [OPTION]... SOURCE... DIRECTORY
  or:  install [OPTION]... -t DIRECTORY SOURCE...
  or:  install [OPTION]... -d DIRECTORY...

This install program copies files (often just compiled) into destination
locations you choose.  If you want to download and install a ready-to-use
package on a GNU/Linux system, you should instead be using a package manager
like yum(1) or apt-get(1).

In the first three forms, copy SOURCE to DEST or multiple SOURCE(s) to
the existing DIRECTORY, while setting permission modes and owner/group.
In the 4th form, create all components of the given DIRECTORY(ies).

Mandatory arguments to long options are mandatory for short options too.
      --backup[=CONTROL]  make a backup of each existing destination file
  -b                  like --backup but does not accept an argument
  -c                  (ignored)
  -C, --compare       compare each pair of source and destination files, and
                        in some cases, do not modify the destination at all
  -d, --directory     treat all arguments as directory names; create all
                        components of the specified directories
  -D                  create all leading components of DEST except the last,
                        or all components of --target-directory,
                        then copy SOURCE to DEST
  -g, --group=GROUP   set group ownership, instead of process' current group
  -m, --mode=MODE     set permission mode (as in chmod), instead of rwxr-xr-x
  -o, --owner=OWNER   set ownership (super-user only)
  -p, --preserve-timestamps   apply access/modification times of SOURCE files
                        to corresponding destination files
  -s, --strip         strip symbol tables
      --strip-program=PROGRAM  program used to strip binaries
  -S, --suffix=SUFFIX  override the usual backup suffix
  -t, --target-directory=DIRECTORY  copy all SOURCE arguments into DIRECTORY
  -T, --no-target-directory  treat DEST as a normal file
  -v, --verbose       print the name of each directory as it is created
      --preserve-context  preserve SELinux security context
  -Z                      set SELinux security context of destination
                            file and each created directory to default type
      --context[=CTX]     like -Z, or if CTX is specified then set the
                            SELinux or SMACK security context to CTX
      --help     display this help and exit
      --version  output version information and exit

The backup suffix is '~', unless set with --suffix or SIMPLE_BACKUP_SUFFIX.
The version control method may be selected via the --backup option or through
the VERSION_CONTROL environment variable.  Here are the values:

  none, off       never make backups (even if --backup is given)
  numbered, t     make numbered backups
  existing, nil   numbered if numbered backups exist, simple otherwise
  simple, never   always make simple backups
```
