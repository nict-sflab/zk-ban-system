{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  packages = with pkgs; [ go gcc pkg-config sqlite sqlite.dev ];
  shellHook = ''
    export CGO_ENABLED=1
    cd ./gm/
    go run -tags "debug libsqlite3" . &
    # go -run . &
    sleep 10s
    cd ../verifier
    go run -tags "debug libsqlite3" .
    cd ..
  '';
  }
