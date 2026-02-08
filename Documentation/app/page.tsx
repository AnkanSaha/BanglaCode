"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import { ArrowRight, Code2, Terminal, Cpu, Download } from "lucide-react";
import Typewriter from "typewriter-effect";
import packageJson from "../package.json";

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen">
      {/* Hero Section */}
      <section className="relative flex flex-col items-center justify-center min-h-[80vh] text-center px-4 overflow-hidden">
        {/* Background Grid/Glow */}
        <div className="absolute inset-0 bg-grid-white/[0.02] bg-[length:50px_50px] pointer-events-none" />
        <div className="absolute inset-0 flex items-center justify-center bg-background [mask-image:radial-gradient(ellipse_at_center,transparent_20%,black)] pointer-events-none" />
        <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[500px] h-[500px] bg-primary/20 rounded-full blur-[100px] pointer-events-none" />

        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          className="relative z-10 max-w-4xl mx-auto space-y-6"
        >
          <div className="inline-block px-3 py-1 mb-4 text-xs font-semibold tracking-wider text-primary uppercase bg-primary/10 rounded-full border border-primary/20">
            v{packageJson.version} Now Available
          </div>

          <h1 className="text-5xl md:text-7xl font-bold tracking-tight bg-clip-text text-transparent bg-gradient-to-b from-white to-white/60 pb-2">
            Coding in <span className="text-primary">Bengali</span>
            <br />
            Made Simple.
          </h1>

          <div className="text-xl md:text-2xl text-muted-foreground h-20">
            <Typewriter
              options={{
                strings: [
                  "Write code in your mother tongue.",
                  "Strict syntax, modern features.",
                  "Built for students, ready for logic.",
                ],
                autoStart: true,
                loop: true,
                delay: 50,
              }}
            />
          </div>

          <div className="flex flex-col sm:flex-row items-center justify-center gap-4 mt-8">
            <Link
              href="https://github.com/nexoral/BanglaCode/releases"
              target="_blank"
              className="px-8 py-3 text-lg font-medium text-white bg-green-600 hover:bg-green-700 rounded-full transition-all hover:shadow-[0_0_20px_rgba(34,197,94,0.5)] flex items-center gap-2"
            >
              Download Now <Download className="w-5 h-5" />
            </Link>
            <Link
              href="/docs/installation"
              className="px-8 py-3 text-lg font-medium text-white bg-primary hover:bg-primary/90 rounded-full transition-all hover:shadow-[0_0_20px_rgba(124,58,237,0.5)] flex items-center gap-2"
            >
              Get Started <ArrowRight className="w-5 h-5" />
            </Link>
            <Link
              href="/playground"
              className="px-8 py-3 text-lg font-medium text-foreground bg-secondary hover:bg-secondary/80 rounded-full transition-all flex items-center gap-2"
            >
              Try Playground <Terminal className="w-5 h-5" />
            </Link>
          </div>
        </motion.div>
      </section>

      {/* Features/Why Section */}
      <section className="py-24 bg-accent/5">
        <div className="container mx-auto px-4">
          <motion.div
            initial={{ opacity: 0 }}
            whileInView={{ opacity: 1 }}
            viewport={{ once: true }}
            className="text-center mb-16"
          >
            <h2 className="text-3xl font-bold mb-4">Why BanglaCode?</h2>
            <p className="text-muted-foreground max-w-2xl mx-auto">
              Bridging the gap between logic and language. Designed specifically to make programming concepts accessible to Bengali speakers.
            </p>
          </motion.div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {[
              {
                icon: <Code2 className="w-8 h-8 text-blue-500" />,
                title: "Native Syntax",
                desc: "Keywords like 'dhoro', 'jodi', and 'ghuriye' make understanding control flow intuitive.",
              },
              {
                icon: <Cpu className="w-8 h-8 text-purple-500" />,
                title: "Technically Roboust",
                desc: "Built on Go, compiled to native binaries. Supports strict typing, recursion, and OOP.",
              },
              {
                icon: <Terminal className="w-8 h-8 text-green-500" />,
                title: "Student Friendly",
                desc: "Clear error messages and documentation designed to help you transition to C/C++/Java later.",
              },
            ].map((feature, i) => (
              <motion.div
                key={i}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                transition={{ delay: i * 0.1 }}
                viewport={{ once: true }}
                className="p-6 rounded-2xl bg-card border border-border hover:border-primary/50 transition-colors"
              >
                <div className="mb-4 p-3 bg-secondary rounded-lg inline-block">
                  {feature.icon}
                </div>
                <h3 className="text-xl font-semibold mb-2">{feature.title}</h3>
                <p className="text-muted-foreground">{feature.desc}</p>
              </motion.div>
            ))}
          </div>
        </div>
      </section>

      {/* Scenes/Use Cases */}
      <section className="py-24">
        <div className="container mx-auto px-4">
          <motion.div
            initial={{ opacity: 0 }}
            whileInView={{ opacity: 1 }}
            viewport={{ once: true }}
            className="text-center mb-16"
          >
            <h2 className="text-3xl font-bold mb-4">Where to use it?</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-12 mt-12 text-left">
              <div className="space-y-6">
                <div className="bg-card border border-border rounded-xl p-6 relative overflow-hidden group hover:border-primary/50 transition-colors">
                  <div className="absolute top-0 right-0 p-4 opacity-10 group-hover:opacity-20 transition-opacity">
                    <span className="text-9xl font-bold">1</span>
                  </div>
                  <h3 className="text-2xl font-bold mb-2 text-primary">Education</h3>
                  <p className="text-muted-foreground mb-4">
                    Perfect for computer science practicals in schools and colleges in West Bengal and Bangladesh.
                    Teach logic without the language barrier.
                  </p>
                  <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm">
                    <span className="text-purple-400">dhoro</span> students = [<span className="text-green-400">"Raju"</span>, <span className="text-green-400">"Priya"</span>];<br />
                    <span className="text-blue-400">ghuriye</span> (<span className="text-purple-400">dhoro</span> s : students) &#123;<br />
                    &nbsp;&nbsp;<span className="text-yellow-400">dekho</span>(<span className="text-green-400">"Swagatam, "</span> + s);<br />
                    &#125;
                  </div>
                </div>
              </div>
              <div className="space-y-6">
                <div className="bg-card border border-border rounded-xl p-6 relative overflow-hidden group hover:border-primary/50 transition-colors">
                  <div className="absolute top-0 right-0 p-4 opacity-10 group-hover:opacity-20 transition-opacity">
                    <span className="text-9xl font-bold">2</span>
                  </div>
                  <h3 className="text-2xl font-bold mb-2 text-primary">Algorithms</h3>
                  <p className="text-muted-foreground mb-4">
                    Implement complex algorithms like Sorting, Searching, and Graph traversals.
                    The syntax is close enough to C/Java to make knowledge transferable.
                  </p>
                  <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm">
                    <span className="text-blue-400">kaj</span> factorial(n) &#123;<br />
                    &nbsp;&nbsp;<span className="text-blue-400">jodi</span> (n &lt;= 1) <span className="text-purple-400">ferao</span> 1;<br />
                    &nbsp;&nbsp;<span className="text-purple-400">ferao</span> n * factorial(n - 1);<br />
                    &#125;
                  </div>
                </div>
              </div>
            </div>
          </motion.div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-8 border-t border-border mt-auto">
        <div className="container mx-auto px-4 text-center text-muted-foreground">
          <p>© {new Date().getFullYear()} BanglaCode. Built with ❤️ for the Bengali community.</p>
        </div>
      </footer>
    </div>
  );
}
