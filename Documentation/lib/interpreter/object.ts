// BanglaCode Object System - Runtime values

export type ObjectType =
  | "NUMBER"
  | "STRING"
  | "BOOLEAN"
  | "NULL"
  | "ARRAY"
  | "MAP"
  | "FUNCTION"
  | "BUILTIN"
  | "CLASS"
  | "INSTANCE"
  | "RETURN_VALUE"
  | "ERROR"
  | "BREAK"
  | "CONTINUE";

export interface Obj {
  type(): ObjectType;
  inspect(): string;
}

export class NumberObj implements Obj {
  value: number;

  constructor(value: number) {
    this.value = value;
  }

  type(): ObjectType {
    return "NUMBER";
  }

  inspect(): string {
    return String(this.value);
  }
}

export class StringObj implements Obj {
  value: string;

  constructor(value: string) {
    this.value = value;
  }

  type(): ObjectType {
    return "STRING";
  }

  inspect(): string {
    return this.value;
  }
}

export class BooleanObj implements Obj {
  value: boolean;

  constructor(value: boolean) {
    this.value = value;
  }

  type(): ObjectType {
    return "BOOLEAN";
  }

  inspect(): string {
    return this.value ? "sotti" : "mittha";
  }
}

export class NullObj implements Obj {
  type(): ObjectType {
    return "NULL";
  }

  inspect(): string {
    return "khali";
  }
}

export class ArrayObj implements Obj {
  elements: Obj[];

  constructor(elements: Obj[] = []) {
    this.elements = elements;
  }

  type(): ObjectType {
    return "ARRAY";
  }

  inspect(): string {
    return "[" + this.elements.map(e => e.inspect()).join(", ") + "]";
  }
}

export class MapObj implements Obj {
  pairs: Map<string, Obj>;

  constructor(pairs: Map<string, Obj> = new Map()) {
    this.pairs = pairs;
  }

  type(): ObjectType {
    return "MAP";
  }

  inspect(): string {
    const parts: string[] = [];
    for (const [key, value] of this.pairs) {
      parts.push(`${key}: ${value.inspect()}`);
    }
    return "{" + parts.join(", ") + "}";
  }
}

import type { BlockStatement, Identifier } from "./ast";
import type { Environment } from "./environment";

export class FunctionObj implements Obj {
  name: string | null;
  parameters: Identifier[];
  body: BlockStatement;
  env: Environment;

  constructor(
    parameters: Identifier[],
    body: BlockStatement,
    env: Environment,
    name: string | null = null
  ) {
    this.name = name;
    this.parameters = parameters;
    this.body = body;
    this.env = env;
  }

  type(): ObjectType {
    return "FUNCTION";
  }

  inspect(): string {
    const params = this.parameters.map(p => p.value).join(", ");
    return `kaj ${this.name || ""}(${params}) { ... }`;
  }
}

export type BuiltinFn = (...args: Obj[]) => Obj;

export class BuiltinObj implements Obj {
  fn: BuiltinFn;
  name: string;

  constructor(name: string, fn: BuiltinFn) {
    this.name = name;
    this.fn = fn;
  }

  type(): ObjectType {
    return "BUILTIN";
  }

  inspect(): string {
    return `[builtin: ${this.name}]`;
  }
}

export class ClassObj implements Obj {
  name: string;
  constructor_: FunctionObj | null;
  methods: Map<string, FunctionObj>;

  constructor(
    name: string,
    constructor_: FunctionObj | null,
    methods: Map<string, FunctionObj>
  ) {
    this.name = name;
    this.constructor_ = constructor_;
    this.methods = methods;
  }

  type(): ObjectType {
    return "CLASS";
  }

  inspect(): string {
    return `[class: ${this.name}]`;
  }
}

export class InstanceObj implements Obj {
  class_: ClassObj;
  properties: Map<string, Obj>;

  constructor(class_: ClassObj) {
    this.class_ = class_;
    this.properties = new Map();
  }

  type(): ObjectType {
    return "INSTANCE";
  }

  inspect(): string {
    const props: string[] = [];
    for (const [key, value] of this.properties) {
      props.push(`${key}: ${value.inspect()}`);
    }
    return `${this.class_.name} { ${props.join(", ")} }`;
  }
}

export class ReturnValue implements Obj {
  value: Obj;

  constructor(value: Obj) {
    this.value = value;
  }

  type(): ObjectType {
    return "RETURN_VALUE";
  }

  inspect(): string {
    return this.value.inspect();
  }
}

export class ErrorObj implements Obj {
  message: string;

  constructor(message: string) {
    this.message = message;
  }

  type(): ObjectType {
    return "ERROR";
  }

  inspect(): string {
    return `ERROR: ${this.message}`;
  }
}

export class BreakSignal implements Obj {
  type(): ObjectType {
    return "BREAK";
  }

  inspect(): string {
    return "break";
  }
}

export class ContinueSignal implements Obj {
  type(): ObjectType {
    return "CONTINUE";
  }

  inspect(): string {
    return "continue";
  }
}

// Singleton objects
export const TRUE = new BooleanObj(true);
export const FALSE = new BooleanObj(false);
export const NULL = new NullObj();
export const BREAK = new BreakSignal();
export const CONTINUE = new ContinueSignal();

// Helper to convert native boolean
export function nativeBoolToBooleanObj(value: boolean): BooleanObj {
  return value ? TRUE : FALSE;
}
