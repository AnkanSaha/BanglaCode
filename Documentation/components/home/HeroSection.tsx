"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import { ArrowRight, Download, Sparkles } from "lucide-react";
import Typewriter from "typewriter-effect";
import AnimatedBackground from "./AnimatedBackground";

interface HeroSectionProps {
  version: string;
  stars: number;
  forks: number;
}

export default function HeroSection({ version, stars, forks }: HeroSectionProps) {
  return (
    <section className="relative flex flex-col items-center justify-center min-h-screen text-center px-4 overflow-hidden">
      <AnimatedBackground />

      {/* Floating Code Snippets */}
      <motion.div
        className="absolute left-[5%] top-[20%] hidden lg:block"
        initial={{ opacity: 0, x: -50 }}
        animate={{ opacity: 0.6, x: 0 }}
        transition={{ delay: 1, duration: 1 }}
      >
        <div className="bg-card/80 backdrop-blur-sm border border-border rounded-lg p-4 font-mono text-sm text-left shadow-xl">
          <span className="text-purple-400">dhoro</span> naam = <span className="text-green-400">"BanglaCode"</span>;
        </div>
      </motion.div>

      <motion.div
        className="absolute right-[5%] top-[30%] hidden lg:block"
        initial={{ opacity: 0, x: 50 }}
        animate={{ opacity: 0.6, x: 0 }}
        transition={{ delay: 1.2, duration: 1 }}
      >
        <div className="bg-card/80 backdrop-blur-sm border border-border rounded-lg p-4 font-mono text-sm text-left shadow-xl">
          <span className="text-blue-400">kaj</span> swagatam() &#123;<br />
          &nbsp;&nbsp;<span className="text-yellow-400">dekho</span>(<span className="text-green-400">"Hello!"</span>);<br />
          &#125;
        </div>
      </motion.div>

      <motion.div
        className="absolute left-[10%] bottom-[25%] hidden lg:block"
        initial={{ opacity: 0, y: 50 }}
        animate={{ opacity: 0.5, y: 0 }}
        transition={{ delay: 1.4, duration: 1 }}
      >
        <div className="bg-card/80 backdrop-blur-sm border border-border rounded-lg p-4 font-mono text-sm text-left shadow-xl">
          <span className="text-blue-400">jodi</span> (sotti) &#123; ... &#125;
        </div>
      </motion.div>

      {/* Main Content */}
      <motion.div
        initial={{ opacity: 0, y: 30 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 1 }}
        className="relative z-10 max-w-5xl mx-auto space-y-8"
      >
        {/* Version Badge */}
        <motion.div
          initial={{ opacity: 0, scale: 0.8 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ delay: 0.2, duration: 0.5 }}
          className="inline-flex items-center gap-2 px-4 py-2 text-sm font-semibold text-primary bg-primary/10 rounded-full border border-primary/20"
        >
          <Sparkles className="w-4 h-4" />
          v{version} Now Available
        </motion.div>

        {/* Main Heading */}
        <motion.h1
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.3, duration: 0.8 }}
          className="text-5xl sm:text-6xl md:text-8xl font-black tracking-tight"
        >
          <span className="bg-clip-text text-transparent bg-gradient-to-b from-white via-white to-white/40">
            Coding in
          </span>
          <br />
          <motion.span
            className="bg-clip-text text-transparent bg-gradient-to-r from-purple-400 via-pink-500 to-purple-600"
            animate={{
              backgroundPosition: ["0% 50%", "100% 50%", "0% 50%"],
            }}
            transition={{ duration: 5, repeat: Infinity }}
            style={{ backgroundSize: "200% 200%" }}
          >
            Bengali
          </motion.span>
        </motion.h1>

        {/* Typewriter */}
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 0.5, duration: 0.8 }}
          className="text-xl md:text-2xl lg:text-3xl text-muted-foreground h-16 md:h-20"
        >
          <Typewriter
            options={{
              strings: [
                "Write code in your mother tongue.",
                "Learn programming concepts naturally.",
                "Built for 300 million Bengali speakers.",
                "From West Bengal to Bangladesh.",
              ],
              autoStart: true,
              loop: true,
              delay: 40,
              deleteSpeed: 20,
            }}
          />
        </motion.div>

        {/* Stats */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.7, duration: 0.8 }}
          className="flex flex-wrap items-center justify-center gap-8 text-muted-foreground"
        >
          <div className="flex items-center gap-2">
            <span className="text-3xl font-bold text-yellow-500">{stars}</span>
            <span>GitHub Stars</span>
          </div>
          <div className="w-px h-8 bg-border hidden sm:block" />
          <div className="flex items-center gap-2">
            <span className="text-3xl font-bold text-green-500">{forks}</span>
            <span>Forks</span>
          </div>
          <div className="w-px h-8 bg-border hidden sm:block" />
          <div className="flex items-center gap-2">
            <span className="text-3xl font-bold text-purple-500">27</span>
            <span>Bengali Keywords</span>
          </div>
        </motion.div>

        {/* CTA Buttons */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.9, duration: 0.8 }}
          className="flex flex-col sm:flex-row items-center justify-center gap-4 pt-4"
        >
          <Link
            href="https://github.com/nexoral/BanglaCode/releases"
            target="_blank"
            className="group relative px-8 py-4 text-lg font-semibold text-white bg-gradient-to-r from-green-500 to-emerald-600 rounded-full transition-all duration-300 hover:shadow-[0_0_40px_rgba(34,197,94,0.4)] flex items-center gap-3 overflow-hidden"
          >
            <span className="relative z-10 flex items-center gap-2">
              Download Now <Download className="w-5 h-5 group-hover:animate-bounce" />
            </span>
            <div className="absolute inset-0 bg-gradient-to-r from-emerald-600 to-green-500 opacity-0 group-hover:opacity-100 transition-opacity" />
          </Link>

          <Link
            href="/docs/installation"
            className="group relative px-8 py-4 text-lg font-semibold text-white bg-gradient-to-r from-purple-500 to-pink-600 rounded-full transition-all duration-300 hover:shadow-[0_0_40px_rgba(168,85,247,0.4)] flex items-center gap-3 overflow-hidden"
          >
            <span className="relative z-10 flex items-center gap-2">
              Get Started <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
            </span>
            <div className="absolute inset-0 bg-gradient-to-r from-pink-600 to-purple-500 opacity-0 group-hover:opacity-100 transition-opacity" />
          </Link>
        </motion.div>
      </motion.div>

      {/* Scroll Indicator */}
      <motion.div
        className="absolute bottom-8 left-1/2 -translate-x-1/2"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ delay: 1.5 }}
      >
        <motion.div
          className="w-6 h-10 border-2 border-muted-foreground/30 rounded-full flex justify-center"
          animate={{ y: [0, 10, 0] }}
          transition={{ duration: 2, repeat: Infinity }}
        >
          <motion.div
            className="w-1.5 h-3 bg-primary rounded-full mt-2"
            animate={{ opacity: [1, 0.3, 1], y: [0, 4, 0] }}
            transition={{ duration: 2, repeat: Infinity }}
          />
        </motion.div>
      </motion.div>
    </section>
  );
}
