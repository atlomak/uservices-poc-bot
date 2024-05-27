<template>
  <div class="welcome">
    <h2>Welcome, {{ user }}</h2>
    <button @click="logout">Logout</button>
  </div>
</template>

<script>
// Import the configured Axios instance
import axios from '../axios';

export default {
  data() {
    return {
      user: {},
    };
  },
  async created() {
    try {
      const token = localStorage.getItem('token');
      const response = await axios.get('/api/user', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      this.user = response.data;
    } catch (err) {
      this.$router.push('/login');
    }
  },
  methods: {
    logout() {
      localStorage.removeItem('token');
      this.$router.push('/login');
    },
  },
};
</script>

<style scoped>
.welcome {
  text-align: center;
}
</style>
