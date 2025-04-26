{
  pkgs,
  lib,
  config,
  inputs,
  ...
}: let
  DB_USERNAME = "root";
  DB_PASSWORD = "root1234";
  DB_HOST = "localhost";
  DB_PORT = "3306";
  DB_NAME = "shopdevgo";
  GOOSE_MIGRATION_DIR = "sql/schema";
  GOOSE_DRIVER = "mysql";
  GOOSE_DBSTRING = "${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}";
in {
  # https://devenv.sh/basics/
  env = {
    inherit GOOSE_DRIVER;
    inherit GOOSE_DBSTRING;
    inherit GOOSE_MIGRATION_DIR;
  };

  # https://devenv.sh/packages/
  packages = with pkgs; [
    iredis
    mycli
    sqlc
    goose
    jinja2-cli
  ];
  services = {
    redis.enable = true;
    mysql = {
      enable = true;
      initialDatabases = [
        {
          name = "shopdevgo";
        }
      ];
      ensureUsers = [
        {
          name = "root";
          password = "root1234";
          ensurePermissions = {"shopdevgo.*" = "ALL PRIVILEGES";};
        }
      ];
      settings = {
        mysqld = {
          sql_mode = "STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION";
        };
      };
    };
  };

  # https://devenv.sh/scripts/
  scripts.hello.exec = ''
    echo hello from $GREET
  '';

  enterShell = ''
    export PATH="$PATH:$(go env GOPATH)/bin"
  '';

  # https://devenv.sh/tasks/
  # tasks = {
  #   "myproj:setup".exec = "mytool build";
  #   "devenv:enterShell".after = [ "myproj:setup" ];
  # };

  # https://devenv.sh/tests/
  enterTest = ''
    echo "Running tests"
    git --version | grep --color=auto "${pkgs.git.version}"
  '';

  # https://devenv.sh/git-hooks/
  # git-hooks.hooks.shellcheck.enable = true;

  # See full reference at https://devenv.sh/reference/options/
}
