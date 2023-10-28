package handler

import (
	"commandservice/errs"
	"errors"
	"log"
	"net"

	"github.com/go-sql-driver/mysql"
)

// データベースアクセスエラーのハンドリング
func DBErrHandler(err error) error {
	var opErr *net.OpError
	var driverErr *mysql.MySQLError
	if errors.As(err, &opErr) { // 接続がタイムアウトかネットワーク関連の問題が原因で接続が確立できない?
		log.Println(err.Error())
		return errs.NewInternalError(opErr.Error())
	} else if errors.As(err, &driverErr) { // MySQLドライバエラー?
		log.Printf("Code:%d Message:%s \n", driverErr.Number, driverErr.Message)
		if driverErr.Number == 1062 { // 一意制約違反?
			return errs.NewCRUDError("一意制約違反です。")
		} else {
			return errs.NewInternalError(driverErr.Message)
		}
	} else { // その他のエラー
		log.Println(err.Error())
		return errs.NewInternalError(err.Error())
	}
}
