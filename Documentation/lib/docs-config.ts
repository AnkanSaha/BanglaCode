import { Book, Code, Terminal, GraduationCap, Server, AlertTriangle, Layers } from "lucide-react";

export const DOCS_CONFIG = [
  {
    section: "Getting Started",
    icon: Book,
    items: [
      { name: "Introduction", href: "/docs" },
      { name: "Installation", href: "/docs/installation" },
      { name: "Quick Start", href: "/docs/quick-start" },
    ],
  },
  {
    section: "Core Concepts",
    icon: Layers,
    items: [
      { name: "Syntax & Variables", href: "/docs/syntax" },
      { name: "Control Flow", href: "/docs/control-flow" },
      { name: "Functions", href: "/docs/functions" },
      { name: "OOP (Classes)", href: "/docs/oop" },
    ],
  },
  {
    section: "Advanced",
    icon: GraduationCap,
    items: [
      { name: "Modules", href: "/docs/modules" },
      { name: "Error Handling", href: "/docs/error-handling" },
      { name: "HTTP Server", href: "/docs/http-server" },
    ]
  }
];
