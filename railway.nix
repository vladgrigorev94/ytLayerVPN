{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.ffmpeg
    pkgs.python3
    pkgs.python3Packages.yt-dlp
  ];

  shellHook = ''
    export PATH=$PATH:$(go env GOPATH)/bin
  '';
}