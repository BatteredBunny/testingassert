# testingassert

Simple package providing Assert functions for Golang testing

```
go get github.com/BatteredBunny/testingassert
```

## Usage

```go
func BasicTest(t *testing.T) {
    assert.TestState = t
    assert.HideSuccess = false

	assert.Equals("123", "1" + "2" + "3")
    assert.NotEquals(123, 5+3, "math operation returned wrong answer")

    a := 1 == 1
    assert.Assert(a, "function shouldn't return false!")
}
```
