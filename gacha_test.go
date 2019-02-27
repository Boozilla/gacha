package gacha

import (
	"testing"
)

func TestGacha(t *testing.T) {
	stat := make([]int, 4)
	gacha := New()

	t.Log("Before gacha.Len(): ", gacha.Len())
	t.Log("Before gacha.Empty(): ", gacha.Empty())

	gacha.Put(
		Capsule{
			Item:        0,
			Probability: float64(5),
		},
		Capsule{
			Item:        1,
			Probability: float64(20),
		},
		Capsule{
			Item:        2,
			Probability: float64(40),
		},
		Capsule{
			Item:        3,
			Probability: float64(35),
		},
	)

	t.Log("After gacha.Len(): ", gacha.Len())
	t.Log("After gacha.Empty(): ", gacha.Empty())

	try := 999999

	for i := 0; i < try; i++ {
		ret, _ := gacha.Peek()

		idx := ret.Item.(int)
		stat[idx] = stat[idx] + 1
	}

	for idx, cnt := range stat {
		t.Log("item:", idx, ",", (float64(cnt)/float64(try))*100, "%")
	}

	for i := 0; i < len(stat)+2; i++ {
		t.Log(gacha.Get())
	}

	t.Log("Last gacha.Len(): ", gacha.Len())
	t.Log("Last gacha.Empty(): ", gacha.Empty())
}

func TestLotto(t *testing.T) {
	entry := make([]Capsule, 45)
	for i := 1; i <= 45; i++ {
		entry[i-1] = Capsule{
			Item:        i,
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
