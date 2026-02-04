package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Connect ke DB (schema sudah dimigrasi)
	dsn := "host=localhost user=postgres password=admin dbname=hr_db_dev port=5432 sslmode=disable search_path=hr"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Init generator
	g := gen.NewGenerator(gen.Config{
		OutPath:           "~/bootcamp/quiz/challange_day_04/internal/domain/query", // Output
		ModelPkgPath:      "~/bootcamp/quiz/challange_day_04/internal/domain/models",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true, // Nullable fields jadi pointer
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// Assign DB untuk introspeksi schema
	g.UseDB(db)

	// Exclude tabel dengan strategy (return "" untuk ignore)
	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		if tableName == "schema_migrations" {
			return "" // Ignore tabel ini
		}
		return tableName // Generate yang lain
	})

	// Generate dari tabel spesifik atau semua
	//g.ApplyBasic(g.GenerateModel("users")) // Struct User dari tabel users di schema hr
	//Atau semua:
	g.ApplyBasic(g.GenerateAllTable()...) // Ini akan generate semua tabel di search_path=hr

	// Execute ke file
	g.Execute()

}
