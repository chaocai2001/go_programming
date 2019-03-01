package fast_hashmap

import (
	"strconv"
	"testing"
)

const NumberOfKey = 5000000
const StartNum = 1000000000000000

var hasShown = false

func BenchmarkBuiltInMap(t *testing.B) {
	//debug.SetGCPercent(-1)
	m := make(map[string]int)
	var k string
	for i := StartNum; i < StartNum+NumberOfKey; i++ {
		k = strconv.Itoa(i)
		m[k] = i
	}
	//var ti int

	t.Run("BenchmarkBuiltInMap_Get", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := StartNum; i < StartNum+NumberOfKey; i++ {
				k = strconv.Itoa(i)
				v, _ := m[k]

				if v != i {
					t.Error("Wrong!")
				}
			}
		}
	})
}

// func BenchmarkFastMadeHashMap(t *testing.B) {
// 	debug.SetGCPercent(-1)
// 	m := NewFastMap(NumberOfKey)
// 	var k string
// 	for i := StartNum; i < StartNum+NumberOfKey; i++ {
// 		k = strconv.Itoa(i)
// 		m.Put(k, i)
//
// 	}
// 	//var ti int
// 	t.ResetTimer()
// 	for i := StartNum; i < StartNum+NumberOfKey; i++ {
// 		k = strconv.Itoa(i)
// 		v, _ := m.Get(k)
//
// 		if v.(int) != i {
// 			t.Error("Wrong!")
// 		}
// 	}
// }

func BenchmarkFasterHashMap(t *testing.B) {
	//	debug.SetGCPercent(-1)
	m := NewFasterMap(NumberOfKey)
	var k string
	for i := StartNum; i < StartNum+NumberOfKey; i++ {
		k = strconv.Itoa(i)
		m.Put(k, i)

	}
	t.Run("BenchmarkFasterHashMap_Get", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := StartNum; i < StartNum+NumberOfKey; i++ {
				k = strconv.Itoa(i)
				v, _ := m.Get(k)

				if v.(int) != i {
					t.Error("Wrong!")
				}
			}
			if !hasShown {
				t.Logf("Number of collisions: %d\n", m.numOfCollisions)
				hasShown = true
			}
		}
	})
}
