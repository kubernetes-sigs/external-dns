// Package ini is an LL(1) parser for configuration files.
//
//	Example:
//	sections, err := ini.OpenFile("/path/to/file")
//	if err != nil {
//		panic(err)
//	}
//
//	profile := "foo"
//	section, ok := sections.GetSection(profile)
//	if !ok {
//		fmt.Printf("section %q could not be found", profile)
//	}
//
// Below is the BNF that describes this parser
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//  Grammar:
//  stmt -> section | stmt'
//  stmt' -> epsilon | expr
//  expr -> value (stmt)* | equal_expr (stmt)*
//  equal_expr -> value ( ':' | '=' ) equal_expr'
//  equal_expr' -> number | string | quoted_string
//  quoted_string -> " quoted_string'
//  quoted_string' -> string quoted_string_end
//  quoted_string_end -> "
//
//  section -> [ section'
//  section' -> section_value section_close
//  section_value -> number | string_subset | boolean | quoted_string_subset
//  quoted_string_subset -> " quoted_string_subset'
//  quoted_string_subset' -> string_subset quoted_string_end
//  quoted_string_subset -> "
//  section_close -> ]
//
//  value -> number | string_subset | boolean
//  string -> ? UTF-8 Code-Points except '\n' (U+000A) and '\r\n' (U+000D U+000A) ?
//  string_subset -> ? Code-points excepted by <string> grammar except ':' (U+003A), '=' (U+003D), '[' (U+005B), and ']' (U+005D) ?
//
//  SkipState will skip (NL WS)+
//
//  comment -> # comment' | ; comment'
//  comment' -> epsilon | value
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//	Grammar:
//	stmt -> value stmt'
//	stmt' -> epsilon | op stmt
//	value -> number | string | boolean | quoted_string
||||||| parent of 5ce8c7613 (update vendored files)
//	Grammar:
//	stmt -> value stmt'
//	stmt' -> epsilon | op stmt
//	value -> number | string | boolean | quoted_string
=======
//  Grammar:
//  stmt -> section | stmt'
//  stmt' -> epsilon | expr
//  expr -> value (stmt)* | equal_expr (stmt)*
//  equal_expr -> value ( ':' | '=' ) equal_expr'
//  equal_expr' -> number | string | quoted_string
//  quoted_string -> " quoted_string'
//  quoted_string' -> string quoted_string_end
//  quoted_string_end -> "
>>>>>>> 5ce8c7613 (update vendored files)
//
//  section -> [ section'
//  section' -> section_value section_close
//  section_value -> number | string_subset | boolean | quoted_string_subset
//  quoted_string_subset -> " quoted_string_subset'
//  quoted_string_subset' -> string_subset quoted_string_end
//  quoted_string_subset -> "
//  section_close -> ]
//
//  value -> number | string_subset | boolean
//  string -> ? UTF-8 Code-Points except '\n' (U+000A) and '\r\n' (U+000D U+000A) ?
//  string_subset -> ? Code-points excepted by <string> grammar except ':' (U+003A), '=' (U+003D), '[' (U+005B), and ']' (U+005D) ?
//
<<<<<<< HEAD
//	comment -> # comment' | ; comment'
//	comment' -> epsilon | value
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
//	comment -> # comment' | ; comment'
//	comment' -> epsilon | value
=======
//  SkipState will skip (NL WS)+
//
//  comment -> # comment' | ; comment'
//  comment' -> epsilon | value
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//	Grammar:
//	stmt -> value stmt'
//	stmt' -> epsilon | op stmt
//	value -> number | string | boolean | quoted_string
||||||| parent of 6b7ce455e (update vendored files)
//	Grammar:
//	stmt -> value stmt'
//	stmt' -> epsilon | op stmt
//	value -> number | string | boolean | quoted_string
=======
//  Grammar:
//  stmt -> section | stmt'
//  stmt' -> epsilon | expr
//  expr -> value (stmt)* | equal_expr (stmt)*
//  equal_expr -> value ( ':' | '=' ) equal_expr'
//  equal_expr' -> number | string | quoted_string
//  quoted_string -> " quoted_string'
//  quoted_string' -> string quoted_string_end
//  quoted_string_end -> "
>>>>>>> 6b7ce455e (update vendored files)
//
//  section -> [ section'
//  section' -> section_value section_close
//  section_value -> number | string_subset | boolean | quoted_string_subset
//  quoted_string_subset -> " quoted_string_subset'
//  quoted_string_subset' -> string_subset quoted_string_end
//  quoted_string_subset -> "
//  section_close -> ]
//
//  value -> number | string_subset | boolean
//  string -> ? UTF-8 Code-Points except '\n' (U+000A) and '\r\n' (U+000D U+000A) ?
//  string_subset -> ? Code-points excepted by <string> grammar except ':' (U+003A), '=' (U+003D), '[' (U+005B), and ']' (U+005D) ?
//
<<<<<<< HEAD
//	comment -> # comment' | ; comment'
//	comment' -> epsilon | value
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
//	comment -> # comment' | ; comment'
//	comment' -> epsilon | value
=======
//  SkipState will skip (NL WS)+
//
//  comment -> # comment' | ; comment'
//  comment' -> epsilon | value
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//	Grammar:
//	stmt -> value stmt'
//	stmt' -> epsilon | op stmt
//	value -> number | string | boolean | quoted_string
||||||| parent of 4d7e5ad26 (update vendored files)
//	Grammar:
//	stmt -> value stmt'
//	stmt' -> epsilon | op stmt
//	value -> number | string | boolean | quoted_string
=======
//  Grammar:
//  stmt -> section | stmt'
//  stmt' -> epsilon | expr
//  expr -> value (stmt)* | equal_expr (stmt)*
//  equal_expr -> value ( ':' | '=' ) equal_expr'
//  equal_expr' -> number | string | quoted_string
//  quoted_string -> " quoted_string'
//  quoted_string' -> string quoted_string_end
//  quoted_string_end -> "
>>>>>>> 4d7e5ad26 (update vendored files)
//
//  section -> [ section'
//  section' -> section_value section_close
//  section_value -> number | string_subset | boolean | quoted_string_subset
//  quoted_string_subset -> " quoted_string_subset'
//  quoted_string_subset' -> string_subset quoted_string_end
//  quoted_string_subset -> "
//  section_close -> ]
//
//  value -> number | string_subset | boolean
//  string -> ? UTF-8 Code-Points except '\n' (U+000A) and '\r\n' (U+000D U+000A) ?
//  string_subset -> ? Code-points excepted by <string> grammar except ':' (U+003A), '=' (U+003D), '[' (U+005B), and ']' (U+005D) ?
//
<<<<<<< HEAD
//	comment -> # comment' | ; comment'
//	comment' -> epsilon | value
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
//	comment -> # comment' | ; comment'
//	comment' -> epsilon | value
=======
//  SkipState will skip (NL WS)+
//
//  comment -> # comment' | ; comment'
//  comment' -> epsilon | value
>>>>>>> 4d7e5ad26 (update vendored files)
package ini
