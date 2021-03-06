# make, new 操作

make はビルトイン型(map, slice, channel)のメモリ割り当て。  
new は各型のメモリを割り当てる。

ビルトイン関数 new は他の言語で使われる同名の関数と同じ機能。  
new(T)はゼロサプレスされた T 型のメモリ空間を割り当て、そのアドレスを返す。  
すなわち `*T 型`の値となる。  
Go の専門用語で言えば、ポインタを返すということ。  
新たに割り当てられた型 T はゼロ値となる。 `new はポインタを返す。`

ビルトイン関数 make(T, args)と new(T)は異なる機能をもつ。  
make は slice、map または channel を作成し、初期値（非ゼロ値）を持つ T 型を返すのみ、`*T` ではない。  
本質的には、この３つの型が異なる点はデータ構造を指し示す参照が使用される前に初期化されているということ。  
例えば、データ（内部 array）を指し示すポインタ、長さ，容量による３点で記述される slice の各項目が初期化される前は、slice は nil。  
slice, map, channel にとって、make は内部のデータ構造を初期化し、適当な値で埋め尽くされる。 `make は初期化後の（非ゼロの）値を返す。`
