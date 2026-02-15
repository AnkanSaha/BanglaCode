import { Metadata } from "next";
import HeroSection from "@/components/home/HeroSection";
import PositioningSection from "@/components/home/PositioningSection";
import StorySection from "@/components/home/StorySection";
import FeaturesSection from "@/components/home/FeaturesSection";
import CodeShowcase from "@/components/home/CodeShowcase";
import ContributeSection from "@/components/home/ContributeSection";
import ContributorsSection from "@/components/home/ContributorsSection";
import CreatorSection from "@/components/home/CreatorSection";
import Footer from "@/components/home/Footer";
import packageJson from "../package.json";

export const metadata: Metadata = {
  title: "BanglaCode - Educational Programming Language in Bengali",
  description:
    "BanglaCode is an educational programming language with Bengali syntax. Inspired by BhaiLang & Vedic, but with production-grade features—build backends, connect databases, write modular code. Perfect for 300 million Bengali speakers learning to code.",
  keywords: [
    "BanglaCode",
    "Bengali programming language",
    "educational programming language",
    "learn coding Bengali",
    "Banglish programming",
    "BhaiLang alternative",
    "Vedic programming",
    "Bengali education",
  ],
  openGraph: {
    title: "BanglaCode - Educational Programming Language in Bengali",
    description:
      "An educational language powerful enough for real projects. Learn programming in Bengali with databases, servers, and modules.",
    type: "website",
    url: "https://banglacode.pages.dev",
  },
  twitter: {
    card: "summary_large_image",
    title: "BanglaCode - Educational Programming Language in Bengali",
    description:
      "Educational language with production-grade features. Build backends, connect databases, write modules—all in Bengali.",
    creator: "@theankansaha",
  },
};

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen overflow-x-hidden">
      <HeroSection version={packageJson.version} />
      <PositioningSection />
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
