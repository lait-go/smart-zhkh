export async function createCharge(charge: {
  user_id: number;
  category: string;
  amount: number;
  date: string;
}) {
  const res = await fetch('http://localhost:8080/api/charges', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(charge),
  });

  if (!res.ok) {
    throw new Error('Ошибка создания начисления');
  }

  return await res.json();
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
