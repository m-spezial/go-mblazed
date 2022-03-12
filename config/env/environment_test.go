package env

import (
	"os"
	"testing"
)

type testStruct struct {
	EmptyEnv        string `env:"ENV_TEST_NOT_EXISTS"`
	EmptyDefaultEnv string `env:"ENV_TEST_DEFAULT,default=DefaultVal"`

	String string `env:"ENV_TEST_STRING"`

	Bool bool `env:"ENV_TEST_BOOL"`

	Int64 int64 `env:"ENV_TEST_INT64"`
	Int32 int64 `env:"ENV_TEST_INT32"`
	Int16 int64 `env:"ENV_TEST_INT16"`
	Int8  int64 `env:"ENV_TEST_INT8"`
	Int   int64 `env:"ENV_TEST_INT"`

	UInt64 int64 `env:"ENV_TEST_UINT64"`
	UInt32 int64 `env:"ENV_TEST_UINT32"`
	UInt16 int64 `env:"ENV_TEST_UINT16"`
	UInt8  int64 `env:"ENV_TEST_UINT8"`
	UInt   int64 `env:"ENV_TEST_UINT"`
}

func prepareEnvVariables() {
	os.Setenv("ENV_TEST_STRING", "String")

	os.Setenv("ENV_TEST_BOOL", "true")

	os.Setenv("ENV_TEST_INT64", "64")
	os.Setenv("ENV_TEST_INT32", "32")
	os.Setenv("ENV_TEST_INT16", "16")
	os.Setenv("ENV_TEST_INT8", "8")
	os.Setenv("ENV_TEST_INT", "-32")

	os.Setenv("ENV_TEST_UINT64", "64")
	os.Setenv("ENV_TEST_UINT32", "32")
	os.Setenv("ENV_TEST_UINT16", "16")
	os.Setenv("ENV_TEST_UINT8", "8")
	os.Setenv("ENV_TEST_UINT", "320")

}

func TestReadValues(t *testing.T) {
	prepareEnvVariables()

	testStruct := &testStruct{}

	err := ReadValues(testStruct)

	if err != nil {
		t.Error(err)
		return
	}

	if testStruct.EmptyEnv != "" {
		t.Errorf("EmptyEnv field should be empty, got: %s, want nothing.", testStruct.String)
	}

	if testStruct.String != "String" {
		t.Errorf("String field was incorrect, got: %s, want: %s.", testStruct.String, "String")
	}

	if testStruct.EmptyDefaultEnv != "DefaultVal" {
		t.Errorf("Default field was incorrect, got: %s, want: %s.", testStruct.EmptyDefaultEnv, "DefaultVal")
	}

	if testStruct.Bool != true {
		t.Errorf("Bool field was incorrect, got: %t, want: %t.", testStruct.Bool, true)
	}

	if testStruct.Int64 != 64 {
		t.Errorf("Int64 field was incorrect, got: %d, want: %d.", testStruct.Int64, 64)
	}

	if testStruct.Int32 != 32 {
		t.Errorf("Int32 field was incorrect, got: %d, want: %d.", testStruct.Int32, 32)
	}

	if testStruct.Int16 != 16 {
		t.Errorf("Int16 field was incorrect, got: %d, want: %d.", testStruct.Int16, 16)
	}

	if testStruct.Int8 != 8 {
		t.Errorf("Int8 field was incorrect, got: %d, want: %d.", testStruct.Int8, 8)
	}

	if testStruct.Int != -32 {
		t.Errorf("Int field was incorrect, got: %d, want: %d.", testStruct.Int, -32)
	}

	if testStruct.UInt64 != 64 {
		t.Errorf("UInt64 field was incorrect, got: %d, want: %d.", testStruct.UInt64, 64)
	}

	if testStruct.UInt32 != 32 {
		t.Errorf("UInt32 field was incorrect, got: %d, want: %d.", testStruct.UInt32, 32)
	}

	if testStruct.UInt16 != 16 {
		t.Errorf("UInt16 field was incorrect, got: %d, want: %d.", testStruct.UInt16, 16)
	}

	if testStruct.UInt8 != 8 {
		t.Errorf("UInt8 field was incorrect, got: %d, want: %d.", testStruct.UInt8, 8)
	}

	if testStruct.UInt != 320 {
		t.Errorf("UInt field was incorrect, got: %d, want: %d.", testStruct.UInt, 320)
	}
}

func TestReadValuesRequired(t *testing.T) {
	type testRequired struct {
		RequiredString string `env:"ENV_TEST_REQUIRED,required"`
	}

	testStruct := &testRequired{}

	err := ReadValues(testStruct)

	if err != nil {

		return
	}

	t.Errorf("Want error because required variable")
}
