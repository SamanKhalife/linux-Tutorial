# make targets

In the context of a Makefile, **targets** are the endpoints that you want to create or update. They are typically files that are produced as a result of the build process, but they can also be actions or commands you want to execute.

### Key Concepts of Makefile Targets

1. **Explicit Targets**: These are directly defined in the Makefile and specify the files or actions to be built.
2. **Implicit Targets**: These are patterns used to define rules for building files of a certain type.
3. **Phony Targets**: These are targets that are not files but rather names for a set of commands to be run. They are often used for tasks such as cleaning up build files.

### Basic Structure of a Target

A target in a Makefile generally has the following structure:

```make
target: dependencies
    command
```

- `target`: The name of the file to be created or the action to be performed.
- `dependencies`: Files or targets that the target depends on. If any dependency is newer than the target, the target is rebuilt.
- `command`: The shell command to be executed to build the target. It must be indented with a tab.

### Example of Makefile Targets

Consider a simple project with the following files:
- `main.c`
- `utils.c`
- `utils.h`

An example Makefile with targets might look like this:

```make
# Compiler and flags
CC = gcc
CFLAGS = -Wall -g

# Targets
all: myprogram

myprogram: main.o utils.o
    $(CC) $(CFLAGS) -o myprogram main.o utils.o

main.o: main.c utils.h
    $(CC) $(CFLAGS) -c main.c

utils.o: utils.c utils.h
    $(CC) $(CFLAGS) -c utils.c

clean:
    rm -f myprogram *.o
```

### Description of Targets

- **all**: The default target which builds the final program. It's often used as a convenience target that depends on other targets.
- **myprogram**: This target creates the executable `myprogram`. It depends on `main.o` and `utils.o`.
- **main.o**: This target compiles `main.c` into an object file. It depends on `main.c` and `utils.h`.
- **utils.o**: This target compiles `utils.c` into an object file. It depends on `utils.c` and `utils.h`.
- **clean**: This is a phony target used to clean up the build directory by removing generated files.

### Phony Targets

Phony targets are not actual files but are instead labels for commands to be executed. They are declared using the `.PHONY` directive:

```make
.PHONY: clean all

clean:
    rm -f myprogram *.o
```

### Implicit Targets

Implicit targets (or pattern rules) use `%` as a wildcard to define a rule that applies to multiple files:

```make
%.o: %.c
    $(CC) $(CFLAGS) -c $< -o $@
```

This rule says that any `.o` file depends on its corresponding `.c` file, and it specifies how to build it.

### Special Built-in Targets

There are several special targets with predefined meanings:

- **`.PHONY`**: Declares phony targets.
- **`.SUFFIXES`**: Defines or clears suffixes used in suffix rules.
- **`.DEFAULT`**: Provides default commands for unknown targets.
- **`.PRECIOUS`**: Prevents deletion of intermediate files.
- **`.INTERMEDIATE`**: Marks files as intermediate.
- **`.SECONDARY`**: Similar to `.INTERMEDIATE`, but keeps the files.

### Using Targets

To build a specific target, you run:

```bash
make target
```

For example:

```bash
make myprogram
make clean
```

To build the default target (often the first target in the Makefile, usually `all`):

```bash
make
```

### Conclusion

Understanding and effectively using targets in Makefiles is crucial for automating the build process in software development. By defining clear rules and dependencies, `make` ensures efficient and accurate builds, rebuilds only what's necessary, and supports complex build scenarios.
