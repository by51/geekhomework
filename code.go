package practice

import (
	"database/sql"
	_ "errors"
	"fmt"
	"github.com/pkg/errors"
	_ "log"
)

type dt struct{
	name string
}

func queryData(ids []string) ([]dt, error){

	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		return nil, errors.Wrap(err, "Db connect error.")
	}

	defer db.Close()

	rows,err := db.Query("select name from users where id = ?", ids)
	if err != nil{
		return nil, errors.Wrap(err, "Quey data row error.")
	}
	defer rows.Close()

	result := make([]dt, 0)

	for rows.Next(){

		dr := new(dt)

		rErr := rows.Scan(dr.name)

		if rErr != nil{
			if rErr == sql.ErrNoRows{
				continue
			}else {
				return nil, errors.Wrap(rErr,"Scan data error")
			}

		}

		result = append(result, *dr)
	}

	return result, nil
}