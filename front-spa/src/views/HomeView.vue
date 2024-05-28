<template>
  <div class="dashboard">
    <div class="left-section">
      <div class="tile user-tile">
        <h2>Welcome, {{ user.username }}</h2>
        <p>Email: {{ user.email }}</p>
      </div>
      <div class="tile bot-grades-tile">
        <h2>Your grades:</h2>
          <div v-if="grades.length">
            <p v-for="course in grades" :key="course.course_name">
              <strong>{{ course.course_name }}:</strong> {{ course.grades.join(', ') }}
            </p>
          </div>
        <p v-else>Loading grades...</p>
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
import axios from 'axios';

export default {
  data() {
    return {
      user: {},
      grades: [],
    };
  },
  async created() {
    try {
      const token = localStorage.getItem('token');
      
      // Fetch user data
      const userResponse = await axios.get('/api/auth/user', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      this.user = userResponse.data;
      
      // Fetch grades data
      const gradesResponse = await axios.get('/api/data/grades', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      this.grades = gradesResponse.data.courses;
    } catch (err) {
      console.log(err)
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

h3 {
  margin-top: 1rem;
  color: #555;
}

p {
  color: #666;
}

ul {
  margin: 0;
  padding: 0;
  list-style-type: none;
}

li {
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
