<template>
  <div class="charges-container">
    <div class="charges-header">
      <h2>Ваши начисления</h2>
      <RouterLink to="/charges/add" class="add-button">+ Добавить</RouterLink>
    </div>

    <div v-if="isLoading">Загрузка...</div>
    <div v-else-if="filteredCharges.length === 0">Пока нет начислений</div>

    <div class="card-grid" v-else>
      <div v-for="charge in filteredCharges" :key="charge.id" class="charge-card">
        <h3>{{ charge.category }}</h3>
        <p>Сумма: {{ charge.amount }} ₽</p>
        <p>Дата: {{ formatDate(charge.date) }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { fetchCharges } from '@/shared/api/charges';

const auth = useAuthStore();
const charges = ref<any[]>([]);
const isLoading = ref(true);

onMounted(async () => {
  try {
    const data = await fetchCharges();
    charges.value = data;
    console.log('Загружены начисления:', data);
    console.log('auth.userId:', auth.userId);
  } catch (e) {
    console.error('Ошибка загрузки:', e);
  } finally {
    isLoading.value = false;
  }
});

// фильтруем только начисления текущего пользователя
const filteredCharges = computed(() =>
  charges.value.filter((c) => Number(c.user_id) === Number(auth.userId)),
);

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('ru-RU');
}
</script>

<style scoped>
.charges-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.04);
}

.charges-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.add-button {
  padding: 0.5rem 1rem;
  background: #2563eb;
  color: white;
  border-radius: 6px;
  text-decoration: none;
  font-weight: 500;
}
.add-button:hover {
  background: #1d4ed8;
}

.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 1rem;
  margin-top: 1rem;
}

.charge-card {
  background: #f1f5f9;
  border-radius: 8px;
  padding: 1rem;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}
</style>
