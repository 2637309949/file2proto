package main

import (
	"log"
	"sort"
	"strings"

	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/schemas"
)

type sql2proto struct {
}

func (to *sql2proto) Transform(uri string) ([]*message, error) {
	tables, err := loadTables(uri)
	if err != nil {
		log.Fatalf("error loadTables: %s", err)
		return []*message{}, err
	}
	msgs := tablesToMessages(tables)
	return msgs, nil
}

func (to *sql2proto) Check(uri string) bool {
	return strings.Contains(uri, "mysql://")
}

func tablesToMessages(tables []*schemas.Table) []*message {
	var out []*message
	seen := map[string]struct{}{}
	for _, t := range tables {
		if _, ok := seen[t.Name]; ok {
			continue
		}
		seen[t.Name] = struct{}{}
		out = appendTablesMessage(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func appendTablesMessage(out []*message, t *schemas.Table) []*message {
	msg := &message{
		Name:   t.Name,
		Fields: []*field{},
	}
	lc := len(t.Columns())
	for i := 0; i < lc; i++ {
		f := t.Columns()[i]
		newField := &field{
			Name:       toProtoFieldName(f.Name),
			TypeName:   toProtoFieldTypeNameBySql(f.SQLType),
			IsRepeated: false,
			Order:      i + 1,
		}
		msg.Fields = append(msg.Fields, newField)
	}
	out = append(out, msg)
	return out
}

func toProtoFieldTypeNameBySql(f schemas.SQLType) string {
	switch f.Name {
	case "VARCHAR", "TEXT", "LONGTEXT", "CHAR", "MEDIUMTEXT", "TINYTEXT":
		return "string"
	case "DATETIME", "TIMESTAMP", "ENUM", "INT", "SMALLINT", "BIGINT", "TINYINT":
		return "int64"
	case "DECIMAL":
		return "string"
	case "BOOLEAN":
		return "bool"
	case "FLOAT", "DOUBLE":
		return "float64"
	case "MEDIUMBLOB", "BLOB":
		return "string"
	}
	return "string"
}

func loadTables(uri string) ([]*schemas.Table, error) {
	engine, err := xorm.NewEngine("mysql", strings.ReplaceAll(uri, "mysql://", ""))
	if err != nil {
		return []*schemas.Table{}, err
	}
	return engine.DBMetas()
}

func init() {
	register("sql2proto", new(sql2proto))
}
