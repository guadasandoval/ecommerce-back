package errors

import (
	"testing"

	"go.uber.org/zap"
)

func TestNewPanic(t *testing.T) {
	tt := []struct {
		name        string
		err         error
		shouldPanic bool
	}{
		{
			name:        "Test 1",
			err:         nil,
			shouldPanic: true,
		},
		{
			name:        "Test 2",
			err:         Msg("error"),
			shouldPanic: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				// Check if panic
				r := recover()

				if r == nil && tc.shouldPanic {
					t.Fatalf("%s: Test should panic", tc.name)
				}

				if r != nil && !tc.shouldPanic {
					t.Fatalf("%s: Test shouldn't panic", tc.name)
				}
			}()

			eerr := NewInfo(Operation("op"), KindUnexpected, tc.err)
			_ = eerr // Ignore error, just for testing (avoid lint error)
		})
	}
}

func TestLevel(t *testing.T) {
	t.Run("Common error", func(t *testing.T) {
		level := Level(Msg("error"))
		if level != LevelError {
			t.Fatalf("Common error: Level result %v, expected level %v", level, LevelError)
		}
	})

	tt := []struct {
		name          string
		level         LevelType
		err           error
		expectedLevel LevelType
	}{
		{
			name:          "Test 1",
			level:         LevelInfo,
			err:           Msg("error"),
			expectedLevel: LevelInfo,
		},
		{
			name:          "Test 2",
			level:         LevelInfo,
			err:           NewError(Operation("op"), KindUnexpected, Msg("error")),
			expectedLevel: LevelInfo,
		},
		{
			name:          "Test 3",
			level:         LevelError,
			err:           NewInfo(Operation("op"), KindUnexpected, Msg("error")),
			expectedLevel: LevelError,
		},
		{
			name:  "Test 4",
			level: LevelInfo,
			err: NewError(Operation("op"), KindUnexpected,
				NewInfo(Operation("op"), KindUnexpected, Msg("error")),
			),
			expectedLevel: LevelInfo,
		},
		{
			name:  "Test 5",
			level: LevelInfo,
			err: NewInfo(Operation("op"), KindUnexpected,
				NewError(Operation("op"), KindUnexpected, Msg("error")),
			),
			expectedLevel: LevelInfo,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := new(tc.level, Operation("op"), KindUnexpected, tc.err)
			level := Level(err)
			if level != tc.expectedLevel {
				t.Fatalf("%s: Level result %v, expected level %v", tc.name, level, tc.expectedLevel)
			}
		})
	}
}

func TestKind(t *testing.T) {
	t.Run("Common error", func(t *testing.T) {
		kind := Kind(Msg("error"))
		if kind != KindUnexpected {
			t.Fatalf("Common error: Kind result %v, expected kind %v", kind, KindUnexpected)
		}
	})

	tt := []struct {
		name         string
		kind         KindType
		err          error
		expectedKind KindType
	}{
		{
			name:         "Test 1",
			kind:         KindNotFound,
			err:          Msg("error"),
			expectedKind: KindNotFound,
		},
		{
			name:         "Test 2",
			kind:         KindNotFound,
			err:          NewError(Operation("op"), KindDBQuery, Msg("error")),
			expectedKind: KindNotFound,
		},
		{
			name:         "Test 3",
			kind:         0,
			err:          NewInfo(Operation("op"), KindUnexpected, Msg("error")),
			expectedKind: KindUnexpected,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := NewError(Operation("op"), tc.kind, tc.err)
			kind := Kind(err)
			if kind != tc.expectedKind {
				t.Fatalf("%s: Kind result %v, expected kind %v", tc.name, kind, tc.expectedKind)
			}
		})
	}
}

func TestOperations(t *testing.T) {
	t.Run("Common error", func(t *testing.T) {
		operations := Operations(Msg("error"))
		if operations == nil || len(operations) != 0 {
			t.Fatalf("Common error: Operations result %v, expected operations []", operations)
		}
	})

	tt := []struct {
		name       string
		op         Operation
		err        error
		expectedOp []Operation
	}{
		{
			name:       "Test 1",
			op:         Operation("op1"),
			err:        Msg("error"),
			expectedOp: []Operation{Operation("op1")},
		},
		{
			name:       "Test 2",
			op:         Operation("op1"),
			err:        NewError(Operation("op2"), KindDBQuery, Msg("error")),
			expectedOp: []Operation{Operation("op1"), Operation("op2")},
		},
		{
			name: "Test 3",
			op:   Operation("op1"),
			err: NewError(Operation("op2"), KindDBQuery, NewError(
				Operation("op3"), KindDBQuery, Msg("error"),
			)),
			expectedOp: []Operation{Operation("op1"), Operation("op2"), Operation("op3")},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := NewError(tc.op, KindUnexpected, tc.err)
			operations := Operations(err)
			if len(operations) != len(tc.expectedOp) {
				t.Fatalf("%s: Operations result len %d, expected operations len %d", tc.name, len(operations),
					len(tc.expectedOp))
			}
			for i := range operations {
				if operations[i] != tc.expectedOp[i] {
					t.Fatalf("%s: Index %d: Operation result %v, expected operation %v", tc.name, i, operations[i],
						tc.expectedOp[i])
				}
			}
		})
	}
}

func TestFields(t *testing.T) {
	t.Run("Common error", func(t *testing.T) {
		fields := Fields(Msg("error"))
		if fields == nil || len(fields) != 0 {
			t.Fatalf("Common error: Fields result %v, expected fields []", fields)
		}
	})

	tt := []struct {
		name           string
		op             Operation
		err            error
		fields         []zap.Field
		expectedFields []zap.Field
	}{
		{
			name:           "Test 1",
			op:             Operation("op1"),
			err:            Msg("error"),
			fields:         []zap.Field{zap.String("a", "aa"), zap.String("b", "bb")},
			expectedFields: []zap.Field{zap.String("op1:a", "aa"), zap.String("op1:b", "bb")},
		},
		{
			name:   "Test 2",
			op:     Operation("op1"),
			err:    NewError(Operation("op2"), KindDBQuery, Msg("error"), zap.String("a", "aa1"), zap.String("b", "bb1")),
			fields: []zap.Field{zap.String("a", "aa2"), zap.String("b", "bb2")},
			expectedFields: []zap.Field{
				zap.String("op1:a", "aa2"), zap.String("op1:b", "bb2"),
				zap.String("op2:a", "aa1"), zap.String("op2:b", "bb1"),
			},
		},
		{
			name: "Test 2",
			op:   Operation("op1"),
			err: NewError(
				Operation("op2"),
				KindDBQuery,
				NewError(
					Operation("op3"),
					KindDBQuery,
					Msg("error"),
					zap.String("c", "aa1"),
					zap.String("d", "bb1"),
				),
				zap.String("a", "aa1"),
				zap.String("b", "bb1"),
			),
			fields: []zap.Field{zap.String("a", "aa2"), zap.String("b", "bb2")},
			expectedFields: []zap.Field{
				zap.String("op1:a", "aa2"), zap.String("op1:b", "bb2"),
				zap.String("op2:a", "aa1"), zap.String("op2:b", "bb1"),
				zap.String("op3:c", "aa1"), zap.String("op3:d", "bb1"),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := NewError(tc.op, KindUnexpected, tc.err, tc.fields...)
			fields := Fields(err)
			if len(fields) != len(tc.expectedFields) {
				t.Fatalf("%s: Fields result len %d, expected fields len %d", tc.name, len(fields), len(tc.expectedFields))
			}
			for i := range fields {
				if fields[i] != tc.expectedFields[i] {
					t.Fatalf("%s: Index %d: Operation result %v, expected operation %v", tc.name, i, fields[i],
						tc.expectedFields[i])
				}
			}
		})
	}
}

func TestIs(t *testing.T) {
	err1 := Msg("error1")
	err2 := Msg("error2")

	tt := []struct {
		name      string
		err       error
		errTarget error
		expected  bool
	}{
		{
			name:      "Test 1",
			err:       err1,
			errTarget: err2,
			expected:  false,
		},
		{
			name:      "Test 2",
			err:       err1,
			errTarget: err1,
			expected:  true,
		},
		{
			name:      "Test 3",
			err:       NewError(Operation("op"), KindDBQuery, err1),
			errTarget: err1,
			expected:  true,
		},
		{
			name:      "Test 4",
			err:       NewError(Operation("op"), KindDBQuery, NewError(Operation("op1"), KindDBQuery, err1)),
			errTarget: err1,
			expected:  true,
		},
		{
			name:      "Test 5",
			err:       NewError(Operation("op"), KindDBQuery, err2),
			errTarget: err1,
			expected:  false,
		},
		{
			name:      "Test 6",
			err:       NewError(Operation("op"), KindDBQuery, NewError(Operation("op1"), KindDBQuery, err2)),
			errTarget: err1,
			expected:  false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if Is(tc.err, tc.errTarget) != tc.expected {
				t.Fatalf("%s: invalid result %v", tc.name, Is(tc.err, tc.errTarget))
			}
		})
	}
}
