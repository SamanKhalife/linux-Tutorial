# addr2line

The addr2line command in Linux is used to convert addresses into file names and line numbers. It can be used to debug programs by finding the source code line that corresponds to a particular address.

Here are some examples of how to use the addr2line command:

```
# To convert the address 0x400512 into a file name and line number:
addr2line 0x400512

# To convert the address 0x400512 and print the names of the functions that contain it:
addr2line -f 0x400512

# To convert the address 0x400512 and print the section name that contains it:
addr2line -s 0x400512

# To convert the address 0x400512 and use the executable file /bin/ls:
addr2line -e /bin/ls 0x400512
```

# help 

```
-a, --always-show-functions   Print the names of all functions, even if they are not in the debug information.
-f, --functions              Print the names of the functions that contain the address.
-h, --help                   Show this help message.
-i, --inlines                Print the names of the inlined functions that contain the address.
-e, --exe=FILE              Specify the executable file to use.
-s, --section=SECTION       Print the section name that contains the address.
-p, --pretty-print           Print the results in a more human-readable format.
-r, --radix=RADIX           Specify the radix to use for addresses. The default radix is 16.
-v, --verbose               Print more verbose output.
```

## breakdown

```
-a, --always-show-functions: This option prints the names of all functions, even if they are not in the debug information.
-f, --functions: This option prints the names of the functions that contain the address.
-h, --help: This option shows this help message.
-i, --inlines: This option prints the names of the inlined functions that contain the address.
-e, --exe=FILE: This option specifies the executable file to use.
-s, --section=SECTION: This option prints the section name that contains the address.
-p, --pretty-print: This option prints the results in a more human-readable format.
-r, --radix=RADIX: This option specifies the radix to use for addresses. The default radix is 16.
-v, --verbose: This option prints more verbose output.
```
