# rpcgen






# help 

```
usage: rpcgen infile
        rpcgen [-abkCLNTM][-Dname[=value]] [-i size] [-I [-K seconds]] [-Y path] infile
        rpcgen [-c | -h | -l | -m | -t | -Sc | -Ss | -Sm] [-o outfile] [infile]
        rpcgen [-s nettype]* [-o outfile] [infile]
        rpcgen [-n netid]* [-o outfile] [infile]
options:
-a              generate all files, including samples
-b              backward compatibility mode (generates code for SunOS 4.1)
-c              generate XDR routines
-C              ANSI C mode
-Dname[=value]  define a symbol (same as #define)
-h              generate header file
-i size         size at which to start generating inline code
-I              generate code for inetd support in server (for SunOS 4.1)
-K seconds      server exits after K seconds of inactivity
-l              generate client side stubs
-L              server errors will be printed to syslog
-m              generate server side stubs
-M              generate MT-safe code
-n netid        generate server code that supports named netid
-N              supports multiple arguments and call-by-value
-o outfile      name of the output file
-s nettype      generate server code that supports named nettype
-Sc             generate sample client code that uses remote procedures
-Ss             generate sample server code that defines remote procedures
-Sm             generate makefile template 
-t              generate RPC dispatch table
-T              generate code to support RPC dispatch tables
-Y path         directory name to find C preprocessor (cpp)
-5              SysVr4 compatibility mode
--help          give this help list
--version       print program version
```
