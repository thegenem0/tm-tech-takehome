{ pkgs, ... }:

{
  packages = with pkgs; [ go ];

  scripts = {
    go-build.exec = ''
      go build -o checkout-app ./cmd/checkout
    '';

    go-test.exec = ''
      go test -v ./...
    '';

    go-coverage.exec = ''
      go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
    '';

    go-clean.exec = ''
      rm -rf ./coverage.out ./.coverage.html ./checkout-app
    '';
  };
}
