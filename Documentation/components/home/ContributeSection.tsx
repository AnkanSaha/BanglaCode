"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import {
  GitBranch,
  Bug,
  BookOpen,
  MessageCircle,
  Star,
  Heart,
  Code,
} from "lucide-react";
import { FaGithub } from "react-icons/fa";
import { VscRepoForked } from "react-icons/vsc";

interface ContributeSectionProps {
  stats: {
    stars: number;
    contributors: number;
    forks: number;
  };
}

const contributionWays = [
  {
    icon: Bug,
    title: "Report Bugs",
    description: "Found an issue? Report it on GitHub Issues with steps to reproduce.",
    link: "https://github.com/nexoral/BanglaCode/issues/new?labels=bug",
    linkText: "Report a Bug",
    color: "text-red-500",
    bgColor: "bg-red-500/10",
  },
  {
    icon: Code,
    title: "Submit Code",
    description: "Fix bugs, add features, or improve performance. All PRs welcome!",
    link: "https://github.com/nexoral/BanglaCode/pulls",
    linkText: "Open a PR",
    color: "text-green-500",
    bgColor: "bg-green-500/10",
  },
  {
    icon: BookOpen,
    title: "Improve Docs",
    description: "Help make documentation clearer. Translate to Bengali or other languages.",
    link: "https://github.com/nexoral/BanglaCode/tree/main/Documentation",
    linkText: "Edit Docs",
    color: "text-blue-500",
    bgColor: "bg-blue-500/10",
  },
  {
    icon: MessageCircle,
    title: "Spread the Word",
    description: "Share BanglaCode with students, teachers, and developers in your network.",
    link: "https://twitter.com/intent/tweet?text=Check%20out%20BanglaCode%20-%20A%20programming%20language%20in%20Bengali!%20%F0%9F%87%A7%D%20https://banglacode.pages.dev",
    linkText: "Share on Twitter",
    color: "text-cyan-500",
    bgColor: "bg-cyan-500/10",
  },
];

export default function ContributeSection({ stats }: ContributeSectionProps) {
  const statsDisplay = [
    { icon: Star, value: stats.stars.toString(), label: "GitHub Stars" },
    { icon: GitBranch, value: stats.contributors.toString(), label: "Contributors" },
    { icon: VscRepoForked, value: stats.forks.toString(), label: "Forks" },
    { icon: Heart, value: "100%", label: "Open Source" },
  ];
  return (
    <section className="py-32 relative overflow-hidden">
      {/* Background */}
      <div className="absolute inset-0">
        <div className="absolute top-0 left-1/4 w-96 h-96 bg-purple-500/5 rounded-full blur-3xl" />
        <div className="absolute bottom-0 right-1/4 w-96 h-96 bg-pink-500/5 rounded-full blur-3xl" />
      </div>

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
            Open Source
          </motion.span>
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Built by the{" "}
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-500">
              Community
            </span>
          </h2>
          <p className="text-xl text-muted-foreground max-w-3xl mx-auto">
            BanglaCode is 100% open source. Every line of code is available on GitHub.
            Join us in making programming accessible to every Bengali speaker.
          </p>
        </motion.div>

        {/* Stats */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="grid grid-cols-2 md:grid-cols-4 gap-6 mb-20"
        >
          {statsDisplay.map((stat, index) => (
            <motion.div
              key={index}
              initial={{ opacity: 0, scale: 0.8 }}
              whileInView={{ opacity: 1, scale: 1 }}
              viewport={{ once: true }}
              transition={{ delay: index * 0.1 }}
              whileHover={{ scale: 1.05, y: -5 }}
              className="bg-card border border-border rounded-2xl p-6 text-center hover:border-primary/50 transition-all duration-300"
            >
              <stat.icon className="w-8 h-8 text-primary mx-auto mb-3" />
              <div className="text-4xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-500">
                {stat.value}
              </div>
              <div className="text-muted-foreground mt-1">{stat.label}</div>
            </motion.div>
          ))}
        </motion.div>

        {/* Contribution Ways */}
        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6 mb-16">
          {contributionWays.map((way, index) => (
            <motion.div
              key={index}
              initial={{ opacity: 0, y: 30 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: index * 0.1 }}
              whileHover={{ y: -10 }}
              className="group"
            >
              <div className="h-full bg-card border border-border rounded-2xl p-6 hover:border-primary/50 transition-all duration-300 hover:shadow-xl flex flex-col">
                <div className={`w-14 h-14 rounded-xl ${way.bgColor} flex items-center justify-center mb-5 group-hover:scale-110 transition-transform`}>
                  <way.icon className={`w-7 h-7 ${way.color}`} />
                </div>
                <h3 className="text-xl font-bold mb-3">{way.title}</h3>
                <p className="text-muted-foreground mb-6 flex-grow">{way.description}</p>
                <Link
                  href={way.link}
                  target="_blank"
                  className={`inline-flex items-center gap-2 ${way.color} hover:underline font-medium`}
                >
                  {way.linkText} →
                </Link>
              </div>
            </motion.div>
          ))}
        </div>

        {/* CTA */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="text-center"
        >
          <Link
            href="https://github.com/nexoral/BanglaCode"
            target="_blank"
            className="inline-flex items-center gap-3 px-10 py-5 text-xl font-semibold text-white bg-gradient-to-r from-gray-800 to-gray-900 hover:from-gray-700 hover:to-gray-800 rounded-full transition-all duration-300 hover:shadow-[0_0_40px_rgba(0,0,0,0.3)] border border-gray-700"
          >
            <FaGithub className="w-7 h-7" />
            Star us on GitHub
          </Link>
          <p className="mt-4 text-muted-foreground">
            Your star helps us reach more students and educators!
          </p>
        </motion.div>
      </div>
    </section>
  );
}
