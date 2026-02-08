import Link from "next/link";
import { ArrowRight } from "lucide-react";

export default function Functions() {
  return (
    <div className="space-y-6">
      <h1>Functions</h1>
      <p className="lead text-xl text-muted-foreground">
        Reusable blocks of code using `kaj`.
      </p>

      <h2>Defining a Function</h2>
      <p>Use the <code>kaj</code> keyword to define a function.</p>

      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">kaj</span> <span className="text-yellow-200">add</span>(a, b) &#123;
          <span className="text-purple-400">ferao</span> a + b;
          &#125;

          <span className="text-blue-400">dhoro</span> result = add(5, 3);
          <span className="text-yellow-200">dekho</span>(result); <span className="text-muted-foreground">// 8</span></code></pre>
      </div>

      <h2>Recursion</h2>
      <p>Functions can call themselves.</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">kaj</span> <span className="text-yellow-200">factorial</span>(n) &#123;
          <span className="text-blue-400">jodi</span> (n &lt;= 1) &#123;
          <span className="text-purple-400">ferao</span> 1;
          &#125;
          <span className="text-purple-400">ferao</span> n * factorial(n - 1);
          &#125;</code></pre>
      </div>

      <div className="pt-8">
        <Link
          href="/docs/oop"
          className="inline-flex items-center gap-2 px-6 py-3 bg-primary text-white rounded-lg hover:bg-primary/90 transition-colors"
        >
          Next: OOP & Classes <ArrowRight className="w-4 h-4" />
        </Link>
      </div>
    </div>
  );
}
