
const SERVER_HOST = "http://localhost:8888";
export const AUTHENTICATION_TOKEN_KEY = "authentication_token";

export async function post(path: string, data = {}) {
  const url = new URL(path, SERVER_HOST);
  const token = localStorage.getItem(AUTHENTICATION_TOKEN_KEY);

  return await fetch(url, {
    method: "POST",
    mode: "cors",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`,
    },
    body: JSON.stringify(data),
  });
}
