import { Metadata, Viewport } from "next";

export const viewport: Viewport = {
  themeColor: "#7c3aed",
};

export const metadata: Metadata = {
  title: {
    default: "BanglaCode - Programming in Bengali",
    template: "%s | BanglaCode",
  },
  description: "A modern, strict, interpreted programming language with Bengali keywords. Built for students to bridge logic and language.",
  keywords: ["Bengali Programming", "Bangla Code", "Learn Coding", "Bangla Programming Language", "West Bengal Coding"],
  authors: [{ name: "Ankan", url: "https://github.com/nexoral" }],
  openGraph: {
    type: "website",
    locale: "bn_IN",
    url: "https://banglacode.pages.dev",
    siteName: "BanglaCode",
    images: [
      {
        url: "/og-image.png", // We'll need to generate this or use a placeholder
        width: 1200,
        height: 630,
        alt: "BanglaCode - Coding in Bengali",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    site: "@banglacode",
    creator: "@nexoral",
  },
  icons: {
    icon: "/favicon.ico",
  },
};
