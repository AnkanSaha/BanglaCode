package parser

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
	"fmt"
	"strconv"
)

// parseExpression parses expressions with precedence climbing
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(lexer.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// noPrefixParseFnError reports an error for missing prefix parse function
func (p *Parser) noPrefixParseFnError(t lexer.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found at line %d, column %d",
		t, p.curToken.Line, p.curToken.Column)
	p.errors = append(p.errors, msg)
}

// ==================== Prefix Expressions ====================

// parseIdentifier parses an identifier
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseNumberLiteral parses a number literal
func (p *Parser) parseNumberLiteral() ast.Expression {
	lit := &ast.NumberLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as number", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

// parseStringLiteral parses a string literal
func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// parseTemplateLiteral parses a template literal with ${expression} interpolation
func (p *Parser) parseTemplateLiteral() ast.Expression {
	return &ast.TemplateLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// parseBooleanLiteral parses sotti/mittha
func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(lexer.SOTTI)}
}

// parseNullLiteral parses khali
func (p *Parser) parseNullLiteral() ast.Expression {
	return &ast.NullLiteral{Token: p.curToken}
}

// parseUnaryExpression parses -x, !x, na x
func (p *Parser) parseUnaryExpression() ast.Expression {
	expression := &ast.UnaryExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// parseGroupedExpression parses (expression)
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	return exp
}

// parseArrayLiteral parses [elements]
func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(lexer.RBRACKET)
	return array
}

// parseMapLiteral parses {key: value}
func (p *Parser) parseMapLiteral() ast.Expression {
	mapLit := &ast.MapLiteral{Token: p.curToken}
	mapLit.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(lexer.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.expectPeek(lexer.COLON) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)

		mapLit.Pairs[key] = value

		if !p.peekTokenIs(lexer.RBRACE) && !p.expectPeek(lexer.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(lexer.RBRACE) {
		return nil
	}

	return mapLit
}

// parseFunctionLiteral parses kaj name(params) { body }
func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	// Check if function has a name
	if p.peekTokenIs(lexer.IDENT) {
		p.nextToken()
		lit.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	params, restParam := p.parseFunctionParametersWithRest()
	lit.Parameters = params
	lit.RestParameter = restParam

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

// parseAsyncFunctionLiteral parses proyash kaj name(params) { body }
func (p *Parser) parseAsyncFunctionLiteral() ast.Expression {
	token := p.curToken // PROYASH token

	// Must be followed by 'kaj'
	if !p.expectPeek(lexer.KAJ) {
		return nil
	}

	lit := &ast.AsyncFunctionLiteral{Token: token}

	// Check if function has a name
	if p.peekTokenIs(lexer.IDENT) {
		p.nextToken()
		lit.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	params, restParam := p.parseFunctionParametersWithRest()
	lit.Parameters = params
	lit.RestParameter = restParam

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

// parseAwaitExpression parses opekha expression
func (p *Parser) parseAwaitExpression() ast.Expression {
	exp := &ast.AwaitExpression{Token: p.curToken}

	p.nextToken()
	exp.Expression = p.parseExpression(PREFIX)

	return exp
}

// parseFunctionParameters parses function parameter list (legacy, kept for compatibility)
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	params, _ := p.parseFunctionParametersWithRest()
	return params
}

// parseFunctionParametersWithRest parses function parameters including rest parameter
func (p *Parser) parseFunctionParametersWithRest() ([]*ast.Identifier, *ast.Identifier) {
	identifiers := []*ast.Identifier{}
	var restParam *ast.Identifier

	if p.peekTokenIs(lexer.RPAREN) {
		p.nextToken()
		return identifiers, nil
	}

	p.nextToken()

	// Check for rest parameter
	if p.curTokenIs(lexer.DOTDOTDOT) {
		if !p.expectPeek(lexer.IDENT) {
			return nil, nil
		}
		restParam = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		if !p.expectPeek(lexer.RPAREN) {
			return nil, nil
		}
		return identifiers, restParam
	}

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()

		// Check for rest parameter
		if p.curTokenIs(lexer.DOTDOTDOT) {
			if !p.expectPeek(lexer.IDENT) {
				return nil, nil
			}
			restParam = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			if !p.expectPeek(lexer.RPAREN) {
				return nil, nil
			}
			return identifiers, restParam
		}

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(lexer.RPAREN) {
		return nil, nil
	}

	return identifiers, restParam
}

// parseNewExpression parses notun ClassName(args)
func (p *Parser) parseNewExpression() ast.Expression {
	exp := &ast.NewExpression{Token: p.curToken}

	p.nextToken()
	exp.Class = p.parseExpression(CALL)

	if p.peekTokenIs(lexer.LPAREN) {
		p.nextToken()
		exp.Arguments = p.parseExpressionList(lexer.RPAREN)
	}

	return exp
}

// ==================== Infix Expressions ====================

// parseBinaryExpression parses binary operators (+, -, *, /, etc.)
func (p *Parser) parseBinaryExpression(left ast.Expression) ast.Expression {
	expression := &ast.BinaryExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

// parseAssignmentExpression parses =, +=, -=, etc.
func (p *Parser) parseAssignmentExpression(left ast.Expression) ast.Expression {
	expression := &ast.AssignmentExpression{
		Token:    p.curToken,
		Name:     left,
		Operator: p.curToken.Literal,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Value = p.parseExpression(precedence)

	return expression
}

// parseCallExpression parses function(args)
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(lexer.RPAREN)
	return exp
}

// parseMemberExpression parses obj.prop or obj[index]
func (p *Parser) parseMemberExpression(left ast.Expression) ast.Expression {
	exp := &ast.MemberExpression{
		Token:  p.curToken,
		Object: left,
	}

	if p.curTokenIs(lexer.LBRACKET) {
		exp.Computed = true
		p.nextToken()
		exp.Property = p.parseExpression(LOWEST)
		if !p.expectPeek(lexer.RBRACKET) {
			return nil
		}
	} else {
		exp.Computed = false
		if !p.expectPeek(lexer.IDENT) {
			return nil
		}
		exp.Property = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	return exp
}

// ==================== Helpers ====================

// parseExpressionList parses a comma-separated list of expressions
func (p *Parser) parseExpressionList(end lexer.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

// parseSpreadElement parses ...expression
func (p *Parser) parseSpreadElement() ast.Expression {
	spread := &ast.SpreadElement{Token: p.curToken}

	p.nextToken()
	spread.Argument = p.parseExpression(LOWEST)

	return spread
}
