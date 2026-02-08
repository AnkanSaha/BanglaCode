"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import clsx from "clsx";

const sidebarLinks = [
  {
    section: "Getting Started",
    items: [
      { name: "Introduction", href: "/docs" },
      { name: "Installation", href: "/docs/installation" },
      { name: "Quick Start", href: "/docs/quick-start" },
    ],
  },
  {
    section: "Core Concepts",
    items: [
      { name: "Syntax & Variables", href: "/docs/syntax" },
      { name: "Control Flow", href: "/docs/control-flow" },
      { name: "Functions", href: "/docs/functions" },
      { name: "OOP (Classes)", href: "/docs/oop" },
    ],
  },
  {
    section: "Advanced",
    items: [
      { name: "Modules", href: "/docs/modules" },
      { name: "Error Handling", href: "/docs/error-handling" },
      { name: "HTTP Server", href: "/docs/http-server" },
    ]
  }
];

export default function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="fixed top-16 left-0 w-64 h-[calc(100vh-4rem)] border-r border-border bg-background/50 backdrop-blur-sm hidden lg:block overflow-y-auto">
      <div className="p-6 space-y-8">
        {sidebarLinks.map((section) => (
          <div key={section.section}>
            <h5 className="font-semibold text-foreground mb-4 text-sm">{section.section}</h5>
            <ul className="space-y-2">
              {section.items.map((item) => (
                <li key={item.href}>
                  <Link
                    href={item.href}
                    className={clsx(
                      "block text-sm transition-colors hover:text-primary",
                      pathname === item.href
                        ? "text-primary font-medium"
                        : "text-muted-foreground"
                    )}
                  >
                    {item.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </aside>
  );
}
