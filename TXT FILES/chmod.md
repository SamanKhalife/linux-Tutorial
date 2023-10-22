# chmod

The `chmod` command in Linux is used to change the permissions of files and directories. The permissions of a file or directory determine who can read, write, and execute it.

The `chmod` command is used in the following syntax:

```
chmod [options] [permissions] [file or directory]
```

The `options` can be used to specify the following:

* `-R` : Change the permissions recursively.
* `-f` : Force the change, even if the file or directory is read-only.

The `permissions` can be specified in three ways:

* Symbolic mode: This is the most common way to specify permissions. It uses a combination of letters to represent the permissions for the owner, group, and others.
* Octal mode: This is a less common way to specify permissions. It uses a three-digit number to represent the permissions for the owner, group, and others.
* Access control list (ACL): This is a more advanced way to specify permissions. It allows you to specify permissions for individual users and groups.

For example, to change the permissions of the file `my_file.txt` to allow the owner to read and write it, and allow everyone else to read it, you would run the following command:

```
chmod 644 my_file.txt
```

This command will change the permissions of the file `my_file.txt` to 644, which in symbolic mode means that the owner has read and write permissions, and everyone else has read permissions.

To change the permissions of the directory `my_directory` recursively to allow the owner to read, write, and execute it, and allow everyone else to read and execute it, you would run the following command:

```
chmod -R 755 my_directory
```

This command will change the permissions of the directory `my_directory` and all of its subdirectories to 755, which in symbolic mode means that the owner has read, write, and execute permissions, and everyone else has read and execute permissions.

To force the change of the permissions of the file `my_file.txt` to 644, even if the file is read-only, you would run the following command:

```
chmod -f 644 my_file.txt
```

This command will change the permissions of the file `my_file.txt` to 644, even if the file is read-only.

The `chmod` command is a powerful tool that can be used to change the permissions of files and directories. It can be used to control who can access files and directories.

Here are some additional things to note about the `chmod` command:

* The `chmod` command is part of the coreutils package.
* The `chmod` command can be used on any system that uses the Linux kernel.
* The `chmod` command can be used to change the permissions of any file or directory that is supported by the Linux kernel.
* The `chmod` command is a safe tool to use. It will not damage any files or directories.

# help

```
Usage: chmod [OPTION]... MODE[,MODE]... FILE...
  or:  chmod [OPTION]... OCTAL-MODE FILE...
  or:  chmod [OPTION]... --reference=RFILE FILE...
Change the mode of each FILE to MODE.
With --reference, change the mode of each FILE to that of RFILE.

  -c, --changes          like verbose but report only when a change is made
  -f, --silent, --quiet  suppress most error messages
  -v, --verbose          output a diagnostic for every file processed
      --no-preserve-root  do not treat '/' specially (the default)
      --preserve-root    fail to operate recursively on '/'
      --reference=RFILE  use RFILE's mode instead of MODE values
  -R, --recursive        change files and directories recursively
      --help     display this help and exit
      --version  output version information and exit

Each MODE is of the form '[ugoa]*([-+=]([rwxXst]*|[ugo]))+|[-+=][0-7]+'.
```
man 
```
NAME
       chmod - change file mode bits

SYNOPSIS
       chmod [OPTION]... MODE[,MODE]... FILE...
       chmod [OPTION]... OCTAL-MODE FILE...
       chmod [OPTION]... --reference=RFILE FILE...
DESCRIPTION
       This  manual  page  documents  the GNU version of chmod.  chmod changes the file mode bits of each given file according to mode,
       which can be either a symbolic representation of changes to make, or an octal number representing the bit pattern  for  the  new
       mode bits.

       The  format of a symbolic mode is [ugoa...][[-+=][perms...]...], where perms is either zero or more letters from the set rwxXst,
       or a single letter from the set ugo.  Multiple symbolic modes can be given, separated by commas.

       A combination of the letters ugoa controls which users' access to the file will be changed: the user  who  owns  it  (u),  other
       users  in  the file's group (g), other users not in the file's group (o), or all users (a).  If none of these are given, the ef‐
       fect is as if (a) were given, but bits that are set in the umask are not affected.

       The operator + causes the selected file mode bits to be added to the existing file mode bits of each file; - causes them  to  be
       removed;  and = causes them to be added and causes unmentioned bits to be removed except that a directory's unmentioned set user
       and group ID bits are not affected.

       The letters rwxXst select file mode bits for the affected users: read (r), write (w), execute (or search for  directories)  (x),
       execute/search only if the file is a directory or already has execute permission for some user (X), set user or group ID on exe‐
       cution (s), restricted deletion flag or sticky bit (t).  Instead of one or more of these letters, you can specify exactly one of
       the  letters ugo: the permissions granted to the user who owns the file (u), the permissions granted to other users who are mem‐
       bers of the file's group (g), and the permissions granted to users that are in neither of the two preceding categories (o).

       A numeric mode is from one to four octal digits (0-7), derived by adding up the bits with values 4, 2, and  1.   Omitted  digits
       are  assumed  to  be leading zeros.  The first digit selects the set user ID (4) and set group ID (2) and restricted deletion or
       sticky (1) attributes.  The second digit selects permissions for the user who owns the file: read (4), write  (2),  and  execute
       (1); the third selects permissions for other users in the file's group, with the same values; and the fourth for other users not
       in the file's group, with the same values.

       chmod never changes the permissions of symbolic links; the chmod system call cannot change their permissions.   This  is  not  a
       problem  since  the  permissions  of symbolic links are never used.  However, for each symbolic link listed on the command line,
       chmod changes the permissions of the pointed-to file.  In contrast, chmod ignores symbolic links  encountered  during  recursive
       directory traversals.

SETUID AND SETGID BITS
       chmod  clears  the set-group-ID bit of a regular file if the file's group ID does not match the user's effective group ID or one
       of the user's supplementary group IDs, unless the user has appropriate privileges.  Additional restrictions may cause  the  set-
       user-ID  and set-group-ID bits of MODE or RFILE to be ignored.  This behavior depends on the policy and functionality of the un‐
       derlying chmod system call.  When in doubt, check the underlying system behavior.

       For directories chmod preserves set-user-ID and set-group-ID bits unless you explicitly specify otherwise.  You can set or clear
       the  bits  with symbolic modes like u+s and g-s.  To clear these bits for directories with a numeric mode requires an additional
       leading zero, or leading = like 00755 , or =755
SETUID AND SETGID BITS
       chmod  clears  the set-group-ID bit of a regular file if the file's group ID does not match the user's effective group ID or one
       of the user's supplementary group IDs, unless the user has appropriate privileges.  Additional restrictions may cause  the  set-
       user-ID  and set-group-ID bits of MODE or RFILE to be ignored.  This behavior depends on the policy and functionality of the un‐
       derlying chmod system call.  When in doubt, check the underlying system behavior.

       For directories chmod preserves set-user-ID and set-group-ID bits unless you explicitly specify otherwise.  You can set or clear
       the  bits  with symbolic modes like u+s and g-s.  To clear these bits for directories with a numeric mode requires an additional
       leading zero, or leading = like 00755 , or =755


RESTRICTED DELETION FLAG OR STICKY BIT
       The restricted deletion flag or sticky bit is a single bit, whose interpretation depends on the file type.  For directories,  it
       prevents unprivileged users from removing or renaming a file in the directory unless they own the file or the directory; this is
       called the restricted deletion flag for the directory, and is commonly found on world-writable directories like /tmp.  For regu‐
       lar  files  on  some  older systems, the bit saves the program's text image on the swap device so it will load more quickly when
       run; this is called the sticky bit.

OPTIONS
       Change the mode of each FILE to MODE.  With --reference, change the mode of each FILE to that of RFILE.

       -c, --changes
              like verbose but report only when a change is made

       -f, --silent, --quiet
              suppress most error messages

       -v, --verbose
              output a diagnostic for every file processed

       --no-preserve-root
              do not treat '/' specially (the default)

       --preserve-root
              fail to operate recursively on '/'

       --reference=RFILE
              use RFILE's mode instead of MODE values

       -R, --recursive
              change files and directories recursively

       --help display this help and exit

       --version
              output version information and exit

       Each MODE is of the form '[ugoa]*([-+=]([rwxXst]*|[ugo]))+|[-+=][0-7]+'.
```
