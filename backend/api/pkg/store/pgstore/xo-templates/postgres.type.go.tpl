{{- $tableStruct := (print .Name "Table" ) -}}
{{- $tupleStruct := (print .Name "Tuple" ) }}
{{- $insertStruct := (print .Name "InsertCommand" ) }}

// {{ .Name }}TableName returns the name of corresponding table
const {{ .Name }}TableName = "{{ .Table.TableName }}"

// {{ .Name }} represents this table meta instance
var {{ .Name }} = {{ $tableStruct }}{}

// {{ $tupleStruct }} represents one row in the "{{ .Table.TableName }}"
type {{ $tupleStruct }} struct {
{{- range .Fields }}

{{- if ne .Type "Jsonb" }}
	{{ .Name }} {{ retype .Type }}
{{ else }}
	{{ .Name }} string
{{ end }}

{{- end }}
}

// Scan fills this tuple with the data read from db. All fields must present in the given row.
func (t *{{ $tupleStruct }}) Scan(row Row) error {
    return row.Scan(
{{- range .Fields }}
		&t.{{ .Name }},
{{- end }}
    )
}

// {{ $insertStruct }} is a helper struct to construct *sql.InsertBuilder that allows to insert one {{ .Name }} row into the db
type {{ $insertStruct }} struct {
{{- range .Fields }}

	{{ .Name }} interface{}

{{- end }}
}

// ToBuilder creates ready-for-use InsertBuilder
func (t *{{ $insertStruct }}) ToBuilder() *sqlbuilder.InsertBuilder {
    builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	return builder.
		InsertInto("{{ .Table.TableName }}").
		Cols(
{{- range .Fields }}
    	    "{{ .Col.ColumnName }}",
{{- end }}
		).
		Values(
{{- range .Fields }}
    		t.{{ .Name }},
{{- end }}
		)
}

// {{ $tableStruct }} exports metadata of corresponding table
type {{ $tableStruct }} struct {}

{{- range .Fields }}

// {{ .Name }} column name
func (t {{ $tableStruct }}) {{ .Name }}() string {
	return "{{ .Col.ColumnName }}"
}

// {{ .Name }}Full is the column name including the table name
func (t {{ $tableStruct }}) {{ .Name }}Full() string {
	return "{{ $.Table.TableName }}.{{ .Col.ColumnName }}"
}

{{- end }}

// SelectByPrimaryKeyBuilder returns ready-for-user sql builder to query the table by primary key. Intended for usage in tests mainly
func (t {{ $tableStruct }}) SelectByPrimaryKeyBuilder(
{{- range $i, $pk :=  .PrimaryKeyFields }}
    id{{ $i }} {{ $pk.Type }},
{{- end }}
) *sqlbuilder.SelectBuilder {
    builder := sqlbuilder.PostgreSQL.NewSelectBuilder()
    return builder.
		Select(
{{- range .Fields }}
            {{ $.Name }}.{{ .Name }}(),
{{- end }}
		).
		From({{ .Name }}TableName).
		Where(
{{- range $i, $pk :=  .PrimaryKeyFields }}
			builder.Equal({{ $.Name }}.{{ .Name }}(), id{{ $i }}),
{{- end }}
		)
}
