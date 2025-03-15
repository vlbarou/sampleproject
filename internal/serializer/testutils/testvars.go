package testutils

// Why Do I declare the structs with capital first letter?

type MockValue struct {
	Param1 string
	Param2 int
	Param3 NestedValue
}

type NestedValue struct {
	NestedParam1 string
	NestedParam2 bool
}
