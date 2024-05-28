<script setup>

</script>

<template>
  <div class="title">
    <h1>登录</h1>
  </div>
  <v-sheet class="constrainsheet">
    <v-form fast-fail width="400px" class="mx-auto">
      <v-text-field rounded="xl" variant="solo-filled" label="用户名" />
      <v-text-field rounded="xl" variant="solo-filled" type="password" label="密码" />
      <v-row justify="end">
        <v-btn type="submit" color="primary" rounded="xl">登录</v-btn>
        <v-btn color="primary" variant="text" rounded="xl" @click="$router.push('/register')">注册</v-btn>
        <v-btn color="primary" variant="text" rounded="xl">忘记密码</v-btn>
      </v-row>
    </v-form>
  </v-sheet>
</template>

<style>

</style>