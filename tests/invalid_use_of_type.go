//po:MsgId "invalid use of type"
//
//PoMessage {
//	#: go/gofrontend/expressions.cc:823
//	#, fuzzy
//	msgid "invalid use of type\n"
//	msgstr "类型使用无效\n"
//}

package p

type T int

func F() func () T {
	return func() T
}
