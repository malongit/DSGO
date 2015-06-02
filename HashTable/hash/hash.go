package hash

<<<<<<< HEAD
=======
type HashTable interface {
	Size() int
	IsEmpty() bool
	Insert(key string) bool
	Search(key string) bool
	Remove(key string) bool
}

//以下算法搜集自网络，基本思路类似，效果相近

>>>>>>> tmp
//by Brian Kernighan & Dennis Ritchie
func BKDRhash(str string) uint {
	const factor uint = 131 //31、131、1313、 13131、131313
	var data = []byte(str)
	var code uint = 0
	for _, ch := range data {
		code = code*factor + uint(ch)
	}
	return code
}

//from SDBM
func SDBMhash(str string) uint {
	var data = []byte(str)
	var code uint = 0
	for _, ch := range data {
		code = code*65599 + uint(ch)
	}
	return code
}

//by Daniel J. Bernstein
func DJBhash(str string) uint {
	var data = []byte(str)
	var code uint = 5381
	for _, ch := range data {
		code = code*33 + uint(ch)
	}
	return code
}
func DJB2hash(str string) uint {
	var data = []byte(str)
	var code uint = 5381
	for _, ch := range data {
		code = code*33 ^ uint(ch)
	}
	return code
}

//from Unix
func FNVhash(str string) uint {
	var data = []byte(str)

	var code uint = 2166136261
	for _, ch := range data {
		code = code*16777619 ^ uint(ch)
	}
	return code
}

//Robert Sedgwicks
func RShash(str string) uint {
	var data = []byte(str)

	var magic uint = 63689
	var code uint = 0
	for _, ch := range data {
		code = code*magic + uint(ch)
		magic *= 378551
	}
	return code
}

//by Justin Sobel
func JShash(str string) uint {
	var data = []byte(str)

	var code uint = 1315423911
	for _, ch := range data {
		code ^= (code << 5) + uint(ch) + (code >> 2)
	}
	return code
}

//by Arash Partow
func APhash(str string) uint {
	var data = []byte(str)

	var code uint = 0
	var tick = true
	for _, ch := range data {
		if tick {
			code ^= (code << 7) ^ uint(ch) ^ (code >> 3)
		} else {
			code ^= ^((code << 11) ^ uint(ch) ^ (code >> 5))
		}
		tick = !tick
	}
	return code
}

/*
Bible.txt
<<<<<<< HEAD
BKDRhash: [max=6] [26056 16619 5205 1065 183 29] [31102/49157]
SDBMhash: [max=6] [26144 16405 5342 1079 160 27] [31102/49157]
DJBhash:  [max=6] [26125 16473 5268 1089 185 17] [31102/49157]
DJB2hash: [max=7] [26161 16448 5235 1104 177 32] [31102/49157]
FNVhash:  [max=5] [26079 16583 5184 1114 176 21] [31102/49157]
RShash:   [max=6] [26099 16527 5251 1077 176 27] [31102/49157]
JShash:   [max=6] [26182 16404 5247 1126 166 32] [31102/49157]
APhash:   [max=6] [26034 16613 5246 1084 159 21] [31102/49157]

圣经.txt
BKDRhash: [max=5] [26203 16350 5296 1097 186 25] [31102/49157]
SDBMhash: [max=6] [26165 16401 5309 1077 177 28] [31102/49157]
DJBhash:  [max=6] [26082 16606 5148 1109 188 24] [31102/49157]
DJB2hash: [max=5] [26197 16390 5234 1118 200 18] [31102/49157]
FNVhash:  [max=8] [26202 16373 5250 1127 183 22] [31102/49157]
RShash:   [max=7] [26086 16581 5198 1088 164 40] [31102/49157]
JShash:   [max=6] [26109 16542 5184 1124 173 25] [31102/49157]
APhash:   [max=6] [26126 16527 5182 1110 183 29] [31102/49157]

圣经UTF8.txt
BKDRhash: [max=7] [26094 16534 5227 1123 153 26] [31102/49157]
SDBMhash: [max=6] [26022 16610 5308 1023 166 28] [31102/49157]
DJBhash:  [max=6] [26066 16551 5255 1119 148 18] [31102/49157]
DJB2hash: [max=7] [26059 16666 5096 1137 170 29] [31102/49157]
FNVhash:  [max=6] [26134 16459 5281 1083 172 28] [31102/49157]
RShash:   [max=6] [26147 16466 5220 1127 174 23] [31102/49157]
JShash:   [max=6] [26013 16641 5239 1093 152 19] [31102/49157]
APhash:   [max=6] [26085 16521 5295 1058 175 23] [31102/49157]
=======
BKDRhash: [max=6] [16429 5204 1170 158 21 3]	[31102/49157]
SDBMhash: [max=7] [16407 5248 1100 188 22 6]	[31102/49157]
DJBhash:  [max=6] [16396 5230 1129 184 21 3]	[31102/49157]
DJB2hash: [max=6] [16422 5211 1127 181 27 3]	[31102/49157]
FNVhash:  [max=7] [16609 5155 1115 168 27 5]	[31102/49157]
RShash:   [max=6] [16727 5100 1127 158 30 2]	[31102/49157]
JShash:   [max=6] [16602 5216 1061 191 23 1]	[31102/49157]
APhash:   [max=6] [16641 5239 1055 171 22 4]	[31102/49157]

圣经.txt
BKDRhash: [max=6] [16416 5215 1120 191 24 2]	[31102/49157]
SDBMhash: [max=6] [16607 5252 1068 164 25 1]	[31102/49157]
DJBhash:  [max=6] [16551 5273 1080 167 17 2]	[31102/49157]
DJB2hash: [max=5] [16353 5291 1107 189 18 0]	[31102/49157]
FNVhash:  [max=6] [16491 5221 1120 168 25 2]	[31102/49157]
RShash:   [max=7] [16403 5308 1092 167 24 3]	[31102/49157]
JShash:   [max=7] [16575 5274 1069 165 21 1]	[31102/49157]
APhash:   [max=6] [16365 5305 1104 171 25 1]	[31102/49157]

圣经UTF8.txt
BKDRhash: [max=6] [16562 5326 1049 154 19 5]	[31102/49157]
SDBMhash: [max=6] [16521 5312 1056 172 19 1]	[31102/49157]
DJBhash:  [max=6] [16688 5219 1071 163 21 1]	[31102/49157]
DJB2hash: [max=6] [16525 5180 1129 167 30 2]	[31102/49157]
FNVhash:  [max=6] [16425 5270 1091 177 24 6]	[31102/49157]
RShash:   [max=6] [16487 5195 1133 173 22 4]	[31102/49157]
JShash:   [max=6] [16543 5264 1098 164 15 1]	[31102/49157]
APhash:   [max=6] [16515 5187 1113 177 32 1]	[31102/49157]
>>>>>>> tmp
*/
