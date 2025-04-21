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
