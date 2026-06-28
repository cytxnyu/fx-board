<template>
  <el-container class="home-container">
    <div class="content-wrapper">
      <p class="eyebrow">FX insights community</p>
      <h1 class="title">FxBoard</h1>
      <p class="description">
        面向个人投资者和外汇学习者的汇率工具与市场观点社区，聚合汇率换算、资讯阅读和互动点赞等核心能力。
      </p>
    </div>
  </el-container>
</template>

<script setup lang="ts">
</script>

<style scoped>
.home-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: calc(100vh - 100px);
  background-color: #f5f7fa;
  padding: 20px;
  box-sizing: border-box;
}

.content-wrapper {
  text-align: center;
  max-width: 820px;
}

.eyebrow {
  color: #409eff;
  font-size: 14px;
  font-weight: 600;
  margin-bottom: 12px;
  text-transform: uppercase;
}

.title {
  color: #1f2d3d;
  font-size: 42px;
  font-weight: bold;
  margin-bottom: 20px;
}

.description {
  color: #5f6f83;
  font-size: 18px;
  line-height: 1.7;
}
</style>
