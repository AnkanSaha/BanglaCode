"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import {
  FaGithub,
  FaTwitter,
  FaLinkedin,
  FaInstagram,
  FaDiscord,
  FaGlobe,
} from "react-icons/fa";
import { MapPin, Briefcase, Code2, Heart, Users, FolderGit2, Star } from "lucide-react";
import { getCreatorData, getNexoralProjects, type CreatorData, type NexoralProject } from "@/lib/github";

const STATIC_DATA = {
  website: "https://ankan.in",
  twitter: "theankansaha",
  linkedin: "theankansaha",
  instagram: "theankansaha",
  discord: "",
  skills: ["Go", "TypeScript", "React", "Node.js", "Python", "System Design", "Open Source"],
};

const FALLBACK_CREATOR: CreatorData = {
  name: "Ankan Saha",
  username: "AnkanSaha",
  bio: "Software Engineer | Building BanglaCode to make programming accessible to Bengali speakers",
  location: "India",
  avatar: "https://avatars.githubusercontent.com/u/90076852",
  website: "https://ankan.in",
  followers: 100,
  following: 0,
  publicRepos: 50,
};

export default function CreatorSection() {
  const [creator, setCreator] = useState<CreatorData>(FALLBACK_CREATOR);
  const [projects, setProjects] = useState<NexoralProject[]>([]);

  useEffect(() => {
    getCreatorData().then(setCreator);
    getNexoralProjects().then(setProjects);
  }, []);

  const socialLinks = [
    { icon: FaGlobe, href: STATIC_DATA.website, label: "Website", color: "hover:text-blue-400" },
    { icon: FaGithub, href: `https://github.com/${creator.username}`, label: "GitHub", color: "hover:text-gray-400" },
    { icon: FaTwitter, href: `https://twitter.com/${STATIC_DATA.twitter}`, label: "Twitter", color: "hover:text-sky-400" },
    { icon: FaLinkedin, href: `https://linkedin.com/in/${STATIC_DATA.linkedin}`, label: "LinkedIn", color: "hover:text-blue-500" },
    { icon: FaInstagram, href: `https://instagram.com/${STATIC_DATA.instagram}`, label: "Instagram", color: "hover:text-pink-500" },
    { icon: FaDiscord, href: "#", label: "Discord", color: "hover:text-indigo-400" },
  ];

  return (
    <section className="py-32 bg-gradient-to-b from-background via-accent/5 to-background relative overflow-hidden">
      {/* Background Elements */}
      <div className="absolute inset-0">
        <div className="absolute top-1/3 left-0 w-[500px] h-[500px] bg-purple-500/10 rounded-full blur-[60px]" />
        <div className="absolute bottom-1/3 right-0 w-[500px] h-[500px] bg-pink-500/10 rounded-full blur-[60px]" />
      </div>

      <div className="container mx-auto px-4 relative z-10">
        <div className="text-center mb-20 animate-fade-in-up">
          <span className="inline-block px-4 py-2 text-sm font-semibold text-primary bg-primary/10 rounded-full border border-primary/20 mb-6">
            Meet the Creator
          </span>
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Built with{" "}
            <span className="inline-flex items-center gap-2">
              <Heart className="w-10 h-10 text-red-500" fill="currentColor" />
            </span>{" "}
            in India
          </h2>
        </div>

        <div className="max-w-5xl mx-auto">
          <div className="relative animate-fade-in-up" style={{ animationDelay: "0.1s" }}>
            {/* Glow Effect */}
            <div className="absolute -inset-1 bg-gradient-to-r from-purple-500 via-pink-500 to-purple-500 rounded-3xl opacity-20 blur-xl" />

            {/* Card */}
            <div className="relative bg-card border border-border rounded-3xl overflow-hidden">
              <div className="grid md:grid-cols-[300px,1fr] gap-0">
                {/* Left - Avatar & Basic Info */}
                <div className="bg-gradient-to-b from-purple-500/20 to-pink-500/20 p-8 flex flex-col items-center justify-center border-r border-border/50">
                  <div className="relative">
                    <div className="absolute -inset-2 bg-gradient-to-r from-purple-500 to-pink-500 rounded-full opacity-50 blur-md" />
                    {/* eslint-disable-next-line @next/next/no-img-element */}
                    <img
                      src={creator.avatar}
                      alt={creator.name}
                      width={180}
                      height={180}
                      className="relative rounded-full border-4 border-background shadow-2xl w-[180px] h-[180px] object-cover"
                      loading="lazy"
                    />
                  </div>

                  <div className="mt-6 text-center">
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
                      <div className="flex items-center gap-1 text-sm text-muted-foreground">
                        <Users className="w-4 h-4" />
                        <span className="font-semibold text-foreground">{creator.followers}</span>
                        <span>followers</span>
                      </div>
                      <div className="flex items-center gap-1 text-sm text-muted-foreground">
                        <FolderGit2 className="w-4 h-4" />
                        <span className="font-semibold text-foreground">{creator.publicRepos}</span>
                        <span>repos</span>
                      </div>
                    </div>
                  </div>

                  {/* Social Links */}
                  <div className="flex gap-3 mt-6">
                    {socialLinks.map((social, index) => (
                      <Link
                        key={index}
                        href={social.href}
                        target="_blank"
                        className={`p-2 bg-background/50 rounded-lg border border-border/50 transition-all duration-200 hover:scale-110 ${social.color}`}
                        aria-label={social.label}
                      >
                        <social.icon className="w-5 h-5" />
                      </Link>
                    ))}
                  </div>
                </div>

                {/* Right - Details */}
                <div className="p-8">
                  <div>
                    <h4 className="text-lg font-semibold mb-3 flex items-center gap-2">
                      <Code2 className="w-5 h-5 text-primary" />
                      About
                    </h4>
                    <p className="text-muted-foreground leading-relaxed mb-6">
                      {creator.bio}
                    </p>

                    <h4 className="text-lg font-semibold mb-3">Skills & Interests</h4>
                    <div className="flex flex-wrap gap-2 mb-8">
                      {STATIC_DATA.skills.map((skill, index) => (
                        <span
                          key={index}
                          className="px-3 py-1 bg-primary/10 text-primary text-sm rounded-full border border-primary/20"
                        >
                          {skill}
                        </span>
                      ))}
                    </div>

                    <h4 className="text-lg font-semibold mb-4">Other Projects</h4>
                    <div className="grid gap-3">
                      {projects.length > 0 ? (
                        projects.map((project, index) => (
                          <Link
                            key={index}
                            href={project.url}
                            target="_blank"
                            className="block p-4 bg-background/50 rounded-xl border border-border/50 hover:border-primary/50 transition-all duration-200 group hover:-translate-y-1"
                          >
                            <div className="flex items-start justify-between">
                              <div>
                                <h5 className="font-semibold group-hover:text-primary transition-colors duration-200">
                                  {project.name}
                                </h5>
                                <p className="text-sm text-muted-foreground mt-1">
                                  {project.description}
                                </p>
                              </div>
                              <div className="flex items-center gap-1 text-muted-foreground">
                                <Star className="w-4 h-4" />
                                <span className="text-sm">{project.stars}</span>
                              </div>
                            </div>
                          </Link>
                        ))
                      ) : (
                        <div className="grid gap-3">
                          {[1, 2, 3].map((i) => (
                            <div key={i} className="p-4 bg-background/50 rounded-xl border border-border/50 animate-pulse">
                              <div className="h-4 bg-muted rounded w-24 mb-2"></div>
                              <div className="h-3 bg-muted rounded w-48"></div>
                            </div>
                          ))}
                        </div>
                      )}
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Message */}
          <div className="mt-12 text-center animate-fade-in-up" style={{ animationDelay: "0.2s" }}>
            <blockquote className="text-xl md:text-2xl font-light italic text-muted-foreground max-w-3xl mx-auto">
              &quot;I believe technology should be accessible to everyone, regardless of their language.
              BanglaCode is my contribution to making this dream a reality for the Bengali-speaking world.&rdquo;
            </blockquote>
            <p className="mt-4 text-primary font-semibold">— {creator.name}</p>
          </div>
        </div>
      </div>
    </section>
  );
}
