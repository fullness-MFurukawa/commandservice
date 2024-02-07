package errs

// ドメインに関連するエラー型
type DomainError struct {
	message string // エラーメッセージ
}

// エラーメッセージを返すメソッド
func (e *DomainError) Error() string {
	return e.message
}

// コンストラクタ
func NewDomainError(message string) *DomainError {
	return &DomainError{message: message}
}
