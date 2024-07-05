# mkdir

The `mkdir` command in Unix and Linux is used to create directories (folders) within the file system. It is a straightforward command for creating new directories at specified locations.

### Basic Usage

The basic syntax for the `mkdir` command is:

```sh
mkdir [options] directory...
```

- **`options`**: Command-line options to control the behavior of `mkdir`.
- **`directory`**: The name(s) of the directory(ies) to be created.

### Examples

#### Creating a Single Directory

To create a single directory:

```sh
mkdir mydir
```

This command creates a directory named `mydir` in the current working directory.

#### Creating Multiple Directories

To create multiple directories at once:

```sh
mkdir dir1 dir2 dir3
```

This command creates directories `dir1`, `dir2`, and `dir3` in the current working directory.

#### Creating Nested Directories

To create nested directories (directories within directories):

```sh
mkdir -p parentdir/subdir
```

This command creates a directory named `parentdir` if it doesn't exist, and within it, creates a directory named `subdir`.

### Options

#### `-p` Option: Create Parent Directories

To create parent directories as needed:

```sh
mkdir -p path/to/parent/dir/newdir
```

This command creates the entire directory structure `path/to/parent/dir/` if it doesn't exist, and then creates `newdir` within it.

#### `-m` Option: Set Directory Permissions

To set permissions for the newly created directory:

```sh
mkdir -m 755 newdir
```

This command creates `newdir` with permissions set to `755` (read, write, execute for owner, read and execute for group and others).

### Practical Use Cases

#### Organizing Project Files

To create directories for organizing project files:

```sh
mkdir project
cd project
mkdir src docs tests
```

This sequence of commands creates a `project` directory and within it, creates `src`, `docs`, and `tests` directories.

#### Creating Temporary Workspaces

To create temporary workspaces for specific tasks:

```sh
mkdir workspaces/task1
mkdir workspaces/task2
```

This command creates separate directories for different tasks within a `workspaces` directory.

### Summary

The `mkdir` command is a simple yet essential tool for creating directories in Unix and Linux environments. Its ability to create nested directories (`-p` option) and set permissions (`-m` option) provides flexibility for various use cases, from organizing files to creating temporary workspaces. Understanding these options and practical use cases can help you efficiently manage directory structures in your file system.

# help

```
Usage: mkdir [OPTION]... DIRECTORY...
Create the DIRECTORY(ies), if they do not already exist.

Mandatory arguments to long options are mandatory for short options too.
  -m, --mode=MODE   set file mode (as in chmod), not a=rwx - umask
  -p, --parents     no error if existing, make parent directories as needed
  -v, --verbose     print a message for each created directory
  -Z                   set SELinux security context of each created directory
                         to the default type
      --context[=CTX]  like -Z, or if CTX is specified then set the SELinux
                         or SMACK security context to CTX
      --help     display this help and exit
      --version  output version information and exit
```




