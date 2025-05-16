package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	clientID := 1

	client, err := selectClient(db, clientID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, clientID, client.ID)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Login)
	assert.NotEmpty(t, client.Birthday)
	assert.NotEmpty(t, client.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	clientID := 1

	client, err := selectClient(db, clientID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, client)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	cl.ID, err = insertClient(db, cl)
	if err != nil {
		t.Fatal(err)
	}
	require.NotEmpty(t, cl.ID)

	client, err := selectClient(db, cl.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, cl.ID, client.ID)
	assert.Equal(t, cl.FIO, client.FIO)
	assert.Equal(t, cl.Login, client.Login)
	assert.Equal(t, cl.Birthday, client.Birthday)
	assert.Equal(t, cl.Email, client.Email)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)
	if err != nil {
		t.Fatal(err)
	}
	require.NotEmpty(t, id)

	_, err = selectClient(db, id)
	if err != nil {
		t.Fatal(err)
	}

	err = deleteClient(db, id)
	if err != nil {
		t.Fatal(err)
	}

	_, err = selectClient(db, id)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
