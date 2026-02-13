// BanglaCode Parser - Builds AST from tokens

import { Lexer, Token, TokenType } from "./lexer";
import * as ast from "./ast";

// Precedence levels
enum Precedence {
  LOWEST = 1,
  ASSIGN,      // =, +=, -=, *=, /=
  OR,          // ba (||)
  AND,         // ebong (&&)
  EQUALS,      // ==, !=
  LESSGREATER, // <, >, <=, >=
  SUM,         // +, -
  PRODUCT,     // *, /, %
  POWER,       // **
  PREFIX,      // -x, !x, na x
  CALL,        // fn()
  INDEX,       // arr[0], obj.prop
}

// Token to precedence mapping
const precedences: Partial<Record<TokenType, Precedence>> = {
  [TokenType.ASSIGN]: Precedence.ASSIGN,
  [TokenType.PLUS_ASSIGN]: Precedence.ASSIGN,
  [TokenType.MINUS_ASSIGN]: Precedence.ASSIGN,
  [TokenType.ASTERISK_ASSIGN]: Precedence.ASSIGN,
  [TokenType.SLASH_ASSIGN]: Precedence.ASSIGN,
  [TokenType.BA]: Precedence.OR,
  [TokenType.EBONG]: Precedence.AND,
  [TokenType.EQ]: Precedence.EQUALS,
  [TokenType.NOT_EQ]: Precedence.EQUALS,
  [TokenType.LT]: Precedence.LESSGREATER,
  [TokenType.GT]: Precedence.LESSGREATER,
  [TokenType.LT_EQ]: Precedence.LESSGREATER,
  [TokenType.GT_EQ]: Precedence.LESSGREATER,
  [TokenType.PLUS]: Precedence.SUM,
  [TokenType.MINUS]: Precedence.SUM,
  [TokenType.ASTERISK]: Precedence.PRODUCT,
  [TokenType.SLASH]: Precedence.PRODUCT,
  [TokenType.PERCENT]: Precedence.PRODUCT,
  [TokenType.POWER]: Precedence.POWER,
  [TokenType.LPAREN]: Precedence.CALL,
  [TokenType.LBRACKET]: Precedence.INDEX,
  [TokenType.DOT]: Precedence.INDEX,
};

type PrefixParseFn = () => ast.Expression | null;
type InfixParseFn = (left: ast.Expression) => ast.Expression | null;

export class Parser {
  private lexer: Lexer;
  private curToken: Token;
  private peekToken: Token;
  errors: string[] = [];

  private prefixParseFns: Map<TokenType, PrefixParseFn> = new Map();
  private infixParseFns: Map<TokenType, InfixParseFn> = new Map();

  constructor(lexer: Lexer) {
    this.lexer = lexer;
    this.curToken = { type: TokenType.EOF, literal: "", line: 0, column: 0 };
    this.peekToken = { type: TokenType.EOF, literal: "", line: 0, column: 0 };

    // Register prefix parsers
    this.registerPrefix(TokenType.IDENTIFIER, this.parseIdentifier.bind(this));
    this.registerPrefix(TokenType.NUMBER, this.parseNumberLiteral.bind(this));
    this.registerPrefix(TokenType.STRING, this.parseStringLiteral.bind(this));
    this.registerPrefix(TokenType.SOTTI, this.parseBooleanLiteral.bind(this));
    this.registerPrefix(TokenType.MITTHA, this.parseBooleanLiteral.bind(this));
    this.registerPrefix(TokenType.KHALI, this.parseNullLiteral.bind(this));
    this.registerPrefix(TokenType.BANG, this.parsePrefixExpression.bind(this));
    this.registerPrefix(TokenType.NA, this.parsePrefixExpression.bind(this));
    this.registerPrefix(TokenType.MINUS, this.parsePrefixExpression.bind(this));
    this.registerPrefix(TokenType.LPAREN, this.parseGroupedExpression.bind(this));
    this.registerPrefix(TokenType.LBRACKET, this.parseArrayLiteral.bind(this));
    this.registerPrefix(TokenType.LBRACE, this.parseMapLiteral.bind(this));
    this.registerPrefix(TokenType.JODI, this.parseIfExpression.bind(this));
    this.registerPrefix(TokenType.KAJ, this.parseFunctionLiteral.bind(this));
    this.registerPrefix(TokenType.NOTUN, this.parseNewExpression.bind(this));
    this.registerPrefix(TokenType.EI, this.parseThisExpression.bind(this));

    // Register infix parsers
    this.registerInfix(TokenType.PLUS, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.MINUS, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.ASTERISK, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.SLASH, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.PERCENT, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.POWER, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.EQ, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.NOT_EQ, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.LT, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.GT, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.LT_EQ, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.GT_EQ, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.EBONG, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.BA, this.parseInfixExpression.bind(this));
    this.registerInfix(TokenType.LPAREN, this.parseCallExpression.bind(this));
    this.registerInfix(TokenType.LBRACKET, this.parseIndexExpression.bind(this));
    this.registerInfix(TokenType.DOT, this.parseMemberExpression.bind(this));
    this.registerInfix(TokenType.ASSIGN, this.parseAssignmentExpression.bind(this));
    this.registerInfix(TokenType.PLUS_ASSIGN, this.parseAssignmentExpression.bind(this));
    this.registerInfix(TokenType.MINUS_ASSIGN, this.parseAssignmentExpression.bind(this));
    this.registerInfix(TokenType.ASTERISK_ASSIGN, this.parseAssignmentExpression.bind(this));
    this.registerInfix(TokenType.SLASH_ASSIGN, this.parseAssignmentExpression.bind(this));

    // Read two tokens to initialize curToken and peekToken
    this.nextToken();
    this.nextToken();
  }

  private registerPrefix(type: TokenType, fn: PrefixParseFn): void {
    this.prefixParseFns.set(type, fn);
  }

  private registerInfix(type: TokenType, fn: InfixParseFn): void {
    this.infixParseFns.set(type, fn);
  }

  private nextToken(): void {
    this.curToken = this.peekToken;
    this.peekToken = this.lexer.nextToken();
  }

  private curTokenIs(type: TokenType): boolean {
    return this.curToken.type === type;
  }

  private peekTokenIs(type: TokenType): boolean {
    return this.peekToken.type === type;
  }

  private expectPeek(type: TokenType): boolean {
    if (this.peekTokenIs(type)) {
      this.nextToken();
      return true;
    }
    this.peekError(type);
    return false;
  }

  private peekError(type: TokenType): void {
    this.errors.push(
      `line ${this.peekToken.line}: expected ${type}, got ${this.peekToken.type} instead`
    );
  }

  private noPrefixParseFnError(type: TokenType): void {
    this.errors.push(`line ${this.curToken.line}: no prefix parse function for ${type}`);
  }

  private peekPrecedence(): Precedence {
    return precedences[this.peekToken.type] || Precedence.LOWEST;
  }

  private curPrecedence(): Precedence {
    return precedences[this.curToken.type] || Precedence.LOWEST;
  }

  parseProgram(): ast.Program {
    const program = new ast.Program();

    while (!this.curTokenIs(TokenType.EOF)) {
      const stmt = this.parseStatement();
      if (stmt) {
        program.statements.push(stmt);
      }
      this.nextToken();
    }

    return program;
  }

  private parseStatement(): ast.Statement | null {
    switch (this.curToken.type) {
      case TokenType.DHORO:
        return this.parseLetStatement();
      case TokenType.FERAO:
        return this.parseReturnStatement();
      case TokenType.JOTOKKHON:
        return this.parseWhileStatement();
      case TokenType.GHURIYE:
        return this.parseForStatement();
      case TokenType.THAMO:
        return this.parseBreakStatement();
      case TokenType.CHHARO:
        return this.parseContinueStatement();
      case TokenType.CHESTA:
        return this.parseTryStatement();
      case TokenType.FELO:
        return this.parseThrowStatement();
      case TokenType.KAJ:
        // Could be function declaration or expression
        if (this.peekTokenIs(TokenType.IDENTIFIER)) {
          return this.parseFunctionDeclaration();
        }
        return this.parseExpressionStatement();
      case TokenType.SRENI:
        return this.parseClassDeclaration();
      case TokenType.ANO:
        return this.parseImportStatement();
      case TokenType.PATHAO:
        return this.parseExportStatement();
      default:
        return this.parseExpressionStatement();
    }
  }

  private parseLetStatement(): ast.LetStatement | null {
    this.nextToken(); // consume 'dhoro'

    if (!this.curTokenIs(TokenType.IDENTIFIER)) {
      this.errors.push(`line ${this.curToken.line}: expected identifier after 'dhoro'`);
      return null;
    }

    const name = new ast.Identifier(this.curToken.literal);

    if (!this.expectPeek(TokenType.ASSIGN)) {
      return null;
    }

    this.nextToken(); // consume '='

    const value = this.parseExpression(Precedence.LOWEST);
    if (!value) return null;

    if (this.peekTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
    }

    return new ast.LetStatement(name, value);
  }

  private parseReturnStatement(): ast.ReturnStatement {
    const stmt = new ast.ReturnStatement();

    this.nextToken(); // consume 'ferao'

    if (!this.curTokenIs(TokenType.SEMICOLON) && !this.curTokenIs(TokenType.RBRACE)) {
      stmt.value = this.parseExpression(Precedence.LOWEST);
    }

    if (this.peekTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
    }

    return stmt;
  }

  private parseWhileStatement(): ast.WhileStatement | null {
    this.nextToken(); // consume 'jotokkhon'

    if (!this.expectPeek(TokenType.LPAREN)) {
      return null;
    }

    this.nextToken(); // skip '('
    const condition = this.parseExpression(Precedence.LOWEST);
    if (!condition) return null;

    if (!this.expectPeek(TokenType.RPAREN)) {
      return null;
    }

    if (!this.expectPeek(TokenType.LBRACE)) {
      return null;
    }

    const body = this.parseBlockStatement();

    return new ast.WhileStatement(condition, body);
  }

  private parseForStatement(): ast.ForStatement | null {
    this.nextToken(); // consume 'ghuriye'

    if (!this.expectPeek(TokenType.LPAREN)) {
      return null;
    }

    const stmt = new ast.ForStatement(new ast.BlockStatement());

    // Parse init
    this.nextToken();
    if (!this.curTokenIs(TokenType.SEMICOLON)) {
      if (this.curTokenIs(TokenType.DHORO)) {
        stmt.init = this.parseLetStatement();
      } else {
        const expr = this.parseExpression(Precedence.LOWEST);
        if (expr) {
          stmt.init = new ast.ExpressionStatement(expr);
        }
      }
    }
    
    // Consume semicolon after init if present
    if (this.curTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
    } else if (this.peekTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
      this.nextToken();
    }

    // Parse condition
    if (!this.curTokenIs(TokenType.SEMICOLON)) {
      stmt.condition = this.parseExpression(Precedence.LOWEST);
    }

    if (!this.expectPeek(TokenType.SEMICOLON)) {
      return null;
    }

    // Parse update
    this.nextToken();
    if (!this.curTokenIs(TokenType.RPAREN)) {
      stmt.update = this.parseExpression(Precedence.LOWEST);
    }

    if (!this.expectPeek(TokenType.RPAREN)) {
      return null;
    }

    if (!this.expectPeek(TokenType.LBRACE)) {
      return null;
    }

    stmt.body = this.parseBlockStatement();

    return stmt;
  }

  private parseBreakStatement(): ast.BreakStatement {
    const stmt = new ast.BreakStatement();
    if (this.peekTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
    }
    return stmt;
  }

  private parseContinueStatement(): ast.ContinueStatement {
    const stmt = new ast.ContinueStatement();
    if (this.peekTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
    }
    return stmt;
  }

  private parseTryStatement(): ast.TryStatement | null {
    this.nextToken(); // consume 'chesta'

    if (!this.expectPeek(TokenType.LBRACE)) {
      return null;
    }

    const tryBlock = this.parseBlockStatement();
    const stmt = new ast.TryStatement(tryBlock);

    // Parse catch
    if (this.peekTokenIs(TokenType.DHORO_BHUL)) {
      this.nextToken(); // consume 'dhoro_bhul'

      if (this.expectPeek(TokenType.LPAREN)) {
        this.nextToken();
        if (this.curTokenIs(TokenType.IDENTIFIER)) {
          stmt.catchParam = new ast.Identifier(this.curToken.literal);
        }
        if (!this.expectPeek(TokenType.RPAREN)) {
          return null;
        }
      }

      if (!this.expectPeek(TokenType.LBRACE)) {
        return null;
      }

      stmt.catchBlock = this.parseBlockStatement();
    }

    // Parse finally
    if (this.peekTokenIs(TokenType.SHESH)) {
      this.nextToken(); // consume 'shesh'

      if (!this.expectPeek(TokenType.LBRACE)) {
        return null;
      }

      stmt.finallyBlock = this.parseBlockStatement();
    }

    return stmt;
  }

  private parseThrowStatement(): ast.ThrowStatement | null {
    this.nextToken(); // consume 'felo'

    const value = this.parseExpression(Precedence.LOWEST);
    if (!value) return null;

    if (this.peekTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
    }

    return new ast.ThrowStatement(value);
  }

  private parseFunctionDeclaration(): ast.FunctionDeclaration | null {
    this.nextToken(); // consume 'kaj'

    if (!this.curTokenIs(TokenType.IDENTIFIER)) {
      return null;
    }

    const name = new ast.Identifier(this.curToken.literal);

    if (!this.expectPeek(TokenType.LPAREN)) {
      return null;
    }

    const parameters = this.parseFunctionParameters();

    if (!this.expectPeek(TokenType.LBRACE)) {
      return null;
    }

    const body = this.parseBlockStatement();
    const decl = new ast.FunctionDeclaration(name, body);
    decl.parameters = parameters;

    return decl;
  }

  private parseClassDeclaration(): ast.ClassDeclaration | null {
    this.nextToken(); // consume 'sreni'

    if (!this.curTokenIs(TokenType.IDENTIFIER)) {
      this.errors.push(`line ${this.curToken.line}: expected class name`);
      return null;
    }

    const name = new ast.Identifier(this.curToken.literal);
    const decl = new ast.ClassDeclaration(name);

    if (!this.expectPeek(TokenType.LBRACE)) {
      return null;
    }

    this.nextToken(); // skip '{'

    while (!this.curTokenIs(TokenType.RBRACE) && !this.curTokenIs(TokenType.EOF)) {
      if (this.curTokenIs(TokenType.SHURU)) {
        // Constructor
        if (!this.expectPeek(TokenType.LPAREN)) {
          return null;
        }

        const params = this.parseFunctionParameters();

        if (!this.expectPeek(TokenType.LBRACE)) {
          return null;
        }

        const body = this.parseBlockStatement();
        const constructor = new ast.FunctionLiteral(body);
        constructor.parameters = params;
        decl.constructor_ = constructor;
      } else if (this.curTokenIs(TokenType.KAJ)) {
        // Method
        this.nextToken(); // skip 'kaj'

        if (!this.curTokenIs(TokenType.IDENTIFIER)) {
          this.errors.push(`line ${this.curToken.line}: expected method name`);
          return null;
        }

        const methodName = this.curToken.literal;

        if (!this.expectPeek(TokenType.LPAREN)) {
          return null;
        }

        const params = this.parseFunctionParameters();

        if (!this.expectPeek(TokenType.LBRACE)) {
          return null;
        }

        const body = this.parseBlockStatement();
        const method = new ast.FunctionLiteral(body);
        method.parameters = params;
        method.name = methodName;
        decl.methods.set(methodName, method);
      }

      this.nextToken();
    }

    return decl;
  }

  private parseImportStatement(): ast.ImportStatement | null {
    this.nextToken(); // consume 'ano'

    if (!this.curTokenIs(TokenType.STRING)) {
      this.errors.push(`line ${this.curToken.line}: expected string after 'ano'`);
      return null;
    }

    const path = new ast.StringLiteral(this.curToken.literal);
    const stmt = new ast.ImportStatement(path);

    if (this.peekTokenIs(TokenType.HISABE)) {
      this.nextToken(); // consume 'hisabe'
      this.nextToken();

      if (!this.curTokenIs(TokenType.IDENTIFIER)) {
        this.errors.push(`line ${this.curToken.line}: expected identifier after 'hisabe'`);
        return null;
      }

      stmt.alias = new ast.Identifier(this.curToken.literal);
    }

    if (this.peekTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
    }

    return stmt;
  }

  private parseExportStatement(): ast.ExportStatement | null {
    this.nextToken(); // consume 'pathao'

    const innerStmt = this.parseStatement();
    if (!innerStmt) return null;

    return new ast.ExportStatement(innerStmt);
  }

  private parseExpressionStatement(): ast.ExpressionStatement | null {
    const expr = this.parseExpression(Precedence.LOWEST);
    if (!expr) return null;

    const stmt = new ast.ExpressionStatement(expr);

    if (this.peekTokenIs(TokenType.SEMICOLON)) {
      this.nextToken();
    }

    return stmt;
  }

  private parseExpression(precedence: Precedence): ast.Expression | null {
    const prefix = this.prefixParseFns.get(this.curToken.type);
    if (!prefix) {
      this.noPrefixParseFnError(this.curToken.type);
      return null;
    }

    let left = prefix();
    if (!left) return null;

    while (!this.peekTokenIs(TokenType.SEMICOLON) && precedence < this.peekPrecedence()) {
      const infix = this.infixParseFns.get(this.peekToken.type);
      if (!infix) {
        return left;
      }

      this.nextToken();
      left = infix(left);
      if (!left) return null;
    }

    return left;
  }

  private parseIdentifier(): ast.Expression {
    return new ast.Identifier(this.curToken.literal);
  }

  private parseNumberLiteral(): ast.Expression {
    return new ast.NumberLiteral(parseFloat(this.curToken.literal));
  }

  private parseStringLiteral(): ast.Expression {
    return new ast.StringLiteral(this.curToken.literal);
  }

  private parseBooleanLiteral(): ast.Expression {
    return new ast.BooleanLiteral(this.curTokenIs(TokenType.SOTTI));
  }

  private parseNullLiteral(): ast.Expression {
    return new ast.NullLiteral();
  }

  private parsePrefixExpression(): ast.Expression | null {
    const operator = this.curToken.literal;
    this.nextToken();
    const right = this.parseExpression(Precedence.PREFIX);
    if (!right) return null;
    return new ast.PrefixExpression(operator, right);
  }

  private parseInfixExpression(left: ast.Expression): ast.Expression | null {
    const operator = this.curToken.literal;
    const precedence = this.curPrecedence();
    this.nextToken();
    const right = this.parseExpression(precedence);
    if (!right) return null;
    return new ast.InfixExpression(left, operator, right);
  }

  private parseAssignmentExpression(left: ast.Expression): ast.Expression | null {
    const operator = this.curToken.literal;
    this.nextToken();
    const value = this.parseExpression(Precedence.LOWEST);
    if (!value) return null;
    return new ast.AssignmentExpression(left, operator, value);
  }

  private parseGroupedExpression(): ast.Expression | null {
    this.nextToken();
    const expr = this.parseExpression(Precedence.LOWEST);
    if (!this.expectPeek(TokenType.RPAREN)) {
      return null;
    }
    return expr;
  }

  private parseArrayLiteral(): ast.Expression {
    const arr = new ast.ArrayLiteral();
    arr.elements = this.parseExpressionList(TokenType.RBRACKET);
    return arr;
  }

  private parseMapLiteral(): ast.Expression {
    const map = new ast.MapLiteral();

    while (!this.peekTokenIs(TokenType.RBRACE)) {
      this.nextToken();

      // Parse key
      let key: ast.Expression;
      if (this.curTokenIs(TokenType.STRING)) {
        key = new ast.StringLiteral(this.curToken.literal);
      } else if (this.curTokenIs(TokenType.IDENTIFIER)) {
        // Allow unquoted keys like { naam: "value" }
        key = new ast.StringLiteral(this.curToken.literal);
      } else {
        key = this.parseExpression(Precedence.LOWEST)!;
      }

      if (!this.expectPeek(TokenType.COLON)) {
        return map;
      }

      this.nextToken();
      const value = this.parseExpression(Precedence.LOWEST);
      if (!value) return map;

      map.pairs.set(key, value);

      if (!this.peekTokenIs(TokenType.RBRACE) && !this.expectPeek(TokenType.COMMA)) {
        return map;
      }
    }

    this.expectPeek(TokenType.RBRACE);
    return map;
  }

  private parseIfExpression(): ast.Expression | null {
    this.nextToken(); // skip 'jodi'

    if (!this.expectPeek(TokenType.LPAREN)) {
      return null;
    }

    this.nextToken();
    const condition = this.parseExpression(Precedence.LOWEST);
    if (!condition) return null;

    if (!this.expectPeek(TokenType.RPAREN)) {
      return null;
    }

    if (!this.expectPeek(TokenType.LBRACE)) {
      return null;
    }

    const consequence = this.parseBlockStatement();
    const ifExpr = new ast.IfExpression(condition, consequence);

    if (this.peekTokenIs(TokenType.NAHOLE)) {
      this.nextToken(); // consume 'nahole'

      if (this.peekTokenIs(TokenType.JODI)) {
        // else if
        this.nextToken();
        ifExpr.alternative = this.parseIfExpression() as ast.IfExpression | null;
      } else if (this.peekTokenIs(TokenType.LBRACE)) {
        // else
        this.nextToken();
        ifExpr.alternative = this.parseBlockStatement();
      }
    }

    return ifExpr;
  }

  private parseFunctionLiteral(): ast.Expression | null {
    // Skip 'kaj' if we're at it
    if (this.curTokenIs(TokenType.KAJ)) {
      this.nextToken();
    }

    // Check for optional name
    let name: string | null = null;
    if (this.curTokenIs(TokenType.IDENTIFIER)) {
      name = this.curToken.literal;
      this.nextToken();
    }

    if (!this.curTokenIs(TokenType.LPAREN)) {
      if (!this.expectPeek(TokenType.LPAREN)) {
        return null;
      }
    }

    const parameters = this.parseFunctionParameters();

    if (!this.expectPeek(TokenType.LBRACE)) {
      return null;
    }

    const body = this.parseBlockStatement();
    const fn = new ast.FunctionLiteral(body);
    fn.name = name;
    fn.parameters = parameters;

    return fn;
  }

  private parseFunctionParameters(): ast.Identifier[] {
    const params: ast.Identifier[] = [];

    if (this.peekTokenIs(TokenType.RPAREN)) {
      this.nextToken();
      return params;
    }

    this.nextToken();
    params.push(new ast.Identifier(this.curToken.literal));

    while (this.peekTokenIs(TokenType.COMMA)) {
      this.nextToken();
      this.nextToken();
      params.push(new ast.Identifier(this.curToken.literal));
    }

    if (!this.expectPeek(TokenType.RPAREN)) {
      return [];
    }

    return params;
  }

  private parseCallExpression(fn: ast.Expression): ast.Expression {
    const call = new ast.CallExpression(fn);
    call.arguments = this.parseExpressionList(TokenType.RPAREN);
    return call;
  }

  private parseExpressionList(end: TokenType): ast.Expression[] {
    const list: ast.Expression[] = [];

    if (this.peekTokenIs(end)) {
      this.nextToken();
      return list;
    }

    this.nextToken();
    const expr = this.parseExpression(Precedence.LOWEST);
    if (expr) list.push(expr);

    while (this.peekTokenIs(TokenType.COMMA)) {
      this.nextToken();
      this.nextToken();
      const e = this.parseExpression(Precedence.LOWEST);
      if (e) list.push(e);
    }

    if (!this.expectPeek(end)) {
      return [];
    }

    return list;
  }

  private parseIndexExpression(left: ast.Expression): ast.Expression | null {
    this.nextToken();
    const index = this.parseExpression(Precedence.LOWEST);
    if (!index) return null;

    if (!this.expectPeek(TokenType.RBRACKET)) {
      return null;
    }

    return new ast.IndexExpression(left, index);
  }

  private parseMemberExpression(left: ast.Expression): ast.Expression | null {
    this.nextToken();

    if (!this.curTokenIs(TokenType.IDENTIFIER)) {
      this.errors.push(`line ${this.curToken.line}: expected identifier after '.'`);
      return null;
    }

    const property = new ast.Identifier(this.curToken.literal);
    return new ast.MemberExpression(left, property);
  }

  private parseNewExpression(): ast.Expression | null {
    this.nextToken(); // skip 'notun'

    const class_ = this.parseExpression(Precedence.CALL);
    if (!class_) return null;

    const newExpr = new ast.NewExpression(class_);

    // If there's a call expression, extract arguments from it
    if (class_ instanceof ast.CallExpression) {
      newExpr.class_ = class_.function;
      newExpr.arguments = class_.arguments;
    }

    return newExpr;
  }

  private parseThisExpression(): ast.Expression {
    return new ast.ThisExpression();
  }

  private parseBlockStatement(): ast.BlockStatement {
    const block = new ast.BlockStatement();

    this.nextToken();

    while (!this.curTokenIs(TokenType.RBRACE) && !this.curTokenIs(TokenType.EOF)) {
      const stmt = this.parseStatement();
      if (stmt) {
        block.statements.push(stmt);
      }
      this.nextToken();
    }

    return block;
  }
}
