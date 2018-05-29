package example

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleHello(t *testing.T) {
	for _, tc := range []struct {
		givenArgument     string
		expectedArgument  string
		givenSuperFlag    string
		expectedSuperFlag string
		err               bool
	}{
		{
			givenArgument:     "ok",
			expectedArgument:  "ok",
			givenSuperFlag:    "ko",
			expectedSuperFlag: "ko",
			err:               false,
		},
		{
			givenArgument:     "ok",
			expectedArgument:  "ok",
			givenSuperFlag:    "",
			expectedSuperFlag: "",
			err:               true,
		},
	} {
		t.Run("", func(t *testing.T) {
			e := &Example{
				Argument:  tc.givenArgument,
				SuperFlag: tc.givenSuperFlag,
			}
			assert.NoError(t, e.Hello())
			assert.Equal(t, tc.expectedArgument, e.GetArgument())
			f, err := e.GetSuperFlag()
			assert.Equal(t, tc.expectedSuperFlag, f)
			assert.Equal(t, tc.err, err != nil)
		})
	}
}
