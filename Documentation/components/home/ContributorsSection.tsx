"use client";

import Link from "next/link";
import { Users, GitCommit, ExternalLink } from "lucide-react";
import { FaGithub } from "react-icons/fa";

interface Contributor {
  login: string;
  avatar_url: string;
  html_url: string;
  contributions: number;
  name?: string;
}

interface ContributorsSectionProps {
  contributors: Contributor[];
}

export default function ContributorsSection({ contributors }: ContributorsSectionProps) {
  if (contributors.length === 0) return null;

  const totalContributions = contributors.reduce((sum, c) => sum + c.contributions, 0);

  return (
    <section className="py-24 relative overflow-hidden">
      {/* Background */}
      <div className="absolute inset-0 bg-gradient-to-b from-accent/5 via-background to-background" />
      <div className="absolute top-0 left-1/4 w-96 h-96 bg-purple-500/5 rounded-full blur-2xl" />
      <div className="absolute bottom-0 right-1/4 w-96 h-96 bg-pink-500/5 rounded-full blur-2xl" />

      <div className="container mx-auto px-4 relative z-10">
        <div className="text-center mb-16 animate-fade-in-up">
          <span className="inline-block px-4 py-2 text-sm font-semibold text-primary bg-primary/10 rounded-full border border-primary/20 mb-6">
            <Users className="w-4 h-4 inline mr-2" />
            Our Heroes
          </span>
          <h2 className="text-4xl md:text-5xl font-bold mb-4">
            Top{" "}
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-500">
              Contributors
            </span>
          </h2>
          <p className="text-xl text-muted-foreground max-w-2xl mx-auto">
            BanglaCode is built by amazing developers from around the world.
            <span className="text-primary font-semibold"> {totalContributions}+ contributions</span> and counting!
          </p>
        </div>

        {/* Contributors Grid */}
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-6 max-w-5xl mx-auto">
          {contributors.map((contributor, index) => (
            <div
              key={contributor.login}
              className="animate-fade-in-up"
              style={{ animationDelay: `${index * 0.03}s` }}
            >
              <Link
                href={contributor.html_url}
                target="_blank"
                className="group block"
              >
                <div className="relative bg-card border border-border rounded-2xl p-6 text-center transition-all duration-200 hover:border-primary/50 hover:shadow-xl hover:shadow-primary/5 hover:-translate-y-1 will-change-transform">
                  {/* Rank Badge for top 3 */}
                  {index < 3 && (
                    <div className={`absolute -top-2 -right-2 w-8 h-8 rounded-full flex items-center justify-center text-sm font-bold shadow-lg ${
                      index === 0 ? "bg-yellow-500 text-yellow-950" :
                      index === 1 ? "bg-gray-300 text-gray-800" :
                      "bg-amber-600 text-amber-100"
                    }`}>
                      {index + 1}
                    </div>
                  )}

                  {/* Avatar */}
                  <div className="relative mx-auto mb-4">
                    <div className="absolute -inset-1 bg-gradient-to-r from-purple-500 to-pink-500 rounded-full opacity-0 group-hover:opacity-50 blur transition-opacity duration-200" />
                    {/* eslint-disable-next-line @next/next/no-img-element */}
                    <img
                      src={contributor.avatar_url}
                      alt={contributor.login}
                      className="relative w-20 h-20 rounded-full border-2 border-border group-hover:border-primary/50 transition-colors duration-200 mx-auto"
                      loading="lazy"
                    />
                  </div>

                  {/* Info */}
                  <h3 className="font-semibold text-foreground group-hover:text-primary transition-colors duration-200 truncate">
                    {contributor.name || contributor.login}
                  </h3>
                  <p className="text-sm text-muted-foreground mt-1">
                    @{contributor.login}
                  </p>

                  {/* Contributions */}
                  <div className="flex items-center justify-center gap-1 mt-3 text-sm">
                    <GitCommit className="w-4 h-4 text-green-500" />
                    <span className="font-semibold text-green-500">{contributor.contributions}</span>
                    <span className="text-muted-foreground">commits</span>
                  </div>

                  {/* External link indicator */}
                  <ExternalLink className="w-4 h-4 absolute top-3 right-3 text-muted-foreground opacity-0 group-hover:opacity-100 transition-opacity duration-200" />
                </div>
              </Link>
            </div>
          ))}
        </div>

        {/* CTA */}
        <div className="text-center mt-12 animate-fade-in-up" style={{ animationDelay: "0.3s" }}>
          <Link
            href="https://github.com/nexoral/BanglaCode/graphs/contributors"
            target="_blank"
            className="inline-flex items-center gap-2 px-6 py-3 bg-card border border-border rounded-full hover:border-primary/50 transition-all duration-200 text-muted-foreground hover:text-foreground hover:-translate-y-1 will-change-transform"
          >
            <FaGithub className="w-5 h-5" />
            View All Contributors
            <ExternalLink className="w-4 h-4" />
          </Link>
        </div>
      </div>
    </section>
  );
}
