%{
// Copyright (c) 2020-2022, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package json

type Placeholder struct {
	Num  int
	Name string
}

type Operator struct {
	Name string
	Params   []any
}

type member struct {
	key   string
	value any
}

func finalize(l yyLexer, value any) {
	l.(*json).value = value
}
%}

%union {
  object map[string]any
  member member
  array  []any
  string string
  value  any
}

%start json

%token <value>   TRUE FALSE NULL NUMBER PLACEHOLDER SUB_PARSER
%token <string>  STRING OPERATOR

%type <object>   object members
%type <member>   member
%type <array>    array elements op_params
%type <value>    json value operator

%%

json: value
                {
                  $$ = $1
                  finalize(yylex, $$)
                }

value:
    object      { $$ = $1 }
  | array       { $$ = $1 }
  | operator    { $$ = $1 }
  | SUB_PARSER  { $$ = $1 }
  | STRING      { $$ = $1 }
  | NUMBER
  | TRUE
  | FALSE
  | NULL
  | PLACEHOLDER
  ;

object: '{' '}'
                {
                  $$ = map[string]any{}
                }
  | '{' members '}'
                {
                  $$ = $2
                }
  | '{' members ',' '}' // not JSON spec but useful
                {
                  $$ = $2
                }

members: member
                {
                  $$ = map[string]any{
                    $1.key: $1.value,
                  }
                }
  | members ',' member
                {
                  $1[$3.key] = $3.value
                  $$ = $1
                }

member: STRING ':' value
                {
                  $$ = member{
                    key:   $1,
                    value: $3,
                  }
                }

array: '[' ']'
                {
                  $$ = []any{}
                }
  | '[' elements ']'
                {
                  $$ = $2
                }
  | '[' elements ',' ']' // not JSON spec but useful
                {
                  $$ = $2
                }

elements: value
                {
                  $$ = []any{$1}
                }
  | elements ',' value
                {
                  $$ = append($1, $3)
                }

op_params: '(' ')'
                {
                  $$ = []any{}
                }
  | '(' elements ')'
                {
                  $$ = $2
                }
  | '(' elements ',' ')'
                {
                  $$ = $2
                }

operator:
  OPERATOR op_params
                {
                  op := yylex.(*json).newOperator($1, $2)
                   if op == nil {
                    return 1
                  }
                  $$ = op
                }
  | OPERATOR
                {
                  op := yylex.(*json).newOperator($1, nil)
                  if op == nil {
                    return 1
                  }
                  $$ = op
                }
%%
