let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-23.11";
  pkgs = import nixpkgs { config = {}; overlays = []; };
in

pkgs.mkShell {
  name = "(nix)-examples-go-debugging";
  packages = with pkgs; [
    go
    zsh
  ];
  shellHook = ''
    export SHELL=$(which zsh)
    export GOBIN=$(pwd)/bin
    export PATH="$PATH:$GOBIN"
    exec $SHELL
  '';
}
