default:
  just --list

alias s := server

server:
  go run ./cmd/server/

up:
  devenv up

[working-directory: 'internal/wire']
wire:
  wire

[positional-arguments]
dump user password database file:
  mysqldump -u$1 -p$2 --databases $3 --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > $4
