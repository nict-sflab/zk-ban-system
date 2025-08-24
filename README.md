# zk-ban-system

Client and server components for zk-ban.  
This repository includes example **GM** (Group Manager) and **Verifier** servers, and **Signer** CLI tool.

## Dependencies

- Go
- GCC

## Running the Examples

### 1. Compile the circuit & set up keys

First, compile the circuit and generate the keys:

```sh
cd ./dump/
go run .
```

If you want prover/verifier keys for a different revocation-list size,
edit `./dump/main.go` before running the command above.

### 2. Example GM server

#### 2.1 Configure the Identity Provider (IdP)

By default, Firebase Authentication is used for registration(`/example/gm/main.go`):

```go
ctx := context.Background()
svc, err := NewFirebaseAuthService(ctx, "serviceAccount.json")
if err != nil {
	log.Fatal(err)
}

gmServ := serv.GMServer[string]{
	GM:        g,
	AuthToken: svc.FirebaseAuth(),
	// AuthToken: allOKAuth,
}
```

To use Firebase, obtain a `serviceAccount.json` following
the [Firebase Admin SDK setup guide](https://firebase.google.com/docs/admin/setup),
and save it to `./example/gm/`.

For local debugging, you can skip user authentication by replacing the snippet above with:

```go
gmServ := serv.GMServer[string]{
	GM:        g,
	// AuthToken: svc.FirebaseAuth(),
	AuthToken: allOKAuth,
}
```

#### 2.2 Run the GM server

```sh
cd ./example/gm/
go run .
```

This deploys the web page for admin (`http://localhost:8080/admin`).
You can revoke a signature with this page.

### 3. Example Verifier server

Run the verifier server with:

```sh
cd ./example/verifier/
go run .
```

Also, this provides the web page `https://localhost:8000/` to sign with the signer app.

### 4. Example Signer CLI

#### 4.1 Register

```sh
cd ./example/signer/
go run . join --token 'token'
```

- If you use Firebase, pass a **Firebase User ID Token** via `--token`.
- Otherwise, any arbitrary string can be used as the token for testing.

#### 4.2 Authenticate

```sh
cd ./example/signer/
go run . sign --count count --message message
```

- `count` is an integer (e.g., `0`).
- `message` is a string (e.g., `"hello"`).

#### 4.3 Update credential (one-shot)

```sh
cd ./example/signer/
go run . update
```

#### 4.4 Update credential (daemon)

```sh
cd ./example/signer/
go run . daemon
```

## Disclaimer

This is research-quality code. **Do not use in production.**
