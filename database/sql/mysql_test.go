package sql

import (
	"database/sql"
	"testing"
)

func TestMysql(t *testing.T) {
	writeKey := "test"
	readKey := "test"
	db := ConnectMysql(writeKey, readKey, "", "")
	defer db.Close()

	testPing(t, db)
	testTable(t, db)
	testExec(t, db)
	testQuery(t, db)
	testQueryRow(t, db)
	testPrepare(t, db)
	testMaster(t, db)
}

func testMaster(t *testing.T, db *DB) {
	master := db.Master()
	if master.read != nil {
		t.Errorf("expect master read conn is 0, get %v", master.read)
	}
	if master.write != db.write {
		t.Errorf("expect master write conn equal to origin db write conn")
	}
}

func (db *DB) Master() *DB {
	if db.master == nil {
		panic("sql: no master instance")
	}
	return db.master
}

func testPrepare(t *testing.T, db *DB) {
	var (
		selsql  = "SELECT name FROM test WHERE name=?"
		execsql = "INSERT INTO test(name) VALUES(?)"
		name    = ""
	)
	selstmt, err := db.read.Prepare(selsql)
	if err != nil {
		t.Errorf("MySQL:Prepare err(%v)", err)
		return
	}
	row := selstmt.QueryRow("noexit")
	if err = row.Scan(&name); err == sql.ErrNoRows {
		t.Logf("MySQL: prepare query error(%v)", err)
	} else {
		t.Errorf("MySQL: prepared query name: noexist")
	}
	rows, err := selstmt.Query("test")
	if err != nil {
		t.Errorf("MySQL:stmt.Query err(%v)", err)
	}
	rows.Close()
	execstmt, err := db.read.Prepare(execsql)
	if err != nil {
		t.Errorf("MySQL:Prepare err(%v)", err)
		return
	}
	if _, err := execstmt.Exec("test"); err != nil {
		t.Errorf("MySQL: stmt.Exec(%v)", err)
	}
}

func testQueryRow(t *testing.T, db *DB) {
	sql := "SELECT name FROM test WHERE name=?"
	name := ""
	row := db.read.QueryRow(sql, "test")

	if err := row.Scan(&name); err != nil {
		t.Errorf("MySQL: queryRow error(%v)", err)
	} else {
		t.Logf("MySQL: queryRow name: %s", name)
	}
}

func testQuery(t *testing.T, db *DB) {
	sql := "SELECT name FROM test WHERE name=?"
	rows, err := db.read.Query(sql, "test")
	if err != nil {
		t.Errorf("MySQL: query error(%v)", err)
	}
	defer rows.Close()
	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			t.Errorf("MySQL: query scan error(%v)", err)
		} else {
			t.Logf("MySQL: query name: %s", name)
		}
	}
}

func testExec(t *testing.T, db *DB) {
	sql := "INSERT INTO test(name) VALUES(?)"
	if _, err := db.write.Exec(sql, "test"); err != nil {
		t.Errorf("MySQL: insert error(%v)", err)
	} else {
		t.Log("MySQL: insert ok")
	}
}

func testTable(t *testing.T, db *DB) {

	table := "CREATE TABLE IF NOT EXISTS `test` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID', `name` varchar(16) NOT NULL DEFAULT '' COMMENT '名称', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8"
	if _, err := db.write.Exec(table); err != nil {
		t.Errorf("MySQL: create table error(%v)", err)
	} else {
		t.Log("MySQL: create table ok")
	}
}

func testPing(t *testing.T, db *DB) {

	if err := db.write.Ping(); err != nil {
		t.Errorf("MySQL: write ping error(%v)", err)
	}

	if err := db.read.Ping(); err != nil {
		t.Errorf("MySQL: read ping error(%v)", err)
	}

	return
}
