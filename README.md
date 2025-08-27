# zk-ban-system

Client and server components for zk-ban.  
This repository includes example **GM** (Group Manager) and **Verifier** servers, and a **Signer** CLI tool.

## Dependencies

- Go
- GCC
- Nix (optional)

## Running the Example Servers

### Using nix-shell

Run the following commands:

```sh
cd ./example
nix-shell
```

### Manually

#### 1. Compile the circuit & set up keys

First, compile the circuit and generate the keys:

```sh
cd ./dump/
go run .
```

If you want prover/verifier keys for a different revocation-list size, edit `./dump/main.go` before running the command above.

#### 2. Example GM server

If you are using Firebase, obtain a `serviceAccount.json` by following the [Firebase Admin SDK setup guide](https://firebase.google.com/docs/admin/setup), and save it to `./example/gm/`.

Then run:

```sh
cd ./example/gm/
go run .              # with Firebase
go -tags debug run .  # for debug
```

This starts an admin web page at `http://localhost:8080/admin`, where you can revoke signatures.

#### 3. Example Verifier server

Run the verifier server with:

```sh
cd ./example/verifier/
go run .
```

This also serves a web page at `https://localhost:8000/` for signing with the Signer app.

## Using the Example Signer CLI

### 1. Register

```sh
cd ./example/signer/
go run . join --token 'token'
```

- If you use Firebase, pass a **Firebase User ID Token** via `--token`.
- Otherwise, any arbitrary string can be used as the token for testing.

### 2. Authenticate

```sh
cd ./example/signer/
go run . sign --count count --message message
```

- `count` is an integer (e.g., `0`).
- `message` is a string (e.g., `"hello"`).

### 3. Update credentials (one-time)

```sh
cd ./example/signer/
go run . update
```

### 4. Update credentials (daemon mode)

```sh
cd ./example/signer/
go run . daemon
```

## Disclaimer

This is research-quality code. **Do not use in production.**
