data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/models",
    "--dialect", "postgres", // | postgres | sqlite | sqlserver
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  url="postgresql://postgres:postgres@localhost:5433/postgres?sslmode=disable"
  // for the migration diff command this will be used so always use the dcoker for this
  dev = "docker://postgres/15/dev"

  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}