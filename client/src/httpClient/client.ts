
const SERVER_HOST = "http://localhost:8888";
export const AUTHENTICATION_TOKEN_KEY = "authentication_token";

type Method = "GET" | "POST" | "PUT" | "DELETE";

const sendRequest = async (method: Method, path: string,  data = {}) => {
  const url = new URL(path, SERVER_HOST);
  const token = localStorage.getItem(AUTHENTICATION_TOKEN_KEY);

  return await fetch(url, {
    method: method,
    mode: "cors",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`,
    },
    body: JSON.stringify(data),
  });
};

export async function Post(path: string, data = {}) {
  return await sendRequest("POST", path, data);
}

export async function Delete(path: string, data = {}) {
  return await sendRequest("DELETE", path, data);
}

