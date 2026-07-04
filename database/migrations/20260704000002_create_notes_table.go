package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cloud-compute/app/facades"
)

type M20260704000002CreateNotesTable struct{}

// Signature The unique signature for the migration.
func (r *M20260704000002CreateNotesTable) Signature() string {
	return "20260704000002_create_notes_table"
}

// Up Run the migrations.
func (r *M20260704000002CreateNotesTable) Up() error {
	if facades.Schema().HasTable("notes") {
		return nil
	}

	return facades.Schema().Create("notes", func(table schema.Blueprint) {
		table.ID()
		table.UnsignedBigInteger("user_id")
		table.String("title")
		table.Text("body")
		table.Timestamps()
		table.Index("user_id")
	})
}

// Down Reverse the migrations.
func (r *M20260704000002CreateNotesTable) Down() error {
	return facades.Schema().DropIfExists("notes")
}
