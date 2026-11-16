// test
package bal

import "testing"

func TestSteps(t *testing.T) { if Steps == "" { t.Fatal("empty") } }
// probe x19: leading-slash CODEOWNERS glob on a _test file
