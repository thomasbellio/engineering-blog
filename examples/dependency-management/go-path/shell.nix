let
    pkgs = import (builtins.fetchTarball {
        url = "https://github.com/NixOS/nixpkgs/archive/77294205ac81810f333e25da2eb876d348fd7edc.tar.gz";
    }) {};

    myPkg = pkgs.go;
in

# Define the shell environment using stdenv.mkDerivation
pkgs.stdenv.mkDerivation {
  name = "go-env-go-path";
  buildInputs = [ myPkg ];
}
