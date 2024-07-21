# doveadm 

The `doveadm` command is a powerful administrative utility provided by Dovecot for managing and querying various aspects of the Dovecot mail server. It helps with tasks like managing mailboxes, inspecting user data, and performing administrative operations.

### Overview of `doveadm`

- **Purpose:** `doveadm` is used for performing administrative tasks on the Dovecot mail server. This includes querying mailbox status, managing mailboxes, testing authentication, and more.
- **Usage:** `doveadm` provides a range of subcommands to interact with different aspects of the Dovecot service.

### Common `doveadm` Subcommands

Here’s an overview of some commonly used `doveadm` subcommands:

1. **`doveadm auth test`**
   - **Purpose:** Test authentication mechanisms for a specific user.
   - **Usage:**
     ```sh
     doveadm auth test <username>
     ```
   - **Example:**
     ```sh
     doveadm auth test user@example.com
     ```
   - **Description:** Tests whether the user can authenticate with the current configuration.

2. **`doveadm auth status`**
   - **Purpose:** Display the status of the authentication service.
   - **Usage:**
     ```sh
     doveadm auth status
     ```

3. **`doveadm auth stats`**
   - **Purpose:** Show statistics related to authentication.
   - **Usage:**
     ```sh
     doveadm auth stats
     ```

4. **`doveadm mailbox status`**
   - **Purpose:** Display status information for mailboxes.
   - **Usage:**
     ```sh
     doveadm mailbox status -u <user>
     ```
   - **Example:**
     ```sh
     doveadm mailbox status -u user@example.com
     ```

5. **`doveadm mailbox list`**
   - **Purpose:** List all mailboxes for a specific user.
   - **Usage:**
     ```sh
     doveadm mailbox list -u <user>
     ```
   - **Example:**
     ```sh
     doveadm mailbox list -u user@example.com
     ```

6. **`doveadm sync`**
   - **Purpose:** Synchronize mailboxes between different servers.
   - **Usage:**
     ```sh
     doveadm sync -u <user>
     ```
   - **Example:**
     ```sh
     doveadm sync -u user@example.com
     ```

7. **`doveadm expunge`**
   - **Purpose:** Expunge (delete) messages from mailboxes.
   - **Usage:**
     ```sh
     doveadm expunge -u <user> <criteria>
     ```
   - **Example:**
     ```sh
     doveadm expunge -u user@example.com before 30d
     ```

8. **`doveadm flush`**
   - **Purpose:** Flush mailbox cache or index files.
   - **Usage:**
     ```sh
     doveadm flush -u <user>
     ```

9. **`doveadm auth test`**
   - **Purpose:** Test authentication for a user.
   - **Usage:**
     ```sh
     doveadm auth test user@example.com
     ```

10. **`doveadm backup`**
    - **Purpose:** Perform backup operations on mailboxes.
    - **Usage:**
      ```sh
      doveadm backup -u <user>
      ```

### Example Commands and Their Outputs

- **Testing Authentication:**
  ```sh
  doveadm auth test user@example.com
  ```
  - **Output:** Displays whether the authentication was successful or if there were errors.

- **Listing Mailboxes:**
  ```sh
  doveadm mailbox list -u user@example.com
  ```
  - **Output:** Lists all mailboxes associated with the specified user.

- **Expunging Messages:**
  ```sh
  doveadm expunge -u user@example.com before 30d
  ```
  - **Output:** Removes messages older than 30 days from the specified user's mailboxes.

### Conclusion

The `doveadm` command is a versatile tool for managing and querying Dovecot’s mail services. It provides administrators with various functionalities to ensure the proper operation and maintenance of mailboxes and user data. Mastery of `doveadm` commands can significantly enhance the management of a Dovecot mail server and simplify administrative tasks.
