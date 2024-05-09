<script setup>
import { ref } from 'vue'
import { NavMenu as TinyNavMenu } from '@opentiny/vue'
import { iconAdministrator,iconSetting } from '@opentiny/vue-icon'
import { Drawer as TinyDrawer, Button as TinyButton } from '@opentiny/vue'

function fn() {
  visible.value = true
}

const visible = ref(false)
const TinyIconAdministrator = iconAdministrator()
const TinyIconSetting = iconSetting()

const dataset = ref([
  { title: '首页', url: '/' },
  { title: '排名', url: '/top'},
  { title: '题目', url: '/problemset'},
  { title: '竞赛', url: '/contest'},
  { title: '关于', url: '/about'},
])
</script>

<template>
  <div class="preview">
    <tiny-nav-menu :data="dataset">
      <template #toolbar >
        <!-- TODO:跳转到登陆、注册页面 -->
          <router-link to="/">
            <tiny-icon-administrator class="icons-nav" ></tiny-icon-administrator>
          </router-link>
          <!-- TODO:打开菜单栏 -->
          <tiny-icon-setting class="icons-nav" @click="fn" type="primary"></tiny-icon-setting>
          <tiny-drawer title="设置" :visible="visible" @update:visible="visible = $event">
            <div>
              <br />
              <br />
              <span>内容区域</span>
            </div>
          </tiny-drawer>
      </template>
    </tiny-nav-menu>
  </div>
</template>

<style scoped>
.preview {
  min-height: 50px;
}

.preview .tiny-nav-menu {
  --ti-common-color-bg-navigation: linear-gradient(134deg, #5167f1 0%, #4e71f6 100%);
  --ti-nav-menu-item-hover-bg-color: rgba(12, 36, 127, 0.12);
  --ti-nav-menu-item-selected-underline-width: 8px;
  --ti-nav-menu-item-selected-underline-height: 2px;
  --ti-nav-menu-item-selected-border-radius: 1px;
  --ti-nav-menu-item-selected-position-bottom: 8px;
  --ti-nav-menu-submenu-padding-horizontal: 24px;
  box-shadow: 0px 7px 8px rgba(0,0,0,0.15);
}

.icons-nav{
  margin-left: 20px;
  margin-top: 10px;
  margin-bottom: 20px;
  fill: white;
  font-size: 28px;
}
</style>
