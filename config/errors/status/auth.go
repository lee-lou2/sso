package status

var (
	NotFoundProvider = []interface{}{
		1000501, 404, "공급사가 존재하지 않습니다",
	}
	InvalidReturnData = []interface{}{
		1000502, 409, "반환된 데이터가 올바르지 않습니다",
	}
	FailedGenerateToken = []interface{}{
		1000503, 500, "토큰 발급을 실패하였습니다",
	}
	GoogleEmailRetrievalError = []interface{}{
		1000504, 500, "구글 이메일 조회간 오류가 발생하였습니다",
	}
)
