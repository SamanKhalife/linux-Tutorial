# install


The `install` command in Linux is used to copy files and set attributes. It is a versatile tool that can also be used to create directories and set permissions. The command is commonly used in scripts and Makefiles to install programs and files to their destination directories.

#### Basic Syntax
```
install [OPTION]... SOURCE... DEST
```

- `SOURCE`: The file or files to be installed.
- `DEST`: The destination directory or file.

#### Common Options

- `-d, --directory`: Treat all arguments as directories, create them if they don't exist.
- `-m, --mode=MODE`: Set the file mode (permissions) to MODE.
- `-o, --owner=OWNER`: Set the ownership of the installed files to OWNER.
- `-g, --group=GROUP`: Set the group ownership of the installed files to GROUP.
- `-p, --preserve-timestamps`: Preserve the modification times of the files.
- `-t, --target-directory=DIRECTORY`: Specify the target directory.
- `-s, --strip`: Strip the symbol table from installed binaries.

#### Examples

1. **Basic File Installation**

   Install a file to a specific directory:
   ```sh
   install file.txt /usr/local/share/
   ```

2. **Set File Permissions**

   Install a file and set its permissions to `755`:
   ```sh
   install -m 755 script.sh /usr/local/bin/
   ```

3. **Install Multiple Files**

   Install multiple files to a directory:
   ```sh
   install file1.txt file2.txt /usr/local/share/
   ```

4. **Create Directories**

   Create a directory and set its permissions:
   ```sh
   install -d -m 755 /usr/local/newdir
   ```

5. **Install with Owner and Group**

   Install a file and set the owner and group:
   ```sh
   sudo install -o user -g group file.txt /usr/local/share/
   ```

6. **Preserve Timestamps**

   Install a file and preserve its modification time:
   ```sh
   install -p file.txt /usr/local/share/
   ```

7. **Use as Part of a Build Process**

   Often used in Makefiles to install binaries after compilation:
   ```Makefile
   install: myprogram
       install -m 755 myprogram /usr/local/bin/
   ```

#### Explanation of Options with Examples

1. **`-d, --directory`**

   Create a directory if it does not exist:
   ```sh
   install -d /usr/local/mydir
   ```

2. **`-m, --mode=MODE`**

   Set the permissions to `644` for the installed file:
   ```sh
   install -m 644 config.conf /etc/myapp/
   ```

3. **`-o, --owner=OWNER`**

   Change the owner to `john` for the installed file:
   ```sh
   sudo install -o john file.txt /usr/local/share/
   ```

4. **`-g, --group=GROUP`**

   Change the group to `staff` for the installed file:
   ```sh
   sudo install -g staff file.txt /usr/local/share/
   ```

5. **`-p, --preserve-timestamps`**

   Preserve the timestamps when installing a file:
   ```sh
   install -p file.txt /usr/local/share/
   ```

6. **`-t, --target-directory=DIRECTORY`**

   Install files to a target directory:
   ```sh
   install -t /usr/local/share/ file1.txt file2.txt
   ```

7. **`-s, --strip`**

   Strip the symbol table from an installed binary:
   ```sh
   install -s myprogram /usr/local/bin/
   ```

The `install` command is powerful and flexible, making it a crucial tool for system administrators and developers for managing file installations and permissions.
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
