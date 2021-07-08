package main

import (
	"testing"
)

func expectString(t *testing.T, input string, expected string) {
	if input != expected {
		t.Errorf("\nExpected:\n'%s'\nGot:\n'%s'", expected, input)
	}
}

func TestInsert(t *testing.T) {
	query := Insert("todo", "created_at", "updated_at", "id").Values(2).Returning("id").Build()
	expected := "INSERT INTO todo ( created_at, updated_at, id ) VALUES ( $1, $2, $3 ), ( $4, $5, $6 ) RETURNING id"
	expectString(t, query, expected)
}

func TestUpdate(t *testing.T) {
	query := Update("todo").Set("title").Where("id").Eq(1).Build()
	expected := "UPDATE todo SET ( title ) = ( $1 ) WHERE id = 1"
	expectString(t, query, expected)
}
func TestUpdate2Fields(t *testing.T) {
	query := Update("todo").Set("title", "text").Where("id").Eq(1).Returning().Build()
	expected := "UPDATE todo SET ( title, text ) = ( $1, $2 ) WHERE id = 1 RETURNING *"
	expectString(t, query, expected)
}

func TestUpdateAndIn(t *testing.T){
	query := Update("todo").Set("title", "text").Where("id").Eq(1).And().In("id", 4).Build()
	expected := "UPDATE todo SET ( title, text ) = ( $1, $2 ) WHERE id = 1 AND id IN ( $3, $4, $5, $6 )"
	expectString(t, query, expected)
}