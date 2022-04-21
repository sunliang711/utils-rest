package types

import "fmt"

const (
	Ok = iota
	ErrGeneral
	ErrEOF
	ErrAdminRole
	ErrUserRole
	ErrInvalidEmailType
	ErrInvalidJsonInput
	ErrInternalDB
	ErrInvalidInput
	ErrParseTemplate
	ErrExecuteTemplate
	ErrWriteFile
	ErrReadFile
	ErrCommandNotFound
	ErrCompile
	ErrCompilerVersion
	ErrParseJwtToken
	ErrJwtType
	ErrInvalidForm
	ErrExceedMaxUpload
	ErrOpenUploadFile
	ErrUploadFile
	ErrRoyaltyLevelContract
	ErrRoyaltyLevelToken
	ErrS3Upload
	ErrVerifyEtherscan
)

var (
	ErrorTable = map[int]string{
		Ok:                      "ok",
		ErrGeneral:              "error",
		ErrEOF:                  "eof error, no input data",
		ErrAdminRole:            "check admin role error",
		ErrUserRole:             "check user role error",
		ErrInvalidEmailType:     "invalid email type",
		ErrInvalidJsonInput:     "invalid json input",
		ErrInternalDB:           "internal db error",
		ErrInvalidInput:         "invalid input",
		ErrParseTemplate:        "parse template file error",
		ErrExecuteTemplate:      "execute template error",
		ErrReadFile:             "read file error",
		ErrWriteFile:            "write file error",
		ErrCommandNotFound:      "command not found",
		ErrCompile:              "compile contract error",
		ErrCompilerVersion:      "get compiler version error",
		ErrParseJwtToken:        "parse jwt token error",
		ErrJwtType:              "jwt token type error",
		ErrInvalidForm:          "invalid form",
		ErrExceedMaxUpload:      "exceed max upload",
		ErrOpenUploadFile:       "open upload file error",
		ErrUploadFile:           "upload file error",
		ErrRoyaltyLevelContract: "royalty level not contract",
		ErrRoyaltyLevelToken:    "royalty level not nft",
		ErrS3Upload:             "s3 upload error",
		ErrVerifyEtherscan:      "verify etherscan error",
	}
)

func ErrorMsg(code int, additionMsg string) string {
	msg := ErrorTable[code]
	if additionMsg != "" {
		msg = fmt.Sprintf("%s, %s", msg, additionMsg)
	}

	return msg
}
