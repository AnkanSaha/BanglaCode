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

// Fetch GitHub repository data
async function getGitHubRepoData() {
  try {
    const [repoResponse, contributorsResponse] = await Promise.all([
      fetch("https://api.github.com/repos/nexoral/BanglaCode", {
        next: { revalidate: 3600 },
        headers: { Accept: "application/vnd.github.v3+json" },
      }),
      fetch("https://api.github.com/repos/nexoral/BanglaCode/contributors?per_page=10", {
        next: { revalidate: 3600 },
        headers: { Accept: "application/vnd.github.v3+json" },
      }),
    ]);

    if (!repoResponse.ok) throw new Error("Failed to fetch repo data");

    const repoData = await repoResponse.json();
    const contributorsData = contributorsResponse.ok
      ? await contributorsResponse.json()
      : [];

    return {
      stars: repoData.stargazers_count || 0,
      forks: repoData.forks_count || 0,
      watchers: repoData.subscribers_count || 0,
      contributors: Array.isArray(contributorsData) ? contributorsData.length : 0,
      openIssues: repoData.open_issues_count || 0,
    };
  } catch {
    return { stars: 0, forks: 0, watchers: 0, contributors: 0, openIssues: 0 };
  }
}

// Fetch top contributors with their profile info
async function getTopContributors() {
  try {
    const contributorsResponse = await fetch(
      "https://api.github.com/repos/nexoral/BanglaCode/contributors?per_page=10",
      {
        next: { revalidate: 3600 },
        headers: { Accept: "application/vnd.github.v3+json" },
      }
    );

    if (!contributorsResponse.ok) throw new Error("Failed to fetch contributors");

    const contributors = await contributorsResponse.json();

    if (!Array.isArray(contributors)) return [];

    // Fetch user details for each contributor to get their name
    const contributorsWithDetails = await Promise.all(
      contributors.map(async (contributor: { login: string; avatar_url: string; html_url: string; contributions: number }) => {
        try {
          const userResponse = await fetch(
            `https://api.github.com/users/${contributor.login}`,
            {
              next: { revalidate: 3600 },
              headers: { Accept: "application/vnd.github.v3+json" },
            }
          );

          if (userResponse.ok) {
            const userData = await userResponse.json();
            return {
              login: contributor.login,
              avatar_url: contributor.avatar_url,
              html_url: contributor.html_url,
              contributions: contributor.contributions,
              name: userData.name || contributor.login,
            };
          }
        } catch {
          // Fall back to basic info if user fetch fails
        }

        return {
          login: contributor.login,
          avatar_url: contributor.avatar_url,
          html_url: contributor.html_url,
          contributions: contributor.contributions,
          name: contributor.login,
        };
      })
    );

    return contributorsWithDetails;
  } catch {
    return [];
  }
}

// Fetch GitHub user data
async function getGitHubUserData() {
  try {
    const response = await fetch("https://api.github.com/users/AnkanSaha", {
      next: { revalidate: 3600 },
      headers: { Accept: "application/vnd.github.v3+json" },
    });

    if (!response.ok) throw new Error("Failed to fetch user data");

    const userData = await response.json();

    return {
      name: userData.name || "Ankan Saha",
      username: userData.login || "AnkanSaha",
      bio: userData.bio || "Software Engineer | Building BanglaCode",
      location: userData.location || "India",
      avatar: userData.avatar_url || "https://avatars.githubusercontent.com/u/AnkanSaha",
      website: userData.blog || "https://ankan.in",
      company: userData.company || "",
      followers: userData.followers || 0,
      following: userData.following || 0,
      publicRepos: userData.public_repos || 0,
    };
  } catch {
    return {
      name: "Ankan Saha",
      username: "AnkanSaha",
      bio: "Software Engineer",
      location: "India",
      avatar: "https://avatars.githubusercontent.com/u/90076852",
      website: "https://ankan.in",
      company: "",
      followers: 0,
      following: 0,
      publicRepos: 0,
    };
  }
}

// Fetch specific repos from Nexoral organization
async function getNexoralProjects() {
  const projectNames = ["AxioDB", "NexoralDNS", "ContainDB", "xpack"];

  try {
    const responses = await Promise.all(
      projectNames.map((name) =>
        fetch(`https://api.github.com/repos/nexoral/${name}`, {
          next: { revalidate: 3600 },
          headers: { Accept: "application/vnd.github.v3+json" },
        })
      )
    );

    const repos = await Promise.all(
      responses.map(async (res) => {
        if (!res.ok) return null;
        return res.json();
      })
    );

    return repos
      .filter((repo) => repo !== null)
      .map((repo: { name: string; description: string; html_url: string; stargazers_count: number }) => ({
        name: repo.name,
        description: repo.description || "No description",
        url: repo.html_url,
        stars: repo.stargazers_count,
      }));
  } catch {
    return [];
  }
}

export default async function Home() {
  // Fetch all data in parallel
  const [repoData, userData, nexoralProjects, topContributors] = await Promise.all([
    getGitHubRepoData(),
    getGitHubUserData(),
    getNexoralProjects(),
    getTopContributors(),
  ]);

  const creatorInfo = {
    name: userData.name,
    username: userData.username,
    bio: userData.bio,
    location: userData.location,
    avatar: userData.avatar,
    website: userData.website,
    twitter: "theankansaha",
    linkedin: "theankansaha",
    instagram: "theankansaha",
    discord: "theankansaha",
    followers: userData.followers,
    publicRepos: userData.publicRepos,
    skills: [
      "Go",
      "TypeScript",
      "Node.js",
      "React",
      "Next.js",
      "Docker",
      "MongoDB",
      "Redis",
      "AWS",
      "Linux",
      "DNS/TCP-IP",
      "Distributed Systems",
    ],
    projects: nexoralProjects.length > 0 ? nexoralProjects : [
      {
        name: "AxioDB",
        description: "Lightweight NoSQL database for Node.js",
        url: "https://github.com/nexoral/AxioDB",
      },
      {
        name: "NexoralDNS",
        description: "Docker-based DNS server for local networks",
        url: "https://github.com/nexoral/NexoralDNS",
      },
      {
        name: "ContainDB",
        description: "CLI tool for containerized database management",
        url: "https://github.com/nexoral/ContainDB",
      },
      {
        name: "xpack",
        description: "Universal Linux package builder",
        url: "https://github.com/nexoral/xpack",
      },
    ],
  };

  const communityStats = {
    stars: repoData.stars,
    contributors: repoData.contributors,
    forks: repoData.forks,
  };

  return (
    <div className="flex flex-col min-h-screen overflow-x-hidden">
      <HeroSection
        version={packageJson.version}
        stars={repoData.stars}
        forks={repoData.forks}
      />
      <StorySection />
      <FeaturesSection />
      <CodeShowcase />
      <ContributeSection stats={communityStats} />
      <ContributorsSection contributors={topContributors} />
      <CreatorSection creator={creatorInfo} />
      <Footer />
    </div>
  );
}
