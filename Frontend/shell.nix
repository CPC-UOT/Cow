{ pkgs ? import <nixpkgs> { } }:

pkgs.mkShell
{
  nativeBuildInputs = with pkgs; [
    nodejs_21
  ];

  shellHook = ''
    source .env
  '';
}
