# cc

The `cc` command in Linux is used to compile C programs. It is a command-line tool that can be used to create executable files from C source code.

The `cc` command is used in the following syntax:

```
cc [options] [source_files]
```

The `options` can be used to specify the following:

* `-c` : Compile the source files, but do not link them.
* `-o` : Output the executable file to the specified filename.
* `-Wall` : Enable all warnings.
* `-Wextra` : Enable extra warnings.
* `-g` : Generate debugging symbols.

For example, to compile the C source file `main.c` into an executable file called `my_program`, you would run the following command:

```
cc -o my_program main.c
```

This command will compile the source file `main.c` and create an executable file called `my_program`.

To enable all warnings, you would run the following command:

```
cc -Wall -o my_program main.c
```

This command will compile the source file `main.c` and create an executable file called `my_program`. It will also enable all warnings.

To enable extra warnings, you would run the following command:

```
cc -Wextra -o my_program main.c
```

This command will compile the source file `main.c` and create an executable file called `my_program`. It will also enable extra warnings.

To generate debugging symbols, you would run the following command:

```
cc -g -o my_program main.c
```

This command will compile the source file `main.c` and create an executable file called `my_program`. It will also generate debugging symbols.

The `cc` command is a powerful tool that can be used to compile C programs. It is a versatile command that can be used to create executable files for a variety of purposes.

Here are some additional things to note about the `cc` command:

* The `cc` command is part of the GNU Compiler Collection (GCC).
* The `cc` command can be used on any system that uses the Linux kernel.
* The `cc` command can be used to compile any C program that is written to the ANSI C standard.
* The `cc` command is a safe tool to use. It will not damage any files on the system.

 
# help

```
Options:
  -pass-exit-codes         Exit with highest error code from a phase.
  --help                   Display this information.
  --target-help            Display target specific command line options (including assembler and linker options).
  --help={common|optimizers|params|target|warnings|[^]{joined|separate|undocumented}}[,...].
                           Display specific types of command line options.
  (Use '-v --help' to display command line options of sub-processes).
  --version                Display compiler version information.
  -dumpspecs               Display all of the built in spec strings.
  -dumpversion             Display the version of the compiler.
  -dumpmachine             Display the compiler's target processor.
  -foffload=<targets>      Specify offloading targets.
  -print-search-dirs       Display the directories in the compiler's search path.
  -print-libgcc-file-name  Display the name of the compiler's companion library.
  -print-file-name=<lib>   Display the full path to library <lib>.
  -print-prog-name=<prog>  Display the full path to compiler component <prog>.
  -print-multiarch         Display the target's normalized GNU triplet, used as
                           a component in the library path.
  -print-multi-directory   Display the root directory for versions of libgcc.
  -print-multi-lib         Display the mapping between command line options and
                           multiple library search directories.
  -print-multi-os-directory Display the relative path to OS libraries.
  -print-sysroot           Display the target libraries directory.
  -print-sysroot-headers-suffix Display the sysroot suffix used to find headers.
  -Wa,<options>            Pass comma-separated <options> on to the assembler.
  -Wp,<options>            Pass comma-separated <options> on to the preprocessor.
  -Wl,<options>            Pass comma-separated <options> on to the linker.
  -Xassembler <arg>        Pass <arg> on to the assembler.
  -Xpreprocessor <arg>     Pass <arg> on to the preprocessor.
  -Xlinker <arg>           Pass <arg> on to the linker.
  -save-temps              Do not delete intermediate files.
  -save-temps=<arg>        Do not delete intermediate files.
  -no-canonical-prefixes   Do not canonicalize paths when building relative
                           prefixes to other gcc components.
  -pipe                    Use pipes rather than intermediate files.
  -time                    Time the execution of each subprocess.
  -specs=<file>            Override built-in specs with the contents of <file>.
  -std=<standard>          Assume that the input sources are for <standard>.
  --sysroot=<directory>    Use <directory> as the root directory for headers
                           and libraries.
  -B <directory>           Add <directory> to the compiler's search paths.
  -v                       Display the programs invoked by the compiler.
  -###                     Like -v but options quoted and commands not executed.
  -E                       Preprocess only; do not compile, assemble or link.
  -S                       Compile only; do not assemble or link.
  -c                       Compile and assemble, but do not link.
  -o <file>                Place the output into <file>.
  -pie                     Create a dynamically linked position independent
                           executable.
  -shared                  Create a shared library.
  -x <language>            Specify the language of the following input files.
                           Permissible languages include: c c++ assembler none
                           'none' means revert to the default behavior of
                           guessing the language based on the file's extension.

Options starting with -g, -f, -m, -O, -W, or --param are automatically
 passed on to the various sub-processes invoked by cc.  In order to pass
 other options on to these processes the -W<letter> options must be used.
```
