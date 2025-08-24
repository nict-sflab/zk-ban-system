{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  packages = with pkgs; [ go gcc ];
  shellHook = ''
    cd ./gm/
    go run -tags debug . &
    # go -run . &
    sleep 10s
    cd ../verifier
    go run -tags debug .
  '';
  }
