{
  description = "simulation devshell";

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
            packages = with pkgs; [ python313 ];

            shellHook = "
							python -m venv .venv
							source .venv/bin/activate
						";
          };
        };
      }
    );
}
