"use client";

import { motion } from "framer-motion";
import { Heart, Lightbulb, Target, Rocket } from "lucide-react";

const timeline = [
  {
    icon: Lightbulb,
    title: "The Spark",
    year: "The Beginning",
    description:
      "Growing up in West Bengal, I noticed how many brilliant minds struggled with English programming syntax. The logic was there, but the language barrier held them back.",
    color: "text-yellow-500",
    bgColor: "bg-yellow-500/10",
  },
  {
    icon: Heart,
    title: "The Passion",
    year: "The Motivation",
    description:
      "What if programming could feel like thinking in your own language? What if 'jodi' (if) and 'nahole' (else) could replace confusing English keywords?",
    color: "text-pink-500",
    bgColor: "bg-pink-500/10",
  },
  {
    icon: Target,
    title: "The Mission",
    year: "The Goal",
    description:
      "BanglaCode was born to democratize programming education. To help students in schools and colleges across Bengal learn coding without fighting a language barrier first.",
    color: "text-blue-500",
    bgColor: "bg-blue-500/10",
  },
  {
    icon: Rocket,
    title: "The Future",
    year: "The Vision",
    description:
      "We're building more than a language. We're building a bridge that connects 300 million Bengali speakers to the world of programming and technology.",
    color: "text-purple-500",
    bgColor: "bg-purple-500/10",
  },
];

export default function StorySection() {
  return (
    <section className="py-32 bg-gradient-to-b from-background via-accent/5 to-background relative overflow-hidden">
      {/* Background Elements */}
      <div className="absolute inset-0 opacity-30">
        <div className="absolute top-1/4 left-0 w-96 h-96 bg-purple-500/10 rounded-full blur-3xl" />
        <div className="absolute bottom-1/4 right-0 w-96 h-96 bg-pink-500/10 rounded-full blur-3xl" />
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
            Our Story
          </motion.span>
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Why{" "}
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-500">
              BanglaCode
            </span>
            ?
          </h2>
          <p className="text-xl text-muted-foreground max-w-3xl mx-auto">
            Every great project starts with a simple question. Ours was:
            <span className="text-primary font-semibold"> "Why should language be a barrier to learning logic?"</span>
          </p>
        </motion.div>

        {/* Timeline */}
        <div className="max-w-4xl mx-auto">
          <div className="relative">
            {/* Vertical Line */}
            <div className="absolute left-8 md:left-1/2 top-0 bottom-0 w-0.5 bg-gradient-to-b from-purple-500 via-pink-500 to-purple-500" />

            {timeline.map((item, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, x: index % 2 === 0 ? -50 : 50 }}
                whileInView={{ opacity: 1, x: 0 }}
                viewport={{ once: true, margin: "-100px" }}
                transition={{ duration: 0.6, delay: index * 0.2 }}
                className={`relative flex items-center mb-12 ${
                  index % 2 === 0 ? "md:flex-row" : "md:flex-row-reverse"
                }`}
              >
                {/* Icon */}
                <motion.div
                  whileHover={{ scale: 1.1, rotate: 360 }}
                  transition={{ duration: 0.5 }}
                  className={`absolute left-8 md:left-1/2 -translate-x-1/2 w-16 h-16 rounded-full ${item.bgColor} border-4 border-background flex items-center justify-center z-10`}
                >
                  <item.icon className={`w-7 h-7 ${item.color}`} />
                </motion.div>

                {/* Content Card */}
                <div
                  className={`ml-24 md:ml-0 md:w-[calc(50%-4rem)] ${
                    index % 2 === 0 ? "md:pr-8 md:text-right" : "md:pl-8 md:text-left"
                  }`}
                >
                  <motion.div
                    whileHover={{ scale: 1.02, y: -5 }}
                    className="bg-card border border-border rounded-2xl p-6 shadow-lg hover:shadow-xl hover:border-primary/30 transition-all duration-300"
                  >
                    <span className={`text-sm font-semibold ${item.color}`}>{item.year}</span>
                    <h3 className="text-2xl font-bold mt-1 mb-3">{item.title}</h3>
                    <p className="text-muted-foreground leading-relaxed">{item.description}</p>
                  </motion.div>
                </div>
              </motion.div>
            ))}
          </div>
        </div>

        {/* Quote */}
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="mt-20 text-center"
        >
          <blockquote className="text-2xl md:text-3xl font-light italic text-muted-foreground max-w-4xl mx-auto">
            "Programming is not about syntax. It's about{" "}
            <span className="text-primary font-semibold not-italic">problem-solving</span>. And problem-solving has no language."
          </blockquote>
          <p className="mt-6 text-lg font-semibold text-primary">— The Philosophy Behind BanglaCode</p>
        </motion.div>
      </div>
    </section>
  );
}
