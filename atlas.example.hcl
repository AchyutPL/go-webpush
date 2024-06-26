variable "DB_PASSWORD" {
  type= string
  default= ""
}
variable "DB_HOST" {
  type= string
  default= ""
}

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
  url="postgresql://postgres:${var.DB_PASSWORD}@${var.DB_PASSWORD}:5432/postgres"
  // for the migration diff command this will be used so always use the docker for this
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