import CodeBlock from "@/components/CodeBlock";
import Link from "next/link";
import { ArrowRight } from "lucide-react";

export default function ControlFlow() {
  return (
    <div className="space-y-6">
      <h1>Control Flow</h1>
      <p className="lead text-xl text-muted-foreground">
        Managing logic with conditionals and loops.
      </p>

      <h2>If / Else (`jodi` / `nahole`)</h2>
      <p>BanglaCode uses <code>jodi</code> for if and <code>nahole</code> for else.</p>

      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-blue-400">dhoro</span> age = 20;

          <span className="text-blue-400">jodi</span> (age &gt;= 18) &#123;
          <span className="text-yellow-200">dekho</span>(<span className="text-green-400">"Adult"</span>);
          &#125; <span className="text-blue-400">nahole</span> &#123;
          <span className="text-yellow-200">dekho</span>(<span className="text-green-400">"Minor"</span>);
          &#125;</code></pre>
      </div>

      <h2>Switch Statement (`bikolpo`)</h2>
      <p>
        Use <code>bikolpo</code> (switch) with <code>khetre</code> (case) clauses and <code>manchito</code> (default):
      </p>

      <CodeBlock
        code={`// Basic switch statement
bikolpo (day) {
    khetre 1 {
        dekho("Monday");
        thamo;  // Break to avoid fall-through
    }
    khetre 2 {
        dekho("Tuesday");
        thamo;
    }
    khetre 3 {
        dekho("Wednesday");
        thamo;
    }
    manchito {
        dekho("Other day");
    }
}

// Switch with expressions
dhoro score = 85;
bikolpo (score / 10) {
    khetre 9 {
        dekho("Grade: A");
        thamo;
    }
    khetre 8 {
        dekho("Grade: B");
        thamo;
    }
    khetre 7 {
        dekho("Grade: C");
        thamo;
    }
    manchito {
        dekho("Grade: F");
    }
}

// Switch with string values
dhoro action = "save";
bikolpo (action) {
    khetre "save" {
        dekho("Saving file...");
        thamo;
    }
    khetre "delete" {
        dekho("Deleting file...");
        thamo;
    }
    khetre "edit" {
        dekho("Editing file...");
        thamo;
    }
    manchito {
        dekho("Unknown action");
    }
}`}
      />

      <h2>Loops</h2>

      <h3>While Loop (`jotokkhon`)</h3>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-blue-400">dhoro</span> i = 0;
          <span className="text-blue-400">jotokkhon</span> (i &lt; 5) &#123;
          <span className="text-yellow-200">dekho</span>(i);
          i = i + 1;
          &#125;</code></pre>
      </div>

      <h3>For Loop (`ghuriye`)</h3>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-blue-400">ghuriye</span> (<span className="text-blue-400">dhoro</span> i = 0; i &lt; 5; i = i + 1) &#123;
          <span className="text-yellow-200">dekho</span>(<span className="text-green-400">"Count: "</span>, i);
          &#125;</code></pre>
      </div>

      <div className="pt-8">
        <Link
          href="/docs/functions"
          className="inline-flex items-center gap-2 px-6 py-3 bg-primary text-white rounded-lg hover:bg-primary/90 transition-colors"
        >
          Next: Functions <ArrowRight className="w-4 h-4" />
        </Link>
      </div>
    </div>
  );
}
