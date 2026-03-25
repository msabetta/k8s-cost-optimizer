const BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export async function fetchCosts() {
  const res = await fetch(`${BASE_URL}/costs`);
  if (!res.ok) throw new Error("Failed to fetch costs");
  return res.json();
}

export async function fetchRecommendations() {
  const res = await fetch(`${BASE_URL}/recommendations`);
  if (!res.ok) throw new Error("Failed to fetch recommendations");
  return res.json();
}

export async function fetchSummary() {
  const res = await fetch(`${BASE_URL}/summary`);
  if (!res.ok) throw new Error("Failed to fetch summary");
  return res.json();
}

export async function fetchForecast() {
  const res = await fetch(`${BASE_URL}/forecast`);
  if (!res.ok) throw new Error("Failed to fetch forecast");
  return res.json();
}
