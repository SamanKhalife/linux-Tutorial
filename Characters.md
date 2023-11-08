
<table class="tg">
<tbody>
<tr>
<th class="tg-yw4l">Char.</th>
<th class="tg-yw4l">Description</th>
</tr>
<tr>
<td class="tg-yw4l">~</td>
<td class="tg-yw4l">(tilde). The path to a user's home directory location.</td>
</tr>
<tr>
<td class="tg-yw4l">-</td>
<td class="tg-yw4l">(hyphen). Go to the previously chosen directory.</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Option flag for a command or filter.</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Arithmetic operator. Minus of arithmetic operations.</td>
</tr>
<tr>
<td class="tg-yw4l">/</td>
<td class="tg-yw4l">Root directory [forward slash]. The path to root directory location.</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Filename path separator.</td>
</tr>
<tr>
<td class="tg-yw4l">\</td>
<td class="tg-yw4l">Escape [backslash]. A quoting mechanism for single characters. It preserves the literal value of the next character that follows, with the exception of newline. </td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Arithmetic operator. Divider of arithmetic operations.</td>
</tr>
<tr>
<td class="tg-yw4l">|</td>
<td class="tg-yw4l">Pipe. This is a method of chaining commands together. Connects the output (stdout) of command1 to the input (stdin) of command2. Each command reads the previous command’s output.</td>
</tr>
<tr>
<td class="tg-yw4l">|&</td>
<td class="tg-yw4l">This operator connects the output (stdout) and error (stderr) of command1 to the input (stdin) of command2.</td>
</tr>
<tr>
<td class="tg-yw4l">||</td>
<td class="tg-yw4l">The OR operator is used to chain commands. It will execute the first command then stop if successful, if not, it will proceed pass failed commands until one is successful and stop.</tr>
<tr>
<td class="tg-yw4l">&&</td>
<td class="tg-yw4l">The AND operator is used to chain commands. It will execute commands only if the first command is successful and proceed until one fails.</td>
</tr>
<tr>
<td class="tg-yw4l">;</td>
<td class="tg-yw4l">Command separator [semicolon]. Used to separate multiple commands and output all successful and failed ones.</td>
</tr>
<tr>
<td class="tg-yw4l">&</td>
<td class="tg-yw4l">Run job in background [and]. A command followed by an & will run in the background.</td>
</tr>
<tr>
<td class="tg-yw4l"> >, >>, <</td>
<td class="tg-yw4l">Redirect a command's standard output (stdout) or input (stdin) into a file.</td>
</tr>
<tr>
<td class="tg-yw4l">&>, >&</td>
<td class="tg-yw4l">Redirects a command's both standard output (stdout) and error (stderr) into a file.</td>
</tr>
<tr>
<td class="tg-yw4l"><&-</td>
<td class="tg-yw4l">Close standard input (stdin) to prevent showing from a file.</td>
</tr>
<tr>
<td class="tg-yw4l">>&-</td>
<td class="tg-yw4l">Close standard output (stdout) to prevent showing from a file.</td>
</tr>
<tr>
<td class="tg-yw4l">>|</td>
<td class="tg-yw4l">Force redirection (even if the noclobber option is set). This will forcibly overwrite an existing file.
</tr>
<tr>
<td class="tg-yw4l">"</td>
<td class="tg-yw4l">Partial quoting [double quotes]. Protects the text inside them from being split into multiple words or arguments, yet allow substitutions to occur, meaning most other special characters is usually prevented.</td>
<tr>
<td class="tg-yw4l">.</td>
<td class="tg-yw4l">Source command [period]. To evaluate commands in the current execution context. This is a bash builtin.</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">"As a component of a filename. When working with filenames, a leading dot is the prefix of a "hidden" file, a file that an ls will not normally show. </td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Character match. When matching characters, as part of a regular expression, a "dot" matches a single character.</td>
</tr>
<tr>
<td class="tg-yw4l">'</td>
<td class="tg-yw4l">Full quoting [single quotes]. Protects the text inside them so that it has a literal meaning. This is a stronger form of quoting than double quotes.</td>
</tr>
<tr>
<td class="tg-yw4l">`</td>
<td class="tg-yw4l">Command substitution [backquotes]. Assign the output of a shell command to a variable.</td>
</tr>
<tr>
<td class="tg-yw4l">#</td>
<td class="tg-yw4l">Comment [number sign]. Lines in files beginning with a # (with the exception of #!) are comments and will not be executed.</td>
</tr>
<tr>
<td class="tg-yw4l">!</td>
<td class="tg-yw4l">Reverse (or negate) [exclamation]. The ! operator inverts the exit status of the command to which it is applied. It also inverts the meaning of a test operator.</td>
</tr>
<tr>
<td class="tg-yw4l">*</td>
<td class="tg-yw4l">Wild card [asterisk]. The * character serves as a "wild card" for filename expansion in globbing By itself, it matches every filename in a given directory.
</td>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Arithmetic operator. Multiplier of arithmetic operations.</td>
</tr>
<tr>
<td class="tg-yw4l">?</td>
<td class="tg-yw4l">Wild card [question mark]. The ? character serves as a single-character "wild card" for filename expansion in globbing, as well as representing one character in an extended regular expression.</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Test operator. Within certain expressions, the ? indicates a test for a condition.</td>
</tr>
<tr>
<td class="tg-yw4l">{ }</td>
<td class="tg-yw4l">Inline group [curly brackets]. Commands inside the curly braces are treated as if they were one command.</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Placeholder for output text.</td>
</tr>
<tr>
<td class="tg-yw4l">( )</td>
<td class="tg-yw4l">Subshell group [parentheses]. Commands within are executed in a subshell (a new process) Used much like a sandbox, if a command causes side effects (like changing variables), it will have no effect on the current shell.</td>
</tr>
<tr>
<td class="tg-yw4l">[ ]</td>
<td class="tg-yw4l">Test expression [brackets]. It is part of the shell builtin test.</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Array element. Brackets set off the numbering of each element.</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">Range of characters. As part of a regular expression, brackets delineate a range of characters to match.</td>
</tr>
<tr>
<td class="tg-yw4l">[[ ]]</td>
<td class="tg-yw4l">Test/Evaluate [double brackets] a condition expression to determine whether true or false. It is more flexible than the single-bracket [ ] test.
</td>
<tr>
<td class="tg-yw4l">(( ))</td>
<td class="tg-yw4l">Arithmetic expression [double parentheses]. Characters such as +, -, *, and / are mathematical operators used for calculations.</td>
</tr>
<tr>
<td class="tg-yw4l">~+</td>
<td class="tg-yw4l">Current working directory.</td>
</tr>
<tr>
<td class="tg-yw4l">~-</td>
<td class="tg-yw4l">Previous working directory.</td>
</tr>
<tr>
<td class="tg-yw4l">:</td>
<td class="tg-yw4l">Null command [colon]. This is the shell equivalent of a "NOP" (no op, a do-nothing operation). It may be considered a synonym for the shell builtin true.</td>
</tr>
<tr>
<td class="tg-yw4l">;;</td>
<td class="tg-yw4l">Terminator [double semicolon]. Only used in case constructs to indicate the end of an alternative.</td>
</tr>
<tr>
<td class="tg-yw4l">" "</td>
<td class="tg-yw4l">Whitespace. This is a tab, newline, vertical tab, form feed, carriage return, or space. Bash uses whitespace to determine where words begin and end.</td>
</tr>
<tr>
<td class="tg-yw4l">,, ,</td>
<td class="tg-yw4l">Lowercase conversion in parameter substitution.</td>
</tr>
<tr>
<td class="tg-yw4l">^, ^^</td>
<td class="tg-yw4l">Uppercase conversion in parameter substitution.</td>
</tr>
<tr>
<td class="tg-yw4l">$</td>
<td class="tg-yw4l">Variable substitution. A $ prefixing a variable name indicates the value the variable holds</td>
</tr>
<tr>
<td class="tg-yw4l">└─></td>
<td class="tg-yw4l">End-of-line. In a regular expression, a "$" addresses the end of a line of text.</td>
</tr>
<tr>
<td class="tg-yw4l">$*</td>
<td class="tg-yw4l">All the arguments are seen as a single word.</td>
</tr>
<tr>
<td class="tg-yw4l">!!</td>
<td class="tg-yw4l">The previous command.</td>
</tr>
<tr>
<td class="tg-yw4l">!$</td>
<td class="tg-yw4l">The last argument to the previous command.</td>
</tr>
<tr>
<td class="tg-yw4l">!*</td>
<td class="tg-yw4l">All the arguments from the previous command.</td>
</tr>
<tr>
<td class="tg-yw4l">$?</td>
<td class="tg-yw4l">The exit status of the last command executed.</td>
</tr>
<tr>
<td class="tg-yw4l">$#</td>
<td class="tg-yw4l">The number of arguments supplied to a script.</td>
</tr>
<tr>
<td class="tg-yw4l">$$</td>
<td class="tg-yw4l">The process number of the current shell. For shell scripts, this is the process ID under which they are executing.</td>
</tr>
<tr>
<td class="tg-yw4l">$!</td>
<td class="tg-yw4l">The process number of the last background command.</td>
</tr>
<tr>
<td class="tg-yw4l">$_</td>
<td class="tg-yw4l">Special variable set to final argument of previous command executed.</td>
</tr>
<tr>
<td class="tg-yw4l">$-</td>
<td class="tg-yw4l">Expands to the current option flags as specified upon invocation, by the  set  builtin command, or those set by the shell itself (such as the -i option).
</td>
<tr>
<td class="tg-yw4l">$0</td>
<td class="tg-yw4l">Used to reference the name of the current shell or current shell script.</td>
</tr>
<tr>
<td class="tg-yw4l">$n</td>
<td class="tg-yw4l">These variables correspond to the arguments with which a script was invoked. Here n is a positive decimal number corresponding to the position of an argument (the first argument is $1, the second argument is $2, and so on).</td>
</tr>
</tbody>
</table>
