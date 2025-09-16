<template>
  <div class="container">
    <div class="card">
      <h2>User Registration</h2>
      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label>Name:</label>
          <input v-model="name" type="text" placeholder="Enter your name" required />
        </div>
        <div class="form-group">
          <label>Email:</label>
          <input v-model="email" type="email" placeholder="Enter your email" required />
        </div>
        <div class="form-group">
          <label>Password:</label>
          <input v-model="password" type="password" placeholder="Enter your password" required />
        </div>
        <button type="submit" class="btn">Register</button>
      </form>
      <p v-if="message" :class="{'success': messageColor==='green', 'error': messageColor==='red'}">
        {{ message }}
      </p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RegisterPage',
  data() {
    return {
      name: '',
      email: '',
      password: '',
      message: '',
      messageColor: 'green'
    }
  },
  methods: {
    async handleRegister() {
      try {
        const res = await axios.post('http://localhost:8080/register', {
          name: this.name,
          email: this.email,
          password: this.password
        })
        this.message = res.data.message
        this.messageColor = 'green'
      } catch (err) {
        this.message = err.response?.data?.message || 'Registration failed'
        this.messageColor = 'red'
      }
    }
  }
}
</script>

<style scoped>
/* Container to center card */
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(to right, #6a11cb, #2575fc);
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

/* Card style */
.card {
  background: white;
  padding: 40px;
  border-radius: 12px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
  width: 350px;
  text-align: center;
}

/* Form groups */
.form-group {
  margin-bottom: 20px;
  text-align: left;
}

/* Labels */
label {
  display: block;
  margin-bottom: 6px;
  font-weight: bold;
  color: #333;
}

/* Inputs */
input {
  width: 100%;
  padding: 10px;
  border-radius: 6px;
  border: 1px solid #ccc;
  box-sizing: border-box;
  transition: border 0.3s;
}

input:focus {
  border-color: #2575fc;
  outline: none;
}

/* Button */
.btn {
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 6px;
  background-color: #2575fc;
  color: white;
  font-size: 16px;
  cursor: pointer;
  transition: background 0.3s;
}

.btn:hover {
  background-color: #6a11cb;
}

/* Message styles */
.success {
  color: green;
  margin-top: 15px;
  font-weight: bold;
}

.error {
  color: red;
  margin-top: 15px;
  font-weight: bold;
}
</style>

