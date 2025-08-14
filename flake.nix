{
  description = "orbtronics-l1-mechatronics-engineer";

  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells = {
          default = pkgs.mkShell {
            packages = with pkgs; [
              # formatters
              black
              isort
              nixfmt-rfc-style
              prettier

              # runtimes & compilers
              python313

              air
              go

              # dev-tools
              atac
              openapi-tui
              posting
              treefmt
            ];
          };
        };
      }
    );
}
