<template>
  <div class="add-charge-container">
    <h2>Добавить начисление</h2>

    <form @submit.prevent="submitForm">
      <div class="form-group">
        <label>Категория</label>
        <input v-model="category" type="text" required />
      </div>

      <div class="form-group">
        <label>Сумма (₽)</label>
        <input v-model.number="amount" type="number" min="0" required />
      </div>

      <div class="form-group">
        <label>Дата</label>
        <input v-model="date" type="date" required />
      </div>

      <button type="submit">Добавить</button>
    </form>

    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { createCharge } from '@/shared/api/charges';
import { useRouter } from 'vue-router';

const router = useRouter();
const auth = useAuthStore();

const category = ref('');
const amount = ref(0);
const date = ref('');
const message = ref('');

async function submitForm() {
  try {
    await createCharge({
      user_id: auth.userId,
      category: category.value,
      amount: amount.value,
      date: date.value,
    });

    message.value = 'Начисление добавлено!';
    setTimeout(() => {
      router.push('/charges');
    }, 1000);
  } catch (e) {
    message.value = 'Ошибка при добавлении начисления';
  }
}
</script>

<style scoped>
.add-charge-container {
  max-width: 500px;
  margin: 2rem auto;
  padding: 2rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}
form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
.form-group {
  display: flex;
  flex-direction: column;
}
label {
  margin-bottom: 0.5rem;
  font-weight: 600;
}
input {
  padding: 0.5rem;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
}
button {
  padding: 0.75rem;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}
button:hover {
  background-color: #1d4ed8;
}
</style>
