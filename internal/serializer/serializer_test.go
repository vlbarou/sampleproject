package serializer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vlbarou/sampleproject/internal/serializer/testutils"
	"testing"
)

func TestMarshal(t *testing.T) {

	// arrange
	m := &testutils.MockValue{
		Param1: "txt param1",
		Param2: 1,
		Param3: testutils.NestedValue{
			NestedParam1: "nested param1 txt",
			NestedParam2: true,
		},
	}

	// act
	str, err := Marshal(m)

	// assert
	if err != nil {
		fmt.Printf("marshal failed for:  %v\n", m)
		t.FailNow()
	}

	assert.Equal(t, "{\"Param1\":\"txt param1\",\"Param2\":1,\"Param3\":{\"NestedParam1\":\"nested param1 txt\",\"NestedParam2\":true}}", str)
}

func TestMarshalUnmarshal(t *testing.T) {
	var mockValue testutils.MockValue

	// arrange
	m := testutils.GetMockStruct()

	// act
	str, err := Marshal(m)

	// assert
	assert.NoError(t, err)
	UnmarshalJSON(str, &mockValue)
	assert.Equal(t, m.Param1, mockValue.Param1)
	assert.Equal(t, m.Param2, mockValue.Param2)
	assert.Equal(t, m.Param3.NestedParam1, mockValue.Param3.NestedParam1)
	assert.Equal(t, m.Param3.NestedParam2, mockValue.Param3.NestedParam2)
}
