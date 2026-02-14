"use client";

import Link from "next/link";
import { FaGithub, FaTwitter, FaLinkedin, FaInstagram, FaDiscord } from "react-icons/fa";
import { Heart, Mail, MapPin } from "lucide-react";

const footerLinks = {
  product: [
    { name: "Documentation", href: "/docs" },
    { name: "Installation", href: "/docs/installation" },
    { name: "Quick Start", href: "/docs/quick-start" },
    { name: "Syntax Guide", href: "/docs/syntax" },
  ],
  resources: [
    { name: "Blog", href: "/blog" },
    { name: "Examples", href: "/docs/examples" },
    { name: "Keywords", href: "/docs/keywords" },
    { name: "Built-in Functions", href: "/docs/builtins" },
  ],
  community: [
    { name: "GitHub", href: "https://github.com/nexoral/BanglaCode" },
    { name: "Discussions", href: "https://github.com/nexoral/BanglaCode/discussions" },
    { name: "Issues", href: "https://github.com/nexoral/BanglaCode/issues" },
    { name: "Contributing", href: "https://github.com/nexoral/BanglaCode/blob/main/CONTRIBUTING.md" },
  ],
  legal: [
    { name: "MIT License", href: "https://github.com/nexoral/BanglaCode/blob/main/LICENSE" },
    { name: "Privacy Policy", href: "/privacy" },
    { name: "Terms of Use", href: "/terms" },
  ],
};

const socialLinks = [
  { icon: FaGithub, href: "https://github.com/nexoral/BanglaCode", label: "GitHub" },
  { icon: FaTwitter, href: "https://twitter.com/theankansaha", label: "Twitter" },
  { icon: FaLinkedin, href: "https://linkedin.com/in/theankansaha", label: "LinkedIn" },
  { icon: FaInstagram, href: "https://instagram.com/theankansaha", label: "Instagram" },
  { icon: FaDiscord, href: "#", label: "Discord" },
];

export default function Footer() {
  return (
    <footer className="relative overflow-hidden border-t border-border">
      {/* Background */}
      <div className="absolute inset-0 bg-gradient-to-b from-background to-accent/5" />
      <div className="absolute bottom-0 left-1/4 w-96 h-96 bg-purple-500/5 rounded-full blur-2xl" />
      <div className="absolute bottom-0 right-1/4 w-96 h-96 bg-pink-500/5 rounded-full blur-2xl" />

      <div className="container mx-auto px-4 py-16 relative z-10">
        <div className="grid grid-cols-2 md:grid-cols-6 gap-8 mb-12">
          {/* Brand */}
          <div className="col-span-2 animate-fade-in-up">
            <Link href="/" className="inline-block">
              <h3 className="text-2xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-purple-500 to-pink-500">
                BanglaCode
              </h3>
            </Link>
            <p className="text-muted-foreground mt-4 max-w-xs">
              A programming language in Bengali (Banglish) syntax, designed to make coding accessible to 300 million Bengali speakers.
            </p>

            <div className="flex items-center gap-2 mt-4 text-sm text-muted-foreground">
              <MapPin className="w-4 h-4" />
              <span>West Bengal, India</span>
            </div>

            <div className="flex items-center gap-2 mt-2 text-sm text-muted-foreground">
              <Mail className="w-4 h-4" />
              <a href="mailto:connect@ankan.in" className="hover:text-primary transition-colors duration-200">
                connect@ankan.in
              </a>
            </div>

            {/* Social Links */}
            <div className="flex gap-3 mt-6">
              {socialLinks.map((social, index) => (
                <a
                  key={index}
                  href={social.href}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="p-2 bg-card border border-border rounded-lg hover:border-primary/50 hover:-translate-y-1 transition-all duration-200 will-change-transform"
                  aria-label={social.label}
                >
                  <social.icon className="w-5 h-5" />
                </a>
              ))}
            </div>
          </div>

          {/* Links */}
          <div className="animate-fade-in-up" style={{ animationDelay: "0.05s" }}>
            <h4 className="font-semibold mb-4">Product</h4>
            <ul className="space-y-3">
              {footerLinks.product.map((link, index) => (
                <li key={index}>
                  <Link
                    href={link.href}
                    className="text-muted-foreground hover:text-primary transition-colors duration-200"
                  >
                    {link.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          <div className="animate-fade-in-up" style={{ animationDelay: "0.1s" }}>
            <h4 className="font-semibold mb-4">Resources</h4>
            <ul className="space-y-3">
              {footerLinks.resources.map((link, index) => (
                <li key={index}>
                  <Link
                    href={link.href}
                    className="text-muted-foreground hover:text-primary transition-colors duration-200"
                  >
                    {link.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          <div className="animate-fade-in-up" style={{ animationDelay: "0.15s" }}>
            <h4 className="font-semibold mb-4">Community</h4>
            <ul className="space-y-3">
              {footerLinks.community.map((link, index) => (
                <li key={index}>
                  <Link
                    href={link.href}
                    target="_blank"
                    className="text-muted-foreground hover:text-primary transition-colors duration-200"
                  >
                    {link.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          <div className="animate-fade-in-up" style={{ animationDelay: "0.2s" }}>
            <h4 className="font-semibold mb-4">Legal</h4>
            <ul className="space-y-3">
              {footerLinks.legal.map((link, index) => (
                <li key={index}>
                  <Link
                    href={link.href}
                    className="text-muted-foreground hover:text-primary transition-colors duration-200"
                  >
                    {link.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>
        </div>

        {/* Bottom Bar */}
        <div className="pt-8 border-t border-border flex flex-col md:flex-row items-center justify-between gap-4 animate-fade-in-up" style={{ animationDelay: "0.25s" }}>
          <p className="text-muted-foreground text-sm text-center md:text-left">
            © {new Date().getFullYear()} BanglaCode. All rights reserved.
          </p>

          <p className="text-muted-foreground text-sm flex items-center gap-2">
            Made with{" "}
            <Heart className="w-4 h-4 text-red-500" fill="currentColor" />
            {" "}in West Bengal, India
          </p>

          <p className="text-muted-foreground text-sm">
            Part of{" "}
            <Link
              href="https://nexoral.in"
              target="_blank"
              className="text-primary hover:underline"
            >
              Nexoral
            </Link>
          </p>
        </div>
      </div>
    </footer>
  );
}
