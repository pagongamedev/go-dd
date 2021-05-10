package sqlx

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	godd "github.com/pagongamedev/go-dd"
)

// Sqlx

func GetRowOne(rows *sqlx.Rows, err error, template interface{}) (interface{}, *godd.Error) {
	if err != nil || rows == nil {
		return false, godd.ErrorNew(http.StatusUnauthorized, err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	isFind := false
	for rows.Next() {
		isFind = true
		err = rows.StructScan(template)
		if err != nil {
			return false, godd.ErrorNew(http.StatusUnauthorized, err)
		}
	}
	if !isFind {
		return nil, nil
	}

	return template, nil
}

func GetRowFunc(rows *sqlx.Rows, err error, fnc func(r *sqlx.Rows) error) *godd.Error {
	if err != nil || rows == nil {
		return godd.ErrorNew(http.StatusUnauthorized, err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	for rows.Next() {

		err = fnc(rows)
		if err != nil {
			return godd.ErrorNew(http.StatusUnauthorized, err)
		}
	}

	return nil
}

func AddValueCount(query string, iMax int) string {
	str := ""

	for i := 1; i <= iMax; i++ {
		str += fmt.Sprintf("$%v", i)
		if i != iMax {
			str += ","
		}
	}
	return strings.Replace(query, "{{Values}}", str, 1)
}

func IsRollbackByError(tx *sqlx.Tx, goddErr *godd.Error) bool {
	return IsRollback(tx, goddErr != nil)
}

func IsRollback(tx *sqlx.Tx, isRollback bool) bool {
	if isRollback {
		err := tx.Rollback()
		if err != nil {
			log.Println("Rollback Error : ", err)
		}
	}
	return isRollback
}

// func helperQueryMany(rows *sqlx.Rows, err error, dList interface{}, template interface{}) *godd.Error {
// 	if err != nil || rows == nil {
// 		return ErrorNew(http.StatusUnauthorized, err)
// 	}
// 	defer func() {
// 		err := rows.Close()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}()
// 	vs := reflect.ValueOf(dList).Elem()
// 	for rows.Next() {

// 		a := reflect.ValueOf(dList).Elem()
// 		fmt.Println("L ", a, a.Type(), a.Kind())

// 		fmt.Println("A ", reflect.ValueOf(dList).Elem().Kind())

// 		err = rows.StructScan(template)
// 		if err != nil {
// 			return ErrorNew(http.StatusUnauthorized, err)
// 		}
// 		vsNew := reflect.Append(vs, reflect.ValueOf(template).Elem())
// 		vs.Set(vsNew)
// 	}

// 	return nil
// }

// ==============================================

func QueryOne(tx *sqlx.Tx, query string, argList []interface{}, responseStuct interface{}) (interface{}, *godd.Error) {
	query = AddValueCount(query, len(argList))
	rows, err := tx.Queryx(
		query,
		argList...,
	)
	return GetRowOne(rows, err, responseStuct)
}

func QueryFunc(tx *sqlx.Tx, query string, argList []interface{}, fnc func(r *sqlx.Rows) error) *godd.Error {
	query = AddValueCount(query, len(argList))
	rows, err := tx.Queryx(
		query,
		argList...,
	)
	return GetRowFunc(rows, err, fnc)
}
