# Sendmail emulation layer commands

The Sendmail emulation layer in Postfix allows administrators to use Postfix as a drop-in replacement for Sendmail. This emulation helps maintain compatibility with existing scripts and configurations that expect Sendmail commands and behavior. 

Here's a list of common Sendmail emulation commands provided by Postfix and how to use them:

### Sendmail Emulation Layer Commands

1. **`sendmail`**
   - **Purpose:** Emulates the Sendmail command for sending email.
   - **Usage:** 
     ```sh
     sendmail recipient@example.com
     ```
   - **Example:**
     ```sh
     echo "Subject: Test Email" | sendmail recipient@example.com
     ```
   - **Notes:** The `sendmail` command sends the email using Postfix but maintains compatibility with Sendmail syntax.

2. **`mailq`**
   - **Purpose:** Displays the mail queue, similar to `sendmail -bp`.
   - **Usage:** 
     ```sh
     mailq
     ```
   - **Example:**
     ```sh
     mailq
     ```
   - **Notes:** Shows the list of emails that are currently in the queue.

3. **`postsuper`**
   - **Purpose:** Provides commands for managing the Postfix mail queue. This is not a direct Sendmail emulation command but is often used in conjunction with `mailq` and `sendmail`.
   - **Usage:**
     ```sh
     postsuper -d ALL
     ```
   - **Example:**
     ```sh
     postsuper -d ALL
     ```
   - **Notes:** Deletes all messages in the queue. Use cautiously.

4. **`sendmail -d0.4`**
   - **Purpose:** Enables debugging output for troubleshooting. This command simulates the Sendmail debugging options.
   - **Usage:**
     ```sh
     sendmail -d0.4 recipient@example.com
     ```
   - **Example:**
     ```sh
     echo "Subject: Test Email" | sendmail -d0.4 recipient@example.com
     ```
   - **Notes:** The `-d` option specifies the debugging level. Adjust the level for more or less detail.

5. **`newaliases`**
   - **Purpose:** Rebuilds the aliases database. This command is used to update the alias mappings.
   - **Usage:**
     ```sh
     newaliases
     ```
   - **Example:**
     ```sh
     newaliases
     ```
   - **Notes:** Updates the alias database used by Postfix. Equivalent to `postmap /etc/aliases`.

6. **`mailstats`**
   - **Purpose:** Provides statistics on mail usage. This is more of a Postfix command but serves a similar purpose as Sendmail’s `mailstats`.
   - **Usage:**
     ```sh
     mailstats
     ```
   - **Example:**
     ```sh
     mailstats
     ```
   - **Notes:** Displays statistics about email delivery and system performance.

7. **`sendmail -bv`**
   - **Purpose:** Verifies the recipient's email address without sending the message.
   - **Usage:**
     ```sh
     sendmail -bv recipient@example.com
     ```
   - **Example:**
     ```sh
     sendmail -bv recipient@example.com
     ```
   - **Notes:** Checks if the recipient address is valid and shows information about its status.

8. **`sendmail -bp`**
   - **Purpose:** Displays the mail queue, similar to `mailq`.
   - **Usage:**
     ```sh
     sendmail -bp
     ```
   - **Example:**
     ```sh
     sendmail -bp
     ```
   - **Notes:** Lists messages in the mail queue. Useful for queue management.

### Compatibility Tips

1. **Aliases Database:** Postfix uses `postalias` to manage the aliases database, while Sendmail uses `newaliases`. The Postfix emulation layer can handle alias updates with `newaliases` for compatibility.

2. **Configuration:** Ensure that Postfix’s configuration files (`/etc/postfix/main.cf` and `/etc/postfix/master.cf`) are correctly set up to handle the emulation of Sendmail commands.

3. **Scripts and Automation:** Scripts written for Sendmail may use commands like `sendmail` or `mailq`. Postfix’s emulation layer allows these scripts to run without modification.

4. **Testing:** Test the emulation commands to ensure they behave as expected. Postfix’s emulation is designed to be as compatible as possible but may have slight differences from Sendmail’s behavior.

5. **Documentation:** Consult Postfix’s documentation for detailed information on compatibility and any nuances between Sendmail and Postfix.

### Conclusion

The Sendmail emulation layer in Postfix provides a way to use Postfix in environments that are accustomed to Sendmail’s command syntax and behavior. By leveraging these emulated commands, administrators can transition to Postfix while maintaining compatibility with existing tools and scripts. Regular testing and verification ensure that the emulation layer meets operational needs effectively.
