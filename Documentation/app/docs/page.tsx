import Link from "next/link";
import { ArrowRight } from "lucide-react";

export default function Introduction() {
  return (
    <div className="space-y-6">
      <h1>Introduction to BanglaCode</h1>
      <p className="lead text-xl text-muted-foreground">
        A programming language that speaks your language.
      </p>

      <div className="p-4 bg-secondary/50 rounded-lg border border-border">
        <p className="font-medium text-primary mb-2">Notice</p>
        <p className="text-sm text-muted-foreground">
          BanglaCode is currently in <strong>Beta</strong>. Syntax and features may evolve.
        </p>
      </div>

      <h2>What is BanglaCode?</h2>
      <p>
        BanglaCode is an interpreted programming language constructed with Bengali keywords.
        It is designed to help native Bengali speakers, especially students in West Bengal and Bangladesh,
        grasp core programming concepts without the initial barrier of English fluency.
      </p>

      <h2>Design Philosophy</h2>
      <ul className="list-disc pl-6 space-y-2">
        <li><strong>Native Syntax:</strong> Uses Banglish (Bengali in English script) for keywords to make them memorable.</li>
        <li><strong>Strict but Simple:</strong> semantic similarities to C and Java to ensure skills are transferable.</li>
        <li><strong>Modern Features:</strong> Supports Objects, Maps, and First-class functions.</li>
      </ul>

      <h2>Code Example</h2>
      <div className="relative rounded-lg overflow-hidden border border-border bg-[#1e1e1e]">
        <div className="flex items-center px-4 py-2 bg-[#2d2d2d] border-b border-border">
          <div className="flex gap-2">
            <div className="w-3 h-3 rounded-full bg-red-500" />
            <div className="w-3 h-3 rounded-full bg-yellow-500" />
            <div className="w-3 h-3 rounded-full bg-green-500" />
          </div>
          <span className="ml-4 text-xs text-muted-foreground">hello.bang</span>
        </div>
        <div className="p-4 font-mono text-sm overflow-x-auto">
          <pre><code><span className="text-purple-400">kaj</span> <span className="text-yellow-200">greet</span>(name) &#123;
            <span className="text-blue-400">jodi</span> (name == <span className="text-green-400">"World"</span>) &#123;
            <span className="text-yellow-200">dekho</span>(<span className="text-green-400">"Namaskar Prithibi!"</span>);
            &#125; <span className="text-blue-400">nahole</span> &#123;
            <span className="text-yellow-200">dekho</span>(<span className="text-green-400">"Nomoshkar "</span> + name);
            &#125;
            &#125;

            greet(<span className="text-green-400">"Bangla"</span>);</code></pre>
        </div>
      </div>

      <div className="pt-8">
        <Link
          href="/docs/installation"
          className="inline-flex items-center gap-2 px-6 py-3 bg-primary text-white rounded-lg hover:bg-primary/90 transition-colors"
        >
          Next: Installation <ArrowRight className="w-4 h-4" />
        </Link>
      </div>
    </div>
  );
}
