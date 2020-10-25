package sec2

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w = iota // グループ化して宣言した場合、w = iotaと同義に扱われるため、 w == 3 となる
)

const v = iota // const キーワードが出現するたびにiotaは置き換わるので v == 0 となる

const (
	e, f, g = iota, iota, iota // e = 0, f = 0, g = 0 iotaの同一行は同じ値となる
)
