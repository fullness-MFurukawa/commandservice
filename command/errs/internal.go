package errs

// 内部エラー型(データベース接続エラーなど)
type InternalError struct {
	message string // エラーメッセージ
}

// エラーメッセージを返すメソッド
func (e *InternalError) Error() string {
	return e.message
}

// コンストラクタ
func NewInternalError(message string) *InternalError {
	return &InternalError{message: message}
}
