"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import {
  FaGithub,
  FaTwitter,
  FaLinkedin,
  FaInstagram,
  FaDiscord,
  FaGlobe,
} from "react-icons/fa";
import { MapPin, Briefcase, Code2, Heart, Users, FolderGit2 } from "lucide-react";

interface CreatorSectionProps {
  creator: {
    name: string;
    username: string;
    bio: string;
    location: string;
    avatar: string;
    website: string;
    twitter: string;
    linkedin: string;
    instagram: string;
    discord: string;
    followers?: number;
    publicRepos?: number;
    skills: string[];
    projects: { name: string; description: string; url: string; stars?: number }[];
  };
}

export default function CreatorSection({ creator }: CreatorSectionProps) {
  const socialLinks = [
    { icon: FaGlobe, href: creator.website, label: "Website", color: "hover:text-blue-400" },
    { icon: FaGithub, href: `https://github.com/${creator.username}`, label: "GitHub", color: "hover:text-gray-400" },
    { icon: FaTwitter, href: `https://twitter.com/${creator.twitter}`, label: "Twitter", color: "hover:text-sky-400" },
    { icon: FaLinkedin, href: `https://linkedin.com/in/${creator.linkedin}`, label: "LinkedIn", color: "hover:text-blue-500" },
    { icon: FaInstagram, href: `https://instagram.com/${creator.instagram}`, label: "Instagram", color: "hover:text-pink-500" },
    { icon: FaDiscord, href: "#", label: "Discord", color: "hover:text-indigo-400" },
  ];

  return (
    <section className="py-32 bg-gradient-to-b from-background via-accent/5 to-background relative overflow-hidden">
      {/* Background Elements */}
      <div className="absolute inset-0">
        <div className="absolute top-1/3 left-0 w-[500px] h-[500px] bg-purple-500/10 rounded-full blur-[100px]" />
        <div className="absolute bottom-1/3 right-0 w-[500px] h-[500px] bg-pink-500/10 rounded-full blur-[100px]" />
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
            Meet the Creator
          </motion.span>
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Built with{" "}
            <span className="inline-flex items-center gap-2">
              <Heart className="w-10 h-10 text-red-500 animate-pulse" fill="currentColor" />
            </span>{" "}
            in India
          </h2>
        </motion.div>

        <div className="max-w-5xl mx-auto">
          <motion.div
            initial={{ opacity: 0, y: 30 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="relative"
          >
            {/* Glow Effect */}
            <div className="absolute -inset-1 bg-gradient-to-r from-purple-500 via-pink-500 to-purple-500 rounded-3xl opacity-20 blur-xl" />

            {/* Card */}
            <div className="relative bg-card border border-border rounded-3xl overflow-hidden">
              <div className="grid md:grid-cols-[300px,1fr] gap-0">
                {/* Left - Avatar & Basic Info */}
                <div className="bg-gradient-to-b from-purple-500/20 to-pink-500/20 p-8 flex flex-col items-center justify-center border-r border-border/50">
                  <motion.div
                    initial={{ scale: 0.8, opacity: 0 }}
                    whileInView={{ scale: 1, opacity: 1 }}
                    viewport={{ once: true }}
                    transition={{ delay: 0.2 }}
                    className="relative"
                  >
                    <div className="absolute -inset-2 bg-gradient-to-r from-purple-500 to-pink-500 rounded-full opacity-50 blur-md animate-pulse" />
                    {/* eslint-disable-next-line @next/next/no-img-element */}
                    <img
                      src={creator.avatar}
                      alt={creator.name}
                      width={180}
                      height={180}
                      className="relative rounded-full border-4 border-background shadow-2xl w-[180px] h-[180px] object-cover"
                    />
                  </motion.div>

                  <motion.div
                    initial={{ opacity: 0, y: 20 }}
                    whileInView={{ opacity: 1, y: 0 }}
                    viewport={{ once: true }}
                    transition={{ delay: 0.3 }}
                    className="mt-6 text-center"
                  >
                    <h3 className="text-2xl font-bold">{creator.name}</h3>
                    <p className="text-primary font-medium">@{creator.username}</p>

                    <div className="flex items-center justify-center gap-2 mt-3 text-muted-foreground">
                      <MapPin className="w-4 h-4" />
                      <span>{creator.location}</span>
                    </div>

                    <div className="flex items-center justify-center gap-2 mt-2 text-muted-foreground">
                      <Briefcase className="w-4 h-4" />
                      <span>Software Engineer</span>
                    </div>

                    {/* GitHub Stats */}
                    <div className="flex items-center justify-center gap-4 mt-4">
                      {creator.followers !== undefined && (
                        <div className="flex items-center gap-1 text-sm text-muted-foreground">
                          <Users className="w-4 h-4" />
                          <span className="font-semibold text-foreground">{creator.followers}</span>
                          <span>followers</span>
                        </div>
                      )}
                      {creator.publicRepos !== undefined && (
                        <div className="flex items-center gap-1 text-sm text-muted-foreground">
                          <FolderGit2 className="w-4 h-4" />
                          <span className="font-semibold text-foreground">{creator.publicRepos}</span>
                          <span>repos</span>
                        </div>
                      )}
                    </div>
                  </motion.div>

                  {/* Social Links */}
                  <motion.div
                    initial={{ opacity: 0 }}
                    whileInView={{ opacity: 1 }}
                    viewport={{ once: true }}
                    transition={{ delay: 0.4 }}
                    className="flex gap-3 mt-6"
                  >
                    {socialLinks.map((social, index) => (
                      <Link
                        key={index}
                        href={social.href}
                        target="_blank"
                        className={`p-2 bg-background/50 rounded-lg border border-border/50 transition-all hover:scale-110 ${social.color}`}
                        aria-label={social.label}
                      >
                        <social.icon className="w-5 h-5" />
                      </Link>
                    ))}
                  </motion.div>
                </div>

                {/* Right - Details */}
                <div className="p-8">
                  <motion.div
                    initial={{ opacity: 0, x: 20 }}
                    whileInView={{ opacity: 1, x: 0 }}
                    viewport={{ once: true }}
                    transition={{ delay: 0.3 }}
                  >
                    <h4 className="text-lg font-semibold mb-3 flex items-center gap-2">
                      <Code2 className="w-5 h-5 text-primary" />
                      About
                    </h4>
                    <p className="text-muted-foreground leading-relaxed mb-6">
                      {creator.bio}
                    </p>

                    <h4 className="text-lg font-semibold mb-3">Skills & Interests</h4>
                    <div className="flex flex-wrap gap-2 mb-8">
                      {creator.skills.map((skill, index) => (
                        <motion.span
                          key={index}
                          initial={{ opacity: 0, scale: 0.8 }}
                          whileInView={{ opacity: 1, scale: 1 }}
                          viewport={{ once: true }}
                          transition={{ delay: 0.4 + index * 0.05 }}
                          className="px-3 py-1 bg-primary/10 text-primary text-sm rounded-full border border-primary/20"
                        >
                          {skill}
                        </motion.span>
                      ))}
                    </div>

                    <h4 className="text-lg font-semibold mb-4">Other Projects</h4>
                    <div className="grid gap-3">
                      {creator.projects.map((project, index) => (
                        <motion.div
                          key={index}
                          initial={{ opacity: 0, x: 20 }}
                          whileInView={{ opacity: 1, x: 0 }}
                          viewport={{ once: true }}
                          transition={{ delay: 0.5 + index * 0.1 }}
                        >
                          <Link
                            href={project.url}
                            target="_blank"
                            className="block p-4 bg-background/50 rounded-xl border border-border/50 hover:border-primary/50 transition-all group"
                          >
                            <div className="flex items-start justify-between">
                              <div>
                                <h5 className="font-semibold group-hover:text-primary transition-colors">
                                  {project.name}
                                </h5>
                                <p className="text-sm text-muted-foreground mt-1">
                                  {project.description}
                                </p>
                              </div>
                              <span className="text-muted-foreground group-hover:text-primary transition-colors">
                                →
                              </span>
                            </div>
                          </Link>
                        </motion.div>
                      ))}
                    </div>
                  </motion.div>
                </div>
              </div>
            </div>
          </motion.div>

          {/* Message */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="mt-12 text-center"
          >
            <blockquote className="text-xl md:text-2xl font-light italic text-muted-foreground max-w-3xl mx-auto">
              "I believe technology should be accessible to everyone, regardless of their language.
              BanglaCode is my contribution to making this dream a reality for the Bengali-speaking world."
            </blockquote>
            <p className="mt-4 text-primary font-semibold">— {creator.name}</p>
          </motion.div>
        </div>
      </div>
    </section>
  );
}
