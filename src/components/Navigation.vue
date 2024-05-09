<script setup>
import { ref } from 'vue'
import { NavMenu as TinyNavMenu } from '@opentiny/vue'
const dataset = ref([
  { title: '首页', url: '#/' },
  { title: '排名', url: '#/top'},
  { title: '关于', url: '#/about'},
])
</script>

<template>
  <div class="preview">
    <tiny-nav-menu :data="dataset">
      <template #toolbar >
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
}
</style>
