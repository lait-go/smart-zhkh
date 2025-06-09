<template>
  <div>
    <header class="site-header">
      <div class="container">
        <h1 class="logo" @click="goHome">🧠 IoT Dashboard</h1>

        <nav v-if="auth.isLoggedIn">
          <RouterLink to="/dashboard">Главная</RouterLink>
          <RouterLink to="/charges">Начисления</RouterLink>
          <RouterLink to="/settings">Настройки</RouterLink>
          <button @click="logout">Выйти</button>
        </nav>

        <nav v-else>
          <RouterLink to="/login">Вход</RouterLink>
          <RouterLink to="/register">Регистрация</RouterLink>
        </nav>
      </div>
    </header>

    <main class="main-content">
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';

const auth = useAuthStore();
const router = useRouter();

function logout() {
  auth.logout();
  router.push('/login');
}

function goHome() {
  if (auth.isLoggedIn) {
    router.push('/dashboard');
  } else {
    router.push('/login');
  }
}
</script>

<style scoped>
.site-header {
  background: #1e293b;
  color: white;
  padding: 1rem 0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
.container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}
.logo {
  font-size: 20px;
  font-weight: bold;
  cursor: pointer;
}
nav {
  display: flex;
  gap: 1rem;
}
nav a,
nav button {
  color: white;
  text-decoration: none;
  background: none;
  border: none;
  font-size: 16px;
  cursor: pointer;
}
nav a:hover,
nav button:hover {
  text-decoration: underline;
}
.main-content {
  padding: 2rem;
  background: #f9fafb;
  min-height: 80vh;
}
</style>
