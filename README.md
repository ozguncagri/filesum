# filesum

Filesum is file hash calculator for checking integrity of data.

Supported hashing functions (for now);

- MD5
- SHA1
- SHA256 (SHA2)
- SHA512 (SHA3)

# Usage

Calculate MD5 of file

```sh
filesum md5 <path_of_file>
```

Calculate SHA1 of file

```sh
filesum sha1 <path_of_file>
```

Calculate SHA256 (SHA2)

```sh
filesum sha256 <path_of_file>
# OR
filesum sha2 <path_of_file>
```

Calculate SHA512 (SHA3)

```sh
filesum sha512 <path_of_file>
# OR
filesum sha3 <path_of_file>
```