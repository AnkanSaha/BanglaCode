const GITHUB_API = "https://api.github.com";
const REPO_OWNER = "nexoral";
const REPO_NAME = "BanglaCode";
const CREATOR_USERNAME = "AnkanSaha";
const CACHE_TTL = 1000 * 60 * 30; // 30 minutes cache

export interface RepoStats {
  stars: number;
  forks: number;
  watchers: number;
  openIssues: number;
}

export interface Contributor {
  login: string;
  avatar_url: string;
  html_url: string;
  contributions: number;
  name: string;
}

export interface CreatorData {
  name: string;
  username: string;
  bio: string;
  location: string;
  avatar: string;
  website: string;
  followers: number;
  following: number;
  publicRepos: number;
}

export interface NexoralProject {
  name: string;
  description: string;
  url: string;
  stars: number;
}

interface CacheEntry<T> {
  data: T;
  timestamp: number;
}

function getCache<T>(key: string): T | null {
  if (typeof window === "undefined") return null;
  try {
    const cached = localStorage.getItem(`github_cache_${key}`);
    if (!cached) return null;
    const entry: CacheEntry<T> = JSON.parse(cached);
    if (Date.now() - entry.timestamp > CACHE_TTL) {
      localStorage.removeItem(`github_cache_${key}`);
      return null;
    }
    return entry.data;
  } catch {
    return null;
  }
}

function setCache<T>(key: string, data: T): void {
  if (typeof window === "undefined") return;
  try {
    const entry: CacheEntry<T> = { data, timestamp: Date.now() };
    localStorage.setItem(`github_cache_${key}`, JSON.stringify(entry));
  } catch {
    // localStorage full or disabled
  }
}

async function fetchJSON<T>(url: string): Promise<T | null> {
  try {
    const res = await fetch(url, {
      headers: { Accept: "application/vnd.github.v3+json" },
    });
    if (!res.ok) return null;
    return res.json();
  } catch {
    return null;
  }
}

export async function getRepoStats(): Promise<RepoStats> {
  const cached = getCache<RepoStats>("repo_stats");
  if (cached) return cached;

  const repo = await fetchJSON<{
    stargazers_count: number;
    forks_count: number;
    subscribers_count: number;
    open_issues_count: number;
  }>(`${GITHUB_API}/repos/${REPO_OWNER}/${REPO_NAME}`);

  const stats = {
    stars: repo?.stargazers_count ?? 0,
    forks: repo?.forks_count ?? 0,
    watchers: repo?.subscribers_count ?? 0,
    openIssues: repo?.open_issues_count ?? 0,
  };
  
  if (repo) setCache("repo_stats", stats);
  return stats;
}

export async function getContributorsCount(): Promise<number> {
  const cached = getCache<number>("contributors_count");
  if (cached !== null) return cached;

  const contributors = await fetchJSON<unknown[]>(
    `${GITHUB_API}/repos/${REPO_OWNER}/${REPO_NAME}/contributors?per_page=100`
  );
  const count = Array.isArray(contributors) ? contributors.length : 0;
  
  if (contributors) setCache("contributors_count", count);
  return count;
}

export async function getTopContributors(): Promise<Contributor[]> {
  const cached = getCache<Contributor[]>("top_contributors");
  if (cached) return cached;

  const contributors = await fetchJSON<
    { login: string; avatar_url: string; html_url: string; contributions: number }[]
  >(`${GITHUB_API}/repos/${REPO_OWNER}/${REPO_NAME}/contributors?per_page=10`);

  if (!Array.isArray(contributors)) return [];

  // Don't fetch individual user details to save API calls - use login as name
  const withNames = contributors.map((c) => ({
    login: c.login,
    avatar_url: c.avatar_url,
    html_url: c.html_url,
    contributions: c.contributions,
    name: c.login,
  }));

  setCache("top_contributors", withNames);
  return withNames;
}

export async function getCreatorData(): Promise<CreatorData> {
  const cached = getCache<CreatorData>("creator_data");
  if (cached) return cached;

  const user = await fetchJSON<{
    name: string | null;
    login: string;
    bio: string | null;
    location: string | null;
    avatar_url: string;
    blog: string | null;
    followers: number;
    following: number;
    public_repos: number;
  }>(`${GITHUB_API}/users/${CREATOR_USERNAME}`);

  const data = {
    name: user?.name ?? "Ankan Saha",
    username: user?.login ?? CREATOR_USERNAME,
    bio: user?.bio ?? "Software Engineer | Building BanglaCode",
    location: user?.location ?? "India",
    avatar: user?.avatar_url ?? `https://avatars.githubusercontent.com/u/90076852`,
    website: user?.blog ?? "https://ankan.in",
    followers: user?.followers ?? 0,
    following: user?.following ?? 0,
    publicRepos: user?.public_repos ?? 0,
  };

  if (user) setCache("creator_data", data);
  return data;
}

export async function getNexoralProjects(): Promise<NexoralProject[]> {
  const cached = getCache<NexoralProject[]>("nexoral_projects");
  if (cached) return cached;

  const projectNames = ["AxioDB", "NexoralDNS", "ContainDB", "xpack"];

  const results = await Promise.allSettled(
    projectNames.map(async (name) => {
      const repo = await fetchJSON<{
        name: string;
        description: string | null;
        html_url: string;
        stargazers_count: number;
      }>(`${GITHUB_API}/repos/${REPO_OWNER}/${name}`);

      if (!repo) return null;
      return {
        name: repo.name,
        description: repo.description ?? "No description",
        url: repo.html_url,
        stars: repo.stargazers_count,
      };
    })
  );

  const projects = results
    .filter((r): r is PromiseFulfilledResult<NexoralProject | null> => r.status === "fulfilled")
    .map((r) => r.value)
    .filter((r): r is NexoralProject => r !== null);

  if (projects.length > 0) setCache("nexoral_projects", projects);
  return projects;
}
