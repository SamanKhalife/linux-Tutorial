



# help 

```
Usage: readelf <option(s)> elf-file(s)
 Display information about the contents of ELF format files
 Options are:
  -a --all               Equivalent to: -h -l -S -s -r -d -V -A -I
  -h --file-header       Display the ELF file header
  -l --program-headers   Display the program headers
     --segments          An alias for --program-headers
  -S --section-headers   Display the sections' header
     --sections          An alias for --section-headers
  -g --section-groups    Display the section groups
  -t --section-details   Display the section details
  -e --headers           Equivalent to: -h -l -S
  -s --syms              Display the symbol table
     --symbols           An alias for --syms
     --dyn-syms          Display the dynamic symbol table
     --lto-syms          Display LTO symbol tables
     --sym-base=[0|8|10|16] 
                         Force base for symbol sizes.  The options are 
                         mixed (the default), octal, decimal, hexadecimal.
  -C --demangle[=STYLE]  Decode mangled/processed symbol names
                           STYLE can be "none", "auto", "gnu-v3", "java",
                           "gnat", "dlang", "rust"
     --no-demangle       Do not demangle low-level symbol names.  (default)
     --recurse-limit     Enable a demangling recursion limit.  (default)
     --no-recurse-limit  Disable a demangling recursion limit
     -U[dlexhi] --unicode=[default|locale|escape|hex|highlight|invalid]
                         Display unicode characters as determined by the current locale
                          (default), escape sequences, "<hex sequences>", highlighted
                          escape sequences, or treat them as invalid and display as
                          "{hex sequences}"
  -n --notes             Display the core notes (if present)
  -r --relocs            Display the relocations (if present)
  -u --unwind            Display the unwind info (if present)
  -d --dynamic           Display the dynamic section (if present)
  -V --version-info      Display the version sections (if present)
  -A --arch-specific     Display architecture specific information (if any)
  -c --archive-index     Display the symbol/file index in an archive
  -D --use-dynamic       Use the dynamic section info when displaying symbols
  -L --lint|--enable-checks
                         Display warning messages for possible problems
  -x --hex-dump=<number|name>
                         Dump the contents of section <number|name> as bytes
  -p --string-dump=<number|name>
                         Dump the contents of section <number|name> as strings
  -R --relocated-dump=<number|name>
                         Dump the relocated contents of section <number|name>
  -z --decompress        Decompress section before dumping it
  -w --debug-dump[a/=abbrev, A/=addr, r/=aranges, c/=cu_index, L/=decodedline,
                  f/=frames, F/=frames-interp, g/=gdb_index, i/=info, o/=loc,
                  m/=macro, p/=pubnames, t/=pubtypes, R/=Ranges, l/=rawline,
                  s/=str, O/=str-offsets, u/=trace_abbrev, T/=trace_aranges,
                  U/=trace_info]
                         Display the contents of DWARF debug sections
  -wk --debug-dump=links Display the contents of sections that link to separate
                          debuginfo files
  -P --process-links     Display the contents of non-debug sections in separate
                          debuginfo files.  (Implies -wK)
  -wK --debug-dump=follow-links
                         Follow links to separate debug info files (default)
  -wN --debug-dump=no-follow-links
                         Do not follow links to separate debug info files
  --dwarf-depth=N        Do not display DIEs at depth N or greater
  --dwarf-start=N        Display DIEs starting at offset N
  --ctf=<number|name>    Display CTF info from section <number|name>
  --ctf-parent=<name>    Use CTF archive member <name> as the CTF parent
  --ctf-symbols=<number|name>
                         Use section <number|name> as the CTF external symtab
  --ctf-strings=<number|name>
                         Use section <number|name> as the CTF external strtab
  -I --histogram         Display histogram of bucket list lengths
  -W --wide              Allow output width to exceed 80 characters
  -T --silent-truncation If a symbol name is truncated, do not add [...] suffix
  @<file>                Read options from <file>
  -H --help              Display this information
  -v --version           Display the version number of readelf
Report bugs to <https://sourceware.org/bugzilla/>
```



## breakdown

```

```