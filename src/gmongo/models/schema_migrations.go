package models

type SchemaMigrations struct {
	Version string `orm:"column(version);size(255)"`
}
