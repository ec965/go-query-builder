package main

import (
	"fmt"
	"strings"
)

type Builder struct {
	query      string
	valCount   int
	inputCount int
}

func (b *Builder) Build() string {
	return b.query
}

// INSERT tableName ( fields... )
func Insert(tableName string, fieldNames ...string) *Builder {
	fieldStr := strings.Join(fieldNames, ", ")
	b := Builder{}
	b.query += fmt.Sprintf("INSERT INTO %s ( %s )", tableName, fieldStr)
	b.valCount = len(fieldNames)
	return &b
}

// VALUES ( $1, $2 ), ( $3, $4 )
func (b *Builder) Values(numOfRows int) *Builder {
	var strArr []string
	for i := 0; i < numOfRows; i++ {
		var fieldsArr []string
		for j := 0; j < b.valCount; j++ {
			b.inputCount++
			fieldsArr = append(fieldsArr, fmt.Sprintf("$%d", b.inputCount))
		}
		fieldStr := fmt.Sprintf("( %s )", strings.Join(fieldsArr, ", "))
		strArr = append(strArr, fieldStr)
	}
	b.query += fmt.Sprintf(" VALUES %s", strings.Join(strArr, ", "))
	return b
}

// RETURNING
// if no parameter is specified, all values are returned (e.g. '*')
func (b *Builder) Returning(fields ...string) *Builder {
	str := "*"
	if len(fields) != 0 {
		str = strings.Join(fields, ", ")
	}
	b.query += fmt.Sprintf(" RETURNING %s", str)
	return b
}

// UPDATE
func Update(tableName string) *Builder {
	b := Builder{}
	b.query = fmt.Sprintf("UPDATE %s", tableName)
	return &b
}

// SET ( field ) = ( $... )
func (b *Builder) Set(fields ...string) *Builder {
	var dollarArr []string
	for i := 0; i < len(fields); i++ {
		b.inputCount++
		dollarArr = append(dollarArr, fmt.Sprintf("$%d", b.inputCount))
	}
	b.query += fmt.Sprintf(" SET ( %s ) = ( %s )", strings.Join(fields, ", "), strings.Join(dollarArr, ", "))
	return b
}

// WHERE field
func (b *Builder) Where(field string) *Builder {
	b.query += fmt.Sprintf(" WHERE %s", field)
	return b
}

// = $1
func (b *Builder) Eq(value interface{}) *Builder {
	b.query += fmt.Sprintf(" = %v", value)
	return b
}

func (b *Builder) And() *Builder {
	b.query += " AND"
	return b
}

func (b *Builder) Or() *Builder {
	b.query += " OR"
	return b
}

func (b *Builder) In(fieldName string, count int) *Builder {
	var dollarArr []string
	for i := 1; i < count+1; i++ {
		b.inputCount++
		dollarArr = append(dollarArr, fmt.Sprintf("$%d", b.inputCount))
	}
	b.query += fmt.Sprintf(" %s IN ( %s )", fieldName, strings.Join(dollarArr, ", "))
	return b
}

func main() {
	s := Update("todo").Set("title", "text").Where("id").Eq(1).And().In("id", 4).Build()
	fmt.Println(s)
}
