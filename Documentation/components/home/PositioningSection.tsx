"use client";

import { GraduationCap, Code2, Database, Network, Package } from "lucide-react";

export default function PositioningSection() {
  return (
    <section className="py-20 bg-gradient-to-b from-background to-accent/5 relative overflow-hidden">
      <div className="container mx-auto px-4 relative z-10">
        {/* Main Statement */}
        <div className="max-w-5xl mx-auto text-center mb-16 animate-fade-in-up">
          <div className="inline-flex items-center gap-2 px-4 py-2 text-sm font-semibold text-amber-500 bg-amber-500/10 rounded-full border border-amber-500/20 mb-6">
            <GraduationCap className="w-4 h-4" />
            Educational Language
          </div>

          <h2 className="text-3xl md:text-4xl lg:text-5xl font-bold mb-6">
            <span className="text-muted-foreground">Not Competing With</span>{" "}
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 to-cyan-500">
              JavaScript
            </span>
            {" or "}
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-yellow-400 to-orange-500">
              Python
            </span>
          </h2>

          <p className="text-xl md:text-2xl text-muted-foreground leading-relaxed mb-8">
            BanglaCode is an <span className="text-primary font-semibold">educational programming language</span> designed
            for learning—but it's powerful enough for real projects.
          </p>

          {/* Inspired By */}
          <div className="bg-card/50 border border-border/50 rounded-2xl p-6 md:p-8 backdrop-blur-sm mb-8">
            <p className="text-lg text-muted-foreground mb-4">
              <span className="text-primary font-semibold">Inspired by:</span>{" "}
              <a
                href="https://github.com/DulLabs/bhai-lang"
                target="_blank"
                rel="noopener noreferrer"
                className="text-purple-400 hover:text-purple-300 underline decoration-dotted"
              >
                BhaiLang
              </a>{" "}
              (Hindi) and{" "}
              <a
                href="https://github.com/vedic-lang/vedic"
                target="_blank"
                rel="noopener noreferrer"
                className="text-pink-400 hover:text-pink-300 underline decoration-dotted"
              >
                Vedic
              </a>{" "}
              (Sanskrit) — regional programming languages that brought coding to India.
            </p>
            <p className="text-lg text-muted-foreground">
              <span className="text-amber-500 font-semibold">The Difference:</span> While those are excellent{" "}
              <span className="italic">toy languages</span> with basic features, BanglaCode is a{" "}
              <span className="text-primary font-semibold">full-featured educational language</span>.
            </p>
          </div>
        </div>

        {/* What You Can Build */}
        <div className="max-w-6xl mx-auto">
          <h3 className="text-2xl md:text-3xl font-bold text-center mb-12">
            What You Can Actually Build
          </h3>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {/* Backend Development */}
            <div className="group bg-card border border-border rounded-xl p-6 hover:border-primary/50 hover:shadow-lg hover:-translate-y-1 transition-all duration-200 animate-fade-in-up" style={{ animationDelay: "0.1s" }}>
              <div className="w-12 h-12 rounded-lg bg-gradient-to-br from-blue-500 to-cyan-500 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                <Network className="w-6 h-6 text-white" />
              </div>
              <h4 className="text-xl font-semibold mb-3">Real Backends</h4>
              <ul className="space-y-2 text-muted-foreground">
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>HTTP servers & REST APIs</span>
                </li>
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>WebSocket servers</span>
                </li>
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>TCP/UDP networking</span>
                </li>
              </ul>
            </div>

            {/* Database Connectivity */}
            <div className="group bg-card border border-border rounded-xl p-6 hover:border-primary/50 hover:shadow-lg hover:-translate-y-1 transition-all duration-200 animate-fade-in-up" style={{ animationDelay: "0.2s" }}>
              <div className="w-12 h-12 rounded-lg bg-gradient-to-br from-purple-500 to-pink-500 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                <Database className="w-6 h-6 text-white" />
              </div>
              <h4 className="text-xl font-semibold mb-3">Database Connectors</h4>
              <ul className="space-y-2 text-muted-foreground">
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>PostgreSQL & MySQL</span>
                </li>
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>MongoDB (NoSQL)</span>
                </li>
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>Redis (Caching)</span>
                </li>
              </ul>
            </div>

            {/* Modular Code */}
            <div className="group bg-card border border-border rounded-xl p-6 hover:border-primary/50 hover:shadow-lg hover:-translate-y-1 transition-all duration-200 animate-fade-in-up" style={{ animationDelay: "0.3s" }}>
              <div className="w-12 h-12 rounded-lg bg-gradient-to-br from-green-500 to-emerald-500 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                <Package className="w-6 h-6 text-white" />
              </div>
              <h4 className="text-xl font-semibold mb-3">Modular Code</h4>
              <ul className="space-y-2 text-muted-foreground">
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>Import/Export system</span>
                </li>
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>Code organization</span>
                </li>
                <li className="flex items-start gap-2">
                  <span className="text-primary mt-1">✓</span>
                  <span>Reusable modules</span>
                </li>
              </ul>
            </div>

            {/* Advanced Features */}
            <div className="group bg-card border border-border rounded-xl p-6 hover:border-primary/50 hover:shadow-lg hover:-translate-y-1 transition-all duration-200 animate-fade-in-up md:col-span-2 lg:col-span-3" style={{ animationDelay: "0.4s" }}>
              <div className="w-12 h-12 rounded-lg bg-gradient-to-br from-yellow-500 to-orange-500 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                <Code2 className="w-6 h-6 text-white" />
              </div>
              <h4 className="text-xl font-semibold mb-3">Complex Programming Concepts</h4>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <ul className="space-y-2 text-muted-foreground">
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-1">✓</span>
                    <span>Object-Oriented Programming</span>
                  </li>
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-1">✓</span>
                    <span>Async/Await & Promises</span>
                  </li>
                </ul>
                <ul className="space-y-2 text-muted-foreground">
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-1">✓</span>
                    <span>Error Handling (Try/Catch)</span>
                  </li>
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-1">✓</span>
                    <span>File I/O & System Access</span>
                  </li>
                </ul>
                <ul className="space-y-2 text-muted-foreground">
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-1">✓</span>
                    <span>JSON & HTTP Client</span>
                  </li>
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-1">✓</span>
                    <span>130+ Built-in Functions</span>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        {/* Bottom Statement */}
        <div className="max-w-4xl mx-auto text-center mt-16 animate-fade-in-up" style={{ animationDelay: "0.5s" }}>
          <div className="bg-gradient-to-r from-purple-500/10 via-pink-500/10 to-purple-500/10 border border-purple-500/20 rounded-2xl p-8">
            <p className="text-xl md:text-2xl font-semibold mb-4">
              <span className="text-green-400">Perfect for:</span>{" "}
              <span className="text-muted-foreground">Bengali students learning programming, educators teaching CS, hobbyists building projects</span>
            </p>
            <p className="text-xl md:text-2xl font-semibold">
              <span className="text-red-400">Not a replacement for:</span>{" "}
              <span className="text-muted-foreground">Production enterprise applications—use JavaScript, Python, Go, etc. for that</span>
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}
