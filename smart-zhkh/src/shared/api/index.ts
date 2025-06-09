export async function registerUser(username: string, password: string) {
  const res = await fetch('http://localhost:8080/api/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });

  if (!res.ok) {
    const text = await res.text();
    throw new Error(text);
  }

  return res.text();
}
export async function loginUser(username: string, password: string) {
  const res = await fetch('http://localhost:8080/api/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });

  if (!res.ok) {
    const text = await res.text();
    throw new Error(text);
  }

  return res.json(); // например: { token: "..." }
}

export async function fetchCharges() {
  const res = await fetch('http://localhost:8080/api/charges', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!res.ok) {
    throw new Error('Не удалось загрузить начисления');
  }

  return await res.json();
}
