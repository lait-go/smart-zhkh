<template>
  <div class="dashboard-container">
    <h2>Привет, {{ auth.username }} 👋</h2>
    <p>Добро пожаловать в личный кабинет</p>

    <div class="card-grid">
      <RouterLink to="/charges" class="info-card">💳 Мои начисления</RouterLink>
      <div class="info-card">🏠 Устройства (в разработке)</div>
      <div class="info-card">👤 Профиль (в разработке)</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';

const router = useRouter();
const auth = useAuthStore();

onMounted(() => {
  if (!auth.isLoggedIn) {
    router.push('/login');
  }
});
</script>

<style scoped>
.dashboard-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.04);
}
h2 {
  font-size: 26px;
  margin-bottom: 1rem;
  color: #1e293b;
}
p {
  color: #4b5563;
  margin-bottom: 2rem;
}
.card-grid {
  display: grid;
  gap: 1rem;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
}
.info-card {
  padding: 1.5rem;
  background: #f1f5f9;
  border-radius: 8px;
  text-align: center;
  font-size: 16px;
  font-weight: 500;
  color: #334155;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
  cursor: pointer;
  transition: 0.2s;
}
.info-card:hover {
  background: #e2e8f0;
}
</style>
