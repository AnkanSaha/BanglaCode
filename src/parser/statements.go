package parser

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
)

// parseStatement determines which statement to parse based on the current token
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case lexer.DHORO:
		return p.parseVariableDeclaration(false, false)
	case lexer.STHIR:
		return p.parseVariableDeclaration(true, false)
	case lexer.BISHWO:
		return p.parseVariableDeclaration(false, true)
	case lexer.JODI:
		return p.parseIfStatement()
	case lexer.JOTOKKHON:
		return p.parseWhileStatement()
	case lexer.GHURIYE:
		return p.parseForStatement()
	case lexer.FERAO:
		return p.parseReturnStatement()
	case lexer.SRENI:
		return p.parseClassDeclaration()
	case lexer.THAMO:
		return p.parseBreakStatement()
	case lexer.CHHARO:
		return p.parseContinueStatement()
	case lexer.ANO:
		return p.parseImportStatement()
	case lexer.PATHAO:
		return p.parseExportStatement()
	case lexer.CHESTA:
		return p.parseTryCatchStatement()
	case lexer.FELO:
		return p.parseThrowStatement()
	case lexer.BIKOLPO:
		return p.parseSwitchStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseVariableDeclaration parses "dhoro x = value", "sthir x = value", or "bishwo x = value"
func (p *Parser) parseVariableDeclaration(isConstant bool, isGlobal bool) *ast.VariableDeclaration {
	stmt := &ast.VariableDeclaration{
		Token:      p.curToken,
		IsConstant: isConstant,
		IsGlobal:   isGlobal,
	}

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseIfStatement parses "jodi (condition) { } nahole { }"
func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(lexer.NAHOLE) {
		p.nextToken()

		// Check for "nahole jodi" (else if)
		if p.peekTokenIs(lexer.JODI) {
			p.nextToken()
			// Parse else if as a nested if statement
			ifStmt := p.parseIfStatement()
			stmt.Alternative = &ast.BlockStatement{
				Token:      p.curToken,
				Statements: []ast.Statement{ifStmt},
			}
		} else {
			if !p.expectPeek(lexer.LBRACE) {
				return nil
			}
			stmt.Alternative = p.parseBlockStatement()
		}
	}

	return stmt
}

// parseWhileStatement parses "jotokkhon (condition) { }"
func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseForStatement parses "ghuriye (init; condition; update) { }"
func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	p.nextToken()

	// Parse init statement
	if !p.curTokenIs(lexer.SEMICOLON) {
		stmt.Init = p.parseStatement()
	}

	if !p.curTokenIs(lexer.SEMICOLON) {
		if !p.expectPeek(lexer.SEMICOLON) {
			return nil
		}
	}

	p.nextToken()

	// Parse condition
	if !p.curTokenIs(lexer.SEMICOLON) {
		stmt.Condition = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(lexer.SEMICOLON) {
		return nil
	}

	p.nextToken()

	// Parse update
	if !p.curTokenIs(lexer.RPAREN) {
		stmt.Update = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseReturnStatement parses "ferao value"
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	if !p.curTokenIs(lexer.SEMICOLON) && !p.curTokenIs(lexer.RBRACE) {
		stmt.ReturnValue = p.parseExpression(LOWEST)
	}

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseClassDeclaration parses "sreni ClassName { methods }"
func (p *Parser) parseClassDeclaration() *ast.ClassDeclaration {
	stmt := &ast.ClassDeclaration{Token: p.curToken}

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	p.nextToken()

	stmt.Methods = []*ast.FunctionLiteral{}

	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		if p.curTokenIs(lexer.KAJ) {
			method := p.parseFunctionLiteral()
			if fn, ok := method.(*ast.FunctionLiteral); ok {
				stmt.Methods = append(stmt.Methods, fn)
			}
		} else if p.curTokenIs(lexer.SHURU) {
			// Parse constructor without kaj keyword: shuru(params) { ... }
			constructor := p.parseConstructor()
			if constructor != nil {
				stmt.Methods = append(stmt.Methods, constructor)
			}
		}
		p.nextToken()
	}

	return stmt
}

// parseConstructor parses "shuru(params) { }" constructor syntax
func (p *Parser) parseConstructor() *ast.FunctionLiteral {
	lit := &ast.FunctionLiteral{Token: p.curToken}
	lit.Name = &ast.Identifier{Token: p.curToken, Value: "shuru"}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

// parseBreakStatement parses "thamo"
func (p *Parser) parseBreakStatement() *ast.BreakStatement {
	stmt := &ast.BreakStatement{Token: p.curToken}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseContinueStatement parses "chharo"
func (p *Parser) parseContinueStatement() *ast.ContinueStatement {
	stmt := &ast.ContinueStatement{Token: p.curToken}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseBlockStatement parses "{ statements }"
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

// parseExpressionStatement parses standalone expressions
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseImportStatement parses "ano 'module' hisabe alias"
func (p *Parser) parseImportStatement() *ast.ImportStatement {
	stmt := &ast.ImportStatement{Token: p.curToken}

	if !p.expectPeek(lexer.STRING) {
		return nil
	}

	stmt.Path = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}

	// Check for optional "hisabe alias" syntax (হিসাবে - as)
	if p.peekTokenIs(lexer.HISABE) {
		p.nextToken() // move to "hisabe"
		if !p.expectPeek(lexer.IDENT) {
			return nil
		}
		stmt.Alias = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExportStatement parses "pathao statement"
func (p *Parser) parseExportStatement() *ast.ExportStatement {
	stmt := &ast.ExportStatement{Token: p.curToken}

	p.nextToken()

	// Parse the statement being exported
	stmt.Statement = p.parseStatement()

	return stmt
}

// parseTryCatchStatement parses "chesta { } dhoro_bhul { } shesh { }"
func (p *Parser) parseTryCatchStatement() *ast.TryCatchStatement {
	stmt := &ast.TryCatchStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.TryBlock = p.parseBlockStatement()

	// Parse catch block
	if p.peekTokenIs(lexer.DHORO_BHUL) {
		p.nextToken()

		if p.peekTokenIs(lexer.LPAREN) {
			p.nextToken()
			if !p.expectPeek(lexer.IDENT) {
				return nil
			}
			stmt.CatchParam = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			if !p.expectPeek(lexer.RPAREN) {
				return nil
			}
		}

		if !p.expectPeek(lexer.LBRACE) {
			return nil
		}

		stmt.CatchBlock = p.parseBlockStatement()
	}

	// Parse finally block
	if p.peekTokenIs(lexer.SHESH) {
		p.nextToken()

		if !p.expectPeek(lexer.LBRACE) {
			return nil
		}

		stmt.FinallyBlock = p.parseBlockStatement()
	}

	return stmt
}

// parseThrowStatement parses "felo value"
func (p *Parser) parseThrowStatement() *ast.ThrowStatement {
	stmt := &ast.ThrowStatement{Token: p.curToken}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseSwitchStatement parses "bikolpo (expression) { khetre value: ... manchito: ... }"
func (p *Parser) parseSwitchStatement() *ast.SwitchStatement {
	stmt := &ast.SwitchStatement{Token: p.curToken}

	// Expect opening parenthesis
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	// Parse the expression to match against
	p.nextToken()
	stmt.Expr = p.parseExpression(LOWEST)

	// Expect closing parenthesis
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	// Expect opening brace
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	// Parse cases and default
	stmt.Cases = []*ast.CaseClause{}

	p.nextToken()
	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		if p.curTokenIs(lexer.KHETRE) {
			caseClause := p.parseCaseClause()
			if caseClause != nil {
				stmt.Cases = append(stmt.Cases, caseClause)
			}
			p.nextToken()
		} else if p.curTokenIs(lexer.MANCHITO) {
			// Parse default case (no colon in simplified syntax)
			// Expect opening brace
			if !p.expectPeek(lexer.LBRACE) {
				return nil
			}

			stmt.Default = p.parseBlockStatement()
			p.nextToken()
		} else {
			return nil
		}
	}

	return stmt
}

// parseCaseClause parses a single "khetre value { ... }" clause
func (p *Parser) parseCaseClause() *ast.CaseClause {
	clause := &ast.CaseClause{Token: p.curToken}

	// Move to the value expression
	p.nextToken()
	clause.Value = p.parseExpression(LOWEST)

	// Expect opening brace (no colon in simplified syntax)
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	clause.Body = p.parseBlockStatement()

	return clause
}
