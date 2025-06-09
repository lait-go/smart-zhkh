import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    username: localStorage.getItem('username') || '',
    token: localStorage.getItem('token') || '',
    userId: 0,
  }),

  actions: {
    login(username: string, token: string) {
      this.username = username;
      this.token = token;
      localStorage.setItem('username', username);
      localStorage.setItem('token', token);
    },

    logout() {
      this.username = '';
      this.token = '';
      localStorage.removeItem('username');
      localStorage.removeItem('token');
    },
  },

  getters: {
    isLoggedIn: (state) => !!state.token,
  },
});
