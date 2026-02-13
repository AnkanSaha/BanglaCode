// BanglaCode Environment - Variable scoping

import type { Obj } from "./object";

export class Environment {
  private store: Map<string, Obj> = new Map();
  private outer: Environment | null = null;

  constructor(outer: Environment | null = null) {
    this.outer = outer;
  }

  get(name: string): Obj | undefined {
    let obj = this.store.get(name);
    if (obj === undefined && this.outer !== null) {
      obj = this.outer.get(name);
    }
    return obj;
  }

  set(name: string, value: Obj): Obj {
    this.store.set(name, value);
    return value;
  }

  // Update existing variable in the correct scope
  update(name: string, value: Obj): boolean {
    if (this.store.has(name)) {
      this.store.set(name, value);
      return true;
    }
    if (this.outer !== null) {
      return this.outer.update(name, value);
    }
    return false;
  }

  // Create a new enclosed environment
  static newEnclosed(outer: Environment): Environment {
    return new Environment(outer);
  }
}
