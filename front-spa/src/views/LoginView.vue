<template>
  <div class="login-container">
    <h2 class="main-title">
      <span class="title-part-one">BOT 2024L</span><br>
      <span class="title-part-two">Microservices Vulnerabilities PoC</span>
    </h2>
    <div class="login-box">
    <form @submit.prevent="login">
      <h1>Login</h1>
      <input v-model="username" required placeholder="Username" class="login-input" />
      <input type="password" v-model="password" required placeholder="Password" class="login-input" />
      <button @click="login" class="login-button">Login</button>
    </form>
    <p v-if="error">{{ error }}</p>
    </div>
  </div>
</template>

<script>
import axios from '../axios';

export default {
  data() {
    return {
      username: '',
      password: '',
      error: null,
    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post('/api/auth/login', {
          username: this.username,
          password: this.password,
        });
        localStorage.setItem('token', response.data);
        this.$router.push('/home');
      } catch (err) {
        this.error = err
      }
    },
  },
};
</script>

<style scoped>
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(to right, #6a11cb, #2575fc);
}

.main-title {
  color: #fff;
  font-size: 2rem; /* Slightly larger */
  font-weight: 300; /* Lighter font-weight for a modern look */
  text-align: center;
  text-shadow: 1px 1px 4px rgba(0, 0, 0, 0.5); /* Soft shadow for better readability */
  margin-bottom: 2rem;
}

.title-part-one {
  font-weight: 700; /* Bold for emphasis */
  color: #ffd700; /* Gold color for a premium look */
  display: block; /* Ensures it takes up the full width */
  letter-spacing: 2px; /* More spacing for a refined appearance */
}

.title-part-two {
  font-size: 1.2rem; /* Smaller than the main part */
  color: #eef; /* Light blue for contrast */
  opacity: 0.9; /* Slightly transparent for depth */
}

.login-box {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-width: 400px;
  width: 100%;
}

.login-box h1 {
  margin-bottom: 1.5rem;
  color: #333;
}

.login-input {
  width: calc(100% - 2rem);
  padding: 0.75rem 1rem;
  margin: 0.5rem 0;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

.login-button {
  width: calc(100% - 2rem);
  padding: 0.75rem 1rem;
  margin-top: 1rem;
  border: none;
  border-radius: 4px;
  background: #2575fc;
  color: white;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.3s;
}

.login-button:hover {
  background: #1a5bb8;
}
</style>
