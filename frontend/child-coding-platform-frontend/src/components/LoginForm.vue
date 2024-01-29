<template>
  <div>
    <!-- 登录表单结构 -->
    <form @submit.prevent="login">
      <input v-model="username" type="text" placeholder="Username" />
      <input v-model="password" type="password" placeholder="Password" />
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    login() {
      // 登录逻辑
      console.log(this.username, this.password);
      // 发送请求到后端...
    },
  },
};
</script>

<style scoped>
/* 你的样式 */
</style>
