// BanglaCode Built-in Functions

import {
  Obj,
  NumberObj,
  StringObj,
  BooleanObj,
  ArrayObj,
  MapObj,
  BuiltinObj,
  ErrorObj,
  NULL,
  TRUE,
  FALSE,
  nativeBoolToBooleanObj
} from "./object";

// Helper to create error
function newError(message: string): ErrorObj {
  return new ErrorObj(message);
}

// Output collector for playground
let outputCollector: string[] = [];

export function getOutput(): string[] {
  return outputCollector;
}

export function clearOutput(): void {
  outputCollector = [];
}

export function addOutput(line: string): void {
  outputCollector.push(line);
}

// Built-in functions map
export const builtins: Map<string, BuiltinObj> = new Map([
  // Output - dekho (দেখো - see/show)
  ["dekho", new BuiltinObj("dekho", (...args: Obj[]) => {
    const output = args.map(arg => arg.inspect()).join(" ");
    addOutput(output);
    return NULL;
  })],

  // Length - dorghyo (দৈর্ঘ্য - length)
  ["dorghyo", new BuiltinObj("dorghyo", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    const arg = args[0];
    if (arg instanceof StringObj) {
      return new NumberObj(arg.value.length);
    }
    if (arg instanceof ArrayObj) {
      return new NumberObj(arg.elements.length);
    }
    return newError(`argument to 'dorghyo' not supported, got ${arg.type()}`);
  })],

  // Push - dhokao (ঢোকাও - insert)
  ["dhokao", new BuiltinObj("dhokao", (...args: Obj[]) => {
    if (args.length !== 2) {
      return newError(`wrong number of arguments. got=${args.length}, want=2`);
    }
    if (!(args[0] instanceof ArrayObj)) {
      return newError(`argument to 'dhokao' must be ARRAY, got ${args[0].type()}`);
    }
    args[0].elements.push(args[1]);
    return args[0];
  })],

  // Pop - berKoro (বের করো - take out)
  ["berKoro", new BuiltinObj("berKoro", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof ArrayObj)) {
      return newError(`argument to 'berKoro' must be ARRAY, got ${args[0].type()}`);
    }
    const arr = args[0];
    if (arr.elements.length > 0) {
      return arr.elements.pop()!;
    }
    return NULL;
  })],

  // Keys - chabi (চাবি - keys)
  ["chabi", new BuiltinObj("chabi", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof MapObj)) {
      return newError(`argument to 'chabi' must be MAP, got ${args[0].type()}`);
    }
    const keys = Array.from(args[0].pairs.keys()).map(k => new StringObj(k));
    return new ArrayObj(keys);
  })],

  // Type - dhoron (ধরন - type)
  ["dhoron", new BuiltinObj("dhoron", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    const arg = args[0];
    if (arg instanceof NumberObj) {
      return new StringObj(Number.isInteger(arg.value) ? "int" : "float");
    }
    if (arg instanceof StringObj) return new StringObj("string");
    if (arg instanceof BooleanObj) return new StringObj("boolean");
    if (arg instanceof ArrayObj) return new StringObj("array");
    if (arg instanceof MapObj) return new StringObj("map");
    return new StringObj(arg.type().toLowerCase());
  })],

  // To string - lipi (লিপি - text/script)
  ["lipi", new BuiltinObj("lipi", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    return new StringObj(args[0].inspect());
  })],

  // To number - sonkha (সংখ্যা - number)
  ["sonkha", new BuiltinObj("sonkha", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    const arg = args[0];
    if (arg instanceof NumberObj) return arg;
    if (arg instanceof StringObj) {
      const num = parseFloat(arg.value);
      if (isNaN(num)) {
        return newError(`cannot convert string to number: ${arg.value}`);
      }
      return new NumberObj(num);
    }
    if (arg instanceof BooleanObj) {
      return new NumberObj(arg.value ? 1 : 0);
    }
    return newError(`cannot convert ${arg.type()} to number`);
  })],

  // Square root - borgomul (বর্গমূল)
  ["borgomul", new BuiltinObj("borgomul", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof NumberObj)) {
      return newError(`argument to 'borgomul' must be NUMBER, got ${args[0].type()}`);
    }
    return new NumberObj(Math.sqrt(args[0].value));
  })],

  // Power - ghat (ঘাত)
  ["ghat", new BuiltinObj("ghat", (...args: Obj[]) => {
    if (args.length !== 2) {
      return newError(`wrong number of arguments. got=${args.length}, want=2`);
    }
    if (!(args[0] instanceof NumberObj) || !(args[1] instanceof NumberObj)) {
      return newError("arguments to 'ghat' must be NUMBERs");
    }
    return new NumberObj(Math.pow(args[0].value, args[1].value));
  })],

  // Floor - niche (নিচে - down)
  ["niche", new BuiltinObj("niche", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof NumberObj)) {
      return newError(`argument to 'niche' must be NUMBER, got ${args[0].type()}`);
    }
    return new NumberObj(Math.floor(args[0].value));
  })],

  // Ceil - upore (উপরে - up)
  ["upore", new BuiltinObj("upore", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof NumberObj)) {
      return newError(`argument to 'upore' must be NUMBER, got ${args[0].type()}`);
    }
    return new NumberObj(Math.ceil(args[0].value));
  })],

  // Round - kache (কাছে - near)
  ["kache", new BuiltinObj("kache", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof NumberObj)) {
      return newError(`argument to 'kache' must be NUMBER, got ${args[0].type()}`);
    }
    return new NumberObj(Math.round(args[0].value));
  })],

  // Absolute - niratek (নিরপেক্ষ - absolute)
  ["niratek", new BuiltinObj("niratek", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof NumberObj)) {
      return newError(`argument to 'niratek' must be NUMBER, got ${args[0].type()}`);
    }
    return new NumberObj(Math.abs(args[0].value));
  })],

  // Min - choto (ছোট - small)
  ["choto", new BuiltinObj("choto", (...args: Obj[]) => {
    if (args.length < 2) {
      return newError(`wrong number of arguments. got=${args.length}, want at least 2`);
    }
    let min = Infinity;
    for (const arg of args) {
      if (!(arg instanceof NumberObj)) {
        return newError("all arguments to 'choto' must be NUMBERs");
      }
      if (arg.value < min) min = arg.value;
    }
    return new NumberObj(min);
  })],

  // Max - boro (বড় - big)
  ["boro", new BuiltinObj("boro", (...args: Obj[]) => {
    if (args.length < 2) {
      return newError(`wrong number of arguments. got=${args.length}, want at least 2`);
    }
    let max = -Infinity;
    for (const arg of args) {
      if (!(arg instanceof NumberObj)) {
        return newError("all arguments to 'boro' must be NUMBERs");
      }
      if (arg.value > max) max = arg.value;
    }
    return new NumberObj(max);
  })],

  // Random - lotto (লটো - lottery/random)
  ["lotto", new BuiltinObj("lotto", () => {
    return new NumberObj(Math.random());
  })],

  // Upper - boroHater (বড় হাতের - uppercase)
  ["boroHater", new BuiltinObj("boroHater", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof StringObj)) {
      return newError(`argument to 'boroHater' must be STRING, got ${args[0].type()}`);
    }
    return new StringObj(args[0].value.toUpperCase());
  })],

  // Lower - chotoHater (ছোট হাতের - lowercase)
  ["chotoHater", new BuiltinObj("chotoHater", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof StringObj)) {
      return newError(`argument to 'chotoHater' must be STRING, got ${args[0].type()}`);
    }
    return new StringObj(args[0].value.toLowerCase());
  })],

  // Split - bhag (ভাগ - divide)
  ["bhag", new BuiltinObj("bhag", (...args: Obj[]) => {
    if (args.length !== 2) {
      return newError(`wrong number of arguments. got=${args.length}, want=2`);
    }
    if (!(args[0] instanceof StringObj) || !(args[1] instanceof StringObj)) {
      return newError("arguments to 'bhag' must be STRINGs");
    }
    const parts = args[0].value.split(args[1].value);
    return new ArrayObj(parts.map(p => new StringObj(p)));
  })],

  // Join - joro (জোড়ো - join)
  ["joro", new BuiltinObj("joro", (...args: Obj[]) => {
    if (args.length !== 2) {
      return newError(`wrong number of arguments. got=${args.length}, want=2`);
    }
    if (!(args[0] instanceof ArrayObj)) {
      return newError(`first argument to 'joro' must be ARRAY, got ${args[0].type()}`);
    }
    if (!(args[1] instanceof StringObj)) {
      return newError(`second argument to 'joro' must be STRING, got ${args[1].type()}`);
    }
    const parts = args[0].elements.map(e => e.inspect());
    return new StringObj(parts.join(args[1].value));
  })],

  // Trim - chhanto (ছাঁটো - trim)
  ["chhanto", new BuiltinObj("chhanto", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof StringObj)) {
      return newError(`argument to 'chhanto' must be STRING, got ${args[0].type()}`);
    }
    return new StringObj(args[0].value.trim());
  })],

  // Index of - khojo (খোঁজো - search)
  ["khojo", new BuiltinObj("khojo", (...args: Obj[]) => {
    if (args.length !== 2) {
      return newError(`wrong number of arguments. got=${args.length}, want=2`);
    }
    if (!(args[0] instanceof StringObj) || !(args[1] instanceof StringObj)) {
      return newError("arguments to 'khojo' must be STRINGs");
    }
    return new NumberObj(args[0].value.indexOf(args[1].value));
  })],

  // Substring - angsho (অংশ - portion)
  ["angsho", new BuiltinObj("angsho", (...args: Obj[]) => {
    if (args.length < 2 || args.length > 3) {
      return newError(`wrong number of arguments. got=${args.length}, want=2 or 3`);
    }
    if (!(args[0] instanceof StringObj)) {
      return newError(`first argument to 'angsho' must be STRING, got ${args[0].type()}`);
    }
    if (!(args[1] instanceof NumberObj)) {
      return newError(`second argument to 'angsho' must be NUMBER, got ${args[1].type()}`);
    }
    const str = args[0].value;
    const start = Math.max(0, args[1].value);
    let end = str.length;
    if (args.length === 3) {
      if (!(args[2] instanceof NumberObj)) {
        return newError(`third argument to 'angsho' must be NUMBER, got ${args[2].type()}`);
      }
      end = Math.min(str.length, args[2].value);
    }
    if (start > end) return new StringObj("");
    return new StringObj(str.slice(start, end));
  })],

  // Replace - bodlo (বদলো - change)
  ["bodlo", new BuiltinObj("bodlo", (...args: Obj[]) => {
    if (args.length !== 3) {
      return newError(`wrong number of arguments. got=${args.length}, want=3`);
    }
    if (!(args[0] instanceof StringObj) || !(args[1] instanceof StringObj) || !(args[2] instanceof StringObj)) {
      return newError("all arguments to 'bodlo' must be STRINGs");
    }
    return new StringObj(args[0].value.split(args[1].value).join(args[2].value));
  })],

  // Slice - kato (কাটো - cut)
  ["kato", new BuiltinObj("kato", (...args: Obj[]) => {
    if (args.length < 2 || args.length > 3) {
      return newError(`wrong number of arguments. got=${args.length}, want=2 or 3`);
    }
    if (!(args[0] instanceof ArrayObj)) {
      return newError(`first argument to 'kato' must be ARRAY, got ${args[0].type()}`);
    }
    if (!(args[1] instanceof NumberObj)) {
      return newError(`second argument to 'kato' must be NUMBER, got ${args[1].type()}`);
    }
    const arr = args[0].elements;
    const start = Math.max(0, args[1].value);
    let end = arr.length;
    if (args.length === 3) {
      if (!(args[2] instanceof NumberObj)) {
        return newError(`third argument to 'kato' must be NUMBER, got ${args[2].type()}`);
      }
      end = Math.min(arr.length, args[2].value);
    }
    if (start > end) return new ArrayObj([]);
    return new ArrayObj(arr.slice(start, end));
  })],

  // Reverse - ulto (উল্টো - reverse)
  ["ulto", new BuiltinObj("ulto", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof ArrayObj)) {
      return newError(`argument to 'ulto' must be ARRAY, got ${args[0].type()}`);
    }
    return new ArrayObj([...args[0].elements].reverse());
  })],

  // Includes - ache (আছে - exists)
  ["ache", new BuiltinObj("ache", (...args: Obj[]) => {
    if (args.length !== 2) {
      return newError(`wrong number of arguments. got=${args.length}, want=2`);
    }
    if (!(args[0] instanceof ArrayObj)) {
      return newError(`first argument to 'ache' must be ARRAY, got ${args[0].type()}`);
    }
    const target = args[1];
    for (const el of args[0].elements) {
      if (objectsEqual(el, target)) {
        return TRUE;
      }
    }
    return FALSE;
  })],

  // Sort - saja (সাজা - arrange)
  ["saja", new BuiltinObj("saja", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof ArrayObj)) {
      return newError(`argument to 'saja' must be ARRAY, got ${args[0].type()}`);
    }
    const sorted = [...args[0].elements].sort((a, b) => {
      if (a instanceof NumberObj && b instanceof NumberObj) {
        return a.value - b.value;
      }
      return a.inspect().localeCompare(b.inspect());
    });
    return new ArrayObj(sorted);
  })],

  // Time - somoy (সময় - time)
  ["somoy", new BuiltinObj("somoy", () => {
    return new NumberObj(Date.now());
  })],

  // JSON Parse - json_poro (JSON পড়ো - read JSON)
  ["json_poro", new BuiltinObj("json_poro", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    if (!(args[0] instanceof StringObj)) {
      return newError(`argument to 'json_poro' must be STRING, got ${args[0].type()}`);
    }
    try {
      const data = JSON.parse(args[0].value);
      return jsonToObject(data);
    } catch (e) {
      return newError(`JSON parse error: ${e}`);
    }
  })],

  // JSON Stringify - json_banao (JSON বানাও - make JSON)
  ["json_banao", new BuiltinObj("json_banao", (...args: Obj[]) => {
    if (args.length !== 1) {
      return newError(`wrong number of arguments. got=${args.length}, want=1`);
    }
    try {
      const data = objectToJSON(args[0]);
      return new StringObj(JSON.stringify(data));
    } catch (e) {
      return newError(`JSON stringify error: ${e}`);
    }
  })],

  // Input simulation - nao (নাও - take) - Not interactive in playground
  ["nao", new BuiltinObj("nao", (...args: Obj[]) => {
    if (args.length > 0) {
      addOutput(args[0].inspect());
    }
    return new StringObj("[input not available in playground]");
  })],

  // File operations - not available in playground
  ["poro", new BuiltinObj("poro", () => {
    return newError("File operations not available in playground");
  })],

  ["lekho", new BuiltinObj("lekho", () => {
    return newError("File operations not available in playground");
  })],

  // Sleep - ghum (ঘুম - sleep) - simulated
  ["ghum", new BuiltinObj("ghum", () => {
    // In browser, we can't actually sleep synchronously
    addOutput("[sleep not available in playground]");
    return NULL;
  })],

  // Exit - bondho (বন্ধ - stop/close)
  ["bondho", new BuiltinObj("bondho", (...args: Obj[]) => {
    const code = args.length > 0 && args[0] instanceof NumberObj ? args[0].value : 0;
    addOutput(`[program exited with code ${code}]`);
    return NULL;
  })],

  // HTTP functions - not available in playground
  ["server_chalu", new BuiltinObj("server_chalu", () => {
    return newError("HTTP server not available in playground");
  })],

  ["anun", new BuiltinObj("anun", () => {
    return newError("HTTP requests not available in playground");
  })],

  ["uttor", new BuiltinObj("uttor", () => {
    return newError("HTTP functions not available in playground");
  })],

  ["json_uttor", new BuiltinObj("json_uttor", () => {
    return newError("HTTP functions not available in playground");
  })],
]);

// Helper to compare objects for equality
function objectsEqual(a: Obj, b: Obj): boolean {
  if (a.type() !== b.type()) return false;
  if (a instanceof NumberObj && b instanceof NumberObj) {
    return a.value === b.value;
  }
  if (a instanceof StringObj && b instanceof StringObj) {
    return a.value === b.value;
  }
  if (a instanceof BooleanObj && b instanceof BooleanObj) {
    return a.value === b.value;
  }
  return a === b;
}

// Convert JSON to BanglaCode object
function jsonToObject(data: unknown): Obj {
  if (data === null) return NULL;
  if (typeof data === "boolean") return nativeBoolToBooleanObj(data);
  if (typeof data === "number") return new NumberObj(data);
  if (typeof data === "string") return new StringObj(data);
  if (Array.isArray(data)) {
    return new ArrayObj(data.map(jsonToObject));
  }
  if (typeof data === "object") {
    const pairs = new Map<string, Obj>();
    for (const [key, value] of Object.entries(data as Record<string, unknown>)) {
      pairs.set(key, jsonToObject(value));
    }
    return new MapObj(pairs);
  }
  return NULL;
}

// Convert BanglaCode object to JSON
function objectToJSON(obj: Obj): unknown {
  if (obj instanceof NumberObj) return obj.value;
  if (obj instanceof StringObj) return obj.value;
  if (obj instanceof BooleanObj) return obj.value;
  if (obj === NULL) return null;
  if (obj instanceof ArrayObj) {
    return obj.elements.map(objectToJSON);
  }
  if (obj instanceof MapObj) {
    const result: Record<string, unknown> = {};
    for (const [key, value] of obj.pairs) {
      result[key] = objectToJSON(value);
    }
    return result;
  }
  return obj.inspect();
}
