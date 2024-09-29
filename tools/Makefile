.PHONY: inspect

inspect:
	atlas schema inspect --env local > schema.sql

.PHONY: migrate

migrate:
	atlas migrate hash --env local

.PHONY: generate

generate:
	atlas migrate diff --env local exploration_atlas