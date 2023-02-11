package status

var (
	NotFoundUser = []interface{}{
		1000102, 404, "존재하지 않은 사용자입니다",
	}
	FailedDBConnection = []interface{}{
		1001101, 500, "데이터베이스 연결을 실패하였습니다",
	}
	UnauthorizedUser = []interface{}{
		1000103, 403, "인증이 완료되지 않은 사용자입니다",
	}
	InvalidUserInfo = []interface{}{
		1000104, 403, "사용자 정보가 올바르지 않습니다",
	}
	MismatchedAuthCode = []interface{}{
		1000110, 409, "인증 코드가 일치하지 않습니다",
	}
	UserAlreadyExists = []interface{}{
		1000116, 409, "사용자가 이미 존재합니다",
	}
)
