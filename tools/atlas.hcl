env "local" {
  src = "file://schema.sql"
  dev = "docker://mysql/8/dev"
  url = "file://../internal/db/migrations/sql?format=golang-migrate"

  migration {
    dir = "file://../internal/db/migrations/sql"
    format = "golang-migrate"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}