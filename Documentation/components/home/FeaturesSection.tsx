"use client";

import { motion } from "framer-motion";
import {
  Code2,
  Cpu,
  Terminal,
  Zap,
  Shield,
  BookOpen,
  Layers,
  Globe,
} from "lucide-react";

const features = [
  {
    icon: Code2,
    title: "27 Bengali Keywords",
    description:
      "Natural syntax with keywords like 'dhoro' (let), 'jodi' (if), 'ghuriye' (for) that feel intuitive to Bengali speakers.",
    gradient: "from-blue-500 to-cyan-500",
  },
  {
    icon: Cpu,
    title: "Built on Go",
    description:
      "Lightning-fast execution with a robust tree-walking interpreter. Compiles to native binaries for any platform.",
    gradient: "from-purple-500 to-pink-500",
  },
  {
    icon: Terminal,
    title: "Interactive REPL",
    description:
      "Learn and experiment with an interactive shell. Multi-line support, syntax highlighting, and helpful error messages.",
    gradient: "from-green-500 to-emerald-500",
  },
  {
    icon: Zap,
    title: "50+ Built-in Functions",
    description:
      "Rich standard library with functions for math, strings, arrays, I/O, and more. All documented in Bengali.",
    gradient: "from-yellow-500 to-orange-500",
  },
  {
    icon: Shield,
    title: "Error Handling",
    description:
      "Try-catch-finally blocks with 'chesta', 'dhoro_bhul', and 'shesh'. Handle errors gracefully in your native language.",
    gradient: "from-red-500 to-pink-500",
  },
  {
    icon: BookOpen,
    title: "OOP Support",
    description:
      "Classes, constructors, methods, and 'this' binding with 'sreni', 'shuru', 'kaj', and 'ei'. Learn OOP concepts naturally.",
    gradient: "from-indigo-500 to-purple-500",
  },
  {
    icon: Layers,
    title: "Module System",
    description:
      "Import and export code with 'ano' and 'pathao'. Organize your code into reusable modules.",
    gradient: "from-teal-500 to-cyan-500",
  },
  {
    icon: Globe,
    title: "Cross-Platform",
    description:
      "Run on Windows, macOS, and Linux. One codebase, deploy anywhere. Perfect for educational institutions.",
    gradient: "from-violet-500 to-fuchsia-500",
  },
];

export default function FeaturesSection() {
  return (
    <section className="py-32 relative overflow-hidden">
      {/* Background */}
      <div className="absolute inset-0 bg-[linear-gradient(rgba(124,58,237,0.02)_1px,transparent_1px),linear-gradient(90deg,rgba(124,58,237,0.02)_1px,transparent_1px)] bg-[size:60px_60px]" />

      <div className="container mx-auto px-4 relative z-10">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="text-center mb-20"
        >
          <motion.span
            initial={{ opacity: 0, scale: 0.8 }}
            whileInView={{ opacity: 1, scale: 1 }}
            viewport={{ once: true }}
            className="inline-block px-4 py-2 text-sm font-semibold text-primary bg-primary/10 rounded-full border border-primary/20 mb-6"
          >
            Features
          </motion.span>
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Everything You Need to{" "}
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-500">
              Learn & Build
            </span>
          </h2>
          <p className="text-xl text-muted-foreground max-w-3xl mx-auto">
            A complete programming language with modern features, designed from the ground up for Bengali speakers.
          </p>
        </motion.div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          {features.map((feature, index) => (
            <motion.div
              key={index}
              initial={{ opacity: 0, y: 30 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true, margin: "-50px" }}
              transition={{ duration: 0.5, delay: index * 0.1 }}
              whileHover={{ y: -10, scale: 1.02 }}
              className="group relative"
            >
              <div className="absolute inset-0 bg-gradient-to-r opacity-0 group-hover:opacity-100 blur-xl transition-opacity duration-500 -z-10"
                style={{
                  background: `linear-gradient(135deg, var(--tw-gradient-from), var(--tw-gradient-to))`,
                }}
              />
              <div className="h-full bg-card border border-border rounded-2xl p-6 hover:border-primary/50 transition-all duration-300 hover:shadow-2xl">
                <div
                  className={`w-14 h-14 rounded-xl bg-gradient-to-r ${feature.gradient} flex items-center justify-center mb-5 group-hover:scale-110 transition-transform duration-300`}
                >
                  <feature.icon className="w-7 h-7 text-white" />
                </div>
                <h3 className="text-xl font-bold mb-3 group-hover:text-primary transition-colors">
                  {feature.title}
                </h3>
                <p className="text-muted-foreground leading-relaxed">
                  {feature.description}
                </p>
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
}
