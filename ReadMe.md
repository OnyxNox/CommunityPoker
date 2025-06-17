# Community Poker

This project serves as a learning platform for exploring Go. It is a straightforward multiplayer application that allows players to join and compete in poker matches against one another.

## Development

### Prerequisite Setup

Since `go-sqlite3` is a CGO-enabled package, you must ensure that the environment variable `CGO_ENABLED=1` is set and that a `gcc` compiler is within the shell's path.

1. Download and install [Go](https://go.dev/dl/).
1. Set the `CGO_ENABLED` environment variable to enable CGO support.
    ```shell
    go env -w CGO_ENABLED=1
    ```
1. Download and install [msys2-x86_64 (MinGW-w64)](https://github.com/msys2/msys2-installer/releases) latest release.
1. Open **MinGW-w64** and install `gcc` compiler.
    ```sh
    pacman -Syu
    pacman -S mingw-w64-x86_64-gcc
    ```
1. Add the `gcc` compiler's path (Default: `C:\msys64\mingw64\bin`) to the system's `PATH` environment variable.
1. Open a new PowerShell terminal and verify `gcc` is accessible.
   ```shell
   gcc --version
   ```

### Usage

```shell
go run ./cmd/server.go
```

## References

- [Texas Hold'em Rules](https://officialgamerules.org/game-rules/texas-holdem/)