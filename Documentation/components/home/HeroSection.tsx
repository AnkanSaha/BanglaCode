"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { ArrowRight, Download, Sparkles } from "lucide-react";
import Typewriter from "typewriter-effect";
import AnimatedBackground from "./AnimatedBackground";
import { getRepoStats } from "@/lib/github";

interface HeroSectionProps {
  version: string;
}

export default function HeroSection({ version }: HeroSectionProps) {
  const [stars, setStars] = useState<number | null>(null);
  const [forks, setForks] = useState<number | null>(null);

  useEffect(() => {
    getRepoStats().then((data) => {
      setStars(data.stars);
      setForks(data.forks);
    });
  }, []);

  return (
    <section className="relative flex flex-col items-center justify-center min-h-screen text-center px-4 overflow-hidden">
      <AnimatedBackground />

      {/* Floating Code Snippets - CSS animated */}
      <div className="absolute left-[5%] top-[20%] hidden lg:block opacity-60 animate-fade-in-up" style={{ animationDelay: "0.8s" }}>
        <div className="bg-card/80 backdrop-blur-sm border border-border rounded-lg p-4 font-mono text-sm text-left shadow-xl">
          <span className="text-purple-400">dhoro</span> naam = <span className="text-green-400">&quot;BanglaCode&quot;</span>;
        </div>
      </div>

      <div className="absolute right-[5%] top-[30%] hidden lg:block opacity-60 animate-fade-in-up" style={{ animationDelay: "1s" }}>
        <div className="bg-card/80 backdrop-blur-sm border border-border rounded-lg p-4 font-mono text-sm text-left shadow-xl">
          <span className="text-blue-400">kaj</span> swagatam() &#123;<br />
          &nbsp;&nbsp;<span className="text-yellow-400">dekho</span>(<span className="text-green-400">&quot;Hello!&quot;</span>);<br />
          &#125;
        </div>
      </div>

      <div className="absolute left-[10%] bottom-[25%] hidden lg:block opacity-50 animate-fade-in-up" style={{ animationDelay: "1.2s" }}>
        <div className="bg-card/80 backdrop-blur-sm border border-border rounded-lg p-4 font-mono text-sm text-left shadow-xl">
          <span className="text-blue-400">jodi</span> (sotti) &#123; ... &#125;
        </div>
      </div>

      {/* Main Content */}
      <div className="relative z-10 max-w-5xl mx-auto space-y-8">
        {/* Version Badge */}
        <div className="inline-flex items-center gap-2 px-4 py-2 text-sm font-semibold text-primary bg-primary/10 rounded-full border border-primary/20 animate-fade-in-up">
          <Sparkles className="w-4 h-4" />
          v{version} Now Available
        </div>

        {/* Main Heading */}
        <h1 className="text-5xl sm:text-6xl md:text-8xl font-black tracking-tight animate-fade-in-up" style={{ animationDelay: "0.1s" }}>
          <span className="bg-clip-text text-transparent bg-gradient-to-b from-white via-white to-white/40">
            Coding in
          </span>
          <br />
          <span className="bg-clip-text text-transparent bg-gradient-to-r from-purple-400 via-pink-500 to-purple-600">
            Bengali
          </span>
        </h1>

        {/* Typewriter */}
        <div className="text-xl md:text-2xl lg:text-3xl text-muted-foreground h-16 md:h-20 animate-fade-in-up" style={{ animationDelay: "0.2s" }}>
          <Typewriter
            options={{
              strings: [
                "An educational language for learning.",
                "Inspired by BhaiLang & Vedic, but more powerful.",
                "Build backends. Connect databases. Write modules.",
                "Built for 300 million Bengali speakers.",
                "Made in West Bengal, India 🇮🇳",
              ],
              autoStart: true,
              loop: true,
              delay: 40,
              deleteSpeed: 20,
            }}
          />
        </div>

        {/* Stats */}
        <div className="flex flex-wrap items-center justify-center gap-8 text-muted-foreground animate-fade-in-up" style={{ animationDelay: "0.3s" }}>
          <div className="flex items-center gap-2">
            <span className="text-3xl font-bold text-yellow-500">
              {stars !== null ? stars : <span className="animate-pulse">--</span>}
            </span>
            <span>GitHub Stars</span>
          </div>
          <div className="w-px h-8 bg-border hidden sm:block" />
          <div className="flex items-center gap-2">
            <span className="text-3xl font-bold text-green-500">
              {forks !== null ? forks : <span className="animate-pulse">--</span>}
            </span>
            <span>Forks</span>
          </div>
          <div className="w-px h-8 bg-border hidden sm:block" />
          <div className="flex items-center gap-2">
            <span className="text-3xl font-bold text-purple-500">29</span>
            <span>Bengali Keywords</span>
          </div>
        </div>

        {/* CTA Buttons */}
        <div className="flex flex-col sm:flex-row items-center justify-center gap-4 pt-4 animate-fade-in-up" style={{ animationDelay: "0.4s" }}>
          <Link
            href="https://github.com/nexoral/BanglaCode/releases"
            target="_blank"
            className="group px-8 py-4 text-lg font-semibold text-white bg-gradient-to-r from-green-500 to-emerald-600 rounded-full transition-all duration-200 hover:shadow-[0_0_30px_rgba(34,197,94,0.4)] flex items-center gap-3 hover:-translate-y-1"
          >
            Download Now <Download className="w-5 h-5 group-hover:animate-bounce" />
          </Link>

          <Link
            href="/docs/installation"
            className="group px-8 py-4 text-lg font-semibold text-white bg-gradient-to-r from-purple-500 to-pink-600 rounded-full transition-all duration-200 hover:shadow-[0_0_30px_rgba(168,85,247,0.4)] flex items-center gap-3 hover:-translate-y-1"
          >
            Get Started <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
          </Link>
        </div>
      </div>

      {/* Scroll Indicator */}
      <div className="absolute bottom-8 left-1/2 -translate-x-1/2 animate-fade-in-up" style={{ animationDelay: "1.5s" }}>
        <div className="w-6 h-10 border-2 border-muted-foreground/30 rounded-full flex justify-center pt-2">
          <div className="w-1.5 h-3 bg-primary rounded-full animate-bounce" />
        </div>
      </div>
    </section>
  );
}
