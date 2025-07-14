{ pkgs, ... }:

{
  packages = with pkgs; [ go ];

  scripts = {
    go-test.exec = ''
      go test -v ./...
    '';

    go-coverage.exec = ''
      go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
    '';
  };
}
