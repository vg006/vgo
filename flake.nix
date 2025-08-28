{
  description = "vgo Development Environment";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    unstable.url = "github:nixos/nixpkgs/nixos-unstable";
  };
  outputs = { self, nixpkgs, flake-utils, unstable }:
    flake-utils.lib.eachDefaultSystem (system:
     	let
    		pkgs = nixpkgs.legacyPackages.${system};
    		unstb = unstable.legacyPackages.${system};
    	in {
        devShells.default = pkgs.mkShell {
       	  packages = with pkgs; [
            unstb.go
            gopls
            delve
            golangci-lint
         	];
          NIX_SHELL_NAME="VGO";
        };
      }
    );
}
