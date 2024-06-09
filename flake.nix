{
  description = "a signaling server for cuttlefish";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }: 
  let 
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in
  {

    devShells.${system}.default = pkgs.mkShell {
      buildInputs = [
        pkgs.neovim
        pkgs.lunarvim
        pkgs.lolcat
        pkgs.gopls
        pkgs.go
      ];
      shellHook = ''
        echo entering auth-cuttlefish build environement | lolcat
      '';
    };

  };
}
