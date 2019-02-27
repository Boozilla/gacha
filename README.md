### Usage example (Lotto 6/45)

```go
package main

import "github.com/Boozilla/gacha/gacha"

func main() {
	entry := make([]Capsule, 45)
	for i := 1; i <= 45; i++ {
		entry[i - 1] = Capsule{
			Item: i,
			Probability: 1.0,
		}
	}

	for l := 0; l < 5; l++ {
		gacha := New()
		gacha.Put(entry...)

		for i := 0; i < 6; i++ {
			c, _ := gacha.Get()

			print(c.Item.(int))
			print("\t")
		}

		println()
	}
}
```
