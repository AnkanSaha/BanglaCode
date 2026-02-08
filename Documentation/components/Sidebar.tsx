"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import clsx from "clsx";
import { DOCS_CONFIG } from "@/lib/docs-config";

export default function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="fixed top-16 left-0 w-64 h-[calc(100vh-4rem)] border-r border-border bg-background/50 backdrop-blur-sm hidden lg:block overflow-y-auto">
      <div className="p-6 space-y-8">
        {DOCS_CONFIG.map((section) => (
          <div key={section.section}>
            <div className="flex items-center gap-2 mb-4 text-primary/80">
              <section.icon className="w-4 h-4" />
              <h5 className="font-semibold text-sm uppercase tracking-wide">{section.section}</h5>
            </div>
            <ul className="space-y-1 relative border-l border-border ml-2 pl-4">
              {section.items.map((item) => (
                <li key={item.href}>
                  <Link
                    href={item.href}
                    className={clsx(
                      "block text-sm py-1 transition-all hover:text-primary hover:translate-x-1",
                      pathname === item.href
                        ? "text-primary font-medium translate-x-1"
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
