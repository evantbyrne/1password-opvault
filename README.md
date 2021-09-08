# 1Password Opvault

Utility for exporting data from 1Password's encrypted opvault format. 1Password version 8 removed the ability to use locally-stored password vaults with their move to cloud-only storage.

## Usage

```
1password-opvault --help
```
```
Usage: 1password-opvault <command>

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  export <path>
    Export 1Password opvault.

  profiles <path>
    Show all profiles in vault.

  version
    Show version.

Run "1password-opvault <command> --help" for more information on a command.
```

### Export

Export all items from 1Password vault. Example usage with password prompt:

```
1password-opvault export /path/to/1Password.opvault/ --hide-passwords
```

Example line of output:

```
{"ainfo":"test@example.com","category":"Login","fields":[{"designation":"username","name":"login[username]","type":"T","value":"test@example.com"},{"designation":"password","name":"login[password]","type":"P","value":"******"},{"designation":"","name":"login[remember]","type":"C","value":"✓"}],"tags":[],"title":"Example","trashed":false,"urls":["https://example.com"]}
```

Data is exported in the [JSON Lines](https://jsonlines.org/) format, which has the advantage of being able to be piped to programs like `grep` for further processing. For example:

```
1password-opvault export /path/to/1Password.opvault/ --hide-passwords | grep -i "aws"
```

Options:

- Flag `--hide-passwords` Hide decrypted passwords from output.
- Flag `--profile` Specify the profile. Defaults to `default`.
- Environment variable `ONEPASSWORD_OPVAULT_PASSWORD` Set to bypass the password prompt in non-interactive environments.

### Profiles

Show all profiles in 1Password vault. Shows one profile name per line. For example:

```
1password-opvault profiles /path/to/1Password.opvault/
```
```
default
```

### Version

Show version of 1password-opvault. For example:

```
1password-opvault version
```
```
1.1.0
```

## Install

1Password Opvault may be installed directly from the source repository on a system with Go 1.16 or higher, and may be done like so:

```
go install github.com/evantbyrne/1password-opvault@latest
```
