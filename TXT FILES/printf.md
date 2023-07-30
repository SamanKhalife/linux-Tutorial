

# help 

```
printf: printf [-v var] format [arguments]
    Formats and prints ARGUMENTS under control of the FORMAT.
    
    Options:
      -v var    assign the output to shell variable VAR rather than
                display it on the standard output
    
    FORMAT is a character string which contains three types of objects: plain
    characters, which are simply copied to standard output; character escape
    sequences, which are converted and copied to the standard output; and
    format specifications, each of which causes printing of the next successive
    argument.
    
    In addition to the standard format specifications described in printf(1),
    printf interprets:
    
      %b        expand backslash escape sequences in the corresponding argument
      %q        quote the argument in a way that can be reused as shell input
      %Q        like %q, but apply any precision to the unquoted argument before
                quoting
      %(fmt)T   output the date-time string resulting from using FMT as a format
                string for strftime(3)
    
    The format is re-used as necessary to consume all of the arguments.  If
    there are fewer arguments than the format requires,  extra format
    specifications behave as if a zero value or null string, as appropriate,
    had been supplied.
    
    Exit Status:
    Returns success unless an invalid option is given or a write or assignment
    error occurs.
```

