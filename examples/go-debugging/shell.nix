let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-24.05";
  pkgs = import nixpkgs { config = {}; overlays = [
	(self: super: {
		go = super.go_1_22; # Ensures go Version 1.22
	})

  ]; };
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
