"use client";

import { motion, AnimatePresence } from "framer-motion";
import { useState } from "react";
import { Play, Copy, Check } from "lucide-react";

const codeExamples = [
  {
    title: "Hello World",
    category: "Basics",
    code: `// Your first BanglaCode program
dekho("Namaskar, Prithibi!");

// Variables
dhoro naam = "BanglaCode";
dhoro version = 5.1;

dekho("Welcome to " + naam);`,
    output: `Namaskar, Prithibi!
Welcome to BanglaCode`,
  },
  {
    title: "Functions",
    category: "Core",
    code: `// Define a function
kaj factorial(n) {
    jodi (n <= 1) {
        ferao 1;
    }
    ferao n * factorial(n - 1);
}

// Call the function
dhoro result = factorial(5);
dekho("5! = " + result);`,
    output: `5! = 120`,
  },
  {
    title: "Loops",
    category: "Control Flow",
    code: `// For loop
ghuriye (dhoro i = 1; i <= 5; i = i + 1) {
    dekho("Count: " + i);
}

// While loop
dhoro x = 3;
jotokkhon (x > 0) {
    dekho("Countdown: " + x);
    x = x - 1;
}`,
    output: `Count: 1
Count: 2
Count: 3
Count: 4
Count: 5
Countdown: 3
Countdown: 2
Countdown: 1`,
  },
  {
    title: "Classes",
    category: "OOP",
    code: `// Define a class
sreni Person {
    shuru(naam, boyes) {
        ei.naam = naam;
        ei.boyes = boyes;
    }

    kaj introduce() {
        dekho("Ami " + ei.naam);
        dekho("Boyesh: " + ei.boyes);
    }
}

// Create an instance
dhoro p = notun Person("Ankan", 21);
p.introduce();`,
    output: `Ami Ankan
Boyesh: 21`,
  },
];

// Keywords for syntax highlighting
const keywords = ["dhoro", "jodi", "nahole", "ghuriye", "jotokkhon", "kaj", "ferao", "sreni", "shuru", "notun", "ei", "sotti", "mittha", "khali"];
const builtins = ["dekho", "dorghyo", "dhokao", "bortoman", "tipo"];

function SyntaxHighlightedCode({ code }: { code: string }) {
  const lines = code.split("\n");

  const highlightLine = (line: string, lineIndex: number) => {
    const tokens: React.ReactNode[] = [];
    let remaining = line;
    let keyIndex = 0;

    // Check if it's a comment line
    if (remaining.trim().startsWith("//")) {
      return (
        <div key={lineIndex} className="flex">
          <span className="text-gray-600 w-8 select-none text-right pr-4">{lineIndex + 1}</span>
          <span className="text-gray-500">{line}</span>
        </div>
      );
    }

    const parts: React.ReactNode[] = [];
    let i = 0;

    while (i < line.length) {
      // Check for strings
      if (line[i] === '"') {
        const endQuote = line.indexOf('"', i + 1);
        if (endQuote !== -1) {
          const str = line.slice(i, endQuote + 1);
          parts.push(<span key={keyIndex++} className="text-green-400">{str}</span>);
          i = endQuote + 1;
          continue;
        }
      }

      // Check for keywords and builtins
      let matched = false;
      for (const keyword of keywords) {
        if (line.slice(i).startsWith(keyword) &&
            (i === 0 || !/\w/.test(line[i - 1])) &&
            (i + keyword.length >= line.length || !/\w/.test(line[i + keyword.length]))) {
          parts.push(<span key={keyIndex++} className="text-purple-400">{keyword}</span>);
          i += keyword.length;
          matched = true;
          break;
        }
      }
      if (matched) continue;

      for (const builtin of builtins) {
        if (line.slice(i).startsWith(builtin) &&
            (i === 0 || !/\w/.test(line[i - 1])) &&
            (i + builtin.length >= line.length || !/\w/.test(line[i + builtin.length]))) {
          parts.push(<span key={keyIndex++} className="text-yellow-400">{builtin}</span>);
          i += builtin.length;
          matched = true;
          break;
        }
      }
      if (matched) continue;

      // Check for numbers
      if (/\d/.test(line[i])) {
        let numEnd = i;
        while (numEnd < line.length && /[\d.]/.test(line[numEnd])) {
          numEnd++;
        }
        parts.push(<span key={keyIndex++} className="text-orange-400">{line.slice(i, numEnd)}</span>);
        i = numEnd;
        continue;
      }

      // Regular character
      parts.push(<span key={keyIndex++} className="text-gray-300">{line[i]}</span>);
      i++;
    }

    return (
      <div key={lineIndex} className="flex">
        <span className="text-gray-600 w-8 select-none text-right pr-4">{lineIndex + 1}</span>
        <span>{parts}</span>
      </div>
    );
  };

  return (
    <div className="font-mono text-sm leading-relaxed">
      {lines.map((line, index) => highlightLine(line, index))}
    </div>
  );
}

export default function CodeShowcase() {
  const [activeIndex, setActiveIndex] = useState(0);
  const [copied, setCopied] = useState(false);

  const copyCode = () => {
    navigator.clipboard.writeText(codeExamples[activeIndex].code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <section className="py-32 bg-gradient-to-b from-background via-accent/5 to-background relative overflow-hidden">
      <div className="container mx-auto px-4">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="text-center mb-16"
        >
          <motion.span
            initial={{ opacity: 0, scale: 0.8 }}
            whileInView={{ opacity: 1, scale: 1 }}
            viewport={{ once: true }}
            className="inline-block px-4 py-2 text-sm font-semibold text-primary bg-primary/10 rounded-full border border-primary/20 mb-6"
          >
            See It In Action
          </motion.span>
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Code That{" "}
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-500">
              Speaks Bengali
            </span>
          </h2>
          <p className="text-xl text-muted-foreground max-w-3xl mx-auto">
            Clean, intuitive syntax that feels natural. See how BanglaCode makes programming accessible.
          </p>
        </motion.div>

        <div className="max-w-6xl mx-auto">
          {/* Tab Navigation */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="flex flex-wrap justify-center gap-3 mb-8"
          >
            {codeExamples.map((example, index) => (
              <motion.button
                key={index}
                onClick={() => setActiveIndex(index)}
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
                className={`px-6 py-3 rounded-full font-medium transition-all duration-300 ${
                  activeIndex === index
                    ? "bg-primary text-white shadow-lg shadow-primary/25"
                    : "bg-card border border-border hover:border-primary/50 text-muted-foreground hover:text-foreground"
                }`}
              >
                <span className="hidden sm:inline text-xs opacity-70 mr-2">
                  {example.category}:
                </span>
                {example.title}
              </motion.button>
            ))}
          </motion.div>

          {/* Code Display */}
          <motion.div
            initial={{ opacity: 0, scale: 0.95 }}
            whileInView={{ opacity: 1, scale: 1 }}
            viewport={{ once: true }}
            className="grid md:grid-cols-2 gap-6"
          >
            {/* Code Editor */}
            <div className="relative group">
              <div className="absolute -inset-1 bg-gradient-to-r from-purple-500 to-pink-500 rounded-2xl opacity-20 group-hover:opacity-30 blur transition-opacity" />
              <div className="relative bg-[#1a1a2e] rounded-2xl overflow-hidden border border-border">
                {/* Editor Header */}
                <div className="flex items-center justify-between px-4 py-3 border-b border-border/50 bg-black/20">
                  <div className="flex items-center gap-2">
                    <div className="flex gap-2">
                      <div className="w-3 h-3 rounded-full bg-red-500" />
                      <div className="w-3 h-3 rounded-full bg-yellow-500" />
                      <div className="w-3 h-3 rounded-full bg-green-500" />
                    </div>
                    <span className="text-sm text-muted-foreground ml-3">
                      {codeExamples[activeIndex].title.toLowerCase().replace(" ", "_")}.bang
                    </span>
                  </div>
                  <button
                    onClick={copyCode}
                    className="p-2 hover:bg-white/10 rounded-lg transition-colors"
                  >
                    {copied ? (
                      <Check className="w-4 h-4 text-green-500" />
                    ) : (
                      <Copy className="w-4 h-4 text-muted-foreground" />
                    )}
                  </button>
                </div>

                {/* Code Content */}
                <AnimatePresence mode="wait">
                  <motion.div
                    key={activeIndex}
                    initial={{ opacity: 0, y: 10 }}
                    animate={{ opacity: 1, y: 0 }}
                    exit={{ opacity: 0, y: -10 }}
                    transition={{ duration: 0.3 }}
                    className="p-6 overflow-x-auto"
                  >
                    <SyntaxHighlightedCode code={codeExamples[activeIndex].code} />
                  </motion.div>
                </AnimatePresence>
              </div>
            </div>

            {/* Output */}
            <div className="relative group">
              <div className="absolute -inset-1 bg-gradient-to-r from-green-500 to-emerald-500 rounded-2xl opacity-20 group-hover:opacity-30 blur transition-opacity" />
              <div className="relative bg-[#0d1117] rounded-2xl overflow-hidden border border-border h-full">
                {/* Terminal Header */}
                <div className="flex items-center gap-2 px-4 py-3 border-b border-border/50 bg-black/20">
                  <Play className="w-4 h-4 text-green-500" />
                  <span className="text-sm text-muted-foreground">Output</span>
                </div>

                {/* Output Content */}
                <AnimatePresence mode="wait">
                  <motion.pre
                    key={activeIndex}
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    exit={{ opacity: 0 }}
                    transition={{ duration: 0.3, delay: 0.2 }}
                    className="p-6 text-sm font-mono text-green-400 leading-relaxed"
                  >
                    {codeExamples[activeIndex].output.split("\n").map((line, i) => (
                      <motion.div
                        key={i}
                        initial={{ opacity: 0, x: -10 }}
                        animate={{ opacity: 1, x: 0 }}
                        transition={{ delay: 0.3 + i * 0.1 }}
                      >
                        <span className="text-gray-600 mr-2">&gt;</span>
                        {line}
                      </motion.div>
                    ))}
                  </motion.pre>
                </AnimatePresence>
              </div>
            </div>
          </motion.div>
        </div>
      </div>
    </section>
  );
}
