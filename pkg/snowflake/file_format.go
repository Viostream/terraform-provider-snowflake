package snowflake

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func FileFormat(name string) *Builder {
	return &Builder{
		entityType: FileFormatType,
		name:       name,
	}
}

type fileFormat struct {
	BinaryFormat              sql.NullString `db:"BinaryFormat"`
	Comment                   sql.NullString `db:"comment"`
	Compression               sql.NullString `db:"compression"`
	DateFormat                sql.NullString `db:"date_format"`
	Escape                    sql.NullString `db:"escape"`
	EscapeUnenclosedFields    sql.NullString `db:"escape_unenclosed_fields"`
	FieldDelimiter            sql.NullString `db:"field_delimiter"`
	FieldOptionallyEnclosedBy sql.NullString `db:"field_optionally_enclosed_by"`
	FileExtension             sql.NullString `db:"file_extension"`
	RecordDelimiter           sql.NullString `db:"record_delimiter"`
	SkipBlankLines            bool           `db:"skip_blank_lines"`
	SkipHeader                sql.NullInt64  `db:"skip_header"`
	TimeFormat                sql.NullString `db:"time_format"`
	TimestampFormat           sql.NullString `db:"timestamp_format"`
	TrimSpace                 bool           `db:"trim_space"`
	Type                      sql.NullString `db:"type"`

	DefaultNamespace sql.NullString `db:"default_namespace"`
	DefaultRole      sql.NullString `db:"default_role"`
	DefaultWarehouse sql.NullString `db:"default_warehouse"`
	Disabled         bool           `db:"disabled"`
	HasRsaPublicKey  bool           `db:"has_rsa_public_key"`
	LoginName        sql.NullString `db:"login_name"`
	Name             sql.NullString `db:"name"`
}

func ScanUser(row *sqlx.Row) (*user, error) {
	r := &user{}
	err := row.StructScan(r)
	return r, err
}
