1. You can replace ** REPLACE_WITH_ENCRYPTED_STRING ** manually or later  automate it using shell commands;

2. You can also pass different input strings or files like this:
    ### make encrypt-string STR_INPUT="MySecret"###
    ### make encrypt-files FILES=file1.txt, file2.txt###

3. Plans to extend the project
    | --- Implement AES_GCM for Authenticated Encryption --- Done|
    | --- RSA for asymmetric Encryption --- |
    | --- Hashing Using SHA-256, SHA-512 --- Done|
    | --- Adding Config File Support --- |
    | --- Encrypt/Decrypt Entire Folders --- |
    | --- Password-Derived Key Support --- Done|
    | --- Logging and Verbose Mode --- |
    | --- Output to specific Directories --- |
    | --- Implementing Digital Signatures using RSA --- |
    | --- Metadata Support --- |
    | --- Adding Tests --- |