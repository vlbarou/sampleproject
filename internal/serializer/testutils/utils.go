package testutils

func GetMockStruct() MockValue {
	return MockValue{
		Param1: "txt param1",
		Param2: 1,
		Param3: NestedValue{
			NestedParam1: "nested param1 txt",
			NestedParam2: true,
		},
	}
}
