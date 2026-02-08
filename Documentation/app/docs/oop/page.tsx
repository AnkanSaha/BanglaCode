import Link from "next/link";

export default function OOP() {
  return (
    <div className="space-y-6">
      <h1>Object Oriented Programming</h1>
      <p className="lead text-xl text-muted-foreground">
        Classes and Objects in BanglaCode.
      </p>

      <h2>Classes (`sreni`)</h2>
      <p>Define a class using <code>sreni</code> and stricture with <code>shuru</code> (constructor).</p>

      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">sreni</span> Manush &#123;
          <span className="text-yellow-200">shuru</span>(naam, boyosh) &#123;
          <span className="text-blue-400">ei</span>.naam = naam;
          <span className="text-blue-400">ei</span>.boyosh = boyosh;
          &#125;

          <span className="text-purple-400">kaj</span> <span className="text-yellow-200">porichoy</span>() &#123;
          <span className="text-yellow-200">dekho</span>(<span className="text-green-400">"Amar naam "</span>, <span className="text-blue-400">ei</span>.naam);
          &#125;
          &#125;</code></pre>
      </div>

      <h2>Creating Objects</h2>
      <p>Use <code>notun</code> to create an instance.</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-blue-400">dhoro</span> person = <span className="text-purple-400">notun</span> Manush(<span className="text-green-400">"Ankan"</span>, 25);
          person.porichoy();</code></pre>
      </div>
    </div>
  );
}
