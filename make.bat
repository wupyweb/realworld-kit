@echo off

rem Generating migrations
atlas migrate diff %1 ^
  --dir "file://ent/migrate/migrations" ^
  --to "ent://ent/schema" ^
  --dev-url "sqlite://file?mode=memory&_fk=1"

rem Applying migrations
atlas migrate apply ^
  --dir "file://ent/migrate/migrations" ^
  --url "sqlite://test.sqlite?_fk=1"
