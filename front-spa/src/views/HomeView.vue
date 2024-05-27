<template>
  <div class="dashboard">
    <div class="left-section">
      <div class="tile user-tile">
        <h2>Welcome, {{ user.name }}</h2>
        <p>Email: {{ user.email }}</p>
      </div>
      <div class="tile bot-grades-tile">
        <h2>Your BOT Grades</h2>
        <p>Placeholder for BOT grades...</p>
      </div>
    </div>
    <div class="right-section">
      <div class="tile events-tile">
        <h2>Incoming Events</h2>
        <p>Placeholder for upcoming events...</p>
      </div>
    </div>
    <button class="logout-btn" @click="logout">Logout</button>
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
.dashboard {
  display: grid;
  grid-template-columns: 1fr 3fr;
  gap: 1rem;
  padding: 2rem;
  background-color: #f0f2f5;
  min-height: 100vh;
  box-sizing: border-box;
}

.left-section {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.right-section {
  display: flex;
  flex-direction: column;
}

.tile {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  box-sizing: border-box;
  height: 100%;
}

h2 {
  margin-bottom: 1rem;
  color: #333;
}

p {
  color: #666;
}

.logout-btn {
  margin-top: 1rem;
  padding: 0.5rem 1.5rem;
  font-size: 1rem;
  color: white;
  background-color: #007bff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
  align-self: flex-end;
}

.logout-btn:hover {
  background-color: #0056b3;
}
</style>
