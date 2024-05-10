<script setup>
import { ref } from 'vue'
import { Divider, NavMenu as TinyNavMenu } from '@opentiny/vue'
import { iconAdministrator,iconSetting } from '@opentiny/vue-icon'
import { Drawer as TinyDrawer } from '@opentiny/vue'
import { Button as TinyButton } from '@opentiny/vue'

function fn() {
  visible.value = true
}

const visible = ref(false)
const TinyIconAdministrator = iconAdministrator()
const TinyIconSetting = iconSetting()

const dataset = ref([
  { title: '首页', url: '/' },
  { title: '排名', url: '/rank'},
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
          <router-link to="/login">
            <tiny-icon-administrator class="icons-nav" ></tiny-icon-administrator>
          </router-link>
          <!-- TODO:打开菜单栏 -->
          <tiny-icon-setting class="icons-nav" @click="fn" type="primary"></tiny-icon-setting>
          <tiny-drawer title="设置" :visible="visible" @update:visible="visible = $event" width="35%">
          <div>
            <!-- TODO：设置待处理 -->
            <div style="display: block; margin-top: 20px;">
              <tiny-button type="text" text="后台管理" @click="$router.push('/admin')"> </tiny-button>
            </div>
            <Divider></Divider>
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
  --ti-common-color-bg-navigation: linear-gradient(134deg, #6176ff 0%, #6176ff 100%);
  --ti-nav-menu-item-hover-bg-color: rgba(12, 36, 127, 0.12);
  --ti-nav-menu-item-selected-underline-width: 8px;
  --ti-nav-menu-item-selected-underline-height: 2px;
  --ti-nav-menu-item-selected-border-radius: 1px;
  --ti-nav-menu-item-selected-position-bottom: 8px;
  --ti-nav-menu-popmenu-node-title-padding-bottom: 8px;
  --ti-nav-menu-popmenu-selected-text-color: #5167f1;
  --ti-nav-menu-popmenu-selected-font-weight: bold;
  --ti-nav-menu-popmenu-more-item-active-bg-color: #fff;
  --ti-nav-menu-popmenu-more-menu-border-color-right: transparent;
  --ti-nav-menu-popmenu-more-item-before-width: 3px;
  --ti-nav-menu-popmenu-title-border-weight-bottom: 1px;
  --ti-nav-menu-popmenu-more-item-hover-icon-color: #fff;
  --ti-nav-menu-popmenu-more-item-icon-color: #fff;
  --ti-nav-menu-popmenu-text-color: #8a8e99;
  --ti-nav-menu-popmenu-more-item-text-color: #575d6c;
  --ti-nav-menu-popmenu-more-item-hover-bg-color: #f2f5fc;
  --ti-nav-menu-popmenu-title-border-color-bottom: #ebecf0;
  --ti-nav-menu-popmenu-more-item-selected-bg-color: #fff;
  --ti-nav-menu-popmenu-box-shadow: 0px 0px 14px 0px rgba(6, 19, 54, 0.1);
  --ti-nav-menu-submenu-padding-horizontal: 24px;
  box-shadow: 0px 7px 8px rgba(0,0,0,0.15);
}

.icons-nav{
  margin-left: 20px;
  margin-top: 10px;
  margin-bottom: 20px;
  fill: white;
  font-size: 26px;
  cursor: pointer;
}
</style>
