"use client";

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
    title: "29 Bengali Keywords",
    description:
      "Natural syntax with keywords like 'dhoro' (let), 'jodi' (if), 'ghuriye' (for), 'proyash' (async) that feel intuitive to Bengali speakers.",
    gradient: "from-blue-500 to-cyan-500",
  },
  {
    icon: Cpu,
    title: "Fast & Efficient",
    description:
      "Go-powered tree-walking interpreter for quick execution. Lightweight runtime with efficient memory usage.",
    gradient: "from-purple-500 to-pink-500",
  },
  {
    icon: Terminal,
    title: "130+ Built-in Functions",
    description:
      "Comprehensive library: math, strings, arrays, HTTP servers, JSON, file I/O, database connectors (PostgreSQL, MySQL, MongoDB, Redis), and OS-level access.",
    gradient: "from-yellow-500 to-orange-500",
  },
  {
    icon: Zap,
    title: "Async/Await Support",
    description:
      "Modern async programming with 'proyash' and 'opekha' keywords. Handle promises and asynchronous operations naturally.",
    gradient: "from-green-500 to-emerald-500",
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
      "Import and export code with 'ano' and 'pathao'. Organize your code into reusable modules with aliases.",
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
        <div className="text-center mb-20 animate-fade-in-up">
          <span className="inline-block px-4 py-2 text-sm font-semibold text-primary bg-primary/10 rounded-full border border-primary/20 mb-6">
            Features
          </span>
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Everything You Need to{" "}
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-500">
              Learn & Build
            </span>
          </h2>
          <p className="text-xl text-muted-foreground max-w-3xl mx-auto">
            An educational programming language with production-grade features—perfect for learning, but powerful enough for real projects.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          {features.map((feature, index) => (
            <div
              key={index}
              className="group relative animate-fade-in-up"
              style={{ animationDelay: `${index * 0.05}s` }}
            >
              <div className="absolute inset-0 bg-gradient-to-r opacity-0 group-hover:opacity-100 blur-xl transition-opacity duration-200 -z-10"
                style={{
                  background: `linear-gradient(135deg, var(--tw-gradient-from), var(--tw-gradient-to))`,
                }}
              />
              <div className="h-full bg-card border border-border rounded-2xl p-6 hover:border-primary/50 transition-all duration-200 hover:shadow-2xl hover:-translate-y-2 will-change-transform">
                <div
                  className={`w-14 h-14 rounded-xl bg-gradient-to-r ${feature.gradient} flex items-center justify-center mb-5 group-hover:scale-110 transition-transform duration-200`}
                >
                  <feature.icon className="w-7 h-7 text-white" />
                </div>
                <h3 className="text-xl font-bold mb-3 group-hover:text-primary transition-colors duration-200">
                  {feature.title}
                </h3>
                <p className="text-muted-foreground leading-relaxed">
                  {feature.description}
                </p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
