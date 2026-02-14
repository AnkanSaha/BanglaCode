import Link from "next/link";
import { ArrowRight, Code, Terminal, Sparkles } from "lucide-react";
import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Introduction() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Getting Started
        </span>
      </div>

      <h1>Introduction to BanglaCode</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode is a Bengali-syntax programming language that makes coding accessible
        to Bengali speakers. Write code using familiar Bengali words in English script (Banglish).
      </p>

      <div className="grid sm:grid-cols-3 gap-4 my-8">
        <div className="p-4 rounded-lg border border-border bg-secondary/20">
          <Code className="w-8 h-8 text-primary mb-2" />
          <h3 className="font-semibold mb-1">Bengali Syntax</h3>
          <p className="text-sm text-muted-foreground">
            Use Bengali keywords like <code>dhoro</code>, <code>jodi</code>, <code>kaj</code>
          </p>
        </div>
        <div className="p-4 rounded-lg border border-border bg-secondary/20">
          <Terminal className="w-8 h-8 text-primary mb-2" />
          <h3 className="font-semibold mb-1">Full-Featured</h3>
          <p className="text-sm text-muted-foreground">
            OOP, modules, error handling, HTTP servers
          </p>
        </div>
        <div className="p-4 rounded-lg border border-border bg-secondary/20">
          <Sparkles className="w-8 h-8 text-primary mb-2" />
          <h3 className="font-semibold mb-1">Easy to Learn</h3>
          <p className="text-sm text-muted-foreground">
            Intuitive syntax for Bengali speakers
          </p>
        </div>
      </div>

      <h2>What is BanglaCode?</h2>

      <p>
        BanglaCode is a <strong>tree-walking interpreter</strong> written in Go that allows you to write
        programs using Bengali keywords. The language uses <strong>Banglish</strong> (Bengali words written
        in English/Latin script) to make it easy to type on any keyboard while remaining familiar to
        Bengali speakers.
      </p>

      <h3>Design Philosophy</h3>

      <ul>
        <li><strong>Accessibility</strong> - Lower the barrier for Bengali speakers to learn programming</li>
        <li><strong>Completeness</strong> - Full programming language features, not just a toy</li>
        <li><strong>Familiarity</strong> - Keywords match natural Bengali expressions</li>
        <li><strong>Simplicity</strong> - Clean, minimal syntax without unnecessary complexity</li>
      </ul>

      <h2>Quick Example</h2>

      <p>Here&apos;s a simple BanglaCode program that demonstrates the core syntax:</p>

      <CodeBlock
        filename="hello.bang"
        code={`// Variable declaration
dhoro naam = "Ankan";
dhoro boyosh = 25;

// Print output
dekho("Namaskar! Amar naam", naam);

// Conditional
jodi (boyosh >= 18) {
    dekho(naam, "is an adult");
} nahole {
    dekho(naam, "is a minor");
}

// Loop
ghuriye (dhoro i = 1; i <= 5; i = i + 1) {
    dekho("Count:", i);
}

// Function
kaj square(n) {
    ferao n * n;
}

dekho("5 squared =", square(5));`}
      />

      <h2>Language Features</h2>

      <p>BanglaCode is a complete programming language with:</p>

      <div className="overflow-x-auto">
        <table>
          <thead>
            <tr>
              <th>Feature</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><strong>29 Bengali Keywords</strong></td>
              <td>All control flow, declarations, and async/await in Bengali</td>
            </tr>
            <tr>
              <td><strong>80+ Built-in Functions</strong></td>
              <td>String, array, math, file I/O, HTTP, JSON, and complete OS-level system access</td>
            </tr>
            <tr>
              <td><strong>Object-Oriented</strong></td>
              <td>Classes, constructors, methods, instances</td>
            </tr>
            <tr>
              <td><strong>Module System</strong></td>
              <td>Import/export with namespace support</td>
            </tr>
            <tr>
              <td><strong>Error Handling</strong></td>
              <td>Try-catch-finally with throw</td>
            </tr>
            <tr>
              <td><strong>HTTP Server</strong></td>
              <td>Build web applications</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Architecture</h2>

      <p>
        The interpreter follows a classic pipeline architecture:
      </p>

      <div className="p-4 rounded-lg bg-secondary/30 border border-border font-mono text-sm my-4">
        Source Code → <span className="text-primary">Lexer</span> → Tokens →{" "}
        <span className="text-primary">Parser</span> → AST →{" "}
        <span className="text-primary">Evaluator</span> → Result
      </div>

      <ol>
        <li><strong>Lexer</strong> - Tokenizes source code, recognizes Bengali keywords</li>
        <li><strong>Parser</strong> - Builds Abstract Syntax Tree using Pratt parsing</li>
        <li><strong>Evaluator</strong> - Tree-walking interpreter executes the AST</li>
      </ol>

      <h2>File Extension</h2>

      <p>
        BanglaCode source files use the <code>.bang</code> extension. For example:
        <code>hello.bang</code>, <code>calculator.bang</code>, <code>server.bang</code>.
      </p>

      <h2>Next Steps</h2>

      <p>Ready to start coding in Bengali? Follow these guides:</p>

      <div className="grid sm:grid-cols-2 gap-4 my-6">
        <Link
          href="/docs/installation"
          className="group p-4 rounded-lg border border-border hover:border-primary/50 hover:bg-secondary/30 transition-colors"
        >
          <h4 className="font-medium group-hover:text-primary transition-colors flex items-center gap-2">
            Installation <ArrowRight className="w-4 h-4" />
          </h4>
          <p className="text-sm text-muted-foreground mt-1">
            Build and install the BanglaCode interpreter
          </p>
        </Link>
        <Link
          href="/docs/quick-start"
          className="group p-4 rounded-lg border border-border hover:border-primary/50 hover:bg-secondary/30 transition-colors"
        >
          <h4 className="font-medium group-hover:text-primary transition-colors flex items-center gap-2">
            Quick Start <ArrowRight className="w-4 h-4" />
          </h4>
          <p className="text-sm text-muted-foreground mt-1">
            Write your first BanglaCode program
          </p>
        </Link>
      </div>

      <DocNavigation currentPath="/docs" />
    </div>
  );
}
