<template>
  <div class="page-wrapper">
    <div class="form-card">
      <h1>Регистрация</h1>

      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label for="username">Логин</label>
          <input v-model="username" id="username" type="text" required />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <input v-model="password" id="password" type="password" required />
        </div>

        <div class="form-group">
          <label for="confirm">Повторите пароль</label>
          <input v-model="confirmPassword" id="confirm" type="password" required />
        </div>

        <button type="submit" class="submit-button">Зарегистрироваться</button>
      </form>

      <ul v-if="errors.length" class="errors">
        <li v-for="(err, index) in errors" :key="index">{{ err }}</li>
      </ul>

      <p v-if="message" class="message">{{ message }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { registerUser } from '@/shared/api';
import { useAuthStore } from '@/shared/store/auth';

const router = useRouter();
const auth = useAuthStore();

const username = ref('');
const password = ref('');
const confirmPassword = ref('');
const errors = ref<string[]>([]);
const message = ref('');

async function handleRegister() {
  errors.value = [];

  if (username.value.trim() === '') {
    errors.value.push('Введите логин');
  }
  if (password.value.length < 6) {
    errors.value.push('Пароль должен быть от 6 символов');
  }
  if (password.value !== confirmPassword.value) {
    errors.value.push('Пароли не совпадают');
  }

  if (errors.value.length > 0) return;

  try {
    await registerUser(username.value, password.value);
    message.value = '✅ Регистрация успешна';
    auth.login(username.value, 'fake-token'); // можно заменить на реальный токен
    router.push('/dashboard');
  } catch (err: any) {
    message.value = '❌ ' + (err.message || 'Ошибка регистрации');
  }
}
</script>

<style scoped>
.page-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: #f5f7fa;
}
.form-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  width: 100%;
  max-width: 400px;
}
h1 {
  margin-bottom: 1.5rem;
  text-align: center;
  font-size: 24px;
  color: #333;
}
.form-group {
  margin-bottom: 1rem;
}
label {
  display: block;
  margin-bottom: 0.25rem;
  color: #555;
  font-weight: 500;
}
input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 6px;
  font-size: 16px;
}
input:focus {
  border-color: #3b82f6;
  outline: none;
}
.submit-button {
  width: 100%;
  padding: 0.6rem;
  background-color: #22c55e;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
}
.submit-button:hover {
  background-color: #16a34a;
}
.errors {
  color: #dc2626;
  margin-top: 1rem;
  list-style: none;
  padding-left: 0;
  font-size: 14px;
}
.message {
  margin-top: 1rem;
  color: #16a34a;
  text-align: center;
  font-weight: 500;
}
</style>
