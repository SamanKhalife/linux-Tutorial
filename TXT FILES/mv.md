# mv

The `mv` command in Unix and Linux is used to move or rename files and directories. It is a versatile command that allows you to transfer files from one location to another or change their names.

### Basic Usage

The basic syntax for the `mv` command is:

```sh
mv [options] source destination
```

- **`options`**: Command-line options to control the behavior of `mv`.
- **`source`**: The file(s) or directory(ies) to be moved or renamed.
- **`destination`**: The new location or new name.

### Examples

#### Moving a File

To move a file to a different directory:

```sh
mv file.txt /path/to/directory/
```

This command moves `file.txt` to `/path/to/directory/`.

#### Renaming a File

To rename a file:

```sh
mv oldname.txt newname.txt
```

This command renames `oldname.txt` to `newname.txt`.

#### Moving Multiple Files

To move multiple files to a directory:

```sh
mv file1.txt file2.txt /path/to/directory/
```

This command moves `file1.txt` and `file2.txt` to `/path/to/directory/`.

#### Moving a Directory

To move a directory and its contents:

```sh
mv dir1 /path/to/directory/
```

This command moves `dir1` and all its contents to `/path/to/directory/`.

### Options

#### `-i` Option: Interactive Mode

To prompt for confirmation before overwriting files:

```sh
mv -i file.txt /path/to/directory/
```

This command asks for confirmation before overwriting an existing file in the destination.

#### `-f` Option: Force Mode

To force the move without prompting (this is the default behavior if no other options are specified):

```sh
mv -f file.txt /path/to/directory/
```

This command moves `file.txt` to `/path/to/directory/`, overwriting any existing file without prompting.

#### `-n` Option: No Clobber

To avoid overwriting existing files:

```sh
mv -n file.txt /path/to/directory/
```

This command moves `file.txt` to `/path/to/directory/` only if there is no existing file with the same name.

#### `-v` Option: Verbose Mode

To display detailed information about the move process:

```sh
mv -v file.txt /path/to/directory/
```

This command outputs detailed information about the file being moved.

### Practical Use Cases

#### Organizing Files

To organize files by moving them into specific directories:

```sh
mv report1.txt report2.txt /home/user/reports/
```

This command moves `report1.txt` and `report2.txt` to the `/home/user/reports/` directory.

#### Renaming Files in Bulk

To rename a group of files with a similar pattern:

```sh
for file in *.txt; do mv "$file" "${file%.txt}.bak"; done
```

This command renames all `.txt` files in the current directory to `.bak`.

#### Moving and Renaming a Directory

To move and rename a directory:

```sh
mv /home/user/oldname /home/user/newname
```

This command moves and renames the directory `/home/user/oldname` to `/home/user/newname`.

### Safety Tips

- **Use `mv` with caution**: Moving or renaming files can lead to data loss if done incorrectly.
- **Double-check file paths**: Ensure you specify the correct source and destination paths to avoid accidental data loss.
- **Use interactive mode (`-i`)**: When in doubt, use the `-i` option to prompt for confirmation before overwriting files.

### Summary

The `mv` command is a fundamental tool for moving and renaming files and directories in Unix and Linux environments. Its various options provide flexibility for different move scenarios, such as interactive mode, force mode, and verbose mode. Understanding these options and practical use cases can help you effectively manage files and directories.
```
Usage: mv [OPTION]... [-T] SOURCE DEST
  or:  mv [OPTION]... SOURCE... DIRECTORY
  or:  mv [OPTION]... -t DIRECTORY SOURCE...
Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.

Mandatory arguments to long options are mandatory for short options too.
      --backup[=CONTROL]       make a backup of each existing destination file
  -b                           like --backup but does not accept an argument
  -f, --force                  do not prompt before overwriting
  -i, --interactive            prompt before overwrite
  -n, --no-clobber             do not overwrite an existing file
If you specify more than one of -i, -f, -n, only the final one takes effect.
      --strip-trailing-slashes  remove any trailing slashes from each SOURCE
                                 argument
  -S, --suffix=SUFFIX          override the usual backup suffix
  -t, --target-directory=DIRECTORY  move all SOURCE arguments into DIRECTORY
  -T, --no-target-directory    treat DEST as a normal file
  -u, --update                 move only when the SOURCE file is newer
                                 than the destination file or when the
                                 destination file is missing
  -v, --verbose                explain what is being done
  -Z, --context                set SELinux security context of destination
                                 file to default type
      --help     display this help and exit
      --version  output version information and exit

```


# man 
```
NAME
       mv - move (rename) files

SYNOPSIS
       mv [OPTION]... [-T] SOURCE DEST
       mv [OPTION]... SOURCE... DIRECTORY
       mv [OPTION]... -t DIRECTORY SOURCE...

DESCRIPTION
       Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.

       Mandatory arguments to long options are mandatory for short options too.

       --backup[=CONTROL]
              make a backup of each existing destination file

       -b     like --backup but does not accept an argument

       -f, --force
              do not prompt before overwriting

       -i, --interactive
              prompt before overwrite

       -n, --no-clobber
              do not overwrite an existing file

       If you specify more than one of -i, -f, -n, only the final one takes effect.

       --strip-trailing-slashes
              remove any trailing slashes from each SOURCE argument

       -S, --suffix=SUFFIX
              override the usual backup suffix

       -t, --target-directory=DIRECTORY
              move all SOURCE arguments into DIRECTORY

       -T, --no-target-directory
              treat DEST as a normal file

       -u, --update
              move only when the SOURCE file is newer than the destination file or when the destination file is missing

       -v, --verbose
              explain what is being done

       -Z, --context
              set SELinux security context of destination file to default type

       --help display this help and exit

       --version
              output version information and exit

       The  backup suffix is '~', unless set with --suffix or SIMPLE_BACKUP_SUFFIX.  The version control method may be selected via the
       --backup option or through the VERSION_CONTROL environment variable.  Here are the values:

       none, off              
       never make backups (even if --backup is given)

       numbered, t
              make numbered backups

       existing, nil
              numbered if numbered backups exist, simple otherwise

       simple, never
              always make simple backups
```
