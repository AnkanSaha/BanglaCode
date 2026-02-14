import { Metadata } from "next";
import HeroSection from "@/components/home/HeroSection";
import StorySection from "@/components/home/StorySection";
import FeaturesSection from "@/components/home/FeaturesSection";
import CodeShowcase from "@/components/home/CodeShowcase";
import ContributeSection from "@/components/home/ContributeSection";
import ContributorsSection from "@/components/home/ContributorsSection";
import CreatorSection from "@/components/home/CreatorSection";
import Footer from "@/components/home/Footer";
import packageJson from "../package.json";

export const metadata: Metadata = {
  title: "BanglaCode - Programming Language in Bengali",
  description:
    "BanglaCode is a programming language with Bengali (Banglish) syntax, designed to make coding accessible to 300 million Bengali speakers. Learn programming in your mother tongue.",
  keywords: [
    "BanglaCode",
    "Bengali programming language",
    "Bangla coding",
    "programming in Bengali",
    "learn coding Bengali",
    "Banglish programming",
  ],
  openGraph: {
    title: "BanglaCode - Programming Language in Bengali",
    description:
      "Learn programming in your mother tongue. BanglaCode makes coding accessible to Bengali speakers.",
    type: "website",
    url: "https://banglacode.pages.dev",
  },
  twitter: {
    card: "summary_large_image",
    title: "BanglaCode - Programming Language in Bengali",
    description:
      "Learn programming in your mother tongue. BanglaCode makes coding accessible to Bengali speakers.",
    creator: "@theankansaha",
  },
};

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen overflow-x-hidden">
      <HeroSection version={packageJson.version} />
      <StorySection />
      <FeaturesSection />
      <CodeShowcase />
      <ContributeSection />
      <ContributorsSection />
      <CreatorSection />
      <Footer />
    </div>
  );
}
