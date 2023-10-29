package handler

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// データベース接続情報
type DBConfig struct {
	Dbname string `toml:"dbname"` //	データベース名
	Host   string `toml:"host"`   //	ホスト名
	Port   int64  `toml:"port"`   //	ポート番号
	User   string `toml:"user"`   //	ユーザー名
	Pass   string `toml:"pass"`   //	パスワード
}

// database.tomlから接続情報を取得してDbConfig型で返す
func tomlRead() (*DBConfig, error) {
	// 環境変数からファイルパスを取得する
	path := os.Getenv("DATABSE_TOML_PATH")
	if path == "" {
		// 環境変数が無い場合のパスを設定する
		path = "infra/sqlboiler/config/database.toml"
	}
	// database.tomlを読取りDBConfigにマッピングする
	m := map[string]DBConfig{}
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}
	config := m["mysql"]
	return &config, nil
}

// データベース接続
func DBConnect() error {

	config, err := tomlRead() // database.tomlの定義内容を読み取る
	if err != nil {
		return DBErrHandler(err)
	}

	// 接続文字列を生成する
	rdbms := "mysql"
	connect_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Pass, config.Host, config.Port, config.Dbname)

	// データベースに接続する
	conn, err := sql.Open(rdbms, connect_str)
	if err != nil {
		return DBErrHandler(err)
	}
	// データベース接続を確認する
	if err = conn.Ping(); err != nil {
		return DBErrHandler(err)
	}

	MAX_IDLE_CONNS := 10                  // 初期接続数
	MAX_OPEN_CONNS := 100                 // 最大接続数
	CONN_MAX_LIFTIME := 300 * time.Second // 最大生存期間

	// コネクションプールの設定
	conn.SetMaxIdleConns(MAX_IDLE_CONNS)      // 初期接続数
	conn.SetMaxOpenConns(MAX_OPEN_CONNS)      // 最大接続数
	conn.SetConnMaxLifetime(CONN_MAX_LIFTIME) // 最大利用生存期間

	boil.SetDB(conn)      // グローバルコネクション設定
	boil.DebugMode = true // デバッグモードに設定 生成されたSQLを出力する
	return nil
}
