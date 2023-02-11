package status

var (
	TokenVerificationFailure = []interface{}{
		1000201, 400, "토큰 검증을 실패하였습니다",
	}
	InvalidToken = []interface{}{
		1000202, 404, "유효하지 않은 토큰입니다",
	}
	InvalidTokenType = []interface{}{
		1000203, 409, "토큰 종류가 올바르지 않습니다",
	}
	NotFoundToken = []interface{}{
		1000207, 404, "토큰이 존재하지 않습니다",
	}
	TokenValidationError = []interface{}{
		1000208, 500, "토큰 검증간 오류가 발생하였습니다",
	}
	TokenMissingRequiredInfo = []interface{}{
		1000209, 400, "토큰에 필수 정보가 포함되어있지 않습니다",
	}
	PageStatusNotSet = []interface{}{
		1000210, 409, "페이지 상태가 설정되어있지 않습니다",
	}
)
