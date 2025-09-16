<template>
  <div class="user-management">
    <h1>User Management</h1>

    <!-- Add User Form -->
    <form @submit.prevent="addUser">
      <input v-model="name" placeholder="Name" required />
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Add User</button>
    </form>

    <!-- User List -->
    <h2>Users</h2>
    <ul>
      <li v-for="user in users" :key="user.id">
        {{ user.name }} - {{ user.email }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "UserManagement",
  data() {
    return {
      users: [],
      name: "",
      email: "",
      password: ""
    };
  },
  methods: {
    // Fetch all users
    async fetchUsers() {
      try {
        const response = await axios.get("http://localhost:8081/users"); // ✅ backend
        this.users = response.data;
      } catch (error) {
        console.error("❌ Error fetching users:", error);
      }
    },

    // Add new user
    async addUser() {
      try {
        await axios.post("http://localhost:8081/users", {
          name: this.name,
          email: this.email,
          password: this.password
        });
        this.name = "";
        this.email = "";
        this.password = "";
        this.fetchUsers(); // Refresh list after adding
      } catch (error) {
        console.error("❌ Error adding user:", error);
      }
    }
  },
  mounted() {
    this.fetchUsers();
  }
};
</script>

<style scoped>
.user-management {
  max-width: 600px;
  margin: auto;
  padding: 20px;
  background: #f9f9f9;
  border-radius: 10px;
}
form {
  margin-bottom: 20px;
}
input {
  margin: 5px;
  padding: 8px;
}
button {
  padding: 8px 15px;
  background: #42b983;
  color: white;
  border: none;
  cursor: pointer;
}
button:hover {
  background: #2c9e70;
}
</style>

