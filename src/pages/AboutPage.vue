<script setup>
function opengithub(){
    window.open('https://github.com/ClaretWheel1481/FeasOJ')
}
</script>

<template>
    <div>
        <h1>关于</h1>
        <p>反馈问题或想加入开发？前往Github!</p>
    </div>
    <div class="parent">
        <div class="githublogo" @click="opengithub"></div>
    </div>
</template>

<style scoped>
.githublogo{
    width: 40px;
    height: 40px;
    background-image: url('https://res.hc-cdn.com/tinyui-design-common/1.0.5.20240401143336/assets/github.svg');
    background-size: cover;
    cursor: pointer;
}
</style>