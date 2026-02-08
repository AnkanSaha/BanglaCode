"use client";

import Link from "next/link";
import { ArrowLeft, ArrowRight } from "lucide-react";

export default function DocsLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex flex-col lg:flex-row min-h-[calc(100vh-4rem)]">
      {/* Sidebar is global but we need to ensure content is pushed */}
      <div className="lg:pl-64 flex-1">
        <div className="max-w-4xl mx-auto px-6 py-12">
          <div className="prose prose-invert prose-purple max-w-none">
            {children}
          </div>

          <div className="mt-16 flex justify-between border-t border-border pt-8">
            {/* Navigation logic can be added here or in individual pages */}
          </div>
        </div>
      </div>
    </div>
  );
}
