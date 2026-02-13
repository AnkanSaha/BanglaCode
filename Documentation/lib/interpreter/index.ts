// BanglaCode Interpreter - Main entry point for browser execution

import { Lexer } from "./lexer";
import { Parser } from "./parser";
import { Evaluator } from "./evaluator";
import { Environment } from "./environment";
import { clearOutput, getOutput, addOutput } from "./builtins";
import type { ErrorObj } from "./object";

export interface ExecutionResult {
  output: string[];
  errors: string[];
  success: boolean;
}

/**
 * Execute BanglaCode source code and return the output
 */
export function execute(source: string): ExecutionResult {
  // Clear previous output
  clearOutput();

  const result: ExecutionResult = {
    output: [],
    errors: [],
    success: false,
  };

  try {
    // Lexer
    const lexer = new Lexer(source);

    // Parser
    const parser = new Parser(lexer);
    const program = parser.parseProgram();

    // Check for parse errors
    if (parser.errors.length > 0) {
      result.errors = parser.errors;
      result.output = getOutput();
      return result;
    }

    // Evaluator
    const env = new Environment();
    const evaluator = new Evaluator();
    const evaluated = evaluator.eval(program, env);

    // Check for runtime errors
    if (evaluated && evaluated.type() === "ERROR") {
      result.errors = [(evaluated as ErrorObj).message];
      result.output = getOutput();
      return result;
    }

    result.output = getOutput();
    result.success = true;
  } catch (e) {
    result.errors = [String(e)];
    result.output = getOutput();
  }

  return result;
}

// Re-export for use in other modules
export { Lexer } from "./lexer";
export { Parser } from "./parser";
export { Evaluator } from "./evaluator";
export { Environment } from "./environment";
export { clearOutput, getOutput, addOutput } from "./builtins";
