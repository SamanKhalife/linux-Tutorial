# mail

The mail command is a Linux command that can be used to send and receive email. It is a simple command to use, but it can be very effective.

Here are some examples of how to use the mail command:


```
# To send an email to the address `user@example.com`:
mail user@example.com

# To send an email with the subject line `Hello!` to the address `user@example.com`:
mail -s "Hello!" user@example.com

# To send a carbon copy of the email to the address `other@example.com`:
mail -c other@example.com user@example.com

# To send a blind carbon copy of the email to the address `secret@example.com`:
mail -b secret@example.com user@example.com
```
# help 

```
mail [options] [recipient]

Send and receive mail.

Options:

-s, --subject=SUBJECT    Set the subject line.
-c, --cc=ADDRESS        Send a carbon copy to ADDRESS.
-b, --bcc=ADDRESS       Send a blind carbon copy to ADDRESS.
-h, --help               Show this help message.

For more information, see the mail man page.
```
## breakdown

```
-s, --subject=SUBJECT: This option tells mail to use the specified subject line for the email.
-c, --cc=ADDRESS: This option tells mail to send a carbon copy of the email to the specified recipient.
-b, --bcc=ADDRESS: This option tells mail to send a blind carbon copy of the email to the specified recipient.
-h, --help: This option shows this help message.
```
