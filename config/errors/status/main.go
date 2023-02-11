package status

var (
	// --- OTP ---

	NotFoundOTPSecretKey = []interface{}{
		1000301, 404, "OTP 시크릿 키가 존재하지 않습니다",
	}

	// --- Security ---

	ShortCipherBlockSize = []interface{}{
		1000401, 500, "Cipher Text 블록 사이즈가 너무 짧습니다",
	}
	ConversionTextNotFound = []interface{}{
		1000402, 404, "변환하려는 텍스트가 존재하지 않습니다",
	}

	// --- Location ---

	NotReceivedLocationInfo = []interface{}{
		1000601, 400, "위치 정보를 받지 못했습니다",
	}
	NotFoundLocationData = []interface{}{
		1000602, 404, "데이터가 존재하지 않습니다",
	}
	LocationRetrievalError = []interface{}{
		1000603, 500, "아이피를 이용한 지역 정보 조회간 오류 발생",
	}
	InvalidInformationError = []interface{}{
		1000604, 400, "조회된 정보 형태가 올바르지 않습니다",
	}

	// --- Client ---

	RequestBodyBindingError = []interface{}{
		1000801, 400, "입력하신 데이터가 올바르지 않습니다",
	}
	ExistsClient = []interface{}{
		1000802, 400, "이미 존재하는 클라이언트입니다",
	}
	FailedSecretKeyGeneration = []interface{}{
		1000803, 500, "시크릿키 생성을 실패하였습니다",
	}
	InvalidClientCode = []interface{}{
		1000804, 409, "클라이언트 코드가 올바르지 않습니다",
	}
	NotFoundClient = []interface{}{
		1000805, 404, "클라이언트가 존재하지 않습니다",
	}
	NotFoundClientGroup = []interface{}{
		1000806, 404, "그룹이 존재하지 않습니다",
	}
	ClientNameRequired = []interface{}{
		1000807, 409, "클라이언트명은 필수 값입니다",
	}

	// --- Cache ---

	CacheListRetrievalError = []interface{}{
		1001001, 409, "캐시 리스트 조회간 오류가 발생하였습니다",
	}
	CacheDataCountRetrievalError = []interface{}{
		1001002, 409, "캐시 데이터 수 조회간 오류가 발생하였습니다",
	}
	CacheDataExtractionError = []interface{}{
		1001003, 409, "캐시 데이터 추출간 오류가 발생하였습니다",
	}
	CacheDataInsertError = []interface{}{
		1001004, 409, "캐시 데이터 삽입간 오류가 발생하였습니다",
	}
	CacheDataConfigError = []interface{}{
		1001006, 409, "캐시 데이터 설정간 오류가 발생하였습니다",
	}
)
