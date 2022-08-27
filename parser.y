%{
package mylang
%}

%union {
  token
  *program
  statement
  expression
}

%token<token>
  IDENT  // foo
  INT    // 12345
  FLOAT  // 123.45

%token<token>
  ADD // +
  SUB // -
  MUL // *
  DIV // /
  REM // %

%token<token>
  ASSIGN // =

%token<token>
  SEMICOLON // ;
  LPAREN // (
  RPAREN // )

%token<token>
  PRINT // print

%type<program> statements
%type<statement> statement
%type<expression> expression
%type<expression> primary

%%
statements:
    statement SEMICOLON {
      $$ = &program{statements: []statement{$1}}
      yylex.(*lexer).result = $$
    }
  |
    statements statement {
        $1.statements = append($1.statements, $2)
        $$ = $1
      }

statement:
    IDENT ASSIGN expression {
      $$ = &assignStatement{identifier: $1, operator: ASSIGN, expr: $3}
    }
  |
    PRINT IDENT {
      $$ = &printStatement{identifier: $2}
    }

expression: expression ADD expression {
    $$ = binaryExpression{$1, ADD, $3}
    }
  |
    expression SUB expression {
      $$ = binaryExpression{$1, SUB, $3}
    }
  |
    expression MUL expression {
      $$ = binaryExpression{$1, MUL, $3}
    }
  |
    expression DIV expression {
      $$ = binaryExpression{$1, DIV, $3}
    }
  |
    expression REM expression {
      $$ = binaryExpression{$1, REM, $3}
    }
  |
    primary

primary: INT {
       $$ = basicLiteral{kind: INT, literal: $1.literal}
    }
  |
    FLOAT {
      $$ = basicLiteral{kind: FLOAT, literal: $1.literal}
    }
  |
    LPAREN expression RPAREN {
      $$ = $1
    }
%%
