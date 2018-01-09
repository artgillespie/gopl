package ex4

// Rotate rotates the slice a by n positions.
func Rotate(a []int, n int) {
	// rotation wraps around at len(a)
	p := n % len(a)
	// no rotation required
	if len(a) == 0 || n == 0 {
		return
	}

	if p < 0 {
		tmp := make([]int, -p)
		copy(tmp, a[:-p])
		copy(a[:], a[-p:])
		copy(a[len(a)+p:], tmp)
	} else {
		tmp := make([]int, p)
		copy(tmp, a[len(a)-p:])
		copy(a[p:], a[:len(a)-p])
		copy(a[:], tmp)
	}

}
