import {
  Book,
  Code,
  GitBranch,
  FunctionSquare,
  GraduationCap,
  List,
  FileText
} from "lucide-react";

export interface DocItem {
  name: string;
  href: string;
  description?: string;
}

export interface DocSection {
  section: string;
  icon: React.ComponentType<{ className?: string }>;
  items: DocItem[];
}

export const DOCS_CONFIG: DocSection[] = [
  {
    section: "Getting Started",
    icon: Book,
    items: [
      { name: "Introduction", href: "/docs", description: "What is BanglaCode?" },
      { name: "Installation", href: "/docs/installation", description: "How to install BanglaCode" },
      { name: "Quick Start", href: "/docs/quick-start", description: "Write your first program" },
    ],
  },
  {
    section: "Language Basics",
    icon: Code,
    items: [
      { name: "Keywords", href: "/docs/keywords", description: "All Bengali keywords" },
      { name: "Variables", href: "/docs/variables", description: "Declaring and using variables" },
      { name: "Data Types", href: "/docs/data-types", description: "Numbers, strings, booleans, etc." },
      { name: "Operators", href: "/docs/operators", description: "Arithmetic, comparison, logical" },
    ],
  },
  {
    section: "Control Flow",
    icon: GitBranch,
    items: [
      { name: "Conditionals", href: "/docs/conditionals", description: "If-else statements" },
      { name: "Loops", href: "/docs/loops", description: "While and for loops" },
      { name: "Break & Continue", href: "/docs/break-continue", description: "Loop control statements" },
    ],
  },
  {
    section: "Functions & OOP",
    icon: FunctionSquare,
    items: [
      { name: "Functions", href: "/docs/functions", description: "Defining and calling functions" },
      { name: "Classes", href: "/docs/classes", description: "Object-oriented programming" },
      { name: "Methods", href: "/docs/methods", description: "Class methods and this" },
    ],
  },
  {
    section: "Data Structures",
    icon: List,
    items: [
      { name: "Arrays", href: "/docs/arrays", description: "Working with arrays" },
      { name: "Maps", href: "/docs/maps", description: "Key-value collections" },
    ],
  },
  {
    section: "Advanced",
    icon: GraduationCap,
    items: [
      { name: "Async/Await", href: "/docs/async-await", description: "Asynchronous programming with promises" },
      { name: "Modules", href: "/docs/modules", description: "Import and export" },
      { name: "Error Handling", href: "/docs/error-handling", description: "Try-catch-finally" },
      { name: "Environment Variables", href: "/docs/environment-variables", description: "Loading and managing .env files" },
      { name: "File I/O", href: "/docs/file-io", description: "Reading and writing files" },
      { name: "HTTP Server", href: "/docs/http-server", description: "Building web servers" },
      { name: "Networking", href: "/docs/networking", description: "TCP, UDP, WebSocket" },
      { name: "Database", href: "/docs/database", description: "PostgreSQL, MySQL, MongoDB, Redis" },
    ],
  },
  {
    section: "Reference",
    icon: FileText,
    items: [
      { name: "Built-in Functions", href: "/docs/builtins", description: "All 135+ built-in functions" },
      { name: "Operator Precedence", href: "/docs/precedence", description: "Order of operations" },
      { name: "Examples", href: "/docs/examples", description: "Complete code examples" },
    ],
  },
];

// Flatten all docs for navigation
export function getAllDocs(): DocItem[] {
  return DOCS_CONFIG.flatMap(section => section.items);
}

// Get previous and next docs for navigation
export function getNavigation(currentHref: string): { prev: DocItem | null; next: DocItem | null } {
  const allDocs = getAllDocs();
  const currentIndex = allDocs.findIndex(doc => doc.href === currentHref);

  return {
    prev: currentIndex > 0 ? allDocs[currentIndex - 1] : null,
    next: currentIndex < allDocs.length - 1 ? allDocs[currentIndex + 1] : null,
  };
}
