// BanglaCode AST - Abstract Syntax Tree nodes

export interface Node {
  tokenLiteral(): string;
}

export interface Statement extends Node {
  statementNode(): void;
}

export interface Expression extends Node {
  expressionNode(): void;
}

// Program is the root node of every AST
export class Program implements Node {
  statements: Statement[] = [];

  tokenLiteral(): string {
    if (this.statements.length > 0) {
      return this.statements[0].tokenLiteral();
    }
    return "";
  }
}

// Identifier
export class Identifier implements Expression {
  value: string;

  constructor(value: string) {
    this.value = value;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return this.value;
  }
}

// Number Literal
export class NumberLiteral implements Expression {
  value: number;

  constructor(value: number) {
    this.value = value;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return String(this.value);
  }
}

// String Literal
export class StringLiteral implements Expression {
  value: string;

  constructor(value: string) {
    this.value = value;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return this.value;
  }
}

// Boolean Literal
export class BooleanLiteral implements Expression {
  value: boolean;

  constructor(value: boolean) {
    this.value = value;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return this.value ? "sotti" : "mittha";
  }
}

// Null Literal
export class NullLiteral implements Expression {
  expressionNode(): void {}
  tokenLiteral(): string {
    return "khali";
  }
}

// Array Literal
export class ArrayLiteral implements Expression {
  elements: Expression[] = [];

  expressionNode(): void {}
  tokenLiteral(): string {
    return "[...]";
  }
}

// Map/Object Literal
export class MapLiteral implements Expression {
  pairs: Map<Expression, Expression> = new Map();

  expressionNode(): void {}
  tokenLiteral(): string {
    return "{...}";
  }
}

// Prefix Expression: !x, -x, na x
export class PrefixExpression implements Expression {
  operator: string;
  right: Expression;

  constructor(operator: string, right: Expression) {
    this.operator = operator;
    this.right = right;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return this.operator;
  }
}

// Infix Expression: a + b, a ebong b
export class InfixExpression implements Expression {
  left: Expression;
  operator: string;
  right: Expression;

  constructor(left: Expression, operator: string, right: Expression) {
    this.left = left;
    this.operator = operator;
    this.right = right;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return this.operator;
  }
}

// Index Expression: arr[0], obj["key"]
export class IndexExpression implements Expression {
  left: Expression;
  index: Expression;

  constructor(left: Expression, index: Expression) {
    this.left = left;
    this.index = index;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return "[";
  }
}

// Member Expression: obj.property
export class MemberExpression implements Expression {
  object: Expression;
  property: Identifier;

  constructor(object: Expression, property: Identifier) {
    this.object = object;
    this.property = property;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return ".";
  }
}

// Call Expression: fn(args)
export class CallExpression implements Expression {
  function: Expression;
  arguments: Expression[] = [];

  constructor(fn: Expression, args: Expression[] = []) {
    this.function = fn;
    this.arguments = args;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return "(";
  }
}

// If Expression: jodi (condition) { } nahole { }
export class IfExpression implements Expression {
  condition: Expression;
  consequence: BlockStatement;
  alternative: BlockStatement | IfExpression | null = null;

  constructor(condition: Expression, consequence: BlockStatement) {
    this.condition = condition;
    this.consequence = consequence;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return "jodi";
  }
}

// Function Literal: kaj name(params) { body }
export class FunctionLiteral implements Expression {
  name: string | null = null;
  parameters: Identifier[] = [];
  body: BlockStatement;

  constructor(body: BlockStatement) {
    this.body = body;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return "kaj";
  }
}

// Class Literal: sreni Name { ... }
export class ClassLiteral implements Expression {
  name: Identifier;
  constructor_: FunctionLiteral | null = null;
  methods: Map<string, FunctionLiteral> = new Map();

  constructor(name: Identifier) {
    this.name = name;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return "sreni";
  }
}

// New Expression: notun ClassName(args)
export class NewExpression implements Expression {
  class_: Expression;
  arguments: Expression[] = [];

  constructor(class_: Expression) {
    this.class_ = class_;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return "notun";
  }
}

// This Expression: ei
export class ThisExpression implements Expression {
  expressionNode(): void {}
  tokenLiteral(): string {
    return "ei";
  }
}

// Assignment Expression: x = 5, x += 1
export class AssignmentExpression implements Expression {
  left: Expression;
  operator: string;
  value: Expression;

  constructor(left: Expression, operator: string, value: Expression) {
    this.left = left;
    this.operator = operator;
    this.value = value;
  }

  expressionNode(): void {}
  tokenLiteral(): string {
    return this.operator;
  }
}

// Block Statement: { statements }
export class BlockStatement implements Statement {
  statements: Statement[] = [];

  statementNode(): void {}
  tokenLiteral(): string {
    return "{";
  }
}

// Let Statement: dhoro x = 5;
export class LetStatement implements Statement {
  name: Identifier;
  value: Expression;

  constructor(name: Identifier, value: Expression) {
    this.name = name;
    this.value = value;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "dhoro";
  }
}

// Return Statement: ferao x;
export class ReturnStatement implements Statement {
  value: Expression | null = null;

  statementNode(): void {}
  tokenLiteral(): string {
    return "ferao";
  }
}

// Expression Statement: just an expression as statement
export class ExpressionStatement implements Statement {
  expression: Expression;

  constructor(expression: Expression) {
    this.expression = expression;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return this.expression.tokenLiteral();
  }
}

// While Statement: jotokkhon (condition) { body }
export class WhileStatement implements Statement {
  condition: Expression;
  body: BlockStatement;

  constructor(condition: Expression, body: BlockStatement) {
    this.condition = condition;
    this.body = body;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "jotokkhon";
  }
}

// For Statement: ghuriye (init; condition; update) { body }
export class ForStatement implements Statement {
  init: Statement | null = null;
  condition: Expression | null = null;
  update: Expression | null = null;
  body: BlockStatement;

  constructor(body: BlockStatement) {
    this.body = body;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "ghuriye";
  }
}

// Break Statement: thamo
export class BreakStatement implements Statement {
  statementNode(): void {}
  tokenLiteral(): string {
    return "thamo";
  }
}

// Continue Statement: chharo
export class ContinueStatement implements Statement {
  statementNode(): void {}
  tokenLiteral(): string {
    return "chharo";
  }
}

// Try Statement: chesta { } dhoro_bhul (e) { } shesh { }
export class TryStatement implements Statement {
  tryBlock: BlockStatement;
  catchParam: Identifier | null = null;
  catchBlock: BlockStatement | null = null;
  finallyBlock: BlockStatement | null = null;

  constructor(tryBlock: BlockStatement) {
    this.tryBlock = tryBlock;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "chesta";
  }
}

// Throw Statement: felo "error"
export class ThrowStatement implements Statement {
  value: Expression;

  constructor(value: Expression) {
    this.value = value;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "felo";
  }
}

// Import Statement: ano "module.bang" [hisabe alias]
export class ImportStatement implements Statement {
  path: StringLiteral;
  alias: Identifier | null = null;

  constructor(path: StringLiteral) {
    this.path = path;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "ano";
  }
}

// Export Statement: pathao kaj/sreni/dhoro
export class ExportStatement implements Statement {
  statement: Statement;

  constructor(statement: Statement) {
    this.statement = statement;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "pathao";
  }
}

// Function Declaration: kaj name(params) { body }
export class FunctionDeclaration implements Statement {
  name: Identifier;
  parameters: Identifier[] = [];
  body: BlockStatement;

  constructor(name: Identifier, body: BlockStatement) {
    this.name = name;
    this.body = body;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "kaj";
  }
}

// Class Declaration: sreni Name { ... }
export class ClassDeclaration implements Statement {
  name: Identifier;
  constructor_: FunctionLiteral | null = null;
  methods: Map<string, FunctionLiteral> = new Map();

  constructor(name: Identifier) {
    this.name = name;
  }

  statementNode(): void {}
  tokenLiteral(): string {
    return "sreni";
  }
}
