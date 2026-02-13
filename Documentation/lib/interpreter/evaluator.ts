// BanglaCode Evaluator - Tree-walking interpreter

import * as ast from "./ast";
import {
  Obj,
  NumberObj,
  StringObj,
  BooleanObj,
  ArrayObj,
  MapObj,
  FunctionObj,
  BuiltinObj,
  ClassObj,
  InstanceObj,
  ReturnValue,
  ErrorObj,
  NULL,
  TRUE,
  FALSE,
  BREAK,
  CONTINUE,
  nativeBoolToBooleanObj,
} from "./object";
import { Environment } from "./environment";
import { builtins } from "./builtins";

export class Evaluator {
  private maxIterations = 100000; // Prevent infinite loops

  eval(node: ast.Node, env: Environment): Obj {
    // Program
    if (node instanceof ast.Program) {
      return this.evalProgram(node, env);
    }

    // Statements
    if (node instanceof ast.ExpressionStatement) {
      return this.eval(node.expression, env);
    }

    if (node instanceof ast.BlockStatement) {
      return this.evalBlockStatement(node, env);
    }

    if (node instanceof ast.LetStatement) {
      const val = this.eval(node.value, env);
      if (this.isError(val)) return val;
      env.set(node.name.value, val);
      return val;
    }

    if (node instanceof ast.ReturnStatement) {
      if (node.value) {
        const val = this.eval(node.value, env);
        if (this.isError(val)) return val;
        return new ReturnValue(val);
      }
      return new ReturnValue(NULL);
    }

    if (node instanceof ast.WhileStatement) {
      return this.evalWhileStatement(node, env);
    }

    if (node instanceof ast.ForStatement) {
      return this.evalForStatement(node, env);
    }

    if (node instanceof ast.BreakStatement) {
      return BREAK;
    }

    if (node instanceof ast.ContinueStatement) {
      return CONTINUE;
    }

    if (node instanceof ast.TryStatement) {
      return this.evalTryStatement(node, env);
    }

    if (node instanceof ast.ThrowStatement) {
      const val = this.eval(node.value, env);
      if (this.isError(val)) return val;
      return new ErrorObj(val.inspect());
    }

    if (node instanceof ast.FunctionDeclaration) {
      const fn = new FunctionObj(node.parameters, node.body, env, node.name.value);
      env.set(node.name.value, fn);
      return fn;
    }

    if (node instanceof ast.ClassDeclaration) {
      return this.evalClassDeclaration(node, env);
    }

    if (node instanceof ast.ImportStatement) {
      // Imports not supported in playground
      return NULL;
    }

    if (node instanceof ast.ExportStatement) {
      return this.eval(node.statement, env);
    }

    // Expressions
    if (node instanceof ast.NumberLiteral) {
      return new NumberObj(node.value);
    }

    if (node instanceof ast.StringLiteral) {
      return new StringObj(node.value);
    }

    if (node instanceof ast.BooleanLiteral) {
      return nativeBoolToBooleanObj(node.value);
    }

    if (node instanceof ast.NullLiteral) {
      return NULL;
    }

    if (node instanceof ast.ArrayLiteral) {
      const elements = this.evalExpressions(node.elements, env);
      if (elements.length === 1 && this.isError(elements[0])) {
        return elements[0];
      }
      return new ArrayObj(elements);
    }

    if (node instanceof ast.MapLiteral) {
      return this.evalMapLiteral(node, env);
    }

    if (node instanceof ast.Identifier) {
      return this.evalIdentifier(node, env);
    }

    if (node instanceof ast.PrefixExpression) {
      const right = this.eval(node.right, env);
      if (this.isError(right)) return right;
      return this.evalPrefixExpression(node.operator, right);
    }

    if (node instanceof ast.InfixExpression) {
      const left = this.eval(node.left, env);
      if (this.isError(left)) return left;
      const right = this.eval(node.right, env);
      if (this.isError(right)) return right;
      return this.evalInfixExpression(node.operator, left, right);
    }

    if (node instanceof ast.AssignmentExpression) {
      return this.evalAssignmentExpression(node, env);
    }

    if (node instanceof ast.IfExpression) {
      return this.evalIfExpression(node, env);
    }

    if (node instanceof ast.FunctionLiteral) {
      return new FunctionObj(node.parameters, node.body, env, node.name);
    }

    if (node instanceof ast.CallExpression) {
      const fn = this.eval(node.function, env);
      if (this.isError(fn)) return fn;

      const args = this.evalExpressions(node.arguments, env);
      if (args.length === 1 && this.isError(args[0])) {
        return args[0];
      }

      return this.applyFunction(fn, args);
    }

    if (node instanceof ast.IndexExpression) {
      const left = this.eval(node.left, env);
      if (this.isError(left)) return left;
      const index = this.eval(node.index, env);
      if (this.isError(index)) return index;
      return this.evalIndexExpression(left, index);
    }

    if (node instanceof ast.MemberExpression) {
      const obj = this.eval(node.object, env);
      if (this.isError(obj)) return obj;
      return this.evalMemberExpression(obj, node.property.value, env);
    }

    if (node instanceof ast.NewExpression) {
      return this.evalNewExpression(node, env);
    }

    if (node instanceof ast.ThisExpression) {
      const thisVal = env.get("ei");
      if (thisVal) return thisVal;
      return new ErrorObj("'ei' used outside of class context");
    }

    return NULL;
  }

  private evalProgram(program: ast.Program, env: Environment): Obj {
    let result: Obj = NULL;

    for (const stmt of program.statements) {
      result = this.eval(stmt, env);

      if (result instanceof ReturnValue) {
        return result.value;
      }

      if (result instanceof ErrorObj) {
        return result;
      }
    }

    return result;
  }

  private evalBlockStatement(block: ast.BlockStatement, env: Environment): Obj {
    let result: Obj = NULL;

    for (const stmt of block.statements) {
      result = this.eval(stmt, env);

      if (result.type() === "RETURN_VALUE" ||
          result.type() === "ERROR" ||
          result.type() === "BREAK" ||
          result.type() === "CONTINUE") {
        return result;
      }
    }

    return result;
  }

  private evalWhileStatement(node: ast.WhileStatement, env: Environment): Obj {
    let iterations = 0;

    while (true) {
      iterations++;
      if (iterations > this.maxIterations) {
        return new ErrorObj("Maximum iterations exceeded. Possible infinite loop.");
      }

      const condition = this.eval(node.condition, env);
      if (this.isError(condition)) return condition;

      if (!this.isTruthy(condition)) {
        break;
      }

      const result = this.eval(node.body, env);

      if (result instanceof ErrorObj) return result;
      if (result instanceof ReturnValue) return result;
      if (result === BREAK) break;
      // CONTINUE just continues the loop
    }

    return NULL;
  }

  private evalForStatement(node: ast.ForStatement, env: Environment): Obj {
    const loopEnv = Environment.newEnclosed(env);
    let iterations = 0;

    // Initialize
    if (node.init) {
      const init = this.eval(node.init, loopEnv);
      if (this.isError(init)) return init;
    }

    // Loop
    while (true) {
      iterations++;
      if (iterations > this.maxIterations) {
        return new ErrorObj("Maximum iterations exceeded. Possible infinite loop.");
      }

      // Check condition
      if (node.condition) {
        const condition = this.eval(node.condition, loopEnv);
        if (this.isError(condition)) return condition;

        if (!this.isTruthy(condition)) {
          break;
        }
      }

      // Execute body
      const result = this.eval(node.body, loopEnv);

      if (result instanceof ErrorObj) return result;
      if (result instanceof ReturnValue) return result;
      if (result === BREAK) break;
      // CONTINUE continues to update

      // Update
      if (node.update) {
        const update = this.eval(node.update, loopEnv);
        if (this.isError(update)) return update;
      }
    }

    return NULL;
  }

  private evalTryStatement(node: ast.TryStatement, env: Environment): Obj {
    let result = this.eval(node.tryBlock, env);

    // If error was thrown
    if (result instanceof ErrorObj && node.catchBlock) {
      const catchEnv = Environment.newEnclosed(env);
      if (node.catchParam) {
        catchEnv.set(node.catchParam.value, new StringObj(result.message));
      }
      result = this.eval(node.catchBlock, catchEnv);
    }

    // Finally always runs
    if (node.finallyBlock) {
      this.eval(node.finallyBlock, env);
    }

    return result instanceof ErrorObj ? NULL : result;
  }

  private evalClassDeclaration(node: ast.ClassDeclaration, env: Environment): Obj {
    let constructor_: FunctionObj | null = null;
    const methods = new Map<string, FunctionObj>();

    if (node.constructor_) {
      constructor_ = new FunctionObj(
        node.constructor_.parameters,
        node.constructor_.body,
        env,
        "shuru"
      );
    }

    for (const [name, method] of node.methods) {
      const fn = new FunctionObj(method.parameters, method.body, env, name);
      methods.set(name, fn);
    }

    const classObj = new ClassObj(node.name.value, constructor_, methods);
    env.set(node.name.value, classObj);
    return classObj;
  }

  private evalIdentifier(node: ast.Identifier, env: Environment): Obj {
    const val = env.get(node.value);
    if (val) return val;

    const builtin = builtins.get(node.value);
    if (builtin) return builtin;

    return new ErrorObj(`identifier not found: ${node.value}`);
  }

  private evalPrefixExpression(operator: string, right: Obj): Obj {
    switch (operator) {
      case "!":
      case "na":
        return this.evalBangOperator(right);
      case "-":
        return this.evalMinusPrefixOperator(right);
      default:
        return new ErrorObj(`unknown operator: ${operator}${right.type()}`);
    }
  }

  private evalBangOperator(right: Obj): Obj {
    if (right === TRUE) return FALSE;
    if (right === FALSE) return TRUE;
    if (right === NULL) return TRUE;
    return FALSE;
  }

  private evalMinusPrefixOperator(right: Obj): Obj {
    if (!(right instanceof NumberObj)) {
      return new ErrorObj(`unknown operator: -${right.type()}`);
    }
    return new NumberObj(-right.value);
  }

  private evalInfixExpression(operator: string, left: Obj, right: Obj): Obj {
    if (left instanceof NumberObj && right instanceof NumberObj) {
      return this.evalNumberInfixExpression(operator, left, right);
    }

    if (left instanceof StringObj && right instanceof StringObj) {
      return this.evalStringInfixExpression(operator, left, right);
    }

    if (left instanceof StringObj || right instanceof StringObj) {
      if (operator === "+") {
        return new StringObj(left.inspect() + right.inspect());
      }
    }

    switch (operator) {
      case "==":
        return nativeBoolToBooleanObj(left === right);
      case "!=":
        return nativeBoolToBooleanObj(left !== right);
      case "ebong":
        return nativeBoolToBooleanObj(this.isTruthy(left) && this.isTruthy(right));
      case "ba":
        return nativeBoolToBooleanObj(this.isTruthy(left) || this.isTruthy(right));
      default:
        return new ErrorObj(`unknown operator: ${left.type()} ${operator} ${right.type()}`);
    }
  }

  private evalNumberInfixExpression(operator: string, left: NumberObj, right: NumberObj): Obj {
    const leftVal = left.value;
    const rightVal = right.value;

    switch (operator) {
      case "+":
        return new NumberObj(leftVal + rightVal);
      case "-":
        return new NumberObj(leftVal - rightVal);
      case "*":
        return new NumberObj(leftVal * rightVal);
      case "/":
        if (rightVal === 0) {
          return new ErrorObj("division by zero");
        }
        return new NumberObj(leftVal / rightVal);
      case "%":
        return new NumberObj(leftVal % rightVal);
      case "**":
        return new NumberObj(Math.pow(leftVal, rightVal));
      case "<":
        return nativeBoolToBooleanObj(leftVal < rightVal);
      case ">":
        return nativeBoolToBooleanObj(leftVal > rightVal);
      case "<=":
        return nativeBoolToBooleanObj(leftVal <= rightVal);
      case ">=":
        return nativeBoolToBooleanObj(leftVal >= rightVal);
      case "==":
        return nativeBoolToBooleanObj(leftVal === rightVal);
      case "!=":
        return nativeBoolToBooleanObj(leftVal !== rightVal);
      default:
        return new ErrorObj(`unknown operator: ${left.type()} ${operator} ${right.type()}`);
    }
  }

  private evalStringInfixExpression(operator: string, left: StringObj, right: StringObj): Obj {
    switch (operator) {
      case "+":
        return new StringObj(left.value + right.value);
      case "==":
        return nativeBoolToBooleanObj(left.value === right.value);
      case "!=":
        return nativeBoolToBooleanObj(left.value !== right.value);
      default:
        return new ErrorObj(`unknown operator: ${left.type()} ${operator} ${right.type()}`);
    }
  }

  private evalAssignmentExpression(node: ast.AssignmentExpression, env: Environment): Obj {
    const value = this.eval(node.value, env);
    if (this.isError(value)) return value;

    // Handle compound assignment
    if (node.operator !== "=") {
      const current = this.evalAssignmentTarget(node.left, env);
      if (this.isError(current)) return current;

      const op = node.operator.slice(0, -1); // Remove '=' from +=, -=, etc.
      const newValue = this.evalInfixExpression(op, current, value);
      if (this.isError(newValue)) return newValue;

      return this.assignValue(node.left, newValue, env);
    }

    return this.assignValue(node.left, value, env);
  }

  private evalAssignmentTarget(target: ast.Expression, env: Environment): Obj {
    if (target instanceof ast.Identifier) {
      return this.evalIdentifier(target, env);
    }

    if (target instanceof ast.IndexExpression) {
      const left = this.eval(target.left, env);
      if (this.isError(left)) return left;
      const index = this.eval(target.index, env);
      if (this.isError(index)) return index;
      return this.evalIndexExpression(left, index);
    }

    if (target instanceof ast.MemberExpression) {
      const obj = this.eval(target.object, env);
      if (this.isError(obj)) return obj;
      return this.evalMemberExpression(obj, target.property.value, env);
    }

    return new ErrorObj("invalid assignment target");
  }

  private assignValue(target: ast.Expression, value: Obj, env: Environment): Obj {
    if (target instanceof ast.Identifier) {
      if (!env.update(target.value, value)) {
        env.set(target.value, value);
      }
      return value;
    }

    if (target instanceof ast.IndexExpression) {
      const left = this.eval(target.left, env);
      if (this.isError(left)) return left;
      const index = this.eval(target.index, env);
      if (this.isError(index)) return index;

      if (left instanceof ArrayObj && index instanceof NumberObj) {
        const idx = Math.floor(index.value);
        if (idx >= 0 && idx < left.elements.length) {
          left.elements[idx] = value;
          return value;
        }
        return new ErrorObj(`array index out of bounds: ${idx}`);
      }

      if (left instanceof MapObj && index instanceof StringObj) {
        left.pairs.set(index.value, value);
        return value;
      }

      return new ErrorObj("index assignment not supported for this type");
    }

    if (target instanceof ast.MemberExpression) {
      const obj = this.eval(target.object, env);
      if (this.isError(obj)) return obj;

      if (obj instanceof MapObj) {
        obj.pairs.set(target.property.value, value);
        return value;
      }

      if (obj instanceof InstanceObj) {
        obj.properties.set(target.property.value, value);
        return value;
      }

      return new ErrorObj("member assignment not supported for this type");
    }

    return new ErrorObj("invalid assignment target");
  }

  private evalIfExpression(node: ast.IfExpression, env: Environment): Obj {
    const condition = this.eval(node.condition, env);
    if (this.isError(condition)) return condition;

    if (this.isTruthy(condition)) {
      return this.eval(node.consequence, env);
    } else if (node.alternative) {
      return this.eval(node.alternative, env);
    }

    return NULL;
  }

  private evalMapLiteral(node: ast.MapLiteral, env: Environment): Obj {
    const pairs = new Map<string, Obj>();

    for (const [keyExpr, valueExpr] of node.pairs) {
      let key: string;
      if (keyExpr instanceof ast.StringLiteral) {
        key = keyExpr.value;
      } else {
        const keyObj = this.eval(keyExpr, env);
        if (this.isError(keyObj)) return keyObj;
        key = keyObj.inspect();
      }

      const value = this.eval(valueExpr, env);
      if (this.isError(value)) return value;

      pairs.set(key, value);
    }

    return new MapObj(pairs);
  }

  private evalExpressions(exprs: ast.Expression[], env: Environment): Obj[] {
    const result: Obj[] = [];

    for (const expr of exprs) {
      const evaluated = this.eval(expr, env);
      if (this.isError(evaluated)) {
        return [evaluated];
      }
      result.push(evaluated);
    }

    return result;
  }

  private applyFunction(fn: Obj, args: Obj[]): Obj {
    if (fn instanceof FunctionObj) {
      const extendedEnv = this.extendFunctionEnv(fn, args);
      const evaluated = this.eval(fn.body, extendedEnv);
      return this.unwrapReturnValue(evaluated);
    }

    if (fn instanceof BuiltinObj) {
      return fn.fn(...args);
    }

    return new ErrorObj(`not a function: ${fn.type()}`);
  }

  private extendFunctionEnv(fn: FunctionObj, args: Obj[]): Environment {
    const env = Environment.newEnclosed(fn.env);

    for (let i = 0; i < fn.parameters.length; i++) {
      env.set(fn.parameters[i].value, args[i] || NULL);
    }

    return env;
  }

  private unwrapReturnValue(obj: Obj): Obj {
    if (obj instanceof ReturnValue) {
      return obj.value;
    }
    return obj;
  }

  private evalIndexExpression(left: Obj, index: Obj): Obj {
    if (left instanceof ArrayObj && index instanceof NumberObj) {
      return this.evalArrayIndexExpression(left, index);
    }

    if (left instanceof MapObj && index instanceof StringObj) {
      return this.evalMapIndexExpression(left, index);
    }

    if (left instanceof StringObj && index instanceof NumberObj) {
      const idx = Math.floor(index.value);
      if (idx >= 0 && idx < left.value.length) {
        return new StringObj(left.value[idx]);
      }
      return NULL;
    }

    return new ErrorObj(`index operator not supported: ${left.type()}`);
  }

  private evalArrayIndexExpression(arr: ArrayObj, index: NumberObj): Obj {
    const idx = Math.floor(index.value);
    const max = arr.elements.length - 1;

    if (idx < 0 || idx > max) {
      return NULL;
    }

    return arr.elements[idx];
  }

  private evalMapIndexExpression(map: MapObj, key: StringObj): Obj {
    const val = map.pairs.get(key.value);
    if (val) return val;
    return NULL;
  }

  private evalMemberExpression(obj: Obj, property: string, env: Environment): Obj {
    if (obj instanceof MapObj) {
      const val = obj.pairs.get(property);
      if (val) return val;
      return NULL;
    }

    if (obj instanceof InstanceObj) {
      // Check properties first
      const propVal = obj.properties.get(property);
      if (propVal) return propVal;

      // Then check methods
      const method = obj.class_.methods.get(property);
      if (method) {
        // Create a bound method
        const boundEnv = Environment.newEnclosed(method.env);
        boundEnv.set("ei", obj);
        return new FunctionObj(method.parameters, method.body, boundEnv, method.name);
      }

      return NULL;
    }

    return new ErrorObj(`member access not supported on ${obj.type()}`);
  }

  private evalNewExpression(node: ast.NewExpression, env: Environment): Obj {
    const classObj = this.eval(node.class_, env);
    if (this.isError(classObj)) return classObj;

    if (!(classObj instanceof ClassObj)) {
      return new ErrorObj(`not a class: ${classObj.type()}`);
    }

    const instance = new InstanceObj(classObj);

    // Call constructor if it exists
    if (classObj.constructor_) {
      const constructorEnv = Environment.newEnclosed(classObj.constructor_.env);
      constructorEnv.set("ei", instance);

      // Set constructor parameters
      const args = this.evalExpressions(node.arguments, env);
      if (args.length === 1 && this.isError(args[0])) {
        return args[0];
      }

      for (let i = 0; i < classObj.constructor_.parameters.length; i++) {
        constructorEnv.set(classObj.constructor_.parameters[i].value, args[i] || NULL);
      }

      const result = this.eval(classObj.constructor_.body, constructorEnv);
      if (this.isError(result)) return result;
    }

    return instance;
  }

  private isTruthy(obj: Obj): boolean {
    if (obj === NULL) return false;
    if (obj === FALSE) return false;
    if (obj instanceof NumberObj && obj.value === 0) return false;
    if (obj instanceof StringObj && obj.value === "") return false;
    return true;
  }

  private isError(obj: Obj): obj is ErrorObj {
    return obj instanceof ErrorObj;
  }
}
