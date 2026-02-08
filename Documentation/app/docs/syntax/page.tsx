import Link from "next/link";
import { ArrowRight } from "lucide-react";

export default function Syntax() {
  return (
    <div className="space-y-6">
      <h1>Syntax Guide</h1>
      <p className="lead text-xl text-muted-foreground">
        A comprehensive guide to BanglaCode keywords and structure.
      </p>

      <h2>Keywords Table</h2>
      <div className="overflow-x-auto border border-border rounded-lg">
        <table className="min-w-full divide-y divide-border">
          <thead className="bg-secondary">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-muted-foreground uppercase tracking-wider">Keyword</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-muted-foreground uppercase tracking-wider">Meaning</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-muted-foreground uppercase tracking-wider">Equivalent</th>
            </tr>
          </thead>
          <tbody className="bg-card divide-y divide-border">
            {[
              { k: "dhoro", m: "hold/var", e: "var/let" },
              { k: "jodi", m: "if", e: "if" },
              { k: "nahole", m: "else", e: "else" },
              { k: "ghuriye", m: "loop", e: "for" },
              { k: "kaj", m: "work", e: "function" },
              { k: "ferao", m: "return", e: "return" },
              { k: "dekho", m: "see/print", e: "print" },
            ].map((row, i) => (
              <tr key={i}>
                <td className="px-6 py-4 whitespace-nowrap font-mono text-primary">{row.k}</td>
                <td className="px-6 py-4 whitespace-nowrap text-muted-foreground">{row.m}</td>
                <td className="px-6 py-4 whitespace-nowrap font-mono text-muted-foreground">{row.e}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      <h2>Variables</h2>
      <p>Variables are declared using <code>dhoro</code>. BanglaCode is dynamically typed.</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        dhoro name = "Rahim";<br />
        dhoro age = 24;<br />
        dhoro isStudent = sotti; <span className="text-muted-foreground">// true</span>
      </div>

      <div className="pt-8">
        <Link
          href="/docs/control-flow"
          className="inline-flex items-center gap-2 px-6 py-3 bg-primary text-white rounded-lg hover:bg-primary/90 transition-colors"
        >
          Next: Control Flow <ArrowRight className="w-4 h-4" />
        </Link>
      </div>
    </div>
  );
}
