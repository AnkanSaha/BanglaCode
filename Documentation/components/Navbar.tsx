"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { motion } from "framer-motion";
import { Menu, X, Github } from "lucide-react";
import { useState } from "react";
import clsx from "clsx";

import { DOCS_CONFIG } from "@/lib/docs-config";

export default function Navbar() {
  const [isOpen, setIsOpen] = useState(false);
  const pathname = usePathname();

  const links = [
    { name: "Docs", href: "/docs" },
    { name: "Blog", href: "/blog" },
  ];

  return (
    <nav className="fixed top-0 left-0 right-0 z-50 bg-background/80 backdrop-blur-md border-b border-border">
      <div className="container mx-auto px-4 h-16 flex items-center justify-between">
        <Link href="/" className="text-xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-purple-500 to-pink-500">
          BanglaCode
        </Link>

        {/* Desktop Menu */}
        <div className="hidden md:flex items-center space-x-8">
          {links.map((link) => (
            <Link
              key={link.href}
              href={link.href}
              className={clsx(
                "text-sm font-medium transition-colors hover:text-primary",
                pathname?.startsWith(link.href) ? "text-primary" : "text-muted-foreground"
              )}
            >
              {link.name}
            </Link>
          ))}
          <a
            href="https://github.com/nexoral/BanglaCode"
            target="_blank"
            rel="noopener noreferrer"
            className="p-2 hover:bg-accent rounded-full transition-colors"
          >
            <Github className="w-5 h-5" />
          </a>
        </div>

        {/* Mobile Menu Button */}
        <button
          className="md:hidden p-2"
          onClick={() => setIsOpen(!isOpen)}
        >
          {isOpen ? <X /> : <Menu />}
        </button>
      </div>

      {/* Mobile Menu */}
      {isOpen && (
        <motion.div
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          className="md:hidden absolute top-16 left-0 right-0 bg-background border-b border-border flex flex-col max-h-[calc(100vh-4rem)] overflow-y-auto"
        >
          <div className="p-4 space-y-4">
            {links.map((link) => (
              <Link
                key={link.href}
                href={link.href}
                className={clsx(
                  "block text-lg font-medium hover:text-primary",
                  pathname === link.href ? "text-primary" : "text-foreground"
                )}
                onClick={() => setIsOpen(false)}
              >
                {link.name}
              </Link>
            ))}
          </div>

          {/* Docs Navigation for Mobile */}
          {pathname?.startsWith("/docs") && (
            <div className="p-4 border-t border-border bg-muted/10">
              <h4 className="text-xs font-semibold text-muted-foreground uppercase tracking-wider mb-4">Documentation</h4>
              <div className="space-y-6">
                {DOCS_CONFIG.map((section) => (
                  <div key={section.section}>
                    <div className="flex items-center gap-2 mb-2 text-primary">
                      <section.icon className="w-4 h-4" />
                      <span className="font-semibold text-sm">{section.section}</span>
                    </div>
                    <div className="pl-6 space-y-2">
                      {section.items.map((item) => (
                        <Link
                          key={item.href}
                          href={item.href}
                          className={clsx(
                            "block text-sm transition-colors",
                            pathname === item.href ? "text-foreground font-medium" : "text-muted-foreground"
                          )}
                          onClick={() => setIsOpen(false)}
                        >
                          {item.name}
                        </Link>
                      ))}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}
        </motion.div>
      )}
    </nav>
  );
}
