package main

func testSync() {
	printSubSection("syncMap", testSyncMap)
	printSubSection("syncPool", testSyncPool)
	printSubSection("singleflight", testSingleflight)
}

func testSyncMap() {
	// panic: go-generics/test_sync.go:10:16: expected ';', found '[' (and 1 more errors)
	// snippet(func() interface{} {
	//   var m syncMap[string, int]
	//   m.Store("foo", 1)
	//   m.Store("bar", 2)
	//
	//   v, _ := m.Load("bar")
	//   return v
	// })
	//
	// snippet(func() interface{} {
	//   var m syncMap[string, int]
	//
	//   v, _ := m.Load("bar")
	//   return v
	// })
}

func testSyncPool() {
	snippet(func() interface{} {
		pool := newSyncPool(func() int { return 1 })

		v := pool.Get()
		defer pool.Put(v)

		return v
	})
}

func testSyncSingleflight() {
	snippet(func() interface{} {
		pool := newSyncPool(func() int { return 1 })

		v := pool.Get()
		defer pool.Put(v)

		return v
	})
}

func testSingleflight() {
	// panic: runtime error: index out of range [1] with length 1
	// snippet(func() interface{} {
	//   var g singleflightGroup[int]
	//
	//   v, _, _ := g.Do("foo", func() (int, error) {
	//     return 1, nil
	//   })
	//
	//   return v
	// })
}
